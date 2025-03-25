package version

// 打印当前的编译信息

// 2025/3/24

import (
	"fmt"
	"runtime"
	"strings"

	"github.com/elastic/beats/v7/libbeat/cmd/instance"
	"github.com/elastic/beats/v7/libbeat/common/cli"
	"github.com/elastic/beats/v7/libbeat/version"
	"github.com/spf13/cobra"
)

var (
	BuildDateTime      = ""
	BuildGitBranch     = ""
	BuildGitCommit     = ""
	BuildPackageModule = ""
)

var BuildInfo = PkgVersion{}

func init() {
	BuildInfo = PkgVersion{
		BuildGitBranch:     BuildGitBranch,
		BuildDateTime:      BuildDateTime,
		BuildGitCommit:     BuildGitCommit,
		BuildPackageModule: BuildPackageModule,
	}
}

type PkgVersion struct {
	BuildDateTime      string
	BuildGitBranch     string
	BuildGitCommit     string
	BuildPackageModule string
}

func (m PkgVersion) String() string {
	return m.BuildDateTime + "-" + m.BuildGitCommit
}

func (m PkgVersion) PrintVersion() {
	fmt.Printf("%s", m.TplVersion())
	fmt.Println()
}

func (m PkgVersion) TplVersion() string {
	str := strings.Builder{}
	str.WriteString(fmt.Sprintf("%15s: %s \n", "BuildTime", BuildDateTime))
	str.WriteString(fmt.Sprintf("%15s: %s \n", "BuildGitBranch", BuildGitBranch))
	str.WriteString(fmt.Sprintf("%15s: %s \n", "BuildGitCommit", BuildGitCommit))
	str.WriteString(fmt.Sprintf("%15s: %s \n", "BuildPkgModule", BuildPackageModule))
	return str.String()
}

func PrintVersion() {
	BuildInfo.PrintVersion()
}

// GenVersionCmd generates the command version for a Beat.
func GenVersionCmd(settings instance.Settings) *cobra.Command {
	return &cobra.Command{
		Use:   "version",
		Short: "Show current version info",
		Run: cli.RunWith(
			func(_ *cobra.Command, args []string) error {
				beat, err := instance.NewBeat(settings.Name, settings.IndexPrefix, settings.Version, settings.ElasticLicensed)
				if err != nil {
					return fmt.Errorf("error initializing beat: %s", err)
				}

				buildTime := "unknown"
				if bt := version.BuildTime(); !bt.IsZero() {
					buildTime = bt.String()
				}
				fmt.Printf("%s version %s (%s), libbeat %s [%s built %s]\n",
					beat.Info.Beat, beat.Info.Version, runtime.GOARCH, version.GetDefaultVersion(),
					version.Commit(), buildTime)

				fmt.Println()

				PrintVersion()

				return nil
			}),
	}
}

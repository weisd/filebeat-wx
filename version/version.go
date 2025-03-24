package version

// 打印当前的编译信息

// 2025/3/24

import (
	"fmt"
	"strings"
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

func (r PkgVersion) String() string {
	return r.BuildDateTime + "-" + r.BuildGitCommit
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

// -------------------------------- main 调用代码---------------------------------------
//package main
//
//import (
//"flag"
//"fmt"
//
//"wps.ktkt.com/debug/debug/version"
//)
//
//var printVersion = flag.Bool("v", false, "show build version for the program")
//
//func main() {
//	flag.Parse()
//
//	if *printVersion {
//		version.PrintVersion()
//	}
//
//	fmt.Println("test ------")
//}
// ** -------------------------------- main 调用代码 END---------------------------------------

// -------------------------------- build bash 代码---------------------------------------

//
// ** -------------------------------- build bash 代码 END---------------------------------------

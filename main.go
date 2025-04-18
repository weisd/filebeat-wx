package main

//import (
//	"os"
//
//	_ "github.com/weisd/filebeat-wx/lib/output/wechat"
//
//	"github.com/elastic/beats/v7/filebeat/cmd"
//)
//
//func main() {
//	if err := cmd.RootCmd.Execute(); err != nil {
//		os.Exit(1)
//	}
//}

import (
	"os"
	_ "time/tzdata" // for timezone handling

	_ "github.com/weisd/filebeat-wx/lib/output/wechat"
	"github.com/weisd/filebeat-wx/version"

	"github.com/elastic/beats/v7/filebeat/cmd"
	inputs "github.com/elastic/beats/v7/filebeat/input/default-inputs"
)

// The basic model of execution:
// - input: finds files in paths/globs to harvest, starts harvesters
// - harvester: reads a file, sends events to the spooler
// - spooler: buffers events until ready to flush to the publisher
// - publisher: writes to the network, notifies registrar
// - registrar: records positions of files read
// Finally, input uses the registrar information, on restart, to
// determine where in each file to restart a harvester.
func main() {
	beatCmd := cmd.Filebeat(inputs.Init, cmd.FilebeatSettings())

	beatCmd.RemoveCommand(beatCmd.VersionCmd)
	beatCmd.AddCommand(version.GenVersionCmd(cmd.FilebeatSettings()))

	if err := beatCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

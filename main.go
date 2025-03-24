package main

//import (
//	"os"
//
//	_ "wps.ktkt.com/kt/wechatoutput/lib/output/wechat"
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
	"flag"
	"os"
	_ "time/tzdata" // for timezone handling
	"wps.ktkt.com/kt/wechatoutput/version"

	_ "wps.ktkt.com/kt/wechatoutput/lib/output/wechat"

	"github.com/elastic/beats/v7/filebeat/cmd"
	inputs "github.com/elastic/beats/v7/filebeat/input/default-inputs"
)

var printVersion = flag.Bool("", false, "show build version for the program")

// The basic model of execution:
// - input: finds files in paths/globs to harvest, starts harvesters
// - harvester: reads a file, sends events to the spooler
// - spooler: buffers events until ready to flush to the publisher
// - publisher: writes to the network, notifies registrar
// - registrar: records positions of files read
// Finally, input uses the registrar information, on restart, to
// determine where in each file to restart a harvester.
func main() {

	flag.Parse()

	if *printVersion {
		version.PrintVersion()
		os.Exit(0)
	}

	if err := cmd.Filebeat(inputs.Init, cmd.FilebeatSettings()).Execute(); err != nil {
		os.Exit(1)
	}
}

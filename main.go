package main

import (
	"os"

	_ "wps.ktkt.com/kt/wechatoutput/lib/output/wechat"

	"github.com/elastic/beats/v7/filebeat/cmd"
)

func main() {
	if err := cmd.RootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

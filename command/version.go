package main

import (
	"bytes"
	"fmt"
	"runtime"
	"strings"
	"text/template"
	"time"

	"github.com/urfave/cli/v2"
)

var (
	version   = "major.minor.patch"
	branch    = "git/branch"
	revision  = "git/revision"
	buildDate = "yyyy-mm-ddThh:mm:ssZ"
	goVersion = runtime.Version()
)

const versionInfoTmpl = `
{{.version}} (branch: {{.branch}}; revision: {{.revision}})
build date:   {{.buildDate}}
go version:   {{.goVersion}}
`

func getVersion() string {
	m := map[string]string{
		"version":   version,
		"branch":    branch,
		"revision":  revision,
		"buildDate": buildDate,
		"goVersion": goVersion,
	}

	tmpl, err := template.Must(template.New("version"), nil).Parse(versionInfoTmpl)
	if err != nil {
		panic(err)
	}

	buf := bytes.Buffer{}
	if err := tmpl.Execute(&buf, m); err != nil {
		panic(err)
	}

	return strings.TrimSpace(buf.String())
}

func versionCommand() *cli.Command {
	return &cli.Command{
		Name:  "version",
		Usage: "输出版本信息",
		Action: func(context *cli.Context) error {
			fmt.Println(getVersion())
			return nil
		},
	}
}

func init() {
	if bt, err := time.Parse("2006-01-02T15:04:05MST", buildDate); err == nil {
		app.Compiled = bt
	}
	app.HideVersion = false
	app.Commands = append(app.Commands, versionCommand())
}

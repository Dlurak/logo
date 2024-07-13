package src

import (
	"fmt"
	"strings"
)

func Main() {
	props := parseProps()

	server(ServerConfig{
		Port:      props.Port,
		UseStdout: props.ShowServerInfo,
	}, func(data RequestBody) {
		line := data.fmtLine(props.ColorFull)

		if props.FilePath == "" {
			fmt.Print(line)
			return
		}

		if strings.HasSuffix(props.FilePath, ".json") {
			json, err := data.json()
			if err != nil {
				return
			}

			writeToFile(props.FilePath, json)
			return
		}

		writeToFile(props.FilePath, line)
	})
}

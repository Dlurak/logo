package src

import "fmt"

func Main() {
	props := parseProps()

	server(ServerConfig{
		Port:      props.Port,
		UseStdout: props.ShowServerInfo,
	}, func(data RequestBody) {
		line := fmt.Sprintf("[%s]: %s\n", fmtLogLevel(data.LogLevel, props.ColorFull), data.Message)

		if props.FilePath == "" {
			fmt.Print(line)
		} else {
			writeToFile(props.FilePath, line)
		}
	})
}

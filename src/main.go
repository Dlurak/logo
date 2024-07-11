package src

import "fmt"

func Main() {
	server(ServerConfig{
		Port:      3000,
		UseStdout: true,
	}, func(data RequestBody) {
		fmt.Printf("[%s]: %s\n", data.LogLevel, data.Message)
	})
}

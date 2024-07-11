package src

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type LogLevel string

const (
	DebugLevel LogLevel = "Debug"
	InfoLevel LogLevel = "Info"
	WarningLevel LogLevel = "Warning"
	ErrorLevel LogLevel = "Error"
	FatalLevel LogLevel = "Fatal"
)

type RequestBody struct {
	LogLevel LogLevel
	Message  string
}

type Callback func(RequestBody)

type ServerConfig struct {
	Port      int
	UseStdout bool
}

func server(config ServerConfig, callback Callback) {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
			return
		}

		var reqBody RequestBody
		decoder := json.NewDecoder(r.Body)
		err := decoder.Decode(&reqBody)
		if err != nil {
			http.Error(w, "Invalid JSON body", http.StatusBadRequest)
			return
		}
		if  !IsValidLogLevel(reqBody.LogLevel) {
			http.Error(w, "Invalid Log Level", http.StatusBadRequest)
			return
		}

		callback(reqBody)

		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "OK\n")
	})

	if config.UseStdout {
		fmt.Println("Starting the server")
	}

	err := http.ListenAndServe(fmt.Sprintf(":%d", config.Port), nil)
	if err != nil {
		fmt.Println("Error starting server", err)
	}
}

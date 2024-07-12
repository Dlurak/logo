package src

import "flag"

func colorAndFile(filepath string, colorfull bool) (string, bool) {
	if filepath == "" {
		return filepath, colorfull
	}

	return filepath, false
}

type Props struct {
	Port           int
	ShowServerInfo bool
	ColorFull      bool
	FilePath       string
}

func parseProps() Props {
	port := flag.Int("port", 3000, "The port to run the server on")
	showServerInfo := flag.Bool(
		"info",
		true,
		"Wheter simple information about the server lifecycle should be printed to stdout",
	)
	colorFull := flag.Bool(
		"color",
		true,
		"Colorfull output",
	)

	filepath := flag.String(
		"file",
		"",
		"The filepath to write the logs to, if empty/ommited stdout will be used",
	)

	flag.Parse()

	filepathValue, colorFullValue := colorAndFile(*filepath, *colorFull)

	return Props{
		Port:           *port,
		ShowServerInfo: *showServerInfo,
		ColorFull:      colorFullValue,
		FilePath:       filepathValue,
	}
}

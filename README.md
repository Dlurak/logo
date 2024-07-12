# Logo

Simple and very easy to use logging API.

---

## Features

- Colorful output
- Logging to file or stdout
- Supports a few flags to change the runtime behaviour

## Usage

```
Usage of logo:
  -color
        Colorfull output (default true)
  -file string
        The filepath to write the logs to, if empty/ommited stdout will be used
  -info
        Wheter simple information about the server lifecycle should be printed to stdout (default true)
  -port int
        The port to run the server on (default 3000)
```

### Building

Simply run

```bash
go build .
```

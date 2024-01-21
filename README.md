# Units Conversion Service

The Unit Conversion Service is a smart, user-friendly, and efficient software solution that enables seamless and accurate conversions between different units of measurement. This project has been deliberately designed to meet the demands of engineers, scientists, students, and professionals who frequently need to convert units during their daily work or complex projects.

## Technology Integration

- **gRPC Protocol**: Utilize gRPC for efficient, fast, and reliable communication between clients and servers.
- **Go**: Leverage the performance and concurrency features of Go for a high-performance server.

## Prerequisites
- Go (1.20 or later)
- Protocol Buffer Compiler (protoc)

## Installation

1. Clone this repository:
```bash
$ git clone https://github.com/ken1009us/unit-conversion-service.git
```

2. Navigate to the project directory:

```bash
$ cd unit-conversion-service
```

3. Install the required packages:

For Go:

```bash
go mod tidy
```

3. The `Makefile` simplifies the process of generating Protobuf and gRPC code. Use the following command:

```bash
make all
```

If you encounter an issue such as protoc-gen-go: program not found or is not executable, it means that the Protocol Buffer compiler cannot find the Go plugins (protoc-gen-go and protoc-gen-go-grpc). To resolve this, follow these steps:

- Install Protocol Compiler Plugins for Go:

Make sure you have the protoc (Protocol Buffer Compiler) installed on your system. Then, install the Go plugins for Protocol Buffers and gRPC:

```bash
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
```

- Ensure the PATH Includes the Go Bin Directory:

The installed plugins will be in your $GOPATH/bin directory or $HOME/go/bin if GOPATH is not set. Ensure this directory is in your system's PATH so that protoc can find the plugins.

You can add the Go bin directory to your PATH by adding the following line to your .bashrc, .zshrc, or equivalent shell configuration file:

```bash
export PATH="$PATH:$(go env GOPATH)/bin"
```

After adding this line, reload your shell configuration:

```bash
source ~/.bashrc  # or source ~/.zshrc
```

### Starting the gRPC Server

After generating the necessary code, you can start the gRPC server with the following command:

Start the gRPC server:

```bash
go run server/server.go
```

## Usage

### Client Library

**Use the Client Library**:

You can use the client library in your Go projects to interact with the conversion service. Here's a basic example:

```go
package main

import (
	"context"
	"fmt"
	"time"

	"github.com/ken1009us/unit-conversion-service/clientlib"
)

func main() {
    ucc, err := clientlib.NewUnitConversionClient()
    if err != nil {
        fmt.Println("Error creating unit conversion client:", err)
        return
    }
    defer ucc.Close()

    ctx, cancel := context.WithTimeout(context.Background(), time.Second)
    defer cancel()

    result, err := ucc.Convert(ctx, "meter", "kilometer", 1000)
    if err != nil {
        fmt.Println("Error during conversion:", err)
        return
    }

    fmt.Printf("Conversion result: %f kilometers\n", result)
}
```

Replace `yourusername` with your actual GitHub username.

### Result

```shell
❯ go run example/main.go
Conversion result: 1.000000 kilometers
```

## Custom Units

To add custom units, modify the conversions.json file in the config directory. The server will load these custom units on startup.

Example conversions.json:

```json
{
    "miG": "G/1024",
    "2miG": "G/512"
}
```

## Project Structure

- clientlib/clientlib.go: Contains the gRPC client library code.
- config/conversions.json: Contains the custom unit definitions.
- server/server.go: Contains the gRPC server code.
- units/units.go: Contains the unit conversion logic.
- pb/: Contains Protocol Buffer files and the generated Go code.
- go.mod & go.sum: Define the module's dependencies.
- Makefile: Contains commands to generate gRPC code and build the project.

## Testing

Run the test:

```bash
$ go test -bench=.
```

## Acknowledgements

### Project Origin

This is an open source project initiated at PhysIQ. The objective was to develop a comprehensive unit conversion service that is versatile, efficient, and user-friendly.

### Personal Contribution
I took the initiative to enhance and complete the project independently. These efforts were aimed at optimizing the service’s performance, extending its functionality, and ensuring its adaptability to diverse unit conversion needs.

### Disclaimer
This project is a new version of the original code and does not represent the entirety of the work done at PhysIQ. It is a demonstration of my individual contributions and enhancements made to improve the project's functionality and performance.

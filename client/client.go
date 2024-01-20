package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/ken1009us/unit-conversion-service/pb"

	"google.golang.org/grpc"
)

const (
	address     = "localhost:50051"
	defaultName = "world"
)

func main() {
    conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
    if err != nil {
        log.Fatalf("did not connect: %v", err)
    }
    defer conn.Close()
    client := pb.NewUnitConversionServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
    defer cancel()

    reader := bufio.NewReader(os.Stdin)
    for {
        fmt.Print("Enter conversion: [FromUnit ToUnit Value] (e.g., 'meter kilometer 1000') or 'exit': ")
        input, _ := reader.ReadString('\n')
        input = strings.TrimSpace(input)

        if input == "exit" {
            break
        }

        parts := strings.Fields(input)
        if len(parts) != 3 {
            fmt.Println("Invalid input. Please follow 'fromUnit toUnit value' format.")
            continue
        }

        fromUnit, toUnit, valueStr := parts[0], parts[1], parts[2]
        value, err := strconv.ParseFloat(valueStr, 64)
        if err != nil {
            fmt.Printf("Invalid value: %s. Please enter a number.\n", valueStr)
            continue
        }

        r, err := client.ConvertUnit(ctx, &pb.UnitConversionRequest{FromUnit: fromUnit, ToUnit: toUnit, Value: value})
        if err != nil {
            log.Fatalf("could not convert: %v", err)
        }

		if r.GetConvertedValue() == 0 {
			fmt.Printf("could not convert!\n")
		} else {
			fmt.Printf("Conversion result: %f\n", r.GetConvertedValue())
		}
    }
}
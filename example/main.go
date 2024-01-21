// example/main.go
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
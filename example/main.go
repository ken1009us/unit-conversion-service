// example/main.go
package main

import (
	"fmt"

	"github.com/ken1009us/unit-conversion-service/clientlib"
)

func main() {
	ucc := clientlib.NewUnitConversionClient()
	defer ucc.Close()  // Ensure the connection is closed when done

	result, err := ucc.Convert("meter", "kilometer", 1000)
	if err != nil {
		fmt.Println("Error during conversion:", err)
		return
	}

	fmt.Printf("Conversion result: %f kilometers\n", result)
}

// units_test.go
package units

import (
	"testing"

	"github.com/ken1009us/unit-conversion-service/units"
)

func BenchmarkSafeUnitConversion(b *testing.B) {
    sc := units.NewSafeConverter()
    fromUnit := "meter"
    toUnit := "kilometer"
    value := 1000.0

    for i := 0; i < b.N; i++ {
        _, err := sc.Convert(value, fromUnit, toUnit)
        if err != nil {
            b.Fatal(err)
        }
    }
}

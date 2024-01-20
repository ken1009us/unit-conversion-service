package units

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"path/filepath"
	"strconv"
	"strings"
)

type Units struct {
    CustomConversions map[string]string
}

func NewUnits() *Units {
    u := &Units{}
    u.loadCustomConversions()
    return u
}

func (u *Units) loadCustomConversions() {
    jsonFile, err := ioutil.ReadFile(filepath.Join("config", "conversions.json"))
    if err != nil {
        log.Fatal(err)
    }

    err = json.Unmarshal(jsonFile, &u.CustomConversions)
    if err != nil {
        log.Fatal(err)
    }
}

func (u *Units) Convert(value float64, fromUnit, toUnit string) (float64, error) {
    log.Printf("Attempting to convert: %f %s to %s", value, fromUnit, toUnit)

    conversionFormula, ok := u.CustomConversions[toUnit]
    if ok {
        log.Printf("Found custom conversion formula for '%s': %s", toUnit, conversionFormula)
        return u.applyCustomConversion(value, conversionFormula)
    }

    log.Printf("Conversion formula not found for: %s to %s", fromUnit, toUnit)
    return 0, errors.New("conversion formula not found")
}

func (u *Units) applyCustomConversion(value float64, formula string) (float64, error) {
    log.Printf("Applying conversion formula: %s to value: %f", formula, value)
    parts := strings.Fields(formula)
    if len(parts) != 3 {
        log.Printf("Invalid custom conversion formula: %s", formula)
        return 0, errors.New("invalid custom conversion formula")
    }

    operand, err := strconv.ParseFloat(parts[2], 64)
    if err != nil {
        log.Printf("Invalid number in conversion formula: %s", parts[2])
        return 0, errors.New("invalid number in conversion formula")
    }

    switch parts[1] {
    case "*":
        result := value * operand
        log.Printf("Dividing: %f / %f = %f (inverse operation)", value, operand, result)
        return result, nil
    case "/":
        result := value / operand
        log.Printf("Multiplying: %f * %f = %f", value, operand, result)
        return result, nil
    default:
        log.Printf("Unsupported operation in conversion formula: %s", parts[1])
        return 0, errors.New("unsupported operation in conversion formula")
    }
}

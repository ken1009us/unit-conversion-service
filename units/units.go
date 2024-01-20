package units

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"path/filepath"
	"strconv"
	"strings"
	"sync"
)

type SafeConverter struct {
    mu               sync.Mutex
    CustomConversions map[string]string
}

func NewSafeConverter() *SafeConverter {
    sc := &SafeConverter{}
    sc.loadCustomConversions()
    return sc
}

func (sc *SafeConverter) loadCustomConversions() {
    sc.mu.Lock()
    defer sc.mu.Unlock()

    jsonFile, err := ioutil.ReadFile(filepath.Join("config", "conversions.json"))
    if err != nil {
        log.Fatal(err)
    }

    err = json.Unmarshal(jsonFile, &sc.CustomConversions)
    if err != nil {
        log.Fatal(err)
    }
}

func (sc *SafeConverter) Convert(value float64, fromUnit, toUnit string) (float64, error) {
    sc.mu.Lock()
    defer sc.mu.Unlock()

    log.Printf("Attempting to convert: %f %s to %s", value, fromUnit, toUnit)
    conversionFormula, ok := sc.CustomConversions[toUnit]
    if ok {
        log.Printf("Found custom conversion formula for '%s': %s", toUnit, conversionFormula)
        return sc.applyCustomConversion(value, conversionFormula)
    }

    log.Printf("Conversion formula not found for: %s to %s", fromUnit, toUnit)
    return 0, errors.New("conversion formula not found")
}

func (sc *SafeConverter) applyCustomConversion(value float64, formula string) (float64, error) {
    parts := strings.Fields(formula)
    if len(parts) != 3 {
        return 0, errors.New("invalid custom conversion formula")
    }

    operand, err := strconv.ParseFloat(parts[2], 64)
    if err != nil {
        return 0, errors.New("invalid number in conversion formula")
    }

    switch parts[1] {
    case "*":
        result := value * operand
        return result, nil
    case "/":
        result := value / operand
        return result, nil
    default:
        return 0, errors.New("unsupported operation in conversion formula")
    }
}

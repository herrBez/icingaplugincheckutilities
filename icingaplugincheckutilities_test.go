package icingaplugincheckutilities

import (
	"strings"
	"testing"
)

func TestPerformanceData(t *testing.T) {
	p := map[string]performanceData{
		"foo": {
			value: 5.123,
			uom:   "s",
		},
		"ba'r": {
			value:    60.5,
			uom:      "b",
			warning:  CreateFloat(50.0),
			critical: CreateFloat(75.0),
			min:      CreateFloat(0.0),
			max:      CreateFloat(100.0),
		},
	}
	s := RenderPerformanceData(p)

	wanted := []string{
		"'foo'=5.12300s;;;;",
		"'ba\"r'=60.50000b;50.00000;75.00000;0.00000;100.00000",
	}

	for _, w := range wanted {
		if !strings.Contains(s, w) {
			t.Fatalf("Given \"%s\" expect %s is contained", s, w)
		}
	}

}

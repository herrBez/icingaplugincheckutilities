package icingaplugincheckutilities

import (
	"fmt"
	"strings"
)

type PerformanceData struct {
	value    float64
	uom      string
	min      *float64
	max      *float64
	warning  *float64
	critical *float64
}

// Help Function that from a literal returns a pointer containing this value
func CreateFloat(value float64) *float64 {
	return &value
}

func renderPerformancePointer(res *float64) string {
	if res == nil {
		return ""
	} else {
		return fmt.Sprintf("%0.5f", *res)
	}
}

func normalizePerformanceDataKey(key string) string {
	tmp := strings.Replace(key, "'", "\"", -1)
	tmp = strings.Replace(tmp, "=", "-equals-", -1)
	return tmp
}

func RenderPerformanceData(performance_data map[string]PerformanceData) string {
	if len(performance_data) == 0 {
		return ""
	}
	output := " | "
	for k, v := range performance_data {
		output += fmt.Sprintf("'%s'=%0.5f%s;%s;%s;%s;%s ",
			normalizePerformanceDataKey(k),
			v.value,
			v.uom,
			renderPerformancePointer(v.warning),
			renderPerformancePointer(v.critical),
			renderPerformancePointer(v.min),
			renderPerformancePointer(v.max),
		)
	}
	return output
}

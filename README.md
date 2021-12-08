# Icinga Plugin Check Utilities

This go module contains set of utilities to simplify the development of Plugin Check Commands.

## Sample Usage

```go
package main

import (
	ipcu "github.com/herrBez/icingaplugincheckutilities"
)

func main() {
	perf := map[string]ipcu.PerformanceData{
		"f'oo": {
			Value:    8.0,
			Uom:      "%",
			Warning:  ipcu.CreateFloat(7.0),
			Critical: ipcu.CreateFloat(9.0),
		},
		"bar": {
			Value: 12.0,
			Uom:   "",
		},
	}
	ipcu.PrintAndExit(ipcu.WARNING, "This is a warning message", perf)
}
```

The result is:

```
This is a warning message | 'bar'=12.00000;;;; 'f"oo'=8.00000%;7.00000;9.00000;;
```
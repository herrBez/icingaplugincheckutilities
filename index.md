## Introduction

A golang module to simplify the development of golang plugin check commands

## Quick start

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
	status, err = icpu.ComputeExitStatus(perf, "f'oo", icpu.StatusFuncGte)
	if err != nil {
		return icpu.PrintAndExit(icpu.UNKNOWN, err, perf)
	}
	ipcu.PrintAndExit(status, "This is a warning message", perf)
}
```







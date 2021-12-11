<!--
SPDX-FileCopyrightText: 2021 Mirko Bez <bez.mirko@gmail.com>

SPDX-License-Identifier: MIT
-->

# Icinga Plugin Check Utilities
[![REUSE status](https://api.reuse.software/badge/github.com/herrBez/icingaplugincheckutilities)](https://api.reuse.software/info/github.com/herrBez/icingaplugincheckutilities)

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
	status, err = icpu.ComputeExitStatus(perf, "f'oo", icpu.StatusFuncGte)
	if err != nil {
		return icpu.PrintAndExit(icpu.UNKNOWN, err, perf)
	}
	ipcu.PrintAndExit(status, "This is a warning message", perf)
}
```

The result is:

```
This is a warning message | 'bar'=12.00000;;;; 'f"oo'=8.00000%;7.00000;9.00000;;
```

And the exit status will be `1`.


## Run the test

```sh
go test
```
/*
 * SPDX-FileCopyrightText: 2021 Mirko Bez <bez.mirko@gmail.com>
 *
 * SPDX-License-Identifier: MIT
 */
package icingaplugincheckutilities

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

type PerformanceData struct {
	Value    float64
	Uom      string
	Min      *float64
	Max      *float64
	Warning  *float64
	Critical *float64
}

const (
	OK       = 0
	WARNING  = 1
	CRITICAL = 2
	UNKNOWN  = 3
)

const (
	UP   = 0
	DOWN = 2
)

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
			v.Value,
			v.Uom,
			renderPerformancePointer(v.Warning),
			renderPerformancePointer(v.Critical),
			renderPerformancePointer(v.Min),
			renderPerformancePointer(v.Max),
		)
	}
	return output
}

type StatusFunc func(PerformanceData) (int, error)

func StatusFuncGte(p PerformanceData) (int, error) {
	if p.Critical == nil || p.Warning == nil {
		return UNKNOWN, errors.New("Critical or Warning Threshold not defined")
	}
	if *p.Critical <= *p.Warning {
		return UNKNOWN, errors.New("Critical Threshold should be greather than Warning Threshold")
	}
	if p.Value >= *p.Critical {
		return CRITICAL, nil
	} else if p.Value >= *p.Warning {
		return WARNING, nil
	} else {
		return OK, nil
	}
}
func StatusFuncLte(p PerformanceData) (int, error) {
	if p.Critical == nil || p.Warning == nil {
		return UNKNOWN, errors.New("Critical or Warning Threshold not defined")
	}
	if *p.Critical >= *p.Warning {
		return UNKNOWN, errors.New("Critical Threshold should be smaller than Warning Threshold")
	}
	if p.Value <= *p.Critical {
		return CRITICAL, nil
	} else if p.Value <= *p.Warning {
		return WARNING, nil
	} else {
		return OK, nil
	}
}

func ComputeExitStatus(perf map[string]PerformanceData, key string, status_func StatusFunc) (int, error) {
	p := perf[key]
	return status_func(p)
}

func PrintAndExit(exit_status int, message string, perf map[string]PerformanceData) {
	fmt.Printf("%s%s", message, RenderPerformanceData(perf))
	os.Exit(exit_status)
}

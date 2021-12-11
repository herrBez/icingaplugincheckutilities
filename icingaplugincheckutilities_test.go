/*
 * SPDX-FileCopyrightText: 2021 Mirko Bez <bez.mirko@gmail.com>
 *
 * SPDX-License-Identifier: MIT
 */
package icingaplugincheckutilities

import (
	"strings"
	"testing"
)

func TestDefaultValue(t *testing.T) {
	p := map[string]PerformanceData{
		"test": {
			Value:    5,
			Uom:      "",
			Warning:  CreateFloat(2.0),
			Critical: CreateFloat(3.0),
		},
	}
	exit_status, err := ComputeExitStatus(p, "test", StatusFuncGte)
	if exit_status != CRITICAL || err != nil {
		t.Fatalf("The test should return CRITICAL")
	}
	exit_status, err = ComputeExitStatus(p, "test", StatusFuncLte)
	if exit_status != UNKNOWN || err == nil {
		t.Fatalf("The test should return Unknown and an error")
	}
}

func TestPerformanceData(t *testing.T) {
	p := map[string]PerformanceData{
		"foo": {
			Value: 5.123,
			Uom:   "s",
		},
		"ba'r": {
			Value:    60.5,
			Uom:      "b",
			Warning:  CreateFloat(50.0),
			Critical: CreateFloat(75.0),
			Min:      CreateFloat(0.0),
			Max:      CreateFloat(100.0),
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

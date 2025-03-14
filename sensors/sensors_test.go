// SPDX-License-Identifier: BSD-3-Clause

package sensors

import (
	"errors"
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/shirou/gopsutil/v4/internal/common"
)

func TestTemperatureStat_String(t *testing.T) {
	v := TemperatureStat{
		SensorKey:   "CPU",
		Temperature: 1.1,
		High:        30.1,
		Critical:    0.1,
	}
	s := `{"sensorKey":"CPU","temperature":1.1,"sensorHigh":30.1,"sensorCritical":0.1}`
	assert.Equalf(t, s, fmt.Sprintf("%v", v), "TemperatureStat string is invalid, %v", fmt.Sprintf("%v", v))
}

func skipIfNotImplementedErr(t *testing.T, err error) {
	if errors.Is(err, common.ErrNotImplementedError) {
		t.Skip("not implemented")
	}
}

func TestTemperatures(t *testing.T) {
	if os.Getenv("CI") != "" {
		t.Skip("Skip CI")
	}
	v, err := SensorsTemperatures()
	skipIfNotImplementedErr(t, err)
	require.NoError(t, err)
	assert.NotEmptyf(t, v, "Could not get temperature %v", v)
	t.Log(v)
}

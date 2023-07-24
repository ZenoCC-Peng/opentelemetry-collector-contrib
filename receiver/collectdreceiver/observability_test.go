// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package collectdreceiver

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.opencensus.io/stats"
	"go.opencensus.io/stats/view"
	"reflect"
	"testing"
)

var testName = "collectdreceiver"

type metricsTestCase struct {
	fn       func()        // function to test updating metrics
	v        *view.View    // view to reference
	m        stats.Measure // expected measure of the view
	calls    int           // number of times to call fn
	expected int           // expected value of reported metric at end of calls
}

func TestRecordMetrics(t *testing.T) {

	metrics := newTestMetrics(t)
	testCases := []metricsTestCase{
		{metrics.recordRequestErrors, metrics.views.viewInvalidRequests, metrics.stats.mErrors, 3, 3},
		{metrics.recordRequestReceived, metrics.views.viewRequestsReceived, metrics.stats.mRequestsReceived, 3, 3},
		{metrics.recordMetricsReceived, metrics.views.viewMetricsReceived, metrics.stats.mMetricsReceived, 3, 3},
		{metrics.recordEventsReceived, metrics.views.viewEventsReceived, metrics.stats.mEventsReceived, 3, 3},
		{metrics.recordDefaultBlankAttrs, metrics.views.viewBlankDefaultAttrs, metrics.stats.mBlankDefaultAttrs, 3, 3},
	}
	for _, tc := range testCases {
		t.Run(tc.m.Name(), func(t *testing.T) {
			for i := 0; i < tc.calls; i++ {
				tc.fn()
			}
			validateMetric(t, tc.v, tc.expected)
		})
	}
}

func validateMetric(t *testing.T, v *view.View, expected interface{}) {
	// hack to reset stats to 0
	defer func() {
		view.Unregister(v)
		err := view.Register(v)
		assert.NoError(t, err)
	}()
	rows, err := view.RetrieveData(v.Name)
	assert.NoError(t, err)
	if expected != nil {
		require.Len(t, rows, 1)
		value := reflect.Indirect(reflect.ValueOf(rows[0].Data)).FieldByName("Value").Interface()
		assert.EqualValues(t, expected, value)
	} else {
		assert.Len(t, rows, 0)
	}
}

// TestRegisterViewsExpectingFailure validates that if an error is returned from view.Register,
// we panic and don't continue with initialization
func TestRegisterViewsExpectingFailure(t *testing.T) {
	statName := "otelcol/collectdreceiver/errors"
	stat := stats.Int64(statName, "", stats.UnitDimensionless)
	err := view.Register(&view.View{
		Name:        buildReceiverCustomMetricName(statName),
		Description: "some description",
		Measure:     stat,
		Aggregation: view.Count(),
	})
	require.NoError(t, err)
	metrics, err := newOpenCensusMetrics(testName)
	assert.Error(t, err)
	assert.Nil(t, metrics)
}

// newTestMetrics builds a new metrics that will cleanup when testing.T completes
func newTestMetrics(t *testing.T) *opencensusMetrics {
	m, err := newOpenCensusMetrics(testName)
	require.NoError(t, err)
	t.Cleanup(func() {
		unregisterMetrics(m)
	})
	return m
}

// unregisterMetrics is used to unregister the metrics for testing purposes
func unregisterMetrics(metrics *opencensusMetrics) {
	view.Unregister(
		metrics.views.viewInvalidRequests,
		metrics.views.viewRequestsReceived,
		metrics.views.viewMetricsReceived,
		metrics.views.viewEventsReceived,
		metrics.views.viewBlankDefaultAttrs,
	)
}

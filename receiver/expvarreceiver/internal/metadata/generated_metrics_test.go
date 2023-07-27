// Code generated by mdatagen. DO NOT EDIT.

package metadata

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/pmetric"
	"go.opentelemetry.io/collector/receiver/receivertest"
	"go.uber.org/zap"
	"go.uber.org/zap/zaptest/observer"
)

type testConfigCollection int

const (
	testSetDefault testConfigCollection = iota
	testSetAll
	testSetNone
)

func TestMetricsBuilder(t *testing.T) {
	tests := []struct {
		name      string
		configSet testConfigCollection
	}{
		{
			name:      "default",
			configSet: testSetDefault,
		},
		{
			name:      "all_set",
			configSet: testSetAll,
		},
		{
			name:      "none_set",
			configSet: testSetNone,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			start := pcommon.Timestamp(1_000_000_000)
			ts := pcommon.Timestamp(1_000_001_000)
			observedZapCore, observedLogs := observer.New(zap.WarnLevel)
			settings := receivertest.NewNopCreateSettings()
			settings.Logger = zap.New(observedZapCore)
			mb := NewMetricsBuilder(loadMetricsBuilderConfig(t, test.name), settings, WithStartTime(start))

			expectedWarnings := 0
			assert.Equal(t, expectedWarnings, observedLogs.Len())

			defaultMetricsCount := 0
			allMetricsCount := 0

			defaultMetricsCount++
			allMetricsCount++
			mb.RecordProcessRuntimeMemstatsBuckHashSysDataPoint(ts, 1)

			defaultMetricsCount++
			allMetricsCount++
			mb.RecordProcessRuntimeMemstatsFreesDataPoint(ts, 1)

			defaultMetricsCount++
			allMetricsCount++
			mb.RecordProcessRuntimeMemstatsGcCPUFractionDataPoint(ts, 1)

			defaultMetricsCount++
			allMetricsCount++
			mb.RecordProcessRuntimeMemstatsGcSysDataPoint(ts, 1)

			defaultMetricsCount++
			allMetricsCount++
			mb.RecordProcessRuntimeMemstatsHeapAllocDataPoint(ts, 1)

			defaultMetricsCount++
			allMetricsCount++
			mb.RecordProcessRuntimeMemstatsHeapIdleDataPoint(ts, 1)

			defaultMetricsCount++
			allMetricsCount++
			mb.RecordProcessRuntimeMemstatsHeapInuseDataPoint(ts, 1)

			defaultMetricsCount++
			allMetricsCount++
			mb.RecordProcessRuntimeMemstatsHeapObjectsDataPoint(ts, 1)

			defaultMetricsCount++
			allMetricsCount++
			mb.RecordProcessRuntimeMemstatsHeapReleasedDataPoint(ts, 1)

			defaultMetricsCount++
			allMetricsCount++
			mb.RecordProcessRuntimeMemstatsHeapSysDataPoint(ts, 1)

			defaultMetricsCount++
			allMetricsCount++
			mb.RecordProcessRuntimeMemstatsLastPauseDataPoint(ts, 1)

			allMetricsCount++
			mb.RecordProcessRuntimeMemstatsLookupsDataPoint(ts, 1)

			defaultMetricsCount++
			allMetricsCount++
			mb.RecordProcessRuntimeMemstatsMallocsDataPoint(ts, 1)

			defaultMetricsCount++
			allMetricsCount++
			mb.RecordProcessRuntimeMemstatsMcacheInuseDataPoint(ts, 1)

			defaultMetricsCount++
			allMetricsCount++
			mb.RecordProcessRuntimeMemstatsMcacheSysDataPoint(ts, 1)

			defaultMetricsCount++
			allMetricsCount++
			mb.RecordProcessRuntimeMemstatsMspanInuseDataPoint(ts, 1)

			defaultMetricsCount++
			allMetricsCount++
			mb.RecordProcessRuntimeMemstatsMspanSysDataPoint(ts, 1)

			defaultMetricsCount++
			allMetricsCount++
			mb.RecordProcessRuntimeMemstatsNextGcDataPoint(ts, 1)

			defaultMetricsCount++
			allMetricsCount++
			mb.RecordProcessRuntimeMemstatsNumForcedGcDataPoint(ts, 1)

			defaultMetricsCount++
			allMetricsCount++
			mb.RecordProcessRuntimeMemstatsNumGcDataPoint(ts, 1)

			defaultMetricsCount++
			allMetricsCount++
			mb.RecordProcessRuntimeMemstatsOtherSysDataPoint(ts, 1)

			defaultMetricsCount++
			allMetricsCount++
			mb.RecordProcessRuntimeMemstatsPauseTotalDataPoint(ts, 1)

			defaultMetricsCount++
			allMetricsCount++
			mb.RecordProcessRuntimeMemstatsStackInuseDataPoint(ts, 1)

			defaultMetricsCount++
			allMetricsCount++
			mb.RecordProcessRuntimeMemstatsStackSysDataPoint(ts, 1)

			defaultMetricsCount++
			allMetricsCount++
			mb.RecordProcessRuntimeMemstatsSysDataPoint(ts, 1)

			allMetricsCount++
			mb.RecordProcessRuntimeMemstatsTotalAllocDataPoint(ts, 1)

			res := pcommon.NewResource()
			res.Attributes().PutStr("k1", "v1")
			metrics := mb.Emit(WithResource(res))

			if test.configSet == testSetNone {
				assert.Equal(t, 0, metrics.ResourceMetrics().Len())
				return
			}

			assert.Equal(t, 1, metrics.ResourceMetrics().Len())
			rm := metrics.ResourceMetrics().At(0)
			assert.Equal(t, res, rm.Resource())
			assert.Equal(t, 1, rm.ScopeMetrics().Len())
			ms := rm.ScopeMetrics().At(0).Metrics()
			if test.configSet == testSetDefault {
				assert.Equal(t, defaultMetricsCount, ms.Len())
			}
			if test.configSet == testSetAll {
				assert.Equal(t, allMetricsCount, ms.Len())
			}
			validatedMetrics := make(map[string]bool)
			for i := 0; i < ms.Len(); i++ {
				switch ms.At(i).Name() {
				case "process.runtime.memstats.buck_hash_sys":
					assert.False(t, validatedMetrics["process.runtime.memstats.buck_hash_sys"], "Found a duplicate in the metrics slice: process.runtime.memstats.buck_hash_sys")
					validatedMetrics["process.runtime.memstats.buck_hash_sys"] = true
					assert.Equal(t, pmetric.MetricTypeSum, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Sum().DataPoints().Len())
					assert.Equal(t, "Bytes of memory in profiling bucket hash tables.", ms.At(i).Description())
					assert.Equal(t, "By", ms.At(i).Unit())
					assert.Equal(t, false, ms.At(i).Sum().IsMonotonic())
					assert.Equal(t, pmetric.AggregationTemporalityCumulative, ms.At(i).Sum().AggregationTemporality())
					dp := ms.At(i).Sum().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
					assert.Equal(t, int64(1), dp.IntValue())
				case "process.runtime.memstats.frees":
					assert.False(t, validatedMetrics["process.runtime.memstats.frees"], "Found a duplicate in the metrics slice: process.runtime.memstats.frees")
					validatedMetrics["process.runtime.memstats.frees"] = true
					assert.Equal(t, pmetric.MetricTypeSum, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Sum().DataPoints().Len())
					assert.Equal(t, "Cumulative count of heap objects freed.", ms.At(i).Description())
					assert.Equal(t, "{objects}", ms.At(i).Unit())
					assert.Equal(t, true, ms.At(i).Sum().IsMonotonic())
					assert.Equal(t, pmetric.AggregationTemporalityCumulative, ms.At(i).Sum().AggregationTemporality())
					dp := ms.At(i).Sum().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
					assert.Equal(t, int64(1), dp.IntValue())
				case "process.runtime.memstats.gc_cpu_fraction":
					assert.False(t, validatedMetrics["process.runtime.memstats.gc_cpu_fraction"], "Found a duplicate in the metrics slice: process.runtime.memstats.gc_cpu_fraction")
					validatedMetrics["process.runtime.memstats.gc_cpu_fraction"] = true
					assert.Equal(t, pmetric.MetricTypeGauge, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Gauge().DataPoints().Len())
					assert.Equal(t, "The fraction of this program's available CPU time used by the GC since the program started.", ms.At(i).Description())
					assert.Equal(t, "1", ms.At(i).Unit())
					dp := ms.At(i).Gauge().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeDouble, dp.ValueType())
					assert.Equal(t, float64(1), dp.DoubleValue())
				case "process.runtime.memstats.gc_sys":
					assert.False(t, validatedMetrics["process.runtime.memstats.gc_sys"], "Found a duplicate in the metrics slice: process.runtime.memstats.gc_sys")
					validatedMetrics["process.runtime.memstats.gc_sys"] = true
					assert.Equal(t, pmetric.MetricTypeSum, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Sum().DataPoints().Len())
					assert.Equal(t, "Bytes of memory in garbage collection metadata.", ms.At(i).Description())
					assert.Equal(t, "By", ms.At(i).Unit())
					assert.Equal(t, false, ms.At(i).Sum().IsMonotonic())
					assert.Equal(t, pmetric.AggregationTemporalityCumulative, ms.At(i).Sum().AggregationTemporality())
					dp := ms.At(i).Sum().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
					assert.Equal(t, int64(1), dp.IntValue())
				case "process.runtime.memstats.heap_alloc":
					assert.False(t, validatedMetrics["process.runtime.memstats.heap_alloc"], "Found a duplicate in the metrics slice: process.runtime.memstats.heap_alloc")
					validatedMetrics["process.runtime.memstats.heap_alloc"] = true
					assert.Equal(t, pmetric.MetricTypeSum, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Sum().DataPoints().Len())
					assert.Equal(t, "Bytes of allocated heap objects.", ms.At(i).Description())
					assert.Equal(t, "By", ms.At(i).Unit())
					assert.Equal(t, false, ms.At(i).Sum().IsMonotonic())
					assert.Equal(t, pmetric.AggregationTemporalityCumulative, ms.At(i).Sum().AggregationTemporality())
					dp := ms.At(i).Sum().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
					assert.Equal(t, int64(1), dp.IntValue())
				case "process.runtime.memstats.heap_idle":
					assert.False(t, validatedMetrics["process.runtime.memstats.heap_idle"], "Found a duplicate in the metrics slice: process.runtime.memstats.heap_idle")
					validatedMetrics["process.runtime.memstats.heap_idle"] = true
					assert.Equal(t, pmetric.MetricTypeSum, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Sum().DataPoints().Len())
					assert.Equal(t, "Bytes in idle (unused) spans.", ms.At(i).Description())
					assert.Equal(t, "By", ms.At(i).Unit())
					assert.Equal(t, false, ms.At(i).Sum().IsMonotonic())
					assert.Equal(t, pmetric.AggregationTemporalityCumulative, ms.At(i).Sum().AggregationTemporality())
					dp := ms.At(i).Sum().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
					assert.Equal(t, int64(1), dp.IntValue())
				case "process.runtime.memstats.heap_inuse":
					assert.False(t, validatedMetrics["process.runtime.memstats.heap_inuse"], "Found a duplicate in the metrics slice: process.runtime.memstats.heap_inuse")
					validatedMetrics["process.runtime.memstats.heap_inuse"] = true
					assert.Equal(t, pmetric.MetricTypeSum, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Sum().DataPoints().Len())
					assert.Equal(t, "Bytes in in-use spans.", ms.At(i).Description())
					assert.Equal(t, "By", ms.At(i).Unit())
					assert.Equal(t, false, ms.At(i).Sum().IsMonotonic())
					assert.Equal(t, pmetric.AggregationTemporalityCumulative, ms.At(i).Sum().AggregationTemporality())
					dp := ms.At(i).Sum().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
					assert.Equal(t, int64(1), dp.IntValue())
				case "process.runtime.memstats.heap_objects":
					assert.False(t, validatedMetrics["process.runtime.memstats.heap_objects"], "Found a duplicate in the metrics slice: process.runtime.memstats.heap_objects")
					validatedMetrics["process.runtime.memstats.heap_objects"] = true
					assert.Equal(t, pmetric.MetricTypeSum, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Sum().DataPoints().Len())
					assert.Equal(t, "Number of allocated heap objects.", ms.At(i).Description())
					assert.Equal(t, "{objects}", ms.At(i).Unit())
					assert.Equal(t, false, ms.At(i).Sum().IsMonotonic())
					assert.Equal(t, pmetric.AggregationTemporalityCumulative, ms.At(i).Sum().AggregationTemporality())
					dp := ms.At(i).Sum().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
					assert.Equal(t, int64(1), dp.IntValue())
				case "process.runtime.memstats.heap_released":
					assert.False(t, validatedMetrics["process.runtime.memstats.heap_released"], "Found a duplicate in the metrics slice: process.runtime.memstats.heap_released")
					validatedMetrics["process.runtime.memstats.heap_released"] = true
					assert.Equal(t, pmetric.MetricTypeSum, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Sum().DataPoints().Len())
					assert.Equal(t, "Bytes of physical memory returned to the OS.", ms.At(i).Description())
					assert.Equal(t, "By", ms.At(i).Unit())
					assert.Equal(t, false, ms.At(i).Sum().IsMonotonic())
					assert.Equal(t, pmetric.AggregationTemporalityCumulative, ms.At(i).Sum().AggregationTemporality())
					dp := ms.At(i).Sum().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
					assert.Equal(t, int64(1), dp.IntValue())
				case "process.runtime.memstats.heap_sys":
					assert.False(t, validatedMetrics["process.runtime.memstats.heap_sys"], "Found a duplicate in the metrics slice: process.runtime.memstats.heap_sys")
					validatedMetrics["process.runtime.memstats.heap_sys"] = true
					assert.Equal(t, pmetric.MetricTypeSum, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Sum().DataPoints().Len())
					assert.Equal(t, "Bytes of heap memory obtained by the OS.", ms.At(i).Description())
					assert.Equal(t, "By", ms.At(i).Unit())
					assert.Equal(t, false, ms.At(i).Sum().IsMonotonic())
					assert.Equal(t, pmetric.AggregationTemporalityCumulative, ms.At(i).Sum().AggregationTemporality())
					dp := ms.At(i).Sum().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
					assert.Equal(t, int64(1), dp.IntValue())
				case "process.runtime.memstats.last_pause":
					assert.False(t, validatedMetrics["process.runtime.memstats.last_pause"], "Found a duplicate in the metrics slice: process.runtime.memstats.last_pause")
					validatedMetrics["process.runtime.memstats.last_pause"] = true
					assert.Equal(t, pmetric.MetricTypeGauge, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Gauge().DataPoints().Len())
					assert.Equal(t, "The most recent stop-the-world pause time.", ms.At(i).Description())
					assert.Equal(t, "ns", ms.At(i).Unit())
					dp := ms.At(i).Gauge().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
					assert.Equal(t, int64(1), dp.IntValue())
				case "process.runtime.memstats.lookups":
					assert.False(t, validatedMetrics["process.runtime.memstats.lookups"], "Found a duplicate in the metrics slice: process.runtime.memstats.lookups")
					validatedMetrics["process.runtime.memstats.lookups"] = true
					assert.Equal(t, pmetric.MetricTypeSum, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Sum().DataPoints().Len())
					assert.Equal(t, "Number of pointer lookups performed by the runtime.", ms.At(i).Description())
					assert.Equal(t, "{lookups}", ms.At(i).Unit())
					assert.Equal(t, false, ms.At(i).Sum().IsMonotonic())
					assert.Equal(t, pmetric.AggregationTemporalityCumulative, ms.At(i).Sum().AggregationTemporality())
					dp := ms.At(i).Sum().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
					assert.Equal(t, int64(1), dp.IntValue())
				case "process.runtime.memstats.mallocs":
					assert.False(t, validatedMetrics["process.runtime.memstats.mallocs"], "Found a duplicate in the metrics slice: process.runtime.memstats.mallocs")
					validatedMetrics["process.runtime.memstats.mallocs"] = true
					assert.Equal(t, pmetric.MetricTypeSum, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Sum().DataPoints().Len())
					assert.Equal(t, "Cumulative count of heap objects allocated.", ms.At(i).Description())
					assert.Equal(t, "{objects}", ms.At(i).Unit())
					assert.Equal(t, true, ms.At(i).Sum().IsMonotonic())
					assert.Equal(t, pmetric.AggregationTemporalityCumulative, ms.At(i).Sum().AggregationTemporality())
					dp := ms.At(i).Sum().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
					assert.Equal(t, int64(1), dp.IntValue())
				case "process.runtime.memstats.mcache_inuse":
					assert.False(t, validatedMetrics["process.runtime.memstats.mcache_inuse"], "Found a duplicate in the metrics slice: process.runtime.memstats.mcache_inuse")
					validatedMetrics["process.runtime.memstats.mcache_inuse"] = true
					assert.Equal(t, pmetric.MetricTypeSum, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Sum().DataPoints().Len())
					assert.Equal(t, "Bytes of allocated mcache structures.", ms.At(i).Description())
					assert.Equal(t, "By", ms.At(i).Unit())
					assert.Equal(t, false, ms.At(i).Sum().IsMonotonic())
					assert.Equal(t, pmetric.AggregationTemporalityCumulative, ms.At(i).Sum().AggregationTemporality())
					dp := ms.At(i).Sum().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
					assert.Equal(t, int64(1), dp.IntValue())
				case "process.runtime.memstats.mcache_sys":
					assert.False(t, validatedMetrics["process.runtime.memstats.mcache_sys"], "Found a duplicate in the metrics slice: process.runtime.memstats.mcache_sys")
					validatedMetrics["process.runtime.memstats.mcache_sys"] = true
					assert.Equal(t, pmetric.MetricTypeSum, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Sum().DataPoints().Len())
					assert.Equal(t, "Bytes of memory obtained from the OS for mcache structures.", ms.At(i).Description())
					assert.Equal(t, "By", ms.At(i).Unit())
					assert.Equal(t, false, ms.At(i).Sum().IsMonotonic())
					assert.Equal(t, pmetric.AggregationTemporalityCumulative, ms.At(i).Sum().AggregationTemporality())
					dp := ms.At(i).Sum().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
					assert.Equal(t, int64(1), dp.IntValue())
				case "process.runtime.memstats.mspan_inuse":
					assert.False(t, validatedMetrics["process.runtime.memstats.mspan_inuse"], "Found a duplicate in the metrics slice: process.runtime.memstats.mspan_inuse")
					validatedMetrics["process.runtime.memstats.mspan_inuse"] = true
					assert.Equal(t, pmetric.MetricTypeSum, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Sum().DataPoints().Len())
					assert.Equal(t, "Bytes of allocated mspan structures.", ms.At(i).Description())
					assert.Equal(t, "By", ms.At(i).Unit())
					assert.Equal(t, false, ms.At(i).Sum().IsMonotonic())
					assert.Equal(t, pmetric.AggregationTemporalityCumulative, ms.At(i).Sum().AggregationTemporality())
					dp := ms.At(i).Sum().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
					assert.Equal(t, int64(1), dp.IntValue())
				case "process.runtime.memstats.mspan_sys":
					assert.False(t, validatedMetrics["process.runtime.memstats.mspan_sys"], "Found a duplicate in the metrics slice: process.runtime.memstats.mspan_sys")
					validatedMetrics["process.runtime.memstats.mspan_sys"] = true
					assert.Equal(t, pmetric.MetricTypeSum, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Sum().DataPoints().Len())
					assert.Equal(t, "Bytes of memory obtained from the OS for mspan structures.", ms.At(i).Description())
					assert.Equal(t, "By", ms.At(i).Unit())
					assert.Equal(t, false, ms.At(i).Sum().IsMonotonic())
					assert.Equal(t, pmetric.AggregationTemporalityCumulative, ms.At(i).Sum().AggregationTemporality())
					dp := ms.At(i).Sum().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
					assert.Equal(t, int64(1), dp.IntValue())
				case "process.runtime.memstats.next_gc":
					assert.False(t, validatedMetrics["process.runtime.memstats.next_gc"], "Found a duplicate in the metrics slice: process.runtime.memstats.next_gc")
					validatedMetrics["process.runtime.memstats.next_gc"] = true
					assert.Equal(t, pmetric.MetricTypeSum, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Sum().DataPoints().Len())
					assert.Equal(t, "The target heap size of the next GC cycle.", ms.At(i).Description())
					assert.Equal(t, "By", ms.At(i).Unit())
					assert.Equal(t, false, ms.At(i).Sum().IsMonotonic())
					assert.Equal(t, pmetric.AggregationTemporalityCumulative, ms.At(i).Sum().AggregationTemporality())
					dp := ms.At(i).Sum().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
					assert.Equal(t, int64(1), dp.IntValue())
				case "process.runtime.memstats.num_forced_gc":
					assert.False(t, validatedMetrics["process.runtime.memstats.num_forced_gc"], "Found a duplicate in the metrics slice: process.runtime.memstats.num_forced_gc")
					validatedMetrics["process.runtime.memstats.num_forced_gc"] = true
					assert.Equal(t, pmetric.MetricTypeSum, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Sum().DataPoints().Len())
					assert.Equal(t, "Number of GC cycles that were forced by the application calling the GC function.", ms.At(i).Description())
					assert.Equal(t, "By", ms.At(i).Unit())
					assert.Equal(t, true, ms.At(i).Sum().IsMonotonic())
					assert.Equal(t, pmetric.AggregationTemporalityCumulative, ms.At(i).Sum().AggregationTemporality())
					dp := ms.At(i).Sum().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
					assert.Equal(t, int64(1), dp.IntValue())
				case "process.runtime.memstats.num_gc":
					assert.False(t, validatedMetrics["process.runtime.memstats.num_gc"], "Found a duplicate in the metrics slice: process.runtime.memstats.num_gc")
					validatedMetrics["process.runtime.memstats.num_gc"] = true
					assert.Equal(t, pmetric.MetricTypeSum, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Sum().DataPoints().Len())
					assert.Equal(t, "Number of completed GC cycles.", ms.At(i).Description())
					assert.Equal(t, "By", ms.At(i).Unit())
					assert.Equal(t, true, ms.At(i).Sum().IsMonotonic())
					assert.Equal(t, pmetric.AggregationTemporalityCumulative, ms.At(i).Sum().AggregationTemporality())
					dp := ms.At(i).Sum().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
					assert.Equal(t, int64(1), dp.IntValue())
				case "process.runtime.memstats.other_sys":
					assert.False(t, validatedMetrics["process.runtime.memstats.other_sys"], "Found a duplicate in the metrics slice: process.runtime.memstats.other_sys")
					validatedMetrics["process.runtime.memstats.other_sys"] = true
					assert.Equal(t, pmetric.MetricTypeSum, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Sum().DataPoints().Len())
					assert.Equal(t, "Bytes of memory in miscellaneous off-heap runtime allocations.", ms.At(i).Description())
					assert.Equal(t, "By", ms.At(i).Unit())
					assert.Equal(t, false, ms.At(i).Sum().IsMonotonic())
					assert.Equal(t, pmetric.AggregationTemporalityCumulative, ms.At(i).Sum().AggregationTemporality())
					dp := ms.At(i).Sum().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
					assert.Equal(t, int64(1), dp.IntValue())
				case "process.runtime.memstats.pause_total":
					assert.False(t, validatedMetrics["process.runtime.memstats.pause_total"], "Found a duplicate in the metrics slice: process.runtime.memstats.pause_total")
					validatedMetrics["process.runtime.memstats.pause_total"] = true
					assert.Equal(t, pmetric.MetricTypeSum, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Sum().DataPoints().Len())
					assert.Equal(t, "The cumulative nanoseconds in GC stop-the-world pauses since the program started.", ms.At(i).Description())
					assert.Equal(t, "By", ms.At(i).Unit())
					assert.Equal(t, true, ms.At(i).Sum().IsMonotonic())
					assert.Equal(t, pmetric.AggregationTemporalityCumulative, ms.At(i).Sum().AggregationTemporality())
					dp := ms.At(i).Sum().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
					assert.Equal(t, int64(1), dp.IntValue())
				case "process.runtime.memstats.stack_inuse":
					assert.False(t, validatedMetrics["process.runtime.memstats.stack_inuse"], "Found a duplicate in the metrics slice: process.runtime.memstats.stack_inuse")
					validatedMetrics["process.runtime.memstats.stack_inuse"] = true
					assert.Equal(t, pmetric.MetricTypeSum, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Sum().DataPoints().Len())
					assert.Equal(t, "Bytes in stack spans.", ms.At(i).Description())
					assert.Equal(t, "By", ms.At(i).Unit())
					assert.Equal(t, false, ms.At(i).Sum().IsMonotonic())
					assert.Equal(t, pmetric.AggregationTemporalityCumulative, ms.At(i).Sum().AggregationTemporality())
					dp := ms.At(i).Sum().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
					assert.Equal(t, int64(1), dp.IntValue())
				case "process.runtime.memstats.stack_sys":
					assert.False(t, validatedMetrics["process.runtime.memstats.stack_sys"], "Found a duplicate in the metrics slice: process.runtime.memstats.stack_sys")
					validatedMetrics["process.runtime.memstats.stack_sys"] = true
					assert.Equal(t, pmetric.MetricTypeSum, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Sum().DataPoints().Len())
					assert.Equal(t, "Bytes of stack memory obtained from the OS.", ms.At(i).Description())
					assert.Equal(t, "By", ms.At(i).Unit())
					assert.Equal(t, false, ms.At(i).Sum().IsMonotonic())
					assert.Equal(t, pmetric.AggregationTemporalityCumulative, ms.At(i).Sum().AggregationTemporality())
					dp := ms.At(i).Sum().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
					assert.Equal(t, int64(1), dp.IntValue())
				case "process.runtime.memstats.sys":
					assert.False(t, validatedMetrics["process.runtime.memstats.sys"], "Found a duplicate in the metrics slice: process.runtime.memstats.sys")
					validatedMetrics["process.runtime.memstats.sys"] = true
					assert.Equal(t, pmetric.MetricTypeSum, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Sum().DataPoints().Len())
					assert.Equal(t, "Total bytes of memory obtained from the OS.", ms.At(i).Description())
					assert.Equal(t, "By", ms.At(i).Unit())
					assert.Equal(t, false, ms.At(i).Sum().IsMonotonic())
					assert.Equal(t, pmetric.AggregationTemporalityCumulative, ms.At(i).Sum().AggregationTemporality())
					dp := ms.At(i).Sum().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
					assert.Equal(t, int64(1), dp.IntValue())
				case "process.runtime.memstats.total_alloc":
					assert.False(t, validatedMetrics["process.runtime.memstats.total_alloc"], "Found a duplicate in the metrics slice: process.runtime.memstats.total_alloc")
					validatedMetrics["process.runtime.memstats.total_alloc"] = true
					assert.Equal(t, pmetric.MetricTypeSum, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Sum().DataPoints().Len())
					assert.Equal(t, "Cumulative bytes allocated for heap objects.", ms.At(i).Description())
					assert.Equal(t, "By", ms.At(i).Unit())
					assert.Equal(t, true, ms.At(i).Sum().IsMonotonic())
					assert.Equal(t, pmetric.AggregationTemporalityCumulative, ms.At(i).Sum().AggregationTemporality())
					dp := ms.At(i).Sum().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
					assert.Equal(t, int64(1), dp.IntValue())
				}
			}
		})
	}
}

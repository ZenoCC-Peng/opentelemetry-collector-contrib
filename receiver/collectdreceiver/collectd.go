// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package collectdreceiver // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/collectdreceiver"

import (
	"encoding/json"
	"fmt"
	"github.com/open-telemetry/opentelemetry-collector-contrib/internal/common/sanitize"
	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/pmetric"
	"strings"
	"time"
)

type TargetMetricType string

const (
	GaugeMetricType      = TargetMetricType("gauge")
	CumulativeMetricType = TargetMetricType("cumulative")
)

const (
	collectDMetricDerive  = "derive"
	collectDMetricCounter = "counter"
)

type collectDRecord struct {
	Dsnames        []*string              `json:"dsnames"`
	Dstypes        []*string              `json:"dstypes"`
	Host           *string                `json:"host"`
	Interval       *json.Number           `json:"interval"`
	Plugin         *string                `json:"plugin"`
	PluginInstance *string                `json:"plugin_instance"`
	Time           *json.Number           `json:"time"`
	TypeS          *string                `json:"type"`
	TypeInstance   *string                `json:"type_instance"`
	Values         []*json.Number         `json:"values"`
	Message        *string                `json:"message"`
	Meta           map[string]interface{} `json:"meta"`
	Severity       *string                `json:"severity"`
}

type createMetricInfo struct {
	Name   string
	DsType *string
	Val    *json.Number
}

func (collectd *collectDRecord) isEvent() bool {
	return collectd.Time != nil && collectd.Severity != nil && collectd.Message != nil
}

func (collectd *collectDRecord) protoTime() pcommon.Timestamp {
	if collectd.Time == nil {
		return pcommon.NewTimestampFromTime(time.Time{})
	}
	collectedTime, _ := parseTime(*collectd.Time)
	timeStamp := time.Unix(0, 0).Add(collectedTime)
	return pcommon.NewTimestampFromTime(timeStamp)
}

func (collectd *collectDRecord) startTimestamp(metricType TargetMetricType) pcommon.Timestamp {

	collectedTime, _ := parseTime(*collectd.Time)
	collectdInterval, _ := parseTime(*collectd.Interval)
	timeDiff := collectedTime - collectdInterval
	if metricType == CumulativeMetricType {
		return pcommon.NewTimestampFromTime(time.Unix(0, 0).Add(timeDiff))
	}
	return pcommon.NewTimestampFromTime(time.Time{})
}

func parseTime(timeValue json.Number) (time.Duration, error) {
	timeStamp := timeValue.String()
	duration, err := time.ParseDuration(timeStamp + "s")
	return duration, err
}

func (collectd *collectDRecord) appendToMetrics(scopeMetrics pmetric.ScopeMetrics, defaultLabels map[string]string) error {

	if collectd.isEvent() {
		recordEventsReceived()
		return nil
	}

	recordMetricsReceived()

	labels := make(map[string]string, len(defaultLabels))

	for k, v := range defaultLabels {
		labels[k] = v
	}

	for i := range collectd.Dsnames {

		if i < len(collectd.Dstypes) && i < len(collectd.Values) && collectd.Values[i] != nil {
			dsType, dsName, val := collectd.Dstypes[i], collectd.Dsnames[i], collectd.Values[i]
			metricName, usedDsName := collectd.getReasonableMetricName(i, labels)
			createMetric := createMetricInfo{
				Name:   metricName,
				DsType: dsType,
				Val:    val,
			}

			addIfNotNullOrEmpty(labels, "plugin", collectd.Plugin)
			parseAndAddLabels(labels, collectd.PluginInstance, collectd.Host)
			if !usedDsName {
				addIfNotNullOrEmpty(labels, "dsname", dsName)
			}

			metric, err := collectd.newMetric(createMetric, labels)
			if err != nil {
				return fmt.Errorf("error processing metric %s: %w", sanitize.String(metricName), err)
			}

			newMetric := scopeMetrics.Metrics().AppendEmpty()
			metric.MoveTo(newMetric)
		}
	}
	return nil
}

// Create new metric, get labels, then setting attribute and metric info
// Returns:
// A new Metric
func (collectd *collectDRecord) newMetric(createMetric createMetricInfo, labels map[string]string) (pmetric.Metric, error) {
	attributes := setAttributes(labels)
	metric, err := collectd.setMetric(createMetric, attributes)
	if err != nil {
		return pmetric.Metric{}, fmt.Errorf("error processing metric %s: %w", createMetric.Name, err)
	}
	return metric, nil
}

// Set new metric info with name, datapoint, time, attributes
// Returns:
// new Metric
func (collectd *collectDRecord) setMetric(createMetric createMetricInfo, attributes pcommon.Map) (pmetric.Metric, error) {
	typ := ""
	metric := pmetric.NewMetric()

	if createMetric.DsType != nil {
		typ = *createMetric.DsType
	}

	metric.SetName(createMetric.Name)
	dataPoint := collectd.setDataPoint(typ, metric)
	// todo: ask from pst to utc is ok???
	dataPoint.SetTimestamp(collectd.protoTime())
	attributes.CopyTo(dataPoint.Attributes())

	if val, err := createMetric.Val.Int64(); err == nil {
		dataPoint.SetIntValue(val)
	} else if val, err := createMetric.Val.Float64(); err == nil {
		dataPoint.SetDoubleValue(val)
	} else {
		return pmetric.Metric{}, fmt.Errorf("value could not be decoded: %w", err)
	}

	return metric, nil
}

func (collectd *collectDRecord) setDataPoint(typ string, metric pmetric.Metric) pmetric.NumberDataPoint {
	var dataPoint pmetric.NumberDataPoint
	switch typ {
	case collectDMetricCounter, collectDMetricDerive:
		sum := metric.SetEmptySum()
		sum.SetIsMonotonic(true)
		dataPoint = sum.DataPoints().AppendEmpty()
	default:
		dataPoint = metric.SetEmptyGauge().DataPoints().AppendEmpty()
	}
	return dataPoint
}

// getReasonableMetricName creates metrics names by joining them (if non empty) type.typeinstance
// if there are more than one dsname append .dsname for the particular uint. if there's only one it
// becomes a dimension.
func (collectd *collectDRecord) getReasonableMetricName(index int, attrs map[string]string) (string, bool) {
	usedDsName := false
	capacity := 0
	if collectd.TypeS != nil {
		capacity += len(*collectd.TypeS)
	}
	if collectd.TypeInstance != nil {
		capacity += len(*collectd.TypeInstance)
	}
	parts := make([]byte, 0, capacity)

	if !isNilOrEmpty(collectd.TypeS) {
		parts = append(parts, *collectd.TypeS...)
	}
	parts = collectd.pointTypeInstance(attrs, parts)
	if collectd.Dsnames != nil && !isNilOrEmpty(collectd.Dsnames[index]) && len(collectd.Dsnames) > 1 {
		if len(parts) > 0 {
			parts = append(parts, '.')
		}
		parts = append(parts, *collectd.Dsnames[index]...)
		usedDsName = true
	}
	return string(parts), usedDsName
}

// pointTypeInstance extracts information from the TypeInstance field and appends to the metric name when possible.
func (collectd *collectDRecord) pointTypeInstance(attrs map[string]string, parts []byte) []byte {
	if isNilOrEmpty(collectd.TypeInstance) {
		return parts
	}

	instanceName, extractedAttrs := LabelsFromName(collectd.TypeInstance)
	if instanceName != "" {
		if len(parts) > 0 {
			parts = append(parts, '.')
		}
		parts = append(parts, instanceName...)
	}
	for k, v := range extractedAttrs {
		if _, exists := attrs[k]; !exists {
			val := v
			addIfNotNullOrEmpty(attrs, k, &val)
		}
	}
	return parts
}

// LabelsFromName tries to pull out dimensions out of name in the format
// "name[k=v,f=x]-more_name".
// For the example above it would return "name-more_name" and extract dimensions
// (k,v) and (f,x).
// If something unexpected is encountered it returns the original metric name.
//
// The code tries to avoid allocation by using local slices and avoiding calls
// to functions like strings.Slice.
func LabelsFromName(val *string) (metricName string, labels map[string]string) {
	metricName = *val

	index := strings.Index(*val, "[")
	if index > -1 {
		left := (*val)[:index]
		rest := (*val)[index+1:]

		index = strings.Index(rest, "]")
		if index > -1 {
			working := make(map[string]string)
			dimensions := rest[:index]
			rest = rest[index+1:]
			cindex := strings.Index(dimensions, ",")
			prev := 0
			for {
				if cindex < prev {
					cindex = len(dimensions)
				}
				piece := dimensions[prev:cindex]
				tindex := strings.Index(piece, "=")
				if tindex == -1 || strings.Contains(piece[tindex+1:], "=") {
					return
				}
				working[piece[:tindex]] = piece[tindex+1:]
				if cindex == len(dimensions) {
					break
				}
				prev = cindex + 1
				cindex = strings.Index(dimensions[prev:], ",") + prev
			}
			labels = working
			metricName = left + rest
		}
	}
	return
}

func isNilOrEmpty(str *string) bool {
	return str == nil || *str == ""
}

func addIfNotNullOrEmpty(m map[string]string, key string, val *string) {
	if val != nil && *val != "" {
		m[key] = *val
	}
}

func parseAndAddLabels(labels map[string]string, pluginInstance *string, host *string) {
	parseNameForLabels(labels, "plugin_instance", pluginInstance)
	parseNameForLabels(labels, "host", host)
}

func parseNameForLabels(labels map[string]string, key string, val *string) {
	instanceName, toAddDims := LabelsFromName(val)

	for k, v := range toAddDims {
		if _, exists := labels[k]; !exists {
			val := v
			addIfNotNullOrEmpty(labels, k, &val)
		}
	}
	addIfNotNullOrEmpty(labels, key, &instanceName)
}

func setAttributes(labels map[string]string) pcommon.Map {

	attributes := pcommon.NewMap()
	for k, v := range labels {
		attributes.PutStr(k, v)
	}
	return attributes
}

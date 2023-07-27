// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package loadscraper // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/hostmetricsreceiver/internal/scraper/loadscraper"

import (
	"context"
	"errors"
	"fmt"
	"go.opentelemetry.io/collector/receiver/scrapererror"
	"runtime"
	"time"

	"github.com/shirou/gopsutil/v3/host"
	"github.com/shirou/gopsutil/v3/load"
	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/pmetric"
	"go.opentelemetry.io/collector/receiver"
	"go.uber.org/zap"

	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/hostmetricsreceiver/internal/perfcounters"
	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/hostmetricsreceiver/internal/scraper/loadscraper/internal/metadata"
)

const metricsLen = 3

// scraper for Load Metrics
type scraper struct {
	settings   receiver.CreateSettings
	config     *Config
	mb         *metadata.MetricsBuilder
	skipScrape bool

	// for mocking
	bootTime func() (uint64, error)
	load     func() (*load.AvgStat, error)
}

// newLoadScraper creates a set of Load related metrics
func newLoadScraper(_ context.Context, settings receiver.CreateSettings, cfg *Config) *scraper {
	return &scraper{settings: settings, config: cfg, bootTime: host.BootTime, load: getSampledLoadAverages}
}

// start
func (s *scraper) start(ctx context.Context, _ component.Host) error {
	bootTime, err := s.bootTime()
	if err != nil {
		return err
	}

	s.mb = metadata.NewMetricsBuilder(s.config.MetricsBuilderConfig, s.settings, metadata.WithStartTime(pcommon.Timestamp(bootTime*1e9)))

	currentTime := time.Now()
	// Define the future time for the range
	// Add 5 mins to the current time
	//futureTime := currentTime.Add(5 * time.Minute)
	//timeRange := int(futureTime.Sub(currentTime).Minutes())
	//maxTime := 5
	//var numCPU int
	//for numCPU < 0 {
	//	numCPU = runtime.NumCPU()
	//	time.Sleep(1 * time.Second)
	//	if timeRange > maxTime {
	//		// If it does, throw an error
	//		err := errors.New("Exceeds the maximum waiting time")
	//		fmt.Println("Error:", err)
	//		return err
	//	}
	//}

	err = startSampling(ctx, s.settings.Logger)

	var initErr *perfcounters.PerfCounterInitError
	switch {
	case errors.As(err, &initErr):
		// This indicates, on Windows, that the performance counters can't be scraped.
		// In order to prevent crashing in a fragile manner, we simply skip scraping.
		s.settings.Logger.Error("Failed to init performance counters, load metrics will not be scraped", zap.Error(err))
		s.skipScrape = true
	case err != nil:
		// Unknown error; fail to start if this is the case
		return err
	}

	return nil
}

// shutdown
func (s *scraper) shutdown(ctx context.Context) error {
	if s.skipScrape {
		// We skipped scraping because the sampler failed to start,
		// so it doesn't need to be shut down.
		return nil
	}
	return stopSampling(ctx)
}

// scrape
func (s *scraper) scrape(_ context.Context) (pmetric.Metrics, error) {
	if s.skipScrape {
		return pmetric.NewMetrics(), nil
	}

	now := pcommon.NewTimestampFromTime(time.Now())

	var avgLoadValues *load.AvgStat
	//avgLoadValues, err := s.load()
	//fmt.Println("env:", runtime.GOOS)
	//if runtime.GOOS == "windows" {
	//	time.Sleep(5 * time.Second)
	//}

	avgLoadValues, err := s.load()
	if err != nil {
		return pmetric.NewMetrics(), scrapererror.NewPartialScrapeError(err, metricsLen)
	}

	for avgLoadValues.Load1 == 0 && avgLoadValues.Load5 == 0 && avgLoadValues.Load15 == 0 {
		avgLoadValues, err = s.load()
		if err != nil {
			return pmetric.NewMetrics(), scrapererror.NewPartialScrapeError(err, metricsLen)
		}

		time.Sleep(1 * time.Second)
	}
	//time.Sleep(5 * time.Second)
	//avgLoadValues, err := s.load()
	//if err != nil {
	//	return pmetric.NewMetrics(), scrapererror.NewPartialScrapeError(err, metricsLen)
	//}

	if s.config.CPUAverage {
		divisor := float64(runtime.NumCPU())
		avgLoadValues.Load1 /= divisor
		avgLoadValues.Load5 /= divisor
		avgLoadValues.Load15 /= divisor
	}

	fmt.Println("load1:", avgLoadValues.Load1)
	fmt.Println("load5:", avgLoadValues.Load5)
	fmt.Println("load15:", avgLoadValues.Load15)

	s.mb.RecordSystemCPULoadAverage1mDataPoint(now, avgLoadValues.Load1)
	s.mb.RecordSystemCPULoadAverage5mDataPoint(now, avgLoadValues.Load5)
	s.mb.RecordSystemCPULoadAverage15mDataPoint(now, avgLoadValues.Load15)

	return s.mb.Emit(), nil
}

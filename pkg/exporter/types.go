package exporter

import "github.com/deifyed/motoblog/pkg/summary"

type Interface interface {
	Export(trip summary.Trip) error
}

package svc

import (
	"net/http"

	"github.com/owncloud/mono/phoenix/pkg/metrics"
)

// NewInstrument returns a service that instruments metrics.
func NewInstrument(next Service, metrics *metrics.Metrics) Service {
	return instrument{
		next:    next,
		metrics: metrics,
	}
}

type instrument struct {
	next    Service
	metrics *metrics.Metrics
}

// ServeHTTP implements the Service interface.
func (i instrument) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	i.next.ServeHTTP(w, r)
}

// Config implements the Service interface.
func (i instrument) Config(w http.ResponseWriter, r *http.Request) {
	i.next.Config(w, r)
}

package goa

import (
	"net/http"
	"time"

	"github.com/Redocly/souin/configurationtypes"
	"github.com/Redocly/souin/pkg/middleware"
	"github.com/Redocly/souin/plugins/souin/storages"
)

var (
	DefaultConfiguration = middleware.BaseConfiguration{
		DefaultCache: &configurationtypes.DefaultCache{
			TTL: configurationtypes.Duration{
				Duration: 10 * time.Second,
			},
			Stale: configurationtypes.Duration{
				Duration: 10 * time.Second,
			},
		},
		LogLevel: "info",
	}
	DevDefaultConfiguration = middleware.BaseConfiguration{
		API: configurationtypes.API{
			BasePath: "/souin-api",
			Prometheus: configurationtypes.APIEndpoint{
				Enable: true,
			},
			Souin: configurationtypes.APIEndpoint{
				Enable: true,
			},
		},
		DefaultCache: &configurationtypes.DefaultCache{
			Regex: configurationtypes.Regex{
				Exclude: "/excluded",
			},
			TTL: configurationtypes.Duration{
				Duration: 5 * time.Second,
			},
			Stale: configurationtypes.Duration{
				Duration: 10 * time.Second,
			},
		},
		LogLevel: "debug",
	}
)

// SouinGoaMiddleware declaration.
type SouinGoaMiddleware struct {
	*middleware.SouinBaseHandler
}

func NewHTTPCache(c middleware.BaseConfiguration) func(http.Handler) http.Handler {
	storages.InitFromConfiguration(&c)
	s := &SouinGoaMiddleware{
		SouinBaseHandler: middleware.NewHTTPCacheHandler(&c),
	}

	return s.handle
}

func (s *SouinGoaMiddleware) handle(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, rq *http.Request) {
		_ = s.SouinBaseHandler.ServeHTTP(rw, rq, func(w http.ResponseWriter, r *http.Request) error {
			next.ServeHTTP(w, r)

			return nil
		})
	})
}

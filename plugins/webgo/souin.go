package webgo

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
		},
		LogLevel: "debug",
	}
)

// SouinWebgoMiddleware declaration.
type SouinWebgoMiddleware struct {
	*middleware.SouinBaseHandler
}

func NewHTTPCache(c middleware.BaseConfiguration) *SouinWebgoMiddleware {
	storages.InitFromConfiguration(&c)
	return &SouinWebgoMiddleware{
		SouinBaseHandler: middleware.NewHTTPCacheHandler(&c),
	}
}

func (s *SouinWebgoMiddleware) Middleware(rw http.ResponseWriter, rq *http.Request, next http.HandlerFunc) {
	_ = s.SouinBaseHandler.ServeHTTP(rw, rq, func(w http.ResponseWriter, r *http.Request) error {
		next(w, r)

		return nil
	})
}

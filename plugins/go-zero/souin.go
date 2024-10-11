package souin

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

// SouinGoZeroMiddleware declaration.
type SouinGozeroMiddleware struct {
	*middleware.SouinBaseHandler
}

func NewHTTPCache(c middleware.BaseConfiguration) *SouinGozeroMiddleware {
	storages.InitFromConfiguration(&c)
	return &SouinGozeroMiddleware{
		SouinBaseHandler: middleware.NewHTTPCacheHandler(&c),
	}
}

func (s *SouinGozeroMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		_ = s.SouinBaseHandler.ServeHTTP(rw, r, func(w http.ResponseWriter, r *http.Request) error {
			next(w, r)

			return nil
		})
	}
}

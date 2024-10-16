package context

import (
	"context"
	"net/http"
	"net/http/httptest"
	"regexp"
	"testing"

	"github.com/caddyserver/caddy/v2"
	"github.com/Redocly/souin/configurationtypes"
)

func Test_KeyContext_SetupContext(t *testing.T) {
	ctx := keyContext{}
	ctx.SetupContext(&testConfiguration{
		defaultCache: &configurationtypes.DefaultCache{
			Key: configurationtypes.Key{},
		},
	})
	if ctx.disable_host {
		t.Errorf("The host must be disabled.")
	}
	if ctx.disable_method {
		t.Errorf("The method must be disabled.")
	}

	rg := configurationtypes.RegValue{
		Regexp: regexp.MustCompile(".*"),
	}
	m := configurationtypes.CacheKeys{
		configurationtypes.CacheKey{
			rg: configurationtypes.Key{
				DisableHost:   true,
				DisableMethod: true,
			},
		},
	}
	ctx.SetupContext(&testConfiguration{
		defaultCache: &configurationtypes.DefaultCache{
			Key: configurationtypes.Key{
				DisableHost:   true,
				DisableMethod: true,
			},
		},
		cacheKeys: m,
	})

	if !ctx.disable_host {
		t.Errorf("The host must be enabled.")
	}
	if !ctx.disable_method {
		t.Errorf("The method must be enabled.")
	}
	if !ctx.overrides[0][rg.Regexp].disable_host {
		t.Errorf("The host must be enabled.")
	}
	if !ctx.overrides[0][rg.Regexp].disable_method {
		t.Errorf("The method must be enabled.")
	}
}

func Test_KeyContext_SetContext(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "http://domain.com", nil)
	ctx := keyContext{
		initializer: func(r *http.Request) *http.Request {
			return r.WithContext(context.WithValue(r.Context(), caddy.ReplacerCtxKey, caddy.NewReplacer()))
		},
	}
	req = ctx.SetContext(req.WithContext(context.WithValue(req.Context(), HashBody, "-with_the_hash")))
	if req.Context().Value(Key).(string) != "GET-http-domain.com--with_the_hash" {
		t.Errorf("The Key context must be equal to GET-http-domain.com--with_the_hash, %s given.", req.Context().Value(Key).(string))
	}

	m := map[*regexp.Regexp]keyContext{
		regexp.MustCompile("/matched"): {
			disable_host:   true,
			disable_method: false,
		},
	}
	ctx2 := keyContext{
		disable_host:   true,
		disable_method: true,
		overrides:      []map[*regexp.Regexp]keyContext{m},
		initializer: func(r *http.Request) *http.Request {
			return r.WithContext(context.WithValue(r.Context(), caddy.ReplacerCtxKey, caddy.NewReplacer()))
		},
	}
	req2 := httptest.NewRequest(http.MethodGet, "http://domain.com/matched", nil)
	req2 = ctx2.SetContext(req2.WithContext(context.WithValue(req2.Context(), HashBody, "")))
	if req2.Context().Value(Key).(string) != "GET-http-/matched" {
		t.Errorf("The Key context must be equal to GET-http-/matched, %s given.", req2.Context().Value(Key).(string))
	}

	m = map[*regexp.Regexp]keyContext{
		regexp.MustCompile("/matched"): {
			disable_host:   true,
			disable_method: false,
		},
	}
	ctx3 := keyContext{
		disable_method: true,
		overrides:      []map[*regexp.Regexp]keyContext{m},
		initializer: func(r *http.Request) *http.Request {
			return r.WithContext(context.WithValue(r.Context(), caddy.ReplacerCtxKey, caddy.NewReplacer()))
		},
	}
	req3 := httptest.NewRequest(http.MethodGet, "http://domain.com/matched", nil)
	req3 = ctx3.SetContext(req3.WithContext(context.WithValue(req3.Context(), HashBody, "")))
	if req3.Context().Value(Key).(string) != "GET-http-/matched" {
		t.Errorf("The Key context must be equal to GET-http-/matched, %s given.", req3.Context().Value(Key).(string))
	}

	req4 := httptest.NewRequest(http.MethodGet, "http://domain.com/something", nil)
	req4 = ctx3.SetContext(req4.WithContext(context.WithValue(req4.Context(), HashBody, "")))
	if req4.Context().Value(Key).(string) != "http-domain.com-/something" {
		t.Errorf("The Key context must be equal to http-domain.com-/something, %s given.", req4.Context().Value(Key).(string))
	}

	// Added tests for disable_query
	ctx4 := keyContext{
		disable_query:  true,
		disable_method: false,
		disable_host:   false,
		initializer: func(r *http.Request) *http.Request {
			return r.WithContext(context.WithValue(r.Context(), caddy.ReplacerCtxKey, caddy.NewReplacer()))
		},
	}
	req5 := httptest.NewRequest(http.MethodGet, "http://domain.com/matched?query=string", nil)
	req5 = ctx4.SetContext(req5.WithContext(context.WithValue(req5.Context(), HashBody, "")))
	if req5.Context().Value(Key).(string) != "GET-http-domain.com-/matched" {
		t.Errorf("The Key context must be equal to GET-http-domain.com-/matched, %s given.", req5.Context().Value(Key).(string))
	}

	ctx5 := keyContext{
		disable_query:  false,
		disable_method: false,
		disable_host:   false,
		initializer: func(r *http.Request) *http.Request {
			return r.WithContext(context.WithValue(r.Context(), caddy.ReplacerCtxKey, caddy.NewReplacer()))
		},
	}
	req6 := httptest.NewRequest(http.MethodGet, "http://domain.com/matched?query=string", nil)
	req6 = ctx5.SetContext(req6.WithContext(context.WithValue(req6.Context(), HashBody, "")))
	if req6.Context().Value(Key).(string) != "GET-http-domain.com-/matched?query=string" {
		t.Errorf("The Key context must be equal to GET-http-domain.com-/matched?query=string, %s given.", req6.Context().Value(Key).(string))
	}

}

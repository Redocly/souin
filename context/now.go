package context

import (
	"context"
	"net/http"
	"time"

	"github.com/Redocly/souin/configurationtypes"
)

const Now ctxKey = "souin_ctx.NOW"

type nowContext struct{}

func (*nowContext) SetContextWithBaseRequest(req *http.Request, _ *http.Request) *http.Request {
	return req
}

func (cc *nowContext) SetupContext(_ configurationtypes.AbstractConfigurationInterface) {}

func (cc *nowContext) SetContext(req *http.Request) *http.Request {
	now := time.Now().UTC()
	req.Header.Set("Date", now.Format(time.RFC1123))
	return req.WithContext(context.WithValue(req.Context(), Now, now))
}

var _ ctx = (*nowContext)(nil)

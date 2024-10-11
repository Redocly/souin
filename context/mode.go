package context

import (
	"context"
	"net/http"

	"github.com/Redocly/souin/configurationtypes"
)

const Mode ctxKey = "souin_ctx.MODE"

type ModeContext struct {
	Strict, Bypass_request, Bypass_response bool
}

func (*ModeContext) SetContextWithBaseRequest(req *http.Request, _ *http.Request) *http.Request {
	return req
}

func (mc *ModeContext) SetupContext(c configurationtypes.AbstractConfigurationInterface) {
	mode := c.GetDefaultCache().GetMode()
	mc.Bypass_request = mode == "bypass" || mode == "bypass_request"
	mc.Bypass_response = mode == "bypass" || mode == "bypass_response"
	mc.Strict = !mc.Bypass_request && !mc.Bypass_response
	c.GetLogger().Debugf("The cache logic will run as %s: %+v", mode, mc)
}

func (mc *ModeContext) SetContext(req *http.Request) *http.Request {
	return req.WithContext(context.WithValue(req.Context(), Mode, mc))
}

var _ ctx = (*ModeContext)(nil)

package requestid

import (
	"github.com/savsgio/atreugo/v11"
	"github.com/savsgio/gotils/uuid"
)

// Config configuration.
type Config struct{}

// New returns the middleware to adds an identifier to the request
// Header name: X-Request-ID
// Header value: uuid4()
//
// It's recomemded to add this middleware at first and before the view execution.
func New(cfg Config) atreugo.Middleware {
	return func(ctx *atreugo.RequestCtx) error {
		if ctx.Request.Header.Peek(atreugo.XRequestIDHeader) == nil {
			ctx.Request.Header.Set(atreugo.XRequestIDHeader, uuid.V4())
		}

		return ctx.Next() // nolint:wrapcheck
	}
}

// Package requestid is used to add/extract a request identifier to/from a
// context.Context.
package requestid

import "context"

type (
	ctxKey struct{}
)

// ToContext adds a request id to a context.
func ToContext(ctx context.Context, id string) context.Context {
	return context.WithValue(ctx, ctxKey{}, id)
}

// FromContext obtains a request id from a context.
func FromContext(ctx context.Context) string {
	id, ok := ctx.Value(ctxKey{}).(string)
	if !ok {
		return ""
	}
	return id
}

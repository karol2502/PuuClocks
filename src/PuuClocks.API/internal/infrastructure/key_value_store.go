package infrastructure

import "context"

type KeyValueStore interface {
	Get(context.Context, string) (string,error)
	Set(context.Context, string, string) error
	Health(context.Context) error
}

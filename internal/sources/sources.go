package sources

import "context"

type PersonSource interface {
	PerformQuery(ammo int, ctx context.Context) error
}

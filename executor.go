package t38c

import (
	"context"

	"github.com/b3q/tile38-client/transport"
)

var _ Executor = (*transport.RadixPool)(nil)

// Executor represents Tile38 connection.
// Communication should be in JSON format only.
type Executor interface {
	Execute(command string, args ...string) ([]byte, error)
	ExecuteStream(ctx context.Context, handler func([]byte) error, command string, args ...string) error
}

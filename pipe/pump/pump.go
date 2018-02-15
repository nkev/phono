package pump

import (
	"context"

	"github.com/dudk/phono"
)

//Pump provides an interface for sources of samples
type Pump interface {
	Pump(ctx context.Context) (out <-chan phono.Message, errc <-chan error, err error)
}

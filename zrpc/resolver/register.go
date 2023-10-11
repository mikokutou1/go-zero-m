package resolver

import (
	"github.com/mikokutou1/go-zero-m/zrpc/resolver/internal"
)

// Register registers schemes defined zrpc.
// Keep it in a separated package to let third party register manually.
func Register() {
	internal.RegisterResolver()
}

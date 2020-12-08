package wtclientrpc

import (
	"github.com/pkt-cash/pktd/lnd/lncfg"
	"github.com/pkt-cash/pktd/lnd/watchtower/wtclient"
)

// Config is the primary configuration struct for the watchtower RPC server. It
// contains all the items required for the RPC server to carry out its duties.
// The fields with struct tags are meant to be parsed as normal configuration
// options, while if able to be populated, the latter fields MUST also be
// specified.
type Config struct {
	// Active indicates if the watchtower client is enabled.
	Active bool

	// Client is the backing watchtower client that we'll interact with
	// through the watchtower RPC subserver.
	Client wtclient.Client

	// Resolver is a custom resolver that will be used to resolve watchtower
	// addresses to ensure we don't leak any information when running over
	// non-clear networks, e.g. Tor, etc.
	Resolver lncfg.TCPResolver
}

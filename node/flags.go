// Copyright (c) 2020 The Meter.io developers

// Distributed under the GNU Lesser General Public License v3.0 software license, see the accompanying
// file LICENSE or <https://www.gnu.org/licenses/lgpl-3.0.html>

package node

import (
	"log/slog"

	cli "gopkg.in/urfave/cli.v1"
)

var (
	dataDirFlag = cli.StringFlag{
		Name:   "data-dir",
		Value:  defaultDataDir(),
		Usage:  "directory for block-chain databases",
		EnvVar: "METER_DATA_DIR",
	}
	apiAddrFlag = cli.StringFlag{
		Name:  "api-addr",
		Value: "localhost:8669",
		Usage: "API service listening address",
	}
	apiCorsFlag = cli.StringFlag{
		Name:  "api-cors",
		Value: "",
		Usage: "comma separated list of domains from which to accept cross origin requests to API",
	}
	apiTimeoutFlag = cli.IntFlag{
		Name:  "api-timeout",
		Value: 10000,
		Usage: "API request timeout value in milliseconds",
	}
	apiBacktraceLimitFlag = cli.IntFlag{
		Name:  "api-backtrace-limit",
		Value: 1000,
		Usage: "limit the distance between 'position' and best block for subscriptions APIs",
	}
	verbosityFlag = cli.IntFlag{
		Name:  "verbosity",
		Value: int(slog.LevelInfo),
		Usage: "log verbosity (0-9)",
	}
	peersFlag = cli.StringSliceFlag{
		Name:  "peers, P",
		Usage: "P2P peers in enode format",
	}
	maxPeersFlag = cli.IntFlag{
		Name:  "max-peers",
		Usage: "maximum number of P2P network peers (P2P network disabled if set to 0)",
		Value: 25,
	}
	p2pPortFlag = cli.IntFlag{
		Name:  "p2p-port",
		Value: 11235,
		Usage: "P2P network listening port",
	}
	natFlag = cli.StringFlag{
		Name:  "nat",
		Value: "any",
		Usage: "port mapping mechanism (any|none|upnp|pmp|extip:<IP>)",
	}
	noDiscoverFlag = cli.BoolFlag{
		Name:  "no-discover",
		Usage: "disable auto discovery mode",
	}
	minCommitteeSizeFlag = cli.IntFlag{
		Name:  "committee-min-size",
		Usage: "committee minimum size",
		Value: 15,
	}
	maxCommitteeSizeFlag = cli.IntFlag{
		Name:  "committee-max-size",
		Usage: "committee maximum size",
		Value: 50,
	}
	discoServerFlag = cli.StringSliceFlag{
		Name:  "disco-server",
		Usage: "override the default discover servers setting",
	}
	discoTopicFlag = cli.StringFlag{
		Name:  "disco-topic",
		Usage: "set the custom discover topics",
		Value: "default-topic",
	}
	epochBlockCountFlag = cli.Int64Flag{
		Name:  "epoch-mblock-count",
		Usage: "mblock count between epochs",
		Value: 1200,
	}
)

var Flags = []cli.Flag{
	dataDirFlag,
	apiAddrFlag,
	apiCorsFlag,
	apiTimeoutFlag,
	apiBacktraceLimitFlag,
	verbosityFlag,
	maxPeersFlag,
	p2pPortFlag,
	natFlag,
	peersFlag,
	noDiscoverFlag,
	minCommitteeSizeFlag,
	maxCommitteeSizeFlag,
	discoServerFlag,
	discoTopicFlag,
	epochBlockCountFlag,
}
package btc

import "runtime"

const(
	SourcesTag = "0.5.5"

	MAX_BLOCK_SIZE = 1e6
	COIN = 1e8
	MAX_MONEY = 21000000 * COIN

	BlockMapInitLen = 300e3

	MessageMagic = "Bitcoin Signed Message:\n"

	MovingCheckopintDepth = 2016  // Do not accept forks that wold go deeper in a past
)

// Increase the number of threads to optimize txs verification time,
// proportionaly among cores, but if you set it too high, the UI and
// network threads may be laggy while parsing blocks.
var UseThreads int = 4 * runtime.NumCPU()

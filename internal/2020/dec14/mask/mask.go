package mask

import (
	"fmt"
)

// LastIndex is a simple const I use for meaningfulness
const LastIndex = uint8(35)

// PowerOf2 returns 2 to the power n
func PowerOf2(n uint8) uint64 {
	if n > LastIndex {
		panic(fmt.Sprintf("that's a very big index: %d, we only tolerate numbers up to %d", n, LastIndex))
	}
	return 1 << n
}

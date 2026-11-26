package uuid

import (
	"encoding/binary"
	"github.com/pkg/errors"
	"strconv"
	"sync"
	"time"
)

var (
	Epoch int64 = 0

	// Number of bits to use for Node
	NodeBits uint8 = 12

	// Number of bits to use for Step
	StepBits uint8 = 10

	nodeMax   int64 = -1 ^ (-1 << NodeBits)
	nodeMask  int64 = nodeMax << StepBits
	stepMask  int64 = -1 ^ (-1 << StepBits)
	timeShift uint8 = NodeBits + StepBits
	nodeShift uint8 = StepBits
)

// A Node struct holds the basic information needed for a snowflake generator
// node
type Node struct {
	mu   sync.Mutex
	time int64
	node int64
	step int64
}

// An ID is a custom type used for a snowflake ID.  This is used so we can
// attach methods onto the ID.
type ID int64

// NewNode returns a new snowflake node that can be used to generate snowflake
// IDs
func NewNode(t, node, step int64) (*Node, error) {

	// re-calc in case custom NodeBits or StepBits were set
	nodeMax = -1 ^ (-1 << NodeBits)
	nodeMask = nodeMax << StepBits
	stepMask = -1 ^ (-1 << StepBits)
	timeShift = NodeBits + StepBits
	nodeShift = StepBits

	now := time.Now().UnixNano() / 1000000
	if t > now || t <= 0 {
		t = now
	}

	if node < 0 || node > nodeMax {
		return nil, errors.Errorf("Node number must be between 0 and %d", nodeMax)
	}

	if step < 0 || step > stepMask {
		return nil, errors.Errorf("Step number must be between 0 and %d", stepMask)
	}

	return &Node{
		time: t,
		node: node,
		step: step,
	}, nil
}

// Generate creates and returns a unique snowflake ID
func (n *Node) Generate() ID {

	n.mu.Lock()

	now := time.Now().UnixNano() / 1000000

	if n.time <= 0 {
		n.time = now
	}

	r := ID((n.time-Epoch)<<timeShift |
		(n.node << nodeShift) |
		(n.step),
	)

	n.mu.Unlock()
	return r
}

// Int64 returns an int64 of the snowflake ID
func (f ID) Int64() int64 {
	return int64(f)
}

// String returns a string of the snowflake ID
func (f ID) String() string {
	return strconv.FormatInt(int64(f), 10)
}

// Bytes returns a byte slice of the snowflake ID
func (f ID) Bytes() []byte {
	return []byte(f.String())
}

// IntBytes returns an array of bytes of the snowflake ID, encoded as a
// big endian integer.
func (f ID) IntBytes() [8]byte {
	var b [8]byte
	binary.BigEndian.PutUint64(b[:], uint64(f))
	return b
}

// Time returns an int64 unix timestamp of the snowflake ID time
func (f ID) Time() int64 {
	return (int64(f) >> timeShift) + Epoch
}

// Node returns an int64 of the snowflake ID node number
func (f ID) Node() int64 {
	return int64(f) & nodeMask >> nodeShift
}

// Step returns an int64 of the snowflake step (or sequence) number
func (f ID) Step() int64 {
	return int64(f) & stepMask
}

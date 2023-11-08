// SEE https://github.com/pascaldekloe/etherstream
// Creduts to [pascaldekloe](https://github.com/pascaldekloe)

package ethereum

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"sort"
	"time"

	// lots of poor naming in go-ethereum 👾
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	ether "github.com/ethereum/go-ethereum/common"
	chain "github.com/ethereum/go-ethereum/core/types"
)

// Reader configures the blockchain connectivity.
//
// Note that users need WebSockets when calling ethclient.DialContext, because
// subscriptions won't work with regular HTTP RPC.
type Reader struct {
	Backend ethereum.LogFilterer // blockchain connection

	// limit for a subscription request (defaults to 2 s)
	SubscribeTimeout time.Duration
	// limit for history retreival (defaults to 7 s)
	FetchTimeout time.Duration

	// idle time on which no new content can be assumed
	// (defaults to 500 ms)
	ReceiveExpire time.Duration
}

// EventsWithHistory resolves all logs matching eventType.
// The history is sorted in ascending order. The first receive from stream
// directly follows the last entry from history, if any.
func (r Reader) EventsWithHistory(ctx context.Context, eventType *abi.Event) (stream <-chan chain.Log, _ ethereum.Subscription, history []chain.Log, _ error) {
	// first topic always is the signature hash of the respective event
	return r.QueryWithHistory(ctx, &ethereum.FilterQuery{Topics: [][]ether.Hash{{eventType.ID}}})
}

// QueryWithHistory resolves all logs matching q.
// The history is sorted in ascending order. The first receive from stream
// directly follows the last entry from history, if any.
func (r Reader) QueryWithHistory(ctx context.Context, q *ethereum.FilterQuery) (stream <-chan chain.Log, _ ethereum.Subscription, history []chain.Log, _ error) {
	// limited retry on chain-reorganisation [errNoOverlap]
	const tryMax = 2
	for tryN := 1; tryN <= tryMax; tryN++ {
		// subscribe live stream
		timeout := r.SubscribeTimeout
		if timeout == 0 {
			timeout = 2 * time.Second
		}
		subscribeCtx, cancel := context.WithTimeout(ctx, timeout)
		defer cancel()
		stream := make(chan chain.Log, 120)
		sub, err := r.Backend.SubscribeFilterLogs(subscribeCtx, *q, stream)
		if err != nil {
			return nil, nil, nil, fmt.Errorf("etherstream: no subscription on log events: %w", err)
		}
		// ⚠️ must Unsubscribe

		history, err := r.linkHistory(ctx, stream, q)
		switch {
		case errors.Is(err, errNoOverlap):
			sub.Unsubscribe()
			continue

		case err != nil:
			sub.Unsubscribe()
			return nil, nil, nil, err
		}

		return stream, sub, history, err // OK
	}

	return nil, nil, nil, fmt.Errorf("etherstream: give up after %d tries: %w ", tryMax, errNoOverlap)
}

var errNoOverlap = errors.New("historic events don't match [overlap] with subscription reception—chain reorganisation assumed")

// LinkHistory returns the full history before the next entry from stream.
func (r *Reader) linkHistory(ctx context.Context, stream <-chan chain.Log, q *ethereum.FilterQuery) ([]chain.Log, error) {
	// fetch history
	timeout := r.FetchTimeout
	if timeout == 0 {
		timeout = 1 * time.Minute
	}
	ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()
	history, err := r.Backend.FilterLogs(ctx, *q)
	if err != nil {
		return nil, fmt.Errorf("etherstream: historic logs unavailable: %w", err)
	}
	if len(history) == 0 {
		return nil, nil
	}
	// lookup not sorted
	sort.Slice(history, func(i, j int) bool { return Order(&history[i], &history[j]) > 0 })

	receiveExpire := r.ReceiveExpire
	if receiveExpire == 0 {
		receiveExpire = time.Second / 2
	}
	expireTimer := time.NewTimer(receiveExpire)
	defer expireTimer.Stop()

	// find overlap in history to ensure no gaps nor duplicates
	// workaround <https://github.com/ethereum/go-ethereum/issues/15063>
	select {
	case first := <-stream:
		cmp := func(i int) int { return Order(&history[i], &first) }
		i, ok := sort.Find(len(history), cmp)
		if ok {
			return history[:i+1], nil
		}
		return nil, errNoOverlap
	case <-expireTimer.C:
		// no reception after history retreival implies no overlap
		return history, nil
	}
}

// Order returns whether b follows a, with positive for yes, zero for equal, or
// negative for no.
func Order(a, b *chain.Log) int {
	diff := int(b.BlockNumber - a.BlockNumber)
	if diff == 0 {
		diff = int(b.TxIndex - a.TxIndex)
	}
	if diff == 0 && a.TxHash != b.TxHash {
		// a and b are on differd chains
		// use an arbitrary but consistent order
		diff = bytes.Compare(b.TxHash[:], a.TxHash[:])
	}
	return diff
}

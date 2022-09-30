package bigNumber

import (
	"encoding/json"
	"fmt"
	"math/big"
	"strings"

	"github.com/yearn/ydaemon/internal/utils/logs"
)

type BigInt struct{ big.Int }
type BigFloat struct{ big.Float }

func FromBigInt(b *big.Int) BigInt {
	return BigInt{*b}
}
func FromBigFloat(b *big.Float) BigFloat {
	return BigFloat{*b}
}

func (b BigInt) MarshalJSON() ([]byte, error) {
	return json.Marshal(b.String())
}

func (b *BigInt) UnmarshalJSON(p []byte) error {
	s := strings.Trim(string(p), "\"")
	if s == "null" {
		return nil
	}
	var z big.Int
	logs.Pretty(p, s)
	_, ok := z.SetString(s, 10)
	if !ok {
		return fmt.Errorf("not a valid big integer: %s", p)
	}
	b.Int = z
	return nil
}

func (b BigFloat) MarshalJSON() ([]byte, error) {
	return json.Marshal(b.String())
}

func (b *BigFloat) UnmarshalJSON(p []byte) error {
	if string(p) == "null" {
		return nil
	}
	var z big.Float
	_, ok := z.SetString(string(p))
	if !ok {
		return fmt.Errorf("not a valid big integer: %s", p)
	}
	b.Float = z
	return nil
}

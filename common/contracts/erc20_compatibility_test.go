package contracts

import (
	"context"
	"fmt"
	"math/big"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
)

// Exercise the checked-in bindings and RPC codecs together, including values
// larger than uint64 and an explicit historical block, without a live node.
func TestERC20BindingTransports(t *testing.T) {
	for _, websocket := range []bool{false, true} {
		name := "HTTP"
		if websocket {
			name = "WebSocket"
		}
		t.Run(name, func(t *testing.T) {
			address := common.HexToAddress("0x1234567890123456789012345678901234567890")
			supply := new(big.Int).Lsh(big.NewInt(1), 200)
			server := rpc.NewServer()
			t.Cleanup(server.Stop)
			if err := server.RegisterName("eth", &erc20RPCFixture{address: address, supply: supply}); err != nil {
				t.Fatal(err)
			}
			var handler http.Handler = server
			if websocket {
				handler = server.WebsocketHandler([]string{"*"})
			}
			httpServer := httptest.NewServer(handler)
			t.Cleanup(httpServer.Close)
			url := httpServer.URL
			if websocket {
				url = "ws" + strings.TrimPrefix(url, "http")
			}
			ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
			defer cancel()
			client, err := ethclient.DialContext(ctx, url)
			if err != nil {
				t.Fatal(err)
			}
			t.Cleanup(client.Close)
			token, err := NewERC20(address, client)
			if err != nil {
				t.Fatal(err)
			}
			opts := &bind.CallOpts{Context: ctx, BlockNumber: big.NewInt(123)}
			decimals, err := token.Decimals(opts)
			if err != nil || decimals != 18 {
				t.Fatalf("decimals = %d, error = %v", decimals, err)
			}
			gotSupply, err := token.TotalSupply(opts)
			if err != nil {
				t.Fatal(err)
			}
			if gotSupply.Cmp(supply) != 0 {
				t.Fatalf("total supply = %s, want %s", gotSupply, supply)
			}
		})
	}
}

type erc20RPCFixture struct {
	address common.Address
	supply  *big.Int
}

func (f *erc20RPCFixture) Call(_ context.Context, args map[string]string, block string) (hexutil.Bytes, error) {
	if !strings.EqualFold(args["to"], f.address.Hex()) || block != "0x7b" {
		return nil, fmt.Errorf("unexpected contract or block: %v at %s", args, block)
	}
	input := args["input"]
	if input == "" {
		input = args["data"]
	}
	switch input {
	case "0x313ce567": // decimals()
		return big.NewInt(18).FillBytes(make([]byte, 32)), nil
	case "0x18160ddd": // totalSupply()
		return f.supply.FillBytes(make([]byte, 32)), nil
	default:
		return nil, fmt.Errorf("unexpected call data: %s", input)
	}
}

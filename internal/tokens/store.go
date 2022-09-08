package tokens

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/yearn/ydaemon/internal/utils/models"
)

type TStore struct {
	// TokenList contains the list of tokens we will need to fetch prices for.
	TokenList map[uint64][]common.Address

	// Tokens contains the list of address for a specific chain and some minimal information about the token
	Tokens map[uint64]map[common.Address]*models.TERC20Token
}

// Store holds the data for the partners for each chain.
var Store = TStore{}

func init() {
	Store.TokenList = make(map[uint64][]common.Address)
	Store.Tokens = make(map[uint64]map[common.Address]*models.TERC20Token)
}

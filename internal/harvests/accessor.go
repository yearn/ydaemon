package harvests

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/yearn/ydaemon/internal/models"
)

func New(harvest *models.THarvest, log types.Log) *models.THarvest {
	harvest.TxHash = log.TxHash
	harvest.BlockHash = log.BlockHash
	harvest.BlockNumber = log.BlockNumber
	harvest.TxIndex = log.TxIndex
	harvest.LogIndex = log.Index
	harvest.Removed = log.Removed
	return harvest
}

var _allHarvests = make(map[common.Address][]models.THarvest)

func GetAll() map[common.Address][]models.THarvest {
	return _allHarvests
}

func Get(vaultAddress common.Address) []models.THarvest {
	if _, ok := _allHarvests[vaultAddress]; !ok {
		return []models.THarvest{}
	}
	return _allHarvests[vaultAddress]
}

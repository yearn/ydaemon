package allocation

import (
	"math/big"

	"github.com/yearn/ydaemon/internal/strategies"
	"github.com/yearn/ydaemon/internal/utils/models"
)

// TStrategyAllocation contains the all the enabled languages for the protocol
type TStrategyAllocation struct {
	Strategy      *models.TStrategyList              `json:"strategy"`
	RiskGroup     *strategies.TStrategyGroupFromRisk `json:"riskGroup"`
	CurrentTVL    *big.Int                           `json:"currentTVL"` // value in WANT
	AvailableTVL  *big.Int                           `json:"availableTVL"`
	CurrentUSDC   *big.Int                           `json:"currencyUSDC"` // value in USDC
	AvailableUSDC *big.Int                           `json:"availableUSDC"`
}

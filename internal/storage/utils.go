package storage

import (
	"reflect"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/yearn/ydaemon/internal/models"
)

type TVersion struct {
	Major int `json:"major"`
	Minor int `json:"minor"`
	Patch int `json:"patch"`
}
type TJsonMetadata struct {
	LastUpdate    time.Time `json:"lastUpdate"`
	Version       TVersion  `json:"version"`
	ShouldRefresh bool      `json:"shouldRefresh"`
}
type TJsonVaultStorage struct {
	TJsonMetadata
	Vaults map[common.Address]models.TVault `json:"vaults"`
}
type TJsonStrategyStorage struct {
	TJsonMetadata
	Strategies map[common.Address]models.TStrategy `json:"strategies"`
}
type TJsonERC20Storage struct {
	TJsonMetadata
	Tokens map[common.Address]models.TERC20Token `json:"tokens"`
}

/**************************************************************************
** If an element is removed, the major version is bumped.
** If an element is added, the minor version is bumped.
** If an element is modified, the patch version is bumped.
**************************************************************************/
func detectVersionUpdate[T any](chainID uint64, currentVersion TVersion, previousElements map[common.Address]T, newElements map[common.Address]T) TVersion {
	shouldBumpMajor := false
	shouldBumpMinor := false
	shouldBumpPatch := false

	for address, element := range previousElements {
		if _, ok := newElements[address]; !ok {
			shouldBumpMinor = true
		} else if !reflect.DeepEqual(newElements[address], element) {
			shouldBumpPatch = true
		}
	}
	if len(previousElements) > len(newElements) {
		shouldBumpMajor = true
	}

	if shouldBumpMajor {
		currentVersion.Major++
		currentVersion.Minor = 0
		currentVersion.Patch = 0
	} else if shouldBumpMinor {
		currentVersion.Minor++
		currentVersion.Patch = 0
	} else if shouldBumpPatch {
		currentVersion.Patch++
	}
	return currentVersion
}

/**************************************************************************
** Little helper to ensure that the sync map is initialized before use.
**************************************************************************/
func safeSyncMap(source map[uint64]*sync.Map, chainID uint64) *sync.Map {
	syncMap := source[chainID]
	if syncMap == nil {
		syncMap = &sync.Map{}
		source[chainID] = syncMap
	}
	return syncMap
}

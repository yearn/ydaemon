package helpers

import (
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/assert"
)

func TestContainAddress(t *testing.T) {
	listOfAddress := []common.Address{
		common.HexToAddress("0x0"),
		common.HexToAddress("0x1"),
		common.HexToAddress("0x2"),
		common.HexToAddress("0x3"),
		common.HexToAddress("0x4"),
		common.HexToAddress("0x5"),
	}

	assert.True(t, ContainsAddress(listOfAddress, common.HexToAddress("0x0")))
	assert.True(t, ContainsAddress(listOfAddress, common.HexToAddress("0x1")))
	assert.True(t, ContainsAddress(listOfAddress, common.HexToAddress("0x2")))
	assert.True(t, ContainsAddress(listOfAddress, common.HexToAddress("0x3")))
	assert.True(t, ContainsAddress(listOfAddress, common.HexToAddress("0x4")))
	assert.True(t, ContainsAddress(listOfAddress, common.HexToAddress("0x5")))
	assert.False(t, ContainsAddress(listOfAddress, common.HexToAddress("0x6")))
	assert.False(t, ContainsAddress(listOfAddress, common.HexToAddress("0x554")))
}

func TestUniqueArrayAddress(t *testing.T) {
	listOfAddress := []common.Address{
		common.HexToAddress("0x0"),
		common.HexToAddress("0x0"),
		common.HexToAddress("0x1"),
		common.HexToAddress("0x1"),
		common.HexToAddress("0x2"),
		common.HexToAddress("0x3"),
		common.HexToAddress("0x3"),
		common.HexToAddress("0x4"),
		common.HexToAddress("0x5"),
		common.HexToAddress("0x5"),
	}
	expectedResultListOfAddress := []common.Address{
		common.HexToAddress("0x0"),
		common.HexToAddress("0x1"),
		common.HexToAddress("0x2"),
		common.HexToAddress("0x3"),
		common.HexToAddress("0x4"),
		common.HexToAddress("0x5"),
	}

	assert.Equal(t, UniqueArrayAddress(listOfAddress), expectedResultListOfAddress)
}

func TestReadAllFilesInDir(t *testing.T) {
	_, _, err := ReadAllFilesInDir(BASE_DATA_PATH+`/meta/strategies/2500/`, `.json`)
	assert.Error(t, err)

	_, _, err = ReadAllFilesInDir(BASE_DATA_PATH+`/meta/protocols/42161/`, `.json`)
	assert.Error(t, err)

	_, _, err = ReadAllFilesInDir(BASE_DATA_PATH+`/meta/protocols/250/0xDAO.json`, `.json`)
	assert.Error(t, err)

	_, _, err = ReadAllFilesInDir(BASE_DATA_PATH+`/meta/protocols/250/`, `.json`)
	assert.Nil(t, err)
}

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

	_, _, err = ReadAllFilesInDir(BASE_DATA_PATH+`/meta/protocols/250/0xDAO.json`, `.json`)
	assert.Error(t, err)

	_, _, err = ReadAllFilesInDir(BASE_DATA_PATH+`/meta/protocols/250/`, `.json`)
	assert.Nil(t, err)
}

func TestSafeString(t *testing.T) {
	assert.Equal(t, SafeString("0", "1"), "0")
	assert.Equal(t, SafeString("", "1"), "1")
	assert.Equal(t, SafeString("\"\"", "1"), "1")
}

func TestFormatUint64(t *testing.T) {
	assert.Equal(t, FormatUint64("0", uint64(0)), uint64(0))
	assert.Equal(t, FormatUint64("1", uint64(0)), uint64(1))
	assert.Equal(t, FormatUint64("0x0", uint64(1)), uint64(1))
	assert.Equal(t, FormatUint64("a", uint64(1)), uint64(1))
	assert.Equal(t, FormatUint64("", uint64(1)), uint64(1))
}

func TestFormatAmount(t *testing.T) {
	fOne, bgOne := FormatAmount("1000000000000000000", 18)
	assert.Equal(t, fOne, 1.0)
	assert.Equal(t, bgOne.String(), "1")

	fSome, _ := FormatAmount("123456789009876543", 18)
	assert.InDelta(t, fSome, 0.123456789, 0.0000001)

	{
		fSome, _ := FormatAmount("123456789009876543", -1)
		assert.InDelta(t, fSome, 1234567890098765430, 0.0000001)
	}

	{
		fSome, _ := FormatAmount("123456789009876543", 0)
		assert.InDelta(t, fSome, 123456789009876543, 0.0000001)
	}

	{
		fSome, _ := FormatAmount("123456789009876543", 3)
		assert.InDelta(t, fSome, 123456789009876.543, 0.0000001)
	}

	{
		fSome, _ := FormatAmount("123456789009876543", 6)
		assert.InDelta(t, fSome, 123456789009.876543, 0.0000001)
	}

	{
		fSome, _ := FormatAmount("123456789009876543", 9)
		assert.InDelta(t, fSome, 123456789.009876543, 0.0000001)
	}

	{
		fSome, _ := FormatAmount("123456789009876543", 12)
		assert.InDelta(t, fSome, 123456.789009876543, 0.0000001)
	}

	{
		fSome, _ := FormatAmount("123456789009876543", 15)
		assert.InDelta(t, fSome, 123.456789009876543, 0.0000001)
	}

	{
		fSome, _ := FormatAmount("123456789009876543", 21)
		assert.InDelta(t, fSome, 0.000123456789009876543, 0.0000001)
	}
}

func TestContainsUint64(t *testing.T) {
	assert.True(t, ContainsUint64([]uint64{1, 2, 3, 4, 5}, 1))
	assert.True(t, ContainsUint64([]uint64{1, 2, 3, 4, 5}, 3))
	assert.True(t, ContainsUint64([]uint64{1, 2, 3, 4, 5}, 5))
	assert.False(t, ContainsUint64([]uint64{1, 2, 3, 4, 5}, 6))
}

func TestAssertChainID(t *testing.T) {
	chainID, ok := AssertChainID("1")
	assert.True(t, ok)
	assert.Equal(t, chainID, uint64(1))

	chainID, ok = AssertChainID("2")
	assert.False(t, ok)
	assert.Equal(t, chainID, uint64(0))

	chainID, ok = AssertChainID("10")
	assert.True(t, ok)
	assert.Equal(t, chainID, uint64(10))

	chainID, ok = AssertChainID("250")
	assert.True(t, ok)
	assert.Equal(t, chainID, uint64(250))

	chainID, ok = AssertChainID("42161")
	assert.True(t, ok)
	assert.Equal(t, chainID, uint64(42161))

	chainID, ok = AssertChainID("a")
	assert.False(t, ok)
	assert.Equal(t, chainID, uint64(0))
}

func TestAssertAddress(t *testing.T) {
	address, ok := AssertAddress("0x0000000000000000000000000000000000000000", 1)
	assert.True(t, ok)
	assert.Equal(t, address, common.HexToAddress("0x0000000000000000000000000000000000000000"))

	address, ok = AssertAddress("0x3B27F92C0e212C671EA351827EDF93DB27cc0c65", 1)
	assert.True(t, ok)
	assert.Equal(t, address, common.HexToAddress("0x3B27F92C0e212C671EA351827EDF93DB27cc0c65"))

	address, ok = AssertAddress("0x3b27f92c0e212c671ea351827edf93db27cc0c65", 1)
	assert.True(t, ok)
	assert.Equal(t, address, common.HexToAddress("0x3B27F92C0e212C671EA351827EDF93DB27cc0c65"))

	address, ok = AssertAddress("0x123", 1)
	assert.False(t, ok)
	assert.Equal(t, address, common.Address{})

	address, ok = AssertAddress("0x662fBF2c1E4b04342EeBA6371ec1C7420042B86F", 1)
	assert.False(t, ok)
	assert.Equal(t, address, common.Address{})

	address, ok = AssertAddress("0x662fBF2c1E4b04342EeBA6371ec1C7420042B86F", 250)
	assert.True(t, ok)
	assert.Equal(t, address, common.HexToAddress("0x662fBF2c1E4b04342EeBA6371ec1C7420042B86F"))
}

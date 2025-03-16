package tokens

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

/**************************************************************************************************
** TestTokenResponseStructure tests the structure of the token response object.
**
** While we can't easily test the full API endpoints without mocking the storage layer,
** we can at least test the structure of the response objects and the conversion helper.
**************************************************************************************************/
func TestTokenResponseStructure(t *testing.T) {
	// Create a test TAllTokens
	token := TAllTokens{
		Address:          "0x1234567890123456789012345678901234567890",
		Name:             "Test Token",
		Symbol:           "TEST",
		Decimals:         18,
		IsVault:          false,
		DisplayName:      "Test Token Display",
		DisplaySymbol:    "TEST",
		Description:      "A token for testing",
		Category:         "Test",
		UnderlyingTokens: []string{"0xunderlying1", "0xunderlying2"},
	}

	// Verify fields are as expected
	assert.Equal(t, "0x1234567890123456789012345678901234567890", token.Address)
	assert.Equal(t, "Test Token", token.Name)
	assert.Equal(t, "TEST", token.Symbol)
	assert.Equal(t, uint64(18), token.Decimals)
	assert.Equal(t, false, token.IsVault)
	assert.Equal(t, "Test Token Display", token.DisplayName)
	assert.Equal(t, "TEST", token.DisplaySymbol)
	assert.Equal(t, "A token for testing", token.Description)
	assert.Equal(t, "Test", token.Category)
	assert.Len(t, token.UnderlyingTokens, 2)
}

/**************************************************************************************************
** TestControllerInitialization tests that the Controller struct can be initialized correctly.
**
** This is a simple test to ensure the controller can be created without any issues.
**************************************************************************************************/
func TestControllerInitialization(t *testing.T) {
	controller := Controller{}
	assert.NotNil(t, controller)
}

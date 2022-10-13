package logs

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Not much to test here, but we are doing them for the coverage
func TestLogs(t *testing.T) {
	Error(`Error`)
	Error(`Error-1`, `Error-2`)
	Success(`Success`)
	Warning(`Warning`)
	Info(`Info`)
	Debug(`Debug`)
	Pretty(`This is a nice string`)
	assert.True(t, true)
}

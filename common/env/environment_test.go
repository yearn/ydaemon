package env

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEnvironment(t *testing.T) {
	os.Clearenv()
	SetEnv(`.noenv`)

	assert.Equal(t, `https://eth.public-rpc.com`, RPC_ENDPOINTS[1])
	assert.Equal(t, `https://rpc.ftm.tools`, RPC_ENDPOINTS[250])
	assert.Equal(t, `https://arbitrum.public-rpc.com`, RPC_ENDPOINTS[42161])

	os.Setenv(`RPC_URI_FOR_1`, `rpc.macarena.finance`)
	os.Setenv(`RPC_URI_FOR_250`, `rpc.flamenco.finance`)
	os.Setenv(`RPC_URI_FOR_42161`, `rpc.rock.finance`)
	SetEnv(`.noenv`)
	assert.Equal(t, `rpc.macarena.finance`, RPC_ENDPOINTS[1])
	assert.Equal(t, `rpc.flamenco.finance`, RPC_ENDPOINTS[250])
	assert.Equal(t, `rpc.rock.finance`, RPC_ENDPOINTS[42161])
}

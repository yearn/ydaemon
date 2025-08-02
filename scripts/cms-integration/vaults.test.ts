import { expect, test } from './test'
import { YDaemonVaultResponseSchema } from './types'

async function fetchVault(host: string, chainId: number, address: string) {
  const response = await fetch(`${host}/747474/vault/${address}`)
  const vault = await response.json()
  return vault
}

test('compare local with live', async () => {
  const liveJson = await fetchVault('https://ydaemon.yearn.fi', 747474, '0x78EC25FBa1bAf6b7dc097Ebb8115A390A2a4Ee12')
  const localJson = await fetchVault('http://localhost:3001', 747474, '0x78EC25FBa1bAf6b7dc097Ebb8115A390A2a4Ee12')
  const local = YDaemonVaultResponseSchema.parse(localJson)
  const live = YDaemonVaultResponseSchema.parse(liveJson)

  expect(local).toAlmostEqual(live, { tolerance: { 
    'tvl/tvl': 10_000 * live.tvl.tvl / local.tvl.tvl,
    'tvl/price': 1e-4,
    'apr/netAPR': 1e-4,
    'apr/forwardAPR/composite/boost': 1e-4,
    'apr/forwardAPR/composite/poolAPY': 1e-4,
    'apr/forwardAPR/composite/boostedAPR': 1e-4,
    'apr/points/weekAgo': 6e-4,
    'apr/points/monthAgo': 1e-4,
    'apr/points/inception': 3e-4,
    'apr/pricePerShare/weekAgo': 2e-5,
    'apr/pricePerShare/monthAgo': 1e-5,
    'featuringScore': 1000
  } })
})

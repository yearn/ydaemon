import { expect, test } from './test'
import { YDaemonVaultResponseSchema } from './types'

async function fetchVaults(host: string, chainId: number) {
  const response = await fetch(`${host}/${chainId}/vaults/all`)
  const json = await response.json()
  return [...(json as any[]).map((vault: any) => vault.address)]
}

async function fetchVault(host: string, chainId: number, address: string) {
  const response = await fetch(`${host}/${chainId}/vault/${address}?strategiesCondition=inQueue`)
  const json = await response.json()
  return YDaemonVaultResponseSchema.parse(json)
}

test('compare local vs live vaults', async () => {
  const chainId = 747474
  const vaults = await fetchVaults('https://ydaemon.yearn.fi', chainId)

  for (const vault of vaults) {
    const live = await fetchVault('https://ydaemon.yearn.fi', chainId, vault)
    const local = await fetchVault('http://localhost:3001', chainId, vault)

    expect(local).toAlmostEqual(live, { tolerance: {
      'tvl/tvl': .15,
      'tvl/totalAssets': .15,
      'tvl/price': .05,
      'pricePerShare': .05,
      'apr/netAPR': .15,
      'apr/fees/performance': 1,
      'apr/forwardAPR/composite/boost': .05,
      'apr/forwardAPR/composite/poolAPY': .05,
      'apr/forwardAPR/composite/boostedAPR': .05,
      'apr/points/weekAgo': .75,
      'apr/points/monthAgo': .15,
      'apr/points/inception': .85,
      'apr/pricePerShare/today': .05,
      'apr/pricePerShare/weekAgo': .05,
      'apr/pricePerShare/monthAgo': .05,
      'featuringScore': .25,
      'info/riskLevel': 1,
    }, message: `chainId: ${chainId} vault: ${vault}` })
  }
}, { timeout: 5_000 })

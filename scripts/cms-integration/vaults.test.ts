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

  expect(local.chainID).toBe(live.chainID)
  expect(local.address).toBe(live.address)
  expect(local.type).toBe(live.type)
  expect(local.kind).toBe(live.kind)
  expect(local.symbol).toBe(live.symbol)
  expect(local.name).toBe(live.name)
  expect(local.category).toBe(live.category)
  expect(local.version).toBe(live.version)
  expect(local.decimals).toBe(live.decimals)
  expect(local.token.address).toBe(live.token.address)
  expect(local.token.name).toBe(live.token.name)
  expect(local.token.symbol).toBe(live.token.symbol)
  expect(local.token.description).toBe(live.token.description)
  expect(local.token.decimals).toBe(live.token.decimals)
  expect(local.tvl.totalAssets).toBe(live.tvl.totalAssets)
  expect(local.tvl.tvl).toAlmostBe(live.tvl.tvl, 1000)
  expect(local.tvl.price).toAlmostBe(live.tvl.price, 1e-4)
  // expect(local.apr.type).toBe(live.apr.type)
  // expect(local.apr.netAPR).toBe(live.apr.netAPR)
  // expect(local.apr.fees.performance).toBe(live.apr.fees.performance)
  // expect(local.apr.fees.management).toBe(live.apr.fees.management)
  // expect(local.apr.forwardAPR.type).toBe(live.apr.forwardAPR.type)
  // expect(local.apr.forwardAPR.netAPR).toBe(live.apr.forwardAPR.netAPR)
  // expect(local.apr.forwardAPR.composite.boost).toBe(live.apr.forwardAPR.composite.boost)
  // expect(local.apr.forwardAPR.composite.poolAPY).toBe(live.apr.forwardAPR.composite.poolAPY)
  // expect(local.apr.forwardAPR.composite.boostedAPR).toBe(live.apr.forwardAPR.composite.boostedAPR)
  // expect(local.apr.forwardAPR.composite.baseAPR).toBe(live.apr.forwardAPR.composite.baseAPR)
  // expect(local.apr.forwardAPR.composite.cvxAPR).toBe(live.apr.forwardAPR.composite.cvxAPR)
  // expect(local.apr.forwardAPR.composite.rewardsAPR).toBe(live.apr.forwardAPR.composite.rewardsAPR)
  // expect(local.apr.extra.stakingRewardsAPR).toBe(live.apr.extra.stakingRewardsAPR)
  // expect(local.apr.extra.gammaRewardAPR).toBe(live.apr.extra.gammaRewardAPR)
})

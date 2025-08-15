import { getAddress } from 'viem'
import { z } from 'zod'

export const zhexstring = z.custom<`0x${string}`>((val: unknown) => 
  typeof val === 'string' && /^0x[a-fA-F0-9]*$/.test(val)
)

export const HexStringSchema = zhexstring.transform(s => s)
export type HexString = z.infer<typeof HexStringSchema>

export const zevmaddressstring = z.custom<`0x${string}`>((val: unknown) => 
  typeof val === 'string' && /^0x[a-fA-F0-9]{40}$/.test(val)
)

export const EvmAddressSchema = zevmaddressstring.transform(s => getAddress(s))
export type EvmAddress = z.infer<typeof EvmAddressSchema>

z.ZodString.prototype.optionalString = function () {
  return this.nullish().or(z.literal(''))
}

const optionalString = (schema: z.ZodType<string>) => schema.nullish().or(z.literal(''))

export const YDaemonVaultResponseSchema = z.object({
  chainID: z.number(),
  address: EvmAddressSchema,
  type: z.string(),
  kind: z.string(),
  symbol: z.string(),
  name: z.string(),
  category: z.string(),
  version: z.string(),
  decimals: z.number(),
  token: z.object({
    address: EvmAddressSchema,
    name: z.string(),
    symbol: z.string(),
    description: z.string(),
    decimals: z.number(),
  }),
  tvl: z.object({
    totalAssets: z.bigint({ coerce: true }),
    tvl: z.number(),
    price: z.number(),
  }),
  apr: z.object({
    type: z.string(),
    netAPR: z.number(),
    fees: z.object({
      performance: z.number(),
      management: z.number(),
    }),
    points: z.object({
      weekAgo: z.number(),
      monthAgo: z.number(),
      inception: z.number(),
    }),
    pricePerShare: z.object({
      today: z.number(),
      weekAgo: z.number(),
      monthAgo: z.number(),
    }),
    extra: z.object({
      stakingRewardsAPR: z.number().nullable(),
      gammaRewardAPR: z.number().nullable(),
    }),
    forwardAPR: z.object({
      type: z.string(),
      netAPR: z.number().nullable(),
      composite: z.object({
        boost: z.number().nullable(),
        poolAPY: z.number().nullable(),
        boostedAPR: z.number().nullable(),
        baseAPR: z.number().nullable(),
        cvxAPR: z.number().nullable(),
        rewardsAPR: z.number().nullable(),
      }),
    }),
  }),
  // strategies: z.array(z.object({
  //   address: EvmAddressSchema,
  //   name: z.string()
  //   ...
  // })),
  staking: z.object({
    address: optionalString(EvmAddressSchema),
    available: z.boolean(),
    source: z.string(),
    rewards: z.number().nullable(),
  }),
  migration: z.object({
    available: z.boolean(),
    address: EvmAddressSchema,
    contract: EvmAddressSchema,
  }),
  featuringScore: z.number(),
  pricePerShare: z.bigint({ coerce: true }),
  info: z.object({
    riskLevel: z.number(),
    isRetired: z.boolean(),
    isBoosted: z.boolean(),
    isHighlighted: z.boolean(),
    riskScore: z.array(z.number()),
  }),
})

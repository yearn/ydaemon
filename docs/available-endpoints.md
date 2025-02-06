# yDaemon Vaults API Documentation

## Vaults

These vaults all share the same query parameters:

| Parameter | Type | Default | Description |
|-----------|------|---------|-------------|
| `orderBy` | string | 'featuringScore' | Determines the order of returned vaults. |
| `orderDirection` | string | 'asc' | Determines the direction of ordering ('asc' or 'desc'). |
| `strategiesCondition` | string | 'debtRatio' | Filters strategies based on specified condition ('inQueue', 'debtLimit', 'debtRatio', 'absolute', 'all') |
| `hideAlways` | boolean | false | If true, hides certain vaults. |
| `migrable` | string | 'none' | Filters vaults based on migration status ('none', 'all', 'nodust', 'ignore') |
| `page` | integer | 1 | Page number for pagination. |
| `limit` | integer | 200 | Number of vaults per page. |
| `chainIDs` | string | - | Comma-separated list of chain IDs to filter vaults. |

---

#### **GET** `/vaults`
Returns all vaults matching the `inclusion.IsYearn` filter.

#### **GET** `/vaults/detected`
Returns all detected vaults.

#### **GET** `/vaults/v2`
Returns all V2 vaults matching the `inclusion.IsYearn` filter.
Filters vaults where the major version number is not `3`.

#### **GET** `/vaults/v3`
Returns all V3 vaults matching the `inclusion.IsYearn` filter.
Filters vaults where the kind is `Multiple` or `Single`, or the major version number is `3`.

#### **GET** `/vaults/juiced`
Returns all vaults matching the `inclusion.IsYearnJuiced` filter.

#### **GET** `/vaults/gimme`
Returns all vaults matching the `inclusion.IsGimme` filter.

#### **GET** `/vaults/retired`
Returns all retired vaults.
Filters vaults where the `Metadata.IsRetired` flag is true.

#### **GET** `/vaults/pendle`
Returns all vaults matching the Pendle category and the `inclusion.IsYearn` filter.

#### **GET** `/vaults/optimism`
Returns all Optimism vaults matching the `inclusion.IsYearn` filter.

#### **GET** `/vaults/pooltogether`
Returns all vaults matching the `inclusion.isPoolTogether` filter.

#### **GET** `/vaults/morpho`
Returns all vaults matching the `inclusion.isMorpho` filter.

#### **GET** `/vaults/ajna`
Returns all vaults with the Ajna category and the `inclusion.IsYearn` filter.

#### **GET** `/vaults/velodrome`
Returns all vaults with the Velodrome category and the `inclusion.IsYearn` filter.

#### **GET** `/vaults/aerodrome`
Returns all vaults with the Aerodrome category and the `inclusion.IsYearn` filter.

#### **GET** `/vaults/curve`
Returns all vaults with the Curve category and the `inclusion.IsYearn` filter.

Note: All endpoints apply additional filtering based on blacklisted vaults, vault visibility, retirement status, and migration availability depending on the query parameters provided.

# Risk Score Processing

This document explains how Yearn's yDaemon ingests risk scores from the security team, caches them, and exposes the data through the public vault endpoints. It complements `docs/risk-score.md`, which describes what each score represents.

## High-Level Flow

1. **Manifest fetch** – every 30 minutes the `SNAPSHOT30M` scheduler job calls `risks.RetrieveAvailableRiskScores`, which attempts to load the aggregated chain manifest (`vaults/<chainID>.json`) from the risk CDN and hydrate both the availability map and cache in one request (`internal/main.go:139-158`, `processes/risks/main.go:128-187`).
2. **Targeted ingestion** – `risks.RetrieveAllRiskScores` reuses any manifest-populated entries and only calls the per-vault CDN endpoint for vaults still missing in the cache (`internal/main.go:153-158`, `processes/risks/main.go:88-133`).
3. **Serving** – all vault endpoints (`external/vaults/...`) read from the in-memory cache via `risks.GetCachedRiskScore` when constructing `TExternalVault.Info` or the simplified vault representations (`external/vaults/models.go:537-546`, `external/vaults/prepare.vaultObject.go:192-210`).

The result is a tight loop that refreshes Yearn's public API within ~30 minutes anytime the security team republishes scores.

## Source of Truth

- **CDN base URL** – `env.RISK_CDN_URL` defaults to `https://risk.yearn.fi/cdn/` but can be overridden via the `RISK_CDN_URL` environment variable (`common/env/constants.go:84-90`, `common/env/environment.go:46-58`). The CDN serves both the chain manifest (`vaults/<chainID>.json`) and the on-demand per-vault files (`vaults/<chainID>/<vaultAddress>.json`).
- **Repository index** – if the manifest fetch fails, the code falls back to the GitHub tree at `https://api.github.com/repos/yearn/risk-score/git/trees/master?recursive=1`, checking both legacy `strategy/<chainID>/` and `vaults/<chainID>/` folders for individual files (`processes/risks/main.go:156-179`).

## Availability Discovery (`RetrieveAvailableRiskScores`)

Located in `processes/risks/main.go:128-187`, this function:

- Attempts to download the chain manifest and, if successful, marks every address in the map as available while seeding `allRisksScores[chainID]` with the parsed `TRiskScoreYsec` structs.
- Falls back to the GitHub tree scan only when the manifest cannot be fetched, preserving legacy behavior for migrations or outages.

## Score Ingestion (`RetrieveAllRiskScores`)

Also in `processes/risks/main.go:35-133`:

- Entries already present in `allRisksScores[chainID]` (typically filled by the manifest) are returned immediately.
- For vaults missing from the cache but flagged as available, the code still hits the per-vault CDN endpoint, trying lowercase then checksummed addresses (`processes/risks/main.go:52-74`).
- Successful fetches populate both the local return map and the shared cache for future callers (`processes/risks/main.go:112-116`).

Individual fetch failures are logged but do not abort the loop, so one missing file does not block the rest of the chain.

## Cache Access (`GetCachedRiskScore`)

`risks.GetCachedRiskScore` is a thin helper that reads from `allRisksScores` using an `RLock`. API handlers rely on this instead of reaching out to the CDN directly, which keeps responses fast and immune to network hiccups (`processes/risks/main.go:172-182`).

- On cache miss it returns an error, letting callers fall back to baked-in metadata (e.g., Kong-provided risk data) when available.

## API Integration

Key touchpoints in `external/vaults`:

- `CreateExternalVault` calls `GetCachedRiskScore` and, when successful, copies `RiskLevel`, the 11-element risk array (via `PopulateRiskScoreArray`), and any comment onto `TExternalVault.Info`. If the cache misses, the function falls back to the vault's metadata-sourced score (`external/vaults/models.go:453-546`).
- `toSimplifiedVersion` repeats the lookup so the lightweight responses served by `/vaults`, `/vaults/v2`, `/vaults/one`, etc., always reflect the most recent cache (`external/vaults/prepare.vaultObject.go:150-237`).
- Bulk endpoints build each simplified vault via `CreateExternalVault` → `toSimplifiedVersion`, so the cached scores propagate automatically without endpoint-specific logic (`external/vaults/prepare.getVaults.go:167-210`, `external/vaults/route.vaults.one.go:82-133`).

## Operational Notes

- **Refresh cadence** – The manifest is reloaded every 30 minutes; manually invoking `RetrieveAvailableRiskScores` + `RetrieveAllRiskScores` for a chain forces an immediate refresh (`internal/main.go:139-158`).
- **Manifest failures** – When the manifest cannot be fetched (network or CDN issue), the code logs a warning and falls back to the slower GitHub tree method, so risk data remains available albeit with per-vault fetching overhead (`processes/risks/main.go:156-179`).
- **Adding new chains** – As soon as a `vaults/<chainID>.json` manifest is published, yDaemon will discover the vaults on that chain automatically; no code changes are required.
- **Troubleshooting** – If a single vault shows zeros while others on the same chain have scores, inspect the logs for `fetchVaultsRiskScore` errors; most often this means the CDN has not generated the per-vault JSON yet even though it appears in the manifest.

With this pipeline in mind, updates to the risk-score repository or CDN are automatically reflected in yDaemon without code changes, so long as the JSON schema matches `TRiskScoreYsec`.

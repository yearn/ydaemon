# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

yDaemon is the next-generation API for Yearn Finance, providing comprehensive data about vaults, strategies, tokens, and APY calculations across multiple blockchain networks (Ethereum, Arbitrum, Optimism, Polygon, Fantom, Base). It serves as the backend API for Yearn's ecosystem, replacing the legacy yearn-exporter.

## Development Commands

### Building and Running
```bash
# Build the project
make build
# or manually:
go build -o yDaemon ./cmd

# Run the daemon
./yDaemon

# Run with Docker
make deploy

# Stop Docker services
make down
```

### Code Quality
```bash
# Format code
make fmt
# or: go fmt ./...

# Lint code
make lint
# or: golint ./...

# Vet code
make vet  
# or: go vet ./...
```

### Testing
```bash
# Run all tests
go test ./...

# Run tests with coverage
go test -v -cover ./...

# Run specific package tests
go test ./external/vaults/
go test ./internal/storage/
```

### Environment Setup
Create a `.env` file with required RPC endpoints:
```
RPC_URI_FOR_1=      # Ethereum mainnet
RPC_URI_FOR_10=     # Optimism
RPC_URI_FOR_137=    # Polygon
RPC_URI_FOR_250=    # Fantom
RPC_URI_FOR_8453=   # Base
RPC_URI_FOR_42161=  # Arbitrum

# Optional
WEBHOOK_SECRET=
GRAPH_API_URI=
SENTRY_DSN=
SENTRY_SAMPLE_RATE=
LOG_LEVEL=          # DEBUG, INFO, WARNING, SUCCESS, ERROR
```

## Architecture Overview

### Core Structure
- **cmd/**: Application entry point and server initialization
- **internal/**: Private application code containing the main business logic
- **external/**: Public API routes and handlers for REST endpoints
- **common/**: Shared utilities, contracts, and helper functions
- **processes/**: Background processes for APR calculation, price fetching, and risk assessment
- **data/**: Static data files including metadata, block times, and configuration

### Key Components

#### Internal Packages
- **fetcher/**: Retrieves and processes data from blockchain and external sources
- **indexer/**: Indexes vaults, strategies, tokens, and staking contracts from registries
- **models/**: Shared data structures and types
- **storage/**: Database layer and caching (using GORM with MySQL/PostgreSQL)
- **multicalls/**: Efficient batch blockchain calls using multicall contracts

#### External API Routes
- **vaults/**: Vault data endpoints (`/[chainID]/vaults/all`, `/[chainID]/vaults/[address]`)
- **strategies/**: Strategy information endpoints
- **prices/**: Token and vault pricing data
- **tokens/**: ERC20 token metadata
- **utils/**: Utility endpoints (chains, blacklisted vaults, TVL)

#### Background Processes
- **apr/**: APR/APY calculations for current, forward, and historical yields
- **prices/**: Price fetching from multiple sources (Lens Oracle, CoinGecko, DeFiLlama)
- **risks/**: Risk score calculation and assessment

### Data Flow
1. **Initialization**: Load vaults from registries, fetch strategies and tokens
2. **Scheduling**: Background processes run on timed intervals (prices every 30s, APY every 10min, metadata every 30min)
3. **Caching**: Results stored in database and in-memory cache for fast API responses
4. **API**: REST endpoints serve cached data with real-time updates

### Supported Networks
- Ethereum (1), Optimism (10), Polygon (137), Fantom (250), Base (8453), Arbitrum (42161)
- Each chain runs independent processes for data fetching and calculations

### Key Dependencies
- **Ethereum**: go-ethereum for blockchain interactions
- **Web Framework**: Gin for REST API
- **Database**: GORM with MySQL/PostgreSQL drivers
- **Scheduling**: gocron for background job management
- **Testing**: testify for unit tests

### Development Notes
- Smart contract interactions use generated Go bindings in `common/contracts/`
- All addresses must be checksummed for consistency
- Price data fetched from multiple sources with fallbacks
- Risk scores calculated using 11-factor assessment system
- APY calculations support multiple strategies: Curve, Convex, Velodrome, Aerodrome, etc.
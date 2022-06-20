# yDaemon
yDaemon is the next-gen API for Yearn. Based on the one from the [exporter](https://github.com/yearn/yearn-exporter), it brings a lot of new features and benefits without breaking the existing system.

## Install
First, ensure [Go](https://go.dev/) is installed on your system. then, clone the repo and create the `.env` file:
```
RPC_URI_FOR_1=
RPC_URI_FOR_250=
RPC_URI_FOR_42161=
```

Then, install, build and run the API:
```bash
go mod vendor
go build -o yDaemon ./cmd
./yDaemon
```

After a few seconds, you should see the API running. You can test it by running the following command:
```bash
curl http://localhost:8080/1/vaults/all
```

## Endpoints
| ✅   `[BASE_URL]/[chainID]/vaults/all`   |
|:------------------------------------------|
This endpoint returns all the vaults of the specified chainID. Supported chains are `1`, `250` and `42161`.


## Data Sources
In order to build this API, data are fetched from a number of Yearn data sources:
- [Yearn Subgraph](https://thegraph.com/explorer/subgraph?id=5xMSe3wTNLgFQqsAc5SCVVwT4MiRb5AogJCuSN9PjzXF) as the base data source.
- [Yearn Meta](https://github.com/yearn/yearn-meta) for some basic data and information updated by the Yearn team.
- [Yearn API](https://api.yearn.finance/) for the APY computation.
- [Yearn Lens Oracle](https://etherscan.io/address/0xca11bde05977b3631167028862be2a173976ca11) for tokens and vault prices.

In order to provide a fast and up-to-date experience, a bunch of daemons are summoned with the API, running in the background, forever and ever.
- Prices from the oracle are updated every 30 seconds for every tokens and vaults, as the price may change at every block.
- APY information is updated every 10 minutes, as the underlying API is updated every 30 minutes
- Meta data are updated every minutes. This will be moved to every 30 minutes in the future, and trust a webhook from the github deployement system to update the data.

## Go Test and Coverage
```
?       github.com/majorfi/ydaemon/cmd  [no test files]
?       github.com/majorfi/ydaemon/internal/contracts   [no test files]
ok      github.com/majorfi/ydaemon/internal/controllers 34.167s coverage: 98.6% of statements
ok      github.com/majorfi/ydaemon/internal/daemons     31.366s coverage: 94.7% of statements
ok      github.com/majorfi/ydaemon/internal/ethereum    30.505s coverage: 78.3% of statements
ok      github.com/majorfi/ydaemon/internal/logs        0.674s  coverage: 100.0% of statements
ok      github.com/majorfi/ydaemon/internal/models      0.521s  coverage: [no statements]
ok      github.com/majorfi/ydaemon/internal/store       0.213s  coverage: [no statements]
ok      github.com/majorfi/ydaemon/internal/utils       0.406s  coverage: 100.0% of statements
```

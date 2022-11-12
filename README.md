# yDaemon

![](./.github/banner.png)

yDaemon is the next-gen API for Yearn. Based on the API from [yearn-exporter](https://github.com/yearn/yearn-exporter), it brings a lot of new features and benefits without breaking the existing system.

See documentation here: https://ydaemon.ycorpo.com/
## Use with Docker
First, ensure [Docker](https://docs.docker.com/get-started/overview/) is installed on your system then, clone the repo and create the '.env' file:
```
RPC_URI_FOR_1=
RPC_URI_FOR_10=
RPC_URI_FOR_250=
RPC_URI_FOR_42161=
```
Then to run, type:
```
make depoly
```
To stop, type:
```
make down
```

## Manual Install
First, ensure [Go](https://go.dev/) is installed on your system. then, clone the repo and create the `.env` file:
```
RPC_URI_FOR_1=
RPC_URI_FOR_10=
RPC_URI_FOR_250=
RPC_URI_FOR_42161=

# Optional
WEBHOOK_SECRET=
GRAPH_API_URI=
SENTRY_DSN=
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
`GET` `[BASE_URL]/[chainID]/vaults/all`  
> This endpoint returns all the vaults of the specified chainID. Supported chains are `1`, `250` and `42161`.  
>  
> **Query**  
`?skip=N` will skip N vault from the graphQL query. Default is `0`  
`?first=N` will limit the result to N vaults on the graphQL query. Default is `1000`  
`?orderBy=S` will order the result by S on the graphQL query. Default is `activation`  
`?orderDirection=asc|desc` will order the result by ascending or descending on the graphQL query. Default is `desc`  
>`?strategiesCondition=debtLimit|inQueue|absolute` will select the "active" strategies based on the specified strategy. Default is `debtLimit`  
>`?strategiesDetails=withDetails|noDetails` indicates if we should also query and serve the details about the strategies. If noDetails is set, the Details field will be ignored. Default is noDetails.  
-------

`GET` `[BASE_URL]/[chainID]/vaults/[address]`  
> This endpoint returns the vault matching the specified address, for the specified chainID. Supported chains are `1`, `10`, `250`, and `42161`.  
>  
> **Query**  
> `?strategiesCondition=debtLimit|inQueue|absolute` will select the "active" strategies based on the specified strategy. Default is `debtLimit`  
>`?strategiesDetails=withDetails|noDetails` indicates if we should also query and serve the details about the strategies. If noDetails is set, the Details field will be ignored. Default is noDetails.  
-------

`GET` `[BASE_URL]/info/chains`  
> This endpoint returns the supported chains for this API.  

-------

`GET` `[BASE_URL]/info/vaults/blacklisted`  
> This endpoint returns the blacklisted vaults for all chains. A blacklisted vault is a vault that will be ignored by the API.  

-------
`GET` `[BASE_URL]/[chainID]/vaults/tvl`  
> This endpoint returns the Total Value Locked for the specified chainID. Does not subtract delegated deposits from one vault to another.  

## Data Sources
To build this API data is fetched from several Yearn data sources:
- [Yearn Subgraph](https://thegraph.com/explorer/subgraph?id=5xMSe3wTNLgFQqsAc5SCVVwT4MiRb5AogJCuSN9PjzXF) as the base data source.
- [Yearn Meta](https://github.com/yearn/yearn-meta) for some basic data and information updated by the Yearn team.
- [Yearn API](https://api.yearn.finance/) for the APY computation.
- [Yearn Lens Oracle](https://etherscan.io/address/0xca11bde05977b3631167028862be2a173976ca11) for tokens and vault prices.

To provide a fast and up-to-date experience, a bunch of daemons are summoned with the API, running in the background, forever and ever.
- Prices from the oracle are updated every 30 seconds for every tokens and vaults, as the price may change at every block.
- APY information is updated every 10 minutes, as the underlying API is updated every 30 minutes
- Meta data is updated every minute. This will be moved to every 30 minutes in the future, and trust a webhook from the github deployement system to update the data.

## Folder and structure
The project is divided as follow:
- `cmd`: contains the `main.go` entry point for this API. It's role is _only_ to init the project.
- `contracts`: contains the `.sol` and `.vy` files used to generate the contract `.go` files with `abigen`
- `internal`: contains the actual code of the API. It's divided in subpackages, each one representing a different part of the API, keeping the code clean and easy to maintain.

The `internal` folder is the _private application and library code_. This is the code we don't want others importing in their applications or libraries. This layout pattern is enforced by the Go compiler itself. Inside of it, we have the following subpackages:
- `contracts`: contains the `abigen` generated contract `.go` files. Theses files should not be altered.
- `controllers`: contains the logic around the `REST` API: routes, server, actions by route, etc. This is where the entry-point and the wrapper around the logic for a specific `REST` endpoint is.
- `daemons`: contains all our background processes. They will run forever in the background to update the data we may need in our application.
- `ethereum`: contains all the logic aroung the Blockchain calls. It's the place where we call the RPC, connect to the client, call the methods, etc.
- `logs`: helper package used to display nice logs.
- `models`: contains all the cross-package types we need.
- `store`: contains all the cross-package global variables.
- `utils`: contains all the cross-package utility functions.

## Docs
To run docs locally use the follow:
```bash
cd docs
yarn
yarn dev
```
## Go Test and Coverage
```
?       github.com/yearn/ydaemon/cmd  [no test files]
?       github.com/yearn/ydaemon/internal/contracts   [no test files]
ok      github.com/yearn/ydaemon/internal/controllers 34.167s coverage: 98.6% of statements
ok      github.com/yearn/ydaemon/internal/daemons     31.366s coverage: 94.7% of statements
ok      github.com/yearn/ydaemon/internal/utils/ethereum    30.505s coverage: 78.3% of statements
ok      github.com/yearn/ydaemon/internal/logs        0.674s  coverage: 100.0% of statements
ok      github.com/yearn/ydaemon/internal/models      0.521s  coverage: [no statements]
ok      github.com/yearn/ydaemon/internal/utils/store       0.213s  coverage: [no statements]
ok      github.com/yearn/ydaemon/internal/utils       0.406s  coverage: 100.0% of statements
```

package main

import (
	"context"
	"strconv"
	"sync"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/yearn/ydaemon/common/logs"
	"github.com/yearn/ydaemon/common/traces"
)

// var chains = env.SUPPORTED_CHAIN_IDS

var chains = []uint64{1}

func waitGroupSummonDaemons(trace *traces.TTrace, wg *sync.WaitGroup, chainID uint64) {
	trace = trace.Child(
		`app.bootstrap.summon.daemon`,
		traces.TTags{Name: `chainID`, Value: strconv.Itoa(int(chainID))},
	)
	defer trace.Finish()

	SummonDaemons(chainID)
	wg.Done()
	logs.Success(`Daemons for chainID ` + strconv.Itoa(int(chainID)) + ` summoned successfully!`)
}

func summonDaemonsForAllChains(trace *traces.TTrace) {
	trace = trace.Child(`app.bootstrap.summon.all`)
	defer trace.Finish()

	var wg sync.WaitGroup
	for _, chainID := range chains {
		wg.Add(1)
		go waitGroupSummonDaemons(trace, &wg, chainID)
	}

	wg.Wait()
}

func waitGroupLoadDaemons(trace *traces.TTrace, wg *sync.WaitGroup, chainID uint64) {
	trace = trace.Child(
		`app.bootstrap.load_state.daemon`,
		traces.TTags{Name: `chainID`, Value: strconv.Itoa(int(chainID))},
	)
	defer trace.Finish()

	LoadDaemons(chainID)
	wg.Done()
	logs.Success(`Store data loaded in yDaemon memory for chainID ` + strconv.Itoa(int(chainID)) + `!`)
}

func loadDaemonsForAllChains(trace *traces.TTrace) {
	trace = trace.Child(`app.bootstrap.load_state.all`)
	defer trace.Finish()

	var wg sync.WaitGroup
	for _, chainID := range chains {
		wg.Add(1)
		go waitGroupLoadDaemons(trace, &wg, chainID)
	}
	wg.Wait()
}

func initialize() {
	traces.SetupSentry()
	trace := traces.Init(`app.bootstrap`).SetTag(`subsystem`, `main`)
	defer trace.Finish()

	logs.Info(`Loading store data to yDaemon memory...`)
	loadDaemonsForAllChains(trace)

	logs.Info(`Summoning yDaemon...`)
	summonDaemonsForAllChains(trace)
}

func main() {
	mai2()
	// go initialize()

	// logs.Success(`Server ready!`)
	// NewRouter().Run()
}

func mai2() {
	client, _ := ethclient.Dial("https://eth-node0.yearn.network")
	rpcclient, _ := rpc.Dial("https://eth-node0.yearn.network")

	block, _ := client.BlockByNumber(context.Background(), nil)
	n := block.Number().Uint64()

	reqs := []rpc.BatchElem{}
	// blocks := []*types.Block{}
	// errorList := []error{}
	for i := uint64(16664093); i <= n; i = i + 100 {
		for j := uint64(1); j <= uint64(100); j++ {
			b := types.Block{}
			var errorr error
			newBarchElemWithError := rpc.BatchElem{
				Method: "eth_getBlockByNumber",
				//args is I + J
				Args:   []interface{}{"0x1b4", true},
				Result: &b,
				Error:  errorr,
				// Error:  &errorList,
			}
			reqs = append(reqs, newBarchElemWithError)
			// blocks = append(blocks, &b)
			// errorList = append(errorList, errorr)
		}
		if err := rpcclient.BatchCallContext(context.Background(), reqs); err != nil {
			logs.Pretty("Error: ", err)
		}
		logs.Pretty(reqs)
		// logs.Pretty("Block len: ", len(blocks))
		// logs.Pretty(blocks)
		// logs.Pretty("Error len: ", len(errorList))
		// logs.Pretty(errorList)
		return
		// if len(blocks) != 0 {
		// 	logs.Pretty(blocks[len(blocks)-1])
		// }

	}
}

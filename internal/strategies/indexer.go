package strategies

func processNewStrategies(chainID uint64, strategiesAddedList []TStrategyAdded) {
	RetrieveAllStrategies(chainID, strategiesAddedList)
}

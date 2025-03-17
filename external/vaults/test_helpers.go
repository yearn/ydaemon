package vaults

/**************************************************************************************************
** This file contains helper functions used specifically for testing purposes.
** These functions should not be used in production code.
**
** When we run tests, we'll use a simplified version of handleError that maintains
** the same error format as before to avoid having to update all test assertions.
**************************************************************************************************/

/**************************************************************************************************
** TestMode controls the behavior of API functions during tests.
**
** When set to true (which happens automatically during test execution):
** 1. Error handlers will use simplified formatting without additional context or user messages
** 2. Error responses will match the format expected by the test assertions
** 3. JSON responses will include only essential fields needed for test verification
**
** This allows tests to verify raw error messages without being affected by formatting changes
** in the production error handling code. The TestMode flag is activated in the test_init.go file
** when running tests, and should remain false in production.
**************************************************************************************************/
var TestMode = false

/**************************************************************************************************
** selectStrategiesCondition is a test helper function to simulate validateStrategyCondition
** without needing the Gin context.
**
** @param input string - The input strategy condition
** @return string - The validated strategy condition
**************************************************************************************************/
func selectStrategiesCondition(input string) string {
	if input == "" {
		return STRATEGY_CONDITION_DEBT_RATIO
	}

	validConditions := map[string]bool{
		STRATEGY_CONDITION_ALL:        true,
		STRATEGY_CONDITION_IN_QUEUE:   true,
		STRATEGY_CONDITION_DEBT_LIMIT: true,
		STRATEGY_CONDITION_DEBT_RATIO: true,
		STRATEGY_CONDITION_ABSOLUTE:   true,
	}

	if validConditions[input] {
		return input
	}

	return STRATEGY_CONDITION_DEBT_RATIO
}

/**************************************************************************************************
** selectMigrableCondition is a test helper function to simulate validateMigrableCondition
** without needing the Gin context.
**
** @param input string - The input migrable condition
** @return string - The validated migrable condition
**************************************************************************************************/
func selectMigrableCondition(input string) string {
	if input == "" {
		return MIGRABLE_CONDITION_NONE
	}

	validConditions := map[string]bool{
		MIGRABLE_CONDITION_NONE:    true,
		MIGRABLE_CONDITION_ALL:     true,
		MIGRABLE_CONDITION_NO_DUST: true,
		MIGRABLE_CONDITION_IGNORE:  true,
	}

	if validConditions[input] {
		return input
	}

	return MIGRABLE_CONDITION_NONE
}

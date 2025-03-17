package ethereum

import (
	"fmt"
	"path/filepath"
	"runtime"

	"github.com/yearn/ydaemon/common/logs"
)

/**************************************************************************************************
** VerboseBlocktime controls whether detailed debug logs are shown for blocktime operations.
** When enabled, all debug log messages are displayed to provide detailed insights into the
** blocktime data processing, API calls, and file operations.
**
** This flag can be controlled via the VERBOSE_BLOCKTIME environment variable or programmatically
** using the EnableVerboseBlocktime and DisableVerboseBlocktime functions.
**************************************************************************************************/
var VerboseBlocktime bool = false

/**************************************************************************************************
** EnableVerboseBlocktime enables detailed debug logging for blocktime operations.
** When enabled, the system will output extensive logging information about blocktime-related
** operations, which is useful for debugging and troubleshooting.
**
** This function sets the VerboseBlocktime flag to true and logs an informational message
** about the change in logging verbosity.
**************************************************************************************************/
func EnableVerboseBlocktime() {
	VerboseBlocktime = true
	logs.Info("Verbose blocktime logging enabled")
}

/**************************************************************************************************
** DisableVerboseBlocktime disables detailed debug logging for blocktime operations.
** When disabled, only essential log messages (warnings and errors) will be displayed,
** reducing the amount of log output during normal operation.
**
** This function sets the VerboseBlocktime flag to false and logs an informational message
** about the change in logging verbosity.
**************************************************************************************************/
func DisableVerboseBlocktime() {
	VerboseBlocktime = false
	logs.Info("Verbose blocktime logging disabled")
}

/**************************************************************************************************
** blocktimeLog logs a message only if verbose logging is enabled.
** This function is used for informational messages that are only relevant during debugging
** or detailed analysis of the system behavior.
**
** @param message The message to log if verbose logging is enabled
**************************************************************************************************/
func blocktimeLog(message string) {
	if VerboseBlocktime {
		logs.Info(message)
	}
}

/**************************************************************************************************
** blocktimeWarning logs a warning message. These messages indicate potential issues that don't
** prevent the system from functioning but may require attention.
**
** When verbose mode is disabled, warnings are still shown but are truncated if longer than 100
** characters to avoid cluttering logs with excessive information.
**
** @param message The warning message to log
**************************************************************************************************/
func blocktimeWarning(message string) {
	if VerboseBlocktime {
		logs.Warning(message)
	} else {
		// For warnings, we still want a minimal log even when not in verbose mode
		if len(message) > 100 {
			logs.Warning(message[:100] + "...")
		}
	}
}

/**************************************************************************************************
** blocktimeSuccess logs a success message only if verbose logging is enabled.
** This function is used to log successful operations that may be helpful during debugging
** but are not necessary during normal operation.
**
** @param message The success message to log if verbose logging is enabled
**************************************************************************************************/
func blocktimeSuccess(message string) {
	if VerboseBlocktime {
		logs.Success(message)
	}
}

/**************************************************************************************************
** blocktimeError logs an error message. These messages indicate critical issues that may affect
** system operation and require immediate attention.
**
** Error messages are always logged regardless of the verbose flag setting, as they represent
** critical information that should never be hidden.
**
** @param message The error message to log
**************************************************************************************************/
func blocktimeError(message string) {
	logs.Error(message)
}

/**************************************************************************************************
** getCallerInfo returns information about the caller of a function to help with debugging.
** It captures the file name, line number, and function name of the caller.
**
** This function uses runtime.Caller with a skip value of 2 to identify the actual caller
** rather than the immediate caller of getCallerInfo.
**
** @return string A string with caller information, or empty string if information can't be retrieved
**************************************************************************************************/
func getCallerInfo() string {
	// Skip 2 frames to get the actual caller (1 would be the caller of getCallerInfo)
	pc, file, line, ok := runtime.Caller(2)
	if !ok {
		return ""
	}

	// Get the function name
	fn := runtime.FuncForPC(pc)
	if fn == nil {
		return fmt.Sprintf("%s:%d", filepath.Base(file), line)
	}

	// Format the caller information
	return fmt.Sprintf("%s:%d %s", filepath.Base(file), line, filepath.Base(fn.Name()))
}

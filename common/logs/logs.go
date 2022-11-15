package logs

import (
	"fmt"
	"os"
	"runtime"
	"strconv"
	"time"

	"github.com/davecgh/go-spew/spew"
	"github.com/fatih/color"
)

var levels = map[string]int{
	"DEBUG":   4,
	"INFO":    3,
	"WARNING": 2,
	"SUCCESS": 1,
	"ERROR":   0,
}

func isLogLevelAtLeast(minimum string) bool {
	LEVEL, exists := os.LookupEnv("LOG_LEVEL")
	if exists {
		return levels[LEVEL] >= levels[minimum]
	}
	return true
}

func shouldUseUI() bool {
	WITH_UI, exists := os.LookupEnv("WITH_UI")
	if !exists {
		return false
	}
	return WITH_UI == "true"
}

var colorGreen = color.New(color.FgGreen).Add(color.Bold).SprintFunc()
var colorRed = color.New(color.FgRed).Add(color.Bold).SprintFunc()
var colorYellow = color.New(color.FgYellow).Add(color.Bold).SprintFunc()
var colorBlue = color.New(color.FgBlue).Add(color.Bold).SprintFunc()
var colorCyan = color.New(color.FgCyan).SprintFunc()
var colorMagenta = color.New(color.FgMagenta).Add(color.Bold).SprintFunc()

// ErrorCrash function logs an error
func Error(err ...interface{}) {
	if shouldUseUI() {
		InfoLogs = append(InfoLogs, TLogs{
			level: "ERROR",
			data:  err,
		})
		return
	}
	pc, _, line, _ := runtime.Caller(1)

	str0 := `[` + strconv.Itoa(runtime.NumGoroutine()) + `]`
	str1 := `[ ERROR ]`
	str2 := `(` + runtime.FuncForPC(pc).Name() + `:` + strconv.Itoa(line) + `)`
	t := time.Now().Format("2006/01/02 15:04:05")

	if len(err) == 1 {
		spew.Config.Indent = "    "
		spew.Printf("%s %-17s %s %s %s\n", t, colorMagenta(str0), colorRed(str1), colorCyan(str2), colorRed(err[0]))
	} else {
		spew.Config.Indent = "    "
		fmt.Printf("%s", colorRed("----------------------------------\n"))
		spew.Printf("%s %-17s %s %s\n", t, colorMagenta(str0), colorRed(str1), colorCyan(str2))
		for _, each := range err {
			spew.Dump(each)
		}
		fmt.Printf("%s", colorRed("----------------------------------\n"))
	}
}

// Success function logs a success message
func Success(success ...interface{}) {
	if shouldUseUI() {
		InfoLogs = append(InfoLogs, TLogs{
			level: "SUCCESS",
			data:  success,
		})
		return
	}
	if !isLogLevelAtLeast("SUCCESS") {
		return
	}

	str0 := `[` + strconv.Itoa(runtime.NumGoroutine()) + `]`
	str1 := `[ OK ] `
	t := time.Now().Format("2006/01/02 15:04:05")

	spew.Printf("%s %-17s %s %s\n", t, colorMagenta(str0), colorGreen(str1), colorCyan(success))
}

// Warning function logs a warning message
func Warning(warning ...interface{}) {
	if shouldUseUI() {
		InfoLogs = append(InfoLogs, TLogs{
			level: "WARNING",
			data:  warning,
		})
		return
	}
	if !isLogLevelAtLeast("WARNING") {
		return
	}

	pc, _, line, _ := runtime.Caller(1)

	str0 := `[` + strconv.Itoa(runtime.NumGoroutine()) + `]`
	str1 := `[WARN] `
	str2 := `(` + runtime.FuncForPC(pc).Name() + `:` + strconv.Itoa(line) + `)`
	t := time.Now().Format("2006/01/02 15:04:05")

	spew.Printf("%s %-17s %s %s %s\n", t, colorMagenta(str0), colorYellow(str1), colorCyan(str2), colorYellow(warning))
}

// Info function logs an info message
func Info(info ...interface{}) {
	if shouldUseUI() {
		InfoLogs = append(InfoLogs, TLogs{
			level: "INFO",
			data:  info,
		})
		return
	}
	if !isLogLevelAtLeast("INFO") {
		return
	}

	str0 := `[` + strconv.Itoa(runtime.NumGoroutine()) + `]`
	str1 := `[INFO] `
	t := time.Now().Format("2006/01/02 15:04:05")

	spew.Printf("%s %-17s %s %s\n", t, colorMagenta(str0), colorBlue(str1), colorBlue(info))
}

// Debug function logs a debug message
func Debug(debug ...interface{}) {
	if shouldUseUI() {
		InfoLogs = append(InfoLogs, TLogs{
			level: "DEBUG",
			data:  debug,
		})
		return
	}
	if !isLogLevelAtLeast("DEBUG") {
		return
	}

	pc, _, line, _ := runtime.Caller(1)

	str0 := `[` + strconv.Itoa(runtime.NumGoroutine()) + `]`
	str1 := `[DEBUG]`
	str2 := `(` + runtime.FuncForPC(pc).Name() + `:` + strconv.Itoa(line) + `)`
	t := time.Now().Format("2006/01/02 15:04:05")

	spew.Printf("%s %-17s %s %s %s\n", t, colorMagenta(str0), colorBlue(str1), colorCyan(str2), colorBlue(debug))
}

func Trace(key string, status int, message string) {
	if shouldUseUI() {
		TraceLogs = append(TraceLogs, TTraces{
			key:     key,
			status:  status,
			message: message,
		})
		return
	}
	if !isLogLevelAtLeast("DEBUG") {
		return
	}

	str0 := `[` + strconv.Itoa(runtime.NumGoroutine()) + `]`
	str1 := `[TRACE]`
	t := time.Now().Format("2006/01/02 15:04:05")

	if status == 0 {
		spew.Printf("%s %-17s %s %s %s\n", t, colorMagenta(str0), colorYellow(str1), colorYellow(key)+` done in`, colorYellow(message))
	}
	if status == 1 {
		spew.Printf("%s %-17s %s %s\n", t, colorMagenta(str0), colorYellow(str1), colorYellow(key)+` strated `)
	}
}

// Pretty function disasemble a variable and display it's struct and values
func Pretty(variable ...interface{}) {
	if shouldUseUI() {
		return
	}
	spew.Config.Indent = "    "
	fmt.Printf("%s", colorYellow("----------------------------------\n"))
	for _, each := range variable {
		spew.Dump(each)
	}
	fmt.Printf("%s", colorYellow("----------------------------------\n"))
}

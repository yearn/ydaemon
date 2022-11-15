package logs

import (
	"context"
	"runtime"
	"strconv"
	"time"

	"github.com/davecgh/go-spew/spew"
	"github.com/mum4k/termdash/terminal/tcell"

	"github.com/mum4k/termdash"
	"github.com/mum4k/termdash/cell"
	"github.com/mum4k/termdash/container"
	"github.com/mum4k/termdash/linestyle"
	"github.com/mum4k/termdash/terminal/terminalapi"
	"github.com/mum4k/termdash/widgets/text"
)

type TLogs struct {
	level string
	data  []interface{}
}

var InfoLogs = []TLogs{}

type TTraces struct {
	key     string
	message string
	status  int
}

var TraceLogs = []TTraces{}

func writeInfoLine(ctx context.Context, t *text.Text) {
	currentIndex := 0

	for {
		for currentIndex < len(InfoLogs) {
			goRoutines := strconv.Itoa(runtime.NumGoroutine())
			now := time.Now().Format("2006/01/02 15:04:05")
			currentLog := InfoLogs[currentIndex]

			color := cell.FgColor(cell.ColorBlue)
			if currentLog.level == "SUCCESS" {
				color = cell.FgColor(cell.ColorGreen)
			}
			if currentLog.level == "ERROR" {
				color = cell.FgColor(cell.ColorRed)
			}
			if currentLog.level == "WARNING" {
				color = cell.FgColor(cell.ColorYellow)
			}
			if currentLog.level == "DEBUG" {
				color = cell.FgColor(cell.ColorYellow)
			}

			t.Write(
				now,
				text.WriteCellOpts(cell.FgColor(cell.ColorWhite)),
			)
			t.Write(
				spew.Sprintf(" [%04s]", goRoutines),
				text.WriteCellOpts(cell.FgColor(cell.ColorMagenta)),
			)
			t.Write(
				` [`+currentLog.level+`] `,
				text.WriteCellOpts(color),
			)
			t.Write(
				spew.Sprintf("%s\n", currentLog.data),
			)
			currentIndex++
		}
	}
}

func writeTraceLine(ctx context.Context, t *text.Text) {
	currentIndex := 0

	for {
		for currentIndex < len(TraceLogs) {
			currentTrace := TraceLogs[currentIndex]

			if currentTrace.status == 1 {
				t.Write(
					spew.Sprintf(" [ADD] "),
					text.WriteCellOpts(cell.FgColor(cell.ColorGreen)),
				)
			} else {
				t.Write(
					spew.Sprintf(" [SUB] "),
					text.WriteCellOpts(cell.FgColor(cell.ColorYellow)),
				)
			}
			t.Write(
				spew.Sprintf("%-42s", currentTrace.key),
				text.WriteCellOpts(cell.FgColor(cell.ColorMagenta)),
			)

			if currentTrace.message != "" {
				t.Write(spew.Sprintf("%s", currentTrace.message))
			}
			t.Write(spew.Sprintf("\n"))
			currentIndex++
		}
	}
}

func Ui() {
	if !shouldUseUI() {
		return
	}
	t, err := tcell.New()
	if err != nil {
		return
	}
	defer t.Close()

	ctx, cancel := context.WithCancel(context.Background())

	//Info panel
	rolledInfo, _ := text.New(text.RollContent(), text.WrapAtWords())
	go writeInfoLine(ctx, rolledInfo)

	//Trace panel
	rolledTrace, _ := text.New(text.RollContent(), text.WrapAtWords())
	go writeTraceLine(ctx, rolledTrace)

	c, err := container.New(
		t,
		container.Border(linestyle.Light),
		container.BorderTitle("PRESS Q TO QUIT"),
		container.SplitVertical(
			container.Left(
				container.PlaceWidget(rolledInfo),
			),
			container.Right(
				container.Border(linestyle.Light),
				container.BorderTitle(" TRACE "),
				container.PlaceWidget(rolledTrace),
			),
			container.SplitPercent(65),
		),
	)
	if err != nil {
		panic(err)
	}

	quit := func(k *terminalapi.Keyboard) {
		if k.Key == 'q' || k.Key == 'Q' {
			cancel()
		}
	}

	if err := termdash.Run(ctx, t, c, termdash.KeyboardSubscriber(quit)); err != nil {
		panic(err)
	}
}

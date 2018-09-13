// Copyright 2013, Örjan Persson. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package logging

import (
	"fmt"
	"io"
	"log"
	"os"
)

type color int

// color values
const (
	ColorBlack = iota + 30
	ColorRed
	ColorGreen
	ColorYellow
	ColorBlue
	ColorMagenta
	ColorCyan
	ColorWhite
)

// background color values
const (
	ColorBlackBg = iota + 40
	ColorRedBg
	ColorGreenBg
	ColorYellowBg
	ColorBlueBg
	ColorMagentaBg
	ColorCyanBg
	ColorWhiteBg
)

var (
	colors = []string{
		CRITICAL: ColorSeq(ColorMagenta),
		ERROR:    ColorSeq(ColorRed),
		WARNING:  ColorSeq(ColorYellow),
		NOTICE:   ColorSeq(ColorGreen),
		INFO:     ColorSeq(ColorGreen),
		DEBUG:    ColorSeq(ColorCyan),
		TRACE:    ColorSeq(ColorCyan),
	}
	bgcolors = []string{
		CRITICAL: ColorSeq(ColorMagentaBg),
		ERROR:    ColorSeq(ColorRedBg),
		WARNING:  ColorSeq(ColorYellowBg),
		NOTICE:   ColorSeq(ColorGreenBg),
		INFO:     ColorSeq(ColorGreenBg),
		DEBUG:    ColorSeq(ColorCyanBg),
		TRACE:    ColorSeq(ColorCyanBg),
	}
	boldcolors = []string{
		CRITICAL: ColorSeqBold(ColorMagenta),
		ERROR:    ColorSeqBold(ColorRed),
		WARNING:  ColorSeqBold(ColorYellow),
		NOTICE:   ColorSeqBold(ColorGreen),
		INFO:     ColorSeqBold(ColorGreen),
		DEBUG:    ColorSeqBold(ColorCyan),
		TRACE:    ColorSeqBold(ColorCyan),
	}
)

// LogBackend utilizes the standard log module.
type LogBackend struct {
	Logger    *log.Logger
	ErrLogger *log.Logger
	Color     bool
}

// NewLogBackend creates a new LogBackend.
func NewLogBackend(out io.Writer, prefix string, flag int, errOut ...io.Writer) *LogBackend {
	b := &LogBackend{Logger: log.New(out, prefix, flag)}
	if len(errOut) > 0 {
		b.ErrLogger = log.New(errOut[0], prefix, flag)
	}
	return b
}

// Log implements the Backend interface.
func (b *LogBackend) Log(calldepth int, rec *Record) {
	var msg string
	if b.Color {
		msg = rec.Formatted(calldepth+1, true)
	} else {
		msg = rec.Formatted(calldepth+1, false)
	}
	var err error
	if rec.Level > ERROR || b.ErrLogger == nil || rec.Level == PRINT {
		err = b.Logger.Output(calldepth+2, msg)
	} else {
		err = b.ErrLogger.Output(calldepth+2, msg)
	}
	if err != nil {
		fmt.Fprintf(os.Stderr, "unable to Console Log msg:%s [error]%s\n", msg, err.Error())
	}
}

// Close closes the log service.
func (b *LogBackend) Close() {}

// ConvertColors takes a list of ints representing colors for log levels and
// converts them into strings for ANSI color formatting
func ConvertColors(colors []int, bold bool) []string {
	converted := []string{}
	for _, i := range colors {
		if bold {
			converted = append(converted, ColorSeqBold(color(i)))
		} else {
			converted = append(converted, ColorSeq(color(i)))
		}
	}

	return converted
}

// ColorSeq adds color identifier
func ColorSeq(color color) string {
	return fmt.Sprintf("\033[%dm", int(color))
}

// ColorSeqBold adds blod color identifier
func ColorSeqBold(color color) string {
	return fmt.Sprintf("\033[%d;1m", int(color))
}

func doFmtVerbLevelColor(layout string, colorful bool, level Level, output io.Writer) {
	if colorful {
		switch layout {
		case "bold":
			output.Write([]byte(boldcolors[level]))
		case "bg":
			output.Write([]byte(bgcolors[level]))
		case "reset":
			output.Write([]byte("\033[0m"))
		default:
			output.Write([]byte(colors[level]))
		}
	}
}

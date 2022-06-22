package log

import (
	"fmt"
	"os"
	"strings"

	"github.com/OhYee/rainbow/color"
	"github.com/OhYee/rainbow/errors"
	"github.com/OhYee/rainbow/log"
)

var (
	debugLogger = log.New().SetPrefix(func(s string) string {
		return fmt.Sprintf(
			"%s %s \n\t",
			color.New().SetFontBold().SetFrontYellow().Colorful("[Debug]"),
			color.New().SetFontWeak().SetFrontYellow().Colorful(getCallStack()),
		)
	})
	infoLogger = log.New().SetPrefix(func(s string) string {
		return fmt.Sprintf(
			"%s %s \n\t",
			color.New().SetFontBold().SetFrontBlue().Colorful("[Info]"),
			color.New().SetFontWeak().SetFrontBlue().Colorful(getCallStack()),
		)
	})
	errLogger = log.New().SetPrefix(func(s string) string {
		return fmt.Sprintf(
			"%s %s \n\t",
			color.New().SetFontBold().SetFrontRed().Colorful("[Error]"),
			color.New().SetFontWeak().SetFrontRed().Colorful(getCallStack()),
		)
	})
)

func init() {
	level := strings.ToLower(os.Getenv("LEVEL"))
	if level == "" {
		level = "info"
	}

	switch level {
	case "debug":
		debugLogger.SetOutputToStdout()
		infoLogger.SetOutputToStdout()
		errLogger.SetOutputToStdout()
	case "info":
		debugLogger.SetOutputToNil()
		infoLogger.SetOutputToStdout()
		errLogger.SetOutputToStdout()
	case "error":
		debugLogger.SetOutputToNil()
		infoLogger.SetOutputToNil()
		errLogger.SetOutputToStdout()
	}
}

func Debugf(fmt string, args ...interface{}) {
	debugLogger.Printf(fmt, args...)
}

func Infof(fmt string, args ...interface{}) {
	infoLogger.Printf(fmt, args...)
}

func Errorf(fmt string, args ...interface{}) {
	errLogger.Printf(fmt, args...)
}

func getCallStack() string {
	stack := errors.GetStack()
	if len(stack) > 4 {
		return stack[4]
	}
	return ""
}

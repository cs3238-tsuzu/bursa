package logtext

// A Logrus hook that adds filename/line number/stack trace to our log outputs

import (
	"runtime"

	"github.com/Sirupsen/logrus"
)

// TargetMaskType is to specialize whether tracing works.
type TargetMaskType int

// DebugOnly makes LogText trace only Debug
const DebugOnly = 1 << logrus.DebugLevel

func NewLogtextTargetMask(levels ...logrus.Level) TargetMaskType {
	var mask TargetMaskType

	for i := range levels {
		mask |= 1 << levels[i]
	}

	return mask
}

// Logtext is a formatter
// Log depth is how many levels to ascend to find where the actual log call occurred
// while debugOnly sets whether or not stack traces should be printed outside of
// debug prints
type Logtext struct {
	Formatter  logrus.Formatter
	LogDepth   int
	TargetMask TargetMaskType
}

func NewLogtext(formatter logrus.Formatter, targetMask TargetMaskType) *Logtext {
	return &Logtext{
		LogDepth:   4,
		Formatter:  formatter,
		TargetMask: targetMask,
	}
}

// Creates a hook to be added to an instance of logger. This is called with
func (hook *Logtext) Format(entry *logrus.Entry) ([]byte, error) {

	if _, file, line, ok := runtime.Caller(hook.LogDepth); ok {
		entry.Data["line"] = line
		entry.Data["file"] = file
	}

	if (hook.TargetMask & (1 << entry.Level)) != 0 {
		stack := getTrace()
		entry.Data["stack"] = stack
	}

	return hook.Formatter.Format(entry)
}

// handles getting the stack trace and returns it as a string
func getTrace() string {
	stack := make([]byte, 2048)
	size := runtime.Stack(stack, false)
	return string(stack[:size])
}

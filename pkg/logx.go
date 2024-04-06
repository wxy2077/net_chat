package pkg

import (
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
	"github.com/uber/jaeger-client-go"
	"net-chat/global"
	"os"
)

var levelMap map[string]logrus.Level

func init() {

	levelMap = make(map[string]logrus.Level)

	levelMap["panic"] = logrus.PanicLevel
	levelMap["fatal"] = logrus.FatalLevel
	levelMap["error"] = logrus.ErrorLevel
	levelMap["warn"] = logrus.WarnLevel
	levelMap["info"] = logrus.InfoLevel
	levelMap["debug"] = logrus.DebugLevel
	levelMap["trace"] = logrus.TraceLevel
}

// LogHook 实现了 logrus.Hook 接口
type LogHook struct {
}

// Fire 方法在记录每条日志时被调用
func (hook *LogHook) Fire(entry *logrus.Entry) error {
	if entry.Context == nil {
		return nil
	}

	fromContextSpan := opentracing.SpanFromContext(entry.Context)
	if fromContextSpan != nil {
		var spanContext = fromContextSpan.Context()
		switch assertData := spanContext.(type) {
		case jaeger.SpanContext:
			traceID := assertData.TraceID().String()
			entry.Data[global.OpenTraceTraceID] = traceID
		}
	}

	return nil
}

// Levels 方法指定了钩子在哪些日志级别下生效
func (hook *LogHook) Levels() []logrus.Level {
	// 在所有日志级别下生效
	return logrus.AllLevels
}

type MyJSONFormatter struct {
}

// NewLog return a logger
func NewLog(level string, formatter string, logPath string) *logrus.Logger {
	GLog := logrus.New()

	GLog.SetLevel(levelMap[level])
	GLog.AddHook(&LogHook{})

	GLog.SetOutput(os.Stdout)

	if formatter == global.DefaultLogFormatterJSON {
		GLog.SetFormatter(&logrus.JSONFormatter{
			TimestampFormat: "2006-01-02 15:04:05",
			//PrettyPrint:     true,
		})
	} else if formatter == global.DefaultLogFormatterTEXT {
		GLog.SetFormatter(&logrus.TextFormatter{})
	}

	GLog.SetReportCaller(true)

	return GLog

}

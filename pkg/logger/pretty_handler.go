package logger

import (
	"context"
	"encoding/json"
	"io"
	"log"
	"log/slog"

	"github.com/fatih/color"
)

type PrettyHandlerOptions struct {
	SlogOpts slog.HandlerOptions
}

type PrettyHandler struct {
	slog.Handler
	l *log.Logger

	serviceName string
}

type Source struct {
	Service  string `json:"service"`
	Function string `json:"function"`
	File     string `json:"file"`
	Line     int    `json:"line"`
}

func (h *PrettyHandler) Handle(ctx context.Context, r slog.Record) error {
	level := r.Level.String() + ":"

	switch r.Level {
	case slog.LevelDebug:
		level = color.MagentaString(level)
	case slog.LevelInfo:
		level = color.BlueString(level)
	case slog.LevelWarn:
		level = color.YellowString(level)
	case slog.LevelError:
		level = color.RedString(level)
	}

	fields := make(map[string]interface{}, r.NumAttrs()+1)

	// // Source
	// fs := runtime.CallersFrames([]uintptr{r.PC})
	// f, _ := fs.Next()
	// source := Source{
	// 	Service:  h.serviceName,
	// 	Function: f.Function,
	// 	File:     f.File,
	// 	Line:     f.Line,
	// }
	// fields["_source"] = source

	r.Attrs(func(a slog.Attr) bool {
		fields[a.Key] = a.Value.Any()

		return true
	})

	b, err := json.MarshalIndent(fields, "", "  ")
	if err != nil {
		return err
	}

	timeStr := r.Time.Format("[2006-01-02 15:04:05]")
	msg := color.CyanString(r.Message)

	h.l.Println(timeStr, level, msg, color.WhiteString(string(b)))

	return nil
}

func NewPrettyHandler(
	serviceName string,
	out io.Writer,
	opts PrettyHandlerOptions,
) *PrettyHandler {
	h := &PrettyHandler{
		serviceName: serviceName,
		Handler:     slog.NewJSONHandler(out, &opts.SlogOpts),
		l:           log.New(out, "", 0),
	}

	return h
}

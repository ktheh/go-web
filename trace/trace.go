package trace

import (
	"fmt"
	"io"
)

// Tracer is struct
type Tracer interface {
	Trace(...interface{})
}

// New should return &tracer
func New(w io.Writer) Tracer {
	return &tracer{out: w}
}

type tracer struct {
	out io.Writer
}

func (t *tracer) Trace(a ...interface{}) {
	t.out.Write([]byte(fmt.Sprint(a...)))
	t.out.Write([]byte("\n"))
}

type nilTracer struct{}

func (t *nilTracer) Trace(a ...interface{}) {}

// Off OffはTraceメソッドの呼び出しを無視するTracerを返します
func Off() Tracer {
	return &nilTracer{}
}

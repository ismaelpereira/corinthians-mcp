package server

import (
	"context"
	"io"
	"os"
	"sync"

	"github.com/metoro-io/mcp-golang/transport"
	"github.com/metoro-io/mcp-golang/transport/stdio"
)

type controlledStdioTransport struct {
	inner *stdio.StdioServerTransport

	ctx    context.Context
	cancel context.CancelFunc

	done     chan struct{}
	doneOnce sync.Once
}

func newControlledStdioTransport() *controlledStdioTransport {
	t := &controlledStdioTransport{
		done: make(chan struct{}),
	}
	t.ctx, t.cancel = context.WithCancel(context.Background())

	reader := &eofNotifyingReader{
		r: os.Stdin,
		notify: func() {
			_ = t.Close()
		},
	}

	t.inner = stdio.NewStdioServerTransportWithIO(reader, os.Stdout)
	return t
}

func (t *controlledStdioTransport) Done() <-chan struct{} {
	return t.done
}

func (t *controlledStdioTransport) Start(_ context.Context) error {
	return t.inner.Start(t.ctx)
}

func (t *controlledStdioTransport) Send(ctx context.Context, message *transport.BaseJsonRpcMessage) error {
	return t.inner.Send(ctx, message)
}

func (t *controlledStdioTransport) Close() error {
	t.doneOnce.Do(func() { close(t.done) })
	t.cancel()
	return t.inner.Close()
}

func (t *controlledStdioTransport) SetCloseHandler(handler func()) {
	t.inner.SetCloseHandler(handler)
}

func (t *controlledStdioTransport) SetErrorHandler(handler func(error)) {
	t.inner.SetErrorHandler(handler)
}

func (t *controlledStdioTransport) SetMessageHandler(handler func(ctx context.Context, message *transport.BaseJsonRpcMessage)) {
	t.inner.SetMessageHandler(handler)
}

type eofNotifyingReader struct {
	r      io.Reader
	notify func()
	once   sync.Once
}

func (r *eofNotifyingReader) Read(p []byte) (int, error) {
	n, err := r.r.Read(p)
	if err != nil {
		r.once.Do(func() {
			if r.notify != nil {
				r.notify()
			}
		})
	}
	return n, err
}

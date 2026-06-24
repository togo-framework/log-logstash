// Package logstash ships togo's slog logs as JSON lines to a Logstash TCP input,
// in addition to the app's existing log output. Install alongside togo-framework/log;
// blank-import registers it.
//
// Env: LOGSTASH_ADDR (host:port of a `tcp { codec => json_lines }` input — required,
// no-op when empty).
package logstash

import (
	"context"
	"log/slog"
	"net"
	"os"
	"sync"
	"time"

	"github.com/togo-framework/togo"
)

func init() {
	togo.RegisterProviderFunc("log-logstash", togo.PriorityService, func(k *togo.Kernel) error {
		addr := os.Getenv("LOGSTASH_ADDR")
		if addr == "" {
			return nil // unconfigured → no-op
		}
		w := &conn{addr: addr}
		jh := slog.NewJSONHandler(w, &slog.HandlerOptions{Level: slog.LevelInfo})
		// Keep the existing handler + also emit JSON lines to Logstash.
		k.Log = slog.New(tee{[]slog.Handler{k.Log.Handler(), jh}})
		return nil
	})
}

// conn is a reconnecting TCP writer: dials lazily and re-dials once on write error.
// Failures are swallowed so logging never breaks the app.
type conn struct {
	addr string
	mu   sync.Mutex
	c    net.Conn
}

func (w *conn) Write(p []byte) (int, error) {
	w.mu.Lock()
	defer w.mu.Unlock()
	if w.c == nil {
		c, err := net.DialTimeout("tcp", w.addr, 3*time.Second)
		if err != nil {
			return len(p), nil
		}
		w.c = c
	}
	if _, err := w.c.Write(p); err != nil {
		w.c.Close()
		w.c = nil
		if c, err2 := net.DialTimeout("tcp", w.addr, 3*time.Second); err2 == nil {
			w.c = c
			_, _ = w.c.Write(p)
		}
	}
	return len(p), nil
}

// tee fans each record out to multiple handlers (the original + Logstash).
type tee struct{ hs []slog.Handler }

func (t tee) Enabled(ctx context.Context, l slog.Level) bool {
	for _, h := range t.hs {
		if h.Enabled(ctx, l) {
			return true
		}
	}
	return false
}
func (t tee) Handle(ctx context.Context, r slog.Record) error {
	for _, h := range t.hs {
		if h.Enabled(ctx, r.Level) {
			_ = h.Handle(ctx, r.Clone())
		}
	}
	return nil
}
func (t tee) WithAttrs(as []slog.Attr) slog.Handler {
	n := make([]slog.Handler, len(t.hs))
	for i, h := range t.hs {
		n[i] = h.WithAttrs(as)
	}
	return tee{n}
}
func (t tee) WithGroup(name string) slog.Handler {
	n := make([]slog.Handler, len(t.hs))
	for i, h := range t.hs {
		n[i] = h.WithGroup(name)
	}
	return tee{n}
}

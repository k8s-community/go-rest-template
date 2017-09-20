package handlers

import (
	"net/http"
	"testing"

	"{[( .projectPath )]}/pkg/config"
	"{[( .projectPath )]}/pkg/logger"
	"{[( .projectPath )]}/pkg/logger/standard"
	"{[( .projectPath )]}/pkg/router/bitroute"
)

func TestReady(t *testing.T) {
	h := New(standard.New(&logger.Config{}), new(config.Config))
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		h.Base(h.Ready)(bitroute.NewControl(w, r))
	})

	testHandler(t, handler, http.StatusOK, http.StatusText(http.StatusOK))
}

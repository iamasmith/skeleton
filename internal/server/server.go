package server

import (
	"context"
	"net/http"
	"sync"
	"time"
)

type ServerState struct {
	service  http.Server
	mux      *http.ServeMux
	shutdown chan struct{}
	wg       sync.WaitGroup
}

func New(listen string) *ServerState {
	var s ServerState
	s.shutdown = make(chan struct{})
	s.mux = http.NewServeMux()
	s.service = http.Server{Addr: listen, Handler: s.mux}
	return &s
}

func (s *ServerState) Mux() *http.ServeMux {
	return s.mux
}

func (s *ServerState) Start() error {
	echan := make(chan error)
	s.wg.Add(1)
	go func() {
		defer s.wg.Done()
		err := s.service.ListenAndServe()
		if err == http.ErrServerClosed {
			err = nil
		}
		echan <- err
	}()
	for {
		select {
		case <-s.shutdown:
			ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
			defer cancel()
			s.service.Shutdown(ctx)
		case err := <-echan:
			return err
		}
	}
}

func (s *ServerState) Stop() {
	close(s.shutdown)
	s.wg.Wait()
}

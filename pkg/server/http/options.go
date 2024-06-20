package httpserver

import (
	"fmt"
	"time"
)

type Option func(*Server)

// Default Validaotr
func WithDefaultValidator() Option {
	return func(s *Server) {
		s.initValidator()
	}
}

// Default Error Handler
func WithDefaultErrorHandlers() Option {
	return func(s *Server) {
		s.initErrorHandler()
	}
}

// Default Middlewares
func WithDefaultMiddlewares() Option {
	return func(s *Server) {
		s.initMiddlewares()
	}
}

// Address
func Address(host, port string) Option {
	return func(s *Server) {
		s.addr = fmt.Sprintf("%s:%s", host, port)
	}
}

// Shutdown Timeout
func ShutdownTimeout(timeout time.Duration) Option {
	return func(s *Server) {
		s.shutdownTimeout = timeout
	}
}

// Set name and version prop
func WithNameAndVersion(name, version string) Option {
	return func(s *Server) {
		s.name = name
		s.version = version
	}
}

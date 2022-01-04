package api_server

import (
	"io"
	"net/http"

	"github.com/Eclipsium/go_rest_api/internal/app/store"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

// APIServer ...
type APIServer struct {
	config *Config
	logger *logrus.Logger
	router *mux.Router
	store  *store.Store
}

func New(config *Config) *APIServer {
	return &APIServer{
		config: config,
		logger: logrus.New(),
		router: mux.NewRouter(),
	}
}

func (s *APIServer) Start() error {
	if err := s.configLogger(); err != nil {
		return err
	}

	s.configRouter()

	if err := s.configStore(); err != nil {
		return err
	}

	s.logger.Info("Starting APIserver")
	return http.ListenAndServe(s.config.BindAddr, s.router)
}

func (s *APIServer) configLogger() error {
	level, err := logrus.ParseLevel(s.config.LogLevel)
	if err != nil {
		return err
	}
	s.logger.SetLevel(level)
	return nil
}

func (s *APIServer) configRouter() {
	s.router.HandleFunc("/hello", s.handleHello())
}

func (s *APIServer) configStore() error {
	st := store.New(s.config.Store)
	if err := st.Open(); err != nil {
		return err
	}

	s.store = st
	return nil
}

func (s *APIServer) handleHello() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "Hello")
	}
}

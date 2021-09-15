package transport

import (
	"context"
	"fmt"
	"net/http"
	"throttle/internal/middleware"
	"throttle/internal/storage"

	"github.com/gorilla/mux"
	"github.com/rs/zerolog"
)

type Server struct {
	router   *mux.Router
	bind     string
	log      zerolog.Logger
	cacheSrv storage.Storage
}

// NewServer - builder for Server type
func NewServer(
	bind string,
	log zerolog.Logger,
	cacheSrv storage.Storage,
) Server {
	server := Server{
		bind:     bind,
		log:      log,
		router:   mux.NewRouter(),
		cacheSrv: cacheSrv,
	}

	server.router.HandleFunc(
		"/v1/payments",
		middleware.Throttle(
			server.getPayments,
			log.With().Str("package", "throttle").Logger(),
			server.cacheSrv,
		)).Methods(http.MethodGet)

	return server
}

// Start - start working server
func (s Server) Start() error {
	fmt.Println("server is starting...")
	return http.ListenAndServe(":"+s.bind, s.router)
}

func (s Server) getPayments(w http.ResponseWriter, r *http.Request) {
	// do something, logic, counts etc.
	err := s.cacheSrv.Increment(context.Background(), fmt.Sprintf("user_%s", r.Header.Get("User-Id")))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	_, err = w.Write([]byte(`{"data": []}`))
	if err != nil {
		s.log.Err(err).Msg("can't write body")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

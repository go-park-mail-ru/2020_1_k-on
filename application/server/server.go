package server

import (
	http1 "2020_1_k-on/application/film/delivery/http"
	"2020_1_k-on/application/film/repository"
	"2020_1_k-on/application/film/usecase"
	"github.com/gorilla/mux"
	"github.com/jackc/pgx/pgxpool"
	"net/http"
)

type server struct {
	port   string
	router *mux.Router
}

func NewServer(port string, connection *pgxpool.Pool) *server {
	router := mux.NewRouter()
	router.Use(Middleware)

	filmrepo := repository.NewPostgresForFilms(connection)
	filmUsecase := usecase.NewUserUsecase(filmrepo)

	http1.NewFilmHandler(router, filmUsecase)

	return &server{
		port:   port,
		router: router,
	}
}

func (serv server) ListenAndServe() error {
	return http.ListenAndServe(serv.port, serv.router)
}

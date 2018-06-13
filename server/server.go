package server

import (
	"net/http"
	"encoding/json"

	"github.com/gorilla/mux"
	"go.uber.org/zap"

	todo "github.com/tomoyat1/yet-another-todo-app"
	"fmt"
)


type Server struct {
	router *mux.Router
	/* TODO: abstract this */
	logger *zap.Logger
	todoItemRepo todo.ItemRepository
}

type status struct {
	Code    int    `json:"status"`
	Message string `json:"message"`
}

func NewServer(repo todo.ItemRepository) (s *Server, err error) {
	s = &Server{
		router: mux.NewRouter(),
		todoItemRepo: repo,
	}

	/* Non-expensive initializations */
	s.routes()
	s.logger, err = zap.NewProduction()
	if err != nil {
		return nil, err
	}
	return
}

func (s *Server) Start(port uint16) error {
	h := http.NewServeMux()
	h.Handle("/", s.router)
	/* TODO: Read port from env */
	return http.ListenAndServe(fmt.Sprintf(":%d", port), h)
}

func (s *Server) routes() {
	s.router.HandleFunc("/healthz", s.handleHealthz())
	s.router.HandleFunc("/todos", s.handlePostTodo()).Methods("POST")
	s.router.HandleFunc("/todos", s.handleListTodos()).Methods("GET")
	//s.router.HandleFunc("/todos/{id}", s.handleDeleteTodos()).Methods("DELETE")
	s.router.NotFoundHandler = http.HandlerFunc(s.handleNotFound())
}

func (s *Server) handleHealthz() http.HandlerFunc {
	status, err := json.Marshal(status{
		Code:    http.StatusOK,
		Message: "all is well :)",
	})
	if err != nil {
		panic("mashalling of json, which should not fail, failed...")
	}

	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write(status)
		return
	}
}

func (s *Server) handleNotFound() http.HandlerFunc {
	status, err := json.Marshal(status{
		Code:    http.StatusNotFound,
		Message: "404 not found",
	})
	if err != nil {
		panic("mashalling of json, which should not fail, failed...")
	}
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		w.Write(status)
		return
	}
}

func (s *Server) handleListTodos() http.HandlerFunc {
	type responseBody struct {
		Count int
		Items []*todo.Item
	}
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		items, err := s.todoItemRepo.List()
		if err != nil {
			/* TODO: this chunk should be reusable */
			status, err := json.Marshal(status{
				Code:    http.StatusInternalServerError,
				Message: "something went wrong: " + err.Error(),
			})
			if err != nil {
				panic("mashalling of json, which should not fail, failed...")
			}
			w.WriteHeader(http.StatusNotFound)
			w.Write(status)
			return
		}
		rb, err := json.Marshal(responseBody{
			Count: len(items),
			Items: items,
		})
		if err != nil {
			/* TODO: this chunk should be reusable */
			status, err := json.Marshal(status{
				Code:    http.StatusInternalServerError,
				Message: "something went wrong: " + err.Error(),
			})
			if err != nil {
				panic("mashalling of json, which should not fail, failed...")
			}
			w.WriteHeader(http.StatusNotFound)
			w.Write(status)
			return
		}
		w.Write(rb)
	}
}

func (s *Server) handlePostTodo() http.HandlerFunc {
	type itemParams struct {
		Title string `json:"title"`
		Details string `json:"details"`
	}
	type requestBody struct {
		Params itemParams `json:"item"`
	}
	type responseBody struct {
		Code int `json:"code"`
		Item *todo.Item `json:"item"`
	}
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		decoder := json.NewDecoder(r.Body)
		decoder.DisallowUnknownFields()
		reqb := requestBody{}
		err := decoder.Decode(&reqb)
		if err != nil {
			/* TODO: this chunk should be reusable */
			status, err := json.Marshal(status{
				Code:    http.StatusInternalServerError,
				Message: "something went wrong: " + err.Error(),
			})
			if err != nil {
				panic("mashalling of json, which should not fail, failed...")
			}
			w.WriteHeader(http.StatusNotFound)
			w.Write(status)
			return
		}
		newItem, err := todo.NewItem(reqb.Params.Title, reqb.Params.Details, false)
		if err != nil {
			/* TODO: this chunk should be reusable */
			status, err := json.Marshal(status{
				Code:    http.StatusInternalServerError,
				Message: "something went wrong: " + err.Error(),
			})
			if err != nil {
				panic("mashalling of json, which should not fail, failed...")
			}
			w.WriteHeader(http.StatusNotFound)
			w.Write(status)
			return
		}
		err = s.todoItemRepo.Save(newItem)
		if err != nil {
			/* TODO: this chunk should be reusable */
			status, err := json.Marshal(status{
				Code:    http.StatusInternalServerError,
				Message: "something went wrong: " + err.Error(),
			})
			if err != nil {
				panic("mashalling of json, which should not fail, failed...")
			}
			w.WriteHeader(http.StatusNotFound)
			w.Write(status)
			return
		}
		resb, err := json.Marshal(responseBody{
			Code: http.StatusOK,
			Item: newItem,
		})
		if err != nil {
			/* TODO: this chunk should be reusable */
			status, err := json.Marshal(status{
				Code:    http.StatusInternalServerError,
				Message: "something went wrong: " + err.Error(),
			})
			if err != nil {
				panic("mashalling of json, which should not fail, failed...")
			}
			w.WriteHeader(http.StatusNotFound)
			w.Write(status)
			return
		}
		w.Write(resb)
	}
}


package server

import (
	"app05/dao"
	"encoding/json"
	"net/http"
)

type Server struct {
	dao *dao.Dao
}

func NewServer(dao *dao.Dao) *Server {
	return &Server{
		dao: dao,
	}
}

type UsersResp struct {
	Users []*dao.User `json:"users"`
}

func (s *Server) GetAllUsersHandler(w http.ResponseWriter, r *http.Request) {
	users, err := s.dao.GetAllUsers()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	json.NewEncoder(w).Encode(UsersResp{
		Users: users,
	})
}

type CreateUserReq struct {
	Name string `json:"name"`
}

func (s *Server) CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	req := CreateUserReq{}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	if err := s.dao.CreateUser(req.Name); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	w.WriteHeader(201)
}

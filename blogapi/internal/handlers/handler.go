package handlers

import (
	"database/sql"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/VelVit24/blogapi/internal/database"
	"github.com/VelVit24/blogapi/internal/models"
)

type Handler struct {
	DB *sql.DB
}

func writeJson(w http.ResponseWriter, status int, data any) {
	w.WriteHeader(status)
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		log.Println(err)
	}
}
func (h *Handler) HandlePosts(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		param := r.URL.Query().Get("term")
		blogs, err := database.GetAllBlogs(h.DB, param)
		if err != nil {
			log.Println(err)
		}
		writeJson(w, 200, blogs)

	case "POST":
		body, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "bad request", http.StatusBadRequest)
			log.Println(err)
			return
		}
		var blog models.Blog
		err = json.Unmarshal(body, &blog)
		if err != nil {
			http.Error(w, "bad request", http.StatusBadRequest)
			log.Println(err)
			return
		}
		err = database.InsertBlog(h.DB, &blog)
		if err != nil {
			http.Error(w, "bad request", http.StatusBadRequest)
			log.Println(err)
			return
		}
		writeJson(w, 201, blog)
	}
}

func (h *Handler) HandlePostsInd(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(strings.TrimPrefix(r.URL.Path, "/posts/"))
	if err != nil {
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}
	switch r.Method {
	case "GET":
		h.HandleGet(id, w)
	case "DELETE":
		h.HandleDelete(id, w)
	case "PUT":
		h.HandlePut(id, w, r)
	}

}

func (h *Handler) HandleGet(id int, w http.ResponseWriter) {
	blog, err := database.GetBlog(h.DB, id)
	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "not found", http.StatusNotFound)
			return
		}
		log.Println(err)
		return
	}
	writeJson(w, 200, blog)
}

func (h *Handler) HandleDelete(id int, w http.ResponseWriter) {
	err := database.DeleteBlog(h.DB, id)
	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "not found", http.StatusNotFound)
			return
		}
		log.Println(err)
		return
	}
	w.WriteHeader(204)
}

func (h *Handler) HandlePut(id int, w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "bad request", http.StatusBadRequest)
		log.Println(err)
		return
	}
	var post models.Blog
	err = json.Unmarshal(body, &post)
	if err != nil {
		http.Error(w, "bad request", http.StatusBadRequest)
		log.Println(err)
		return
	}
	err = database.UpdateBlog(h.DB, id, &post)
	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "not found", http.StatusNotFound)
			return
		}
		log.Println(err)
		return
	}
	writeJson(w, 200, post)
}

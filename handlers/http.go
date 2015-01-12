package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/peleteiro/bandit-server/repository"
	"github.com/peleteiro/bandit-server/strategies"
	"net/http"
	"strings"
)

func toJson(r map[string]string) string {
	b, err := json.Marshal(r)
	if err != nil {
		return ""
	}
	return string(b)
}

func doDefaultHeaders(w http.ResponseWriter, r *http.Request) {
	if origin := r.Header.Get("Origin"); origin != "" {
		w.Header().Set("Access-Control-Allow-Origin", origin)
	}
	w.Header().Set("Access-Control-Allow-Methods", "PUT, GET, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
}

func doGet(strategy strategies.Strategy, repo repository.Repository, w http.ResponseWriter, r *http.Request) {
	doDefaultHeaders(w, r)

	var result map[string]string = make(map[string]string)

	r.ParseForm()
	for context, values := range r.Form {
		experiments := strings.Split(values[0], ",")
		result[context] = strategy.Choose(repo, context, experiments)
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	fmt.Fprint(w, toJson(result))
}

func doPut(repo repository.Repository, w http.ResponseWriter, r *http.Request) {
	doDefaultHeaders(w, r)

	r.ParseForm()
	for context, values := range r.Form {
		repo.Reward(context, values[0])
	}

	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	fmt.Fprint(w, "ok")
}

func NewHttpHandler(strategy strategies.Strategy, repo repository.Repository) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
			doGet(strategy, repo, w, r)
		case "PUT":
			doPut(repo, w, r)
		case "OPTIONS":
			doDefaultHeaders(w, r)
		default:
			http.Error(w, "Method Not Allowed", 405)
		}
	}
}

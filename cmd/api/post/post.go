package post

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"github.com/M1ralai/me-portfolio/cmd/api/db"
	"github.com/M1ralai/me-portfolio/cmd/api/types"
)

var logger = types.NewLogger("post")

func HandlePostRequests(w http.ResponseWriter, r *http.Request) {
	var logStatus int
	var logError string
	switch r.Method {
	case http.MethodGet:
		logStatus, logError = getRequest(w, r)
	case http.MethodPost:
		logStatus, logError = postRequest(w, r)
	case http.MethodDelete:
		logStatus, logError = deleteRequest(w, r)
	default:
		http.Error(w, "401 - unauthorized acces tried", http.StatusUnauthorized)
		logStatus = 401
	}
	logger.Println(r.Method, " done from: ", r.RemoteAddr, "with: ", logStatus, " status code ", logError)
}

// get posts from a sorted by date list and teke it from index a to b ex: 5 to 10
func getRequest(w http.ResponseWriter, r *http.Request) (int, string) {
	var data map[string]interface{}
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		http.Error(w, "invalid json", http.StatusBadRequest)
		return 504, "invalid json"
	}
	from, ok := data["from"].(int)
	if !ok {
		http.Error(w, "'from' key is missing or not integer", http.StatusBadRequest)
		return 504, "from' key is missing or not integer"
	}
	count, ok := data["count"].(int)
	if !ok {
		http.Error(w, "'count' key is missing or not integer", http.StatusBadRequest)
		return 504, "'count' key is missing or not integer"
	}
	res, err := db.GetPostsForIndex(from, count)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return 500, err.Error()
	}
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(res); err != nil {
		http.Error(w, "json encode error", http.StatusInternalServerError)
		return 500, err.Error()
	}
	return 200, ""
}

func postRequest(w http.ResponseWriter, r *http.Request) (int, string) {
	if types.GetEnv("API_KEY") != r.Header.Get("api_key") {
		http.Error(w, "api key is wrong or do not have api key", http.StatusUnauthorized)
		return 501, "unauthorized post request"
	}
	var data map[string]interface{}
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		http.Error(w, "invalid json", http.StatusBadRequest)
		return 504, "invalid json"
	}

	title, ok := data["title"].(string)
	if !ok {
		http.Error(w, "'title' key is missing", http.StatusBadRequest)
		return 504, "'title' key is missing"
	}
	content, ok := data["content"].(string)
	if !ok {
		http.Error(w, "'content' key is missing", http.StatusBadRequest)
		return 504, "'content' key is missing"
	}
	excerpt, ok := data["excerpt"].(string)
	if !ok {
		http.Error(w, "'content' key is missing", http.StatusBadRequest)
		return 504, "'content' key is missing"
	}
	date, ok := data["date"].(string)
	if !ok {
		http.Error(w, "'date' key is missing", http.StatusBadRequest)
		return 504, "'date' key is missing"
	}
	p := types.Post{
		Date:    date,
		Title:   title,
		Content: content,
		Excerpt: excerpt,
	}
	err := db.CreatePost(p)
	if err != nil {
		http.Error(w, "db error while creating a post", http.StatusInternalServerError)
		return 501, "db error while creating a post"
	}
	return 200, ""
}

func deleteRequest(w http.ResponseWriter, r *http.Request) (int, string) {
	if types.GetEnv("API_KEY") != r.Header.Get("api_key") {
		http.Error(w, "api key is wrong or do not have api key", http.StatusUnauthorized)
		return 501, "unauthorized post request"
	}

	var data map[string]interface{}

	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		http.Error(w, "invalid json", http.StatusBadRequest)
		return 504, "invalid json"
	}

	id, ok := data["id"].(int)
	if !ok {
		http.Error(w, "'id' key is missing or not integer", http.StatusBadRequest)
		return 504, "'id' key is missing or not integer"
	}
	err := db.DeletePost(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return 504, err.Error()
	}
	return 200, ""
}

func GetPostById(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "unauthorized access denied ", http.StatusUnauthorized)
		logger.Println(r.Method, " done from: ", r.RemoteAddr, "with: ", 401, " status code unauthorized access denied ")
		return
	}
	parts := strings.Split(r.URL.Path, "/")
	id, err := strconv.Atoi(parts[len(parts)-1])
	if err != nil {
		http.Error(w, " invalid id ", http.StatusBadRequest)
		logger.Println(r.Method, " done from: ", r.RemoteAddr, "with: ", 400, " status code invalid id error occured: ", err)
		return
	}
	res, err := db.GetPostById(id)
	if err != nil {
		http.Error(w, " there is an error occured in db ", http.StatusInternalServerError)
		logger.Println(r.Method, " done from: ", r.RemoteAddr, "with: ", 500, " status code  there is an error occured in db: ", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}

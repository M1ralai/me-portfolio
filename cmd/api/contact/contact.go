package contact

import (
	"net/http"

	"github.com/M1ralai/me-portfolio/cmd/api/types"
)

var logger = types.NewLogger("contact")

func HandleContactRequests(w http.ResponseWriter, r *http.Request) {
	var logError string
	var logStatus int
	switch r.Method {
	case http.MethodGet:
		// logStatus, logError = getContact(w, r)
	case http.MethodPost:
	default:
		http.Error(w, " unauthorized access tried ", http.StatusUnauthorized)
		logStatus = 401
	}
	logger.Println(r.Method, " done from: ", r.RemoteAddr, "with: ", logStatus, " status code ", logError)
}

// func getContact(w http.ResponseWriter, r *http.Request) (int, string) {
// 	var data map[string]interface{}
// 	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
// 		http.Error(w, "invalid json", http.StatusBadRequest)
// 		return 504, "invalid json"
// 	}
// }

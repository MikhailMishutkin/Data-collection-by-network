package handler

import (
	"encoding/json"
	"graduatework/internal/model"
	"net/http"
	"sync"

	"github.com/gorilla/mux"
)

// Handler is the struct with injection of service layer
type Handler struct {
	sm ServiceManager
}

//ServiceManager is the interface that implemented in service layer
type ServiceManager interface {
	SortSMS() [][]model.SMSData
	SortMMS() ([][]model.MMSData, int)
	SortEmailBySpeed(f func() []model.EmailData) map[string][][]model.EmailData
	SortWorkLoad() ([]int, int)
	SortIncident() (sortData []model.IncidentData, respStatusCode int)
	GetResultData(wg *sync.WaitGroup) (r model.ResultSetT)
}

func NewHandler(s ServiceManager) *Handler {
	return &Handler{sm: s}
}

//Register is the Hahdler method  that routing request from user
func (h *Handler) RegisterR(router *mux.Router) {
	router.HandleFunc("/api", h.HandleConnection)
}

//HandeConnection is the method for to handle displaying a data
func (h *Handler) HandleConnection(w http.ResponseWriter, r *http.Request) {

	result := &model.ResultT{}
	var a, b, c int
	_, a = h.sm.SortMMS()
	_, b = h.sm.SortIncident()
	_, c = h.sm.SortWorkLoad()
	if a != 200 || b != 200 || c != 200 {
		result.Status = false
		result.Error = "Error on collect data"
	} else {
		result.Status = true
		result.Data = h.sm.GetResultData(&sync.WaitGroup{})
	}

	response, err := json.Marshal(result)
	if err != nil {
		result.Error = "Error on marshal data"
	}
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Write(response)

}

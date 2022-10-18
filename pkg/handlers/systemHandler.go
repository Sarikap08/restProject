package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/Sarikap08/restProject/pkg/models"
	"github.com/Sarikap08/restProject/pkg/services"
	"github.com/Sarikap08/restProject/pkg/utils"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

type SystemHandler struct {
	storageService services.StorageService
}

func NewSystemHandler() *SystemHandler {
	return &SystemHandler{
		storageService: *services.NewStorageService(),
	}
}

func (s *SystemHandler) RegisterEndpoint(router *mux.Router) {
	r := router.NewRoute().Subrouter()
	r.HandleFunc("/machines", s.GetMachine).Methods(http.MethodGet, http.MethodOptions)
	r.HandleFunc("/machines", s.CreateMachine).Methods(http.MethodPost, http.MethodOptions)
}

func (s *SystemHandler) GetMachine(w http.ResponseWriter, r *http.Request) {
	machineJSON, err := json.Marshal(s.storageService.GetAllMachine())

	if err != nil {
		HandlerMessage := []byte(`{
			"message": "Error while parsing the machine data",
		   }`)

		utils.ReturnJsonResponse(w, http.StatusInternalServerError, HandlerMessage)
		return
	}
	utils.ReturnJsonResponse(w, http.StatusOK, machineJSON)
	return
}

func (s *SystemHandler) CreateMachine(w http.ResponseWriter, r *http.Request) {

	var machine models.Machine
	machine.Id = uuid.New()

	payload := r.Body

	defer payload.Close()
	err := json.NewDecoder(payload).Decode(&machine)

	if err != nil {
		HandlerMessage := []byte(`{			
			"message": "Error parsing the machine data",
		   }`)

		utils.ReturnJsonResponse(w, http.StatusInternalServerError, HandlerMessage)
		return
	}
	s.storageService.AddMachine(&machine)

	HandlerMessage := []byte(`{
			"message": "Machine details are saved successfully",
		}`)

	utils.ReturnJsonResponse(w, http.StatusCreated, HandlerMessage)
}

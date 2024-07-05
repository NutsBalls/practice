package endpoint

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/NutsBalls/practice-project-backend/internal/pkg/model"
)

type Service interface {
	GetStatus() model.Response
	ParseVacancies(jobName string, cityName string) model.Response
	GetRegions() model.Response
	GetVacancies(jobName string, cityName string, minSalary int, currency string, experience int) model.Response
}

type Endpoint struct {
	s Service
}

func New(s Service) *Endpoint {
	return &Endpoint{
		s: s,
	}
}

func displayData(w http.ResponseWriter, resp model.Response) {
	w.WriteHeader(resp.Code)
	data, err := json.Marshal(resp.Data)
	if err != nil {
		w.Write([]byte("error parsing json"))
		log.Println("error occurred marshalling json", err)
	}
	w.Write(data)
}

func (e *Endpoint) Vacancies(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	city := r.URL.Query().Get("city")
	minSalary, err := strconv.Atoi(r.URL.Query().Get("minSalary"))
	if err != nil {
		minSalary = 0
	}
	experience, err := strconv.Atoi(r.URL.Query().Get("experience"))
	if err != nil {
		experience = 0
	}
	currency := r.URL.Query().Get("currency")
	displayData(w, e.s.GetVacancies(name, city, minSalary, currency, experience))
}

func (e *Endpoint) Status(w http.ResponseWriter, r *http.Request) {
	displayData(w, e.s.GetStatus())
}

func (e *Endpoint) Regions(w http.ResponseWriter, r *http.Request) {
	displayData(w, e.s.GetRegions())
}

func (e *Endpoint) Parse(w http.ResponseWriter, r *http.Request) {
	jobName := r.URL.Query().Get("text")
	city := r.URL.Query().Get("city")
	log.Println(jobName, city)
	response := e.s.ParseVacancies(jobName, city)
	displayData(w, response)
}

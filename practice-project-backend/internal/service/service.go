package service

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"strings"

	"github.com/NutsBalls/practice-project-backend/internal/migrations"
	"github.com/NutsBalls/practice-project-backend/internal/pkg/database"
	"github.com/NutsBalls/practice-project-backend/internal/pkg/model"
)

var baseUrl = "https://api.hh.ru"

type Service struct {
	db database.DB
}

func New() *Service {
	db := database.NewConnection()
	migrations.SetupMigrations(db)
	return &Service{
		db: db,
	}
}

func (s *Service) GetVacancies(jobName string, cityName string, minSalary int, currency string, experience int) model.Response {
	resp := s.db.GetFromDB(jobName, cityName, minSalary, currency, experience)
	return model.Response{
		Code: 200,
		Data: resp,
	}
}

func generateErrorResponse() model.Response {
	var respText struct {
		Data string `json:"data"`
	}
	respText.Data = "server error occured"
	return model.Response{
		Code: 500,
		Data: respText,
	}
}

func (s *Service) GetStatus() model.Response {
	var respText struct {
		Data string `json:"data"`
	}

	respText.Data = "server currently working"
	return model.Response{
		Code: 200,
		Data: respText,
	}
}

func (s *Service) GetRegions() model.Response {
	regionsUrl, err := url.Parse(baseUrl + "/areas")
	if err != nil {
		log.Println("error parsing url", err)
		return generateErrorResponse()
	}

	resp, err := http.Get(regionsUrl.String())
	if err != nil {
		log.Println("error getting data", err)
		return generateErrorResponse()
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println("error reading response body", err)
		return generateErrorResponse()
	}

	var areasData model.AreasData
	err = json.Unmarshal(body, &areasData)
	if err != nil {
		log.Println("error unmarshalling json", err)
		return generateErrorResponse()
	}

	return model.Response{
		Code: 200,
		Data: areasData,
	}

}

func GetRegions() model.AreasData {
	var areasData model.AreasData
	regionsUrl, err := url.Parse(baseUrl + "/areas")
	if err != nil {
		log.Println("error parsing url", err)
		return areasData
	}

	resp, err := http.Get(regionsUrl.String())
	if err != nil {
		log.Println("error getting data", err)
		return areasData
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println("error reading response body", err)
		return areasData
	}

	err = json.Unmarshal(body, &areasData)
	if err != nil {
		log.Println("error unmarshalling json", err)
		return areasData
	}

	return areasData
}

func GetRegionIDs(name string) []string {
	var ids []string
	regions := GetRegions()
	if len(regions) == 0 {
		return ids
	}
	for _, region := range regions {
		if region.Name == "Россия" {
			ids = findAreaIDsByName(region.Areas, name)
		}
	}
	if len(ids) == 0 {
		ids = []string{"1"}
	}
	return ids
}

func findAreaIDsByName(areas []model.Area, name string) []string {
	queue := model.Queue{}
	ids := []string{}
	queue.Add(areas)
	for len(queue.Data) != 0 {
		curr := queue.Pop()
		for _, elem := range curr {
			if strings.HasPrefix(strings.ToLower(elem.Name), strings.ToLower(name)) {
				ids = append(ids, elem.ID)
			}
			queue.Add(elem.Areas)
		}
	}
	return ids
}

func (s *Service) ParseVacancies(jobName string, city string) model.Response {
	vacanciesUrl, err := url.Parse(baseUrl + "/vacancies")
	if err != nil {
		return generateErrorResponse()
	}
	q := vacanciesUrl.Query()
	q.Add("text", jobName)
	q.Add("page", "1")

	vacanciesUrl.RawQuery = q.Encode()

	resp, err := http.Get(vacanciesUrl.String())
	if err != nil {
		return generateErrorResponse()
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return generateErrorResponse()
	}
	var data model.JobsData
	err = json.Unmarshal(body, &data)

	if err != nil {
		return generateErrorResponse()
	}

	for i := 2; i <= data.Pages; i++ {
		q.Set("page", fmt.Sprint(i))
		vacanciesUrl.RawQuery = q.Encode()
		r, err := http.Get(vacanciesUrl.String())
		if err != nil {
			return generateErrorResponse()
		}
		defer r.Body.Close()
		b, err := io.ReadAll(r.Body)
		if err != nil {
			return generateErrorResponse()
		}
		var d model.JobsData
		err = json.Unmarshal(b, &d)

		if err != nil {
			return generateErrorResponse()
		}

		data.Jobs = append(data.Jobs, d.Jobs...)
	}

	s.db.SaveToDB(data)
	log.Println("saved to DB")
	return model.Response{
		Code: 200,
		Data: data,
	}

}

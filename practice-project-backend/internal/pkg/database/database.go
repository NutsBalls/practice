package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"sync"

	"github.com/NutsBalls/practice-project-backend/internal/pkg/model"
	"github.com/lib/pq"
)

type DB struct {
	Client *sql.DB
}

func NewConnection() DB {
	var (
		host     = os.Getenv("HOST")
		port     = os.Getenv("PORT")
		user     = os.Getenv("USER")
		password = os.Getenv("PASSWORD")
		dbname   = os.Getenv("DBNAME")
	)
	connString := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s?sslmode=disable", user, password, host, port, dbname)
	conn, err := sql.Open("postgres", connString)

	if err != nil {
		log.Fatal(err)
	}

	return DB{
		Client: conn,
	}
}

func (db *DB) SaveToDB(vacancies model.JobsData) {
	wg := sync.WaitGroup{}
	wg.Add(len(vacancies.Jobs))
	for _, job := range vacancies.Jobs {
		experienceId := -1
		switch job.Experience.Name {
		case `Нет опыта`:
			experienceId = 0
		case `От 1 года до 3 лет`:
			experienceId = 1
		case `От 3 до 6 лет`:
			experienceId = 2
		case `Более 6 лет`:
			experienceId = 3
		}
		query := fmt.Sprintf(`INSERT INTO vacancies
		(id, url, name, salary_from, salary_to, currency, city_name, company_name, company_url, experience)
		VALUES ('%s', '%s', '%s', %d, %d, '%s', '%s',  '%s', '%s', %d);`,
			job.ID, job.URL, job.Name, job.Salary.From, job.Salary.To, job.Salary.Currency, job.Area.Name,
			job.Employer.Name, job.Employer.URL, experienceId)

		_, err := db.Client.Exec(query)

		if err, ok := err.(*pq.Error); ok {
			if err.Code != "23505" {
				log.Println("failed to insert data to DB", err)
			}
		}
	}
}

func (db *DB) GetFromDB(jobName string, cityName string, minSalary int, currency string, experience int) model.JobsData {
	var response model.JobsData

	flag := true
	limit := 20

	getVacanciesQuery := `SELECT * FROM vacancies `

	if jobName != "" {
		if flag {
			getVacanciesQuery += `WHERE LOWER(name) LIKE` + ` LOWER('%` + jobName + `%') `
			flag = false
		} else {
			getVacanciesQuery += `AND LOWER(name) LIKE` + ` LOWER('%` + jobName + `%') `
		}
	}

	if cityName != "" {
		if flag {
			getVacanciesQuery += `WHERE LOWER(city_name) LIKE` + ` LOWER('%` + cityName + `%') `
			flag = false
		} else {
			getVacanciesQuery += `AND LOWER(city_name) LIKE` + ` LOWER('%` + cityName + `%') `
		}
	}

	if minSalary != 0 {
		if flag {
			getVacanciesQuery += fmt.Sprintf(`WHERE (salary_from > %d or salary_to > %d) and currency = '%s' `, minSalary, minSalary, currency)
			flag = false
		} else {
			getVacanciesQuery += fmt.Sprintf(`AND (salary_from > %d or salary_to > %d) and currency = '%s' `, minSalary, minSalary, currency)
		}
	}

	if flag {
		getVacanciesQuery += fmt.Sprintf(`WHERE experience <= %d `, experience)
		flag = false
	} else {
		getVacanciesQuery += fmt.Sprintf(`AND experience <= %d `, experience)
	}

	var count int

	log.Println(getVacanciesQuery)

	rows, err := db.Client.Query(getVacanciesQuery)

	if err != nil {
		log.Println("error getting vacancies", err)
	}

	for rows.Next() {
		count++
		var job model.Job

		err = rows.Scan(&job.ID, &job.URL, &job.Name, &job.Salary.From, &job.Salary.To, &job.Salary.Currency,
			&job.Area.Name, &job.Employer.Name, &job.Employer.URL, &job.Experience.Name)

		if err != nil {
			log.Println("error scanning row", err)
			return model.JobsData{}
		}

		response.Jobs = append(response.Jobs, job)
	}
	response.Pages = count / limit
	if count%limit != 0 {
		response.Pages++
	}
	log.Println(count)
	return response
}

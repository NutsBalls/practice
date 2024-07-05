-- +goose Up
CREATE TABLE vacancies(
    id TEXT UNIQUE PRIMARY KEY,
    url TEXT NOT NULL,
    name TEXT NOT NULL,
    salary_from INTEGER,
    salary_to INTEGER,
    currency TEXT,
    city_name TEXT,
    company_name TEXT,
    company_url TEXT,
    experience INTEGER
);

-- +goose Down
DROP TABLE vacancies;
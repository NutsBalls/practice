package model

type Response struct {
	Code int
	Data interface{}
}

type Job struct {
	ID   string `json:"id"`
	URL  string `json:"alternate_url"`
	Name string `json:"name"`
	Area struct {
		Name string `json:"name"`
	} `json:"area"`
	Salary struct {
		Currency string `json:"currency"`
		From     int    `json:"from"`
		To       int    `json:"to"`
	} `json:"salary"`
	Employer struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"employer"`
	Experience struct {
		Name string `json:"name"`
	} `json:"experience"`
}

type JobsData struct {
	Jobs  []Job `json:"items"`
	Pages int   `json:"pages"`
}

type Area struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Areas []Area `json:"areas"`
}

type AreasData []Area

type Queue struct {
	Data [][]Area
}

func (s *Queue) Add(new []Area) {
	s.Data = append(s.Data, new)
}

func (s *Queue) Pop() []Area {
	toRet := s.Data[0]
	newData := s.Data[1:]
	s.Data = newData
	return toRet
}

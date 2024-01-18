package structures

type PersonFullData struct {
	Person      Person            `json:"person"`
	Gender      PersonGenderize   `json:"gender"`
	Age         PersonAgify       `json:"age"`
	Nationality PersonNationalize `json:"nationality"`
}

type Person struct {
	Name       string `json:"name"`
	Surname    string `json:"surname"`
	Patronymic string `json:"patronymic"`
}

type PersonGenderize struct {
	Count       int     `json:"count"`
	Name        string  `json:"name"`
	Gender      string  `json:"gender"`
	Probability float32 `json:"probability"`
}

type PersonAgify struct {
	Count int    `json:"count"`
	Name  string `json:"name"`
	Age   int    `json:"age"`
}

type PersonNationalize struct {
	Count       int                 `json:"count"`
	Name        string              `json:"name"`
	CountryList []CountryProbablity `json:"country"`
}

type CountryProbablity struct {
	Country_ID  string  `json:"country_id"`
	Probability float32 `json:"probability"`
}

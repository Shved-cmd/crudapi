package movies

type Movie struct {
	ID       string    `json:"id"`
	Isbn     string    `json:"isbn"` // уникальный идентификатор фильма
	Title    string    `json:"title"`
	Director *Director `json:"director"`
}

type Director struct {
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
}

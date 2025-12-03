package request

type CreateBook struct {
	Title         string `json:"title"`
	Author        string `json:"author"`
	ISBN          string `json:"isbn"`
	PublishedYear int    `json:"publishedYear"`
	Genre         string `json:"genre"`
}

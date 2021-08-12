package data

type Food struct {
	ID        int    `json:"food-id"`
	WhereItIs string `json:"where-it-is"`
	ImageURL  string `json:"image-url"`
	Name      string `json:"name"`
	Cost      int    `json:"cost"`
}

type Api map[string][]Food

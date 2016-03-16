package apptweak

import "encoding/json"

type App struct {
	Developer string   `json:"developer"`
	Genres    []string `json:"genres"`
	Icon      string   `json:"icon"`
	Id        string   `json:"id"` // same as package_name
	Price     string   `json:"price"`
	Title     string   `json:"title"`
}

func (a App) String() string {
	jsonBytes, err := json.Marshal(a)
	if err != nil {
		return ""
	}
	return string(jsonBytes)
}

type AppResponse struct {
	Apps  []App `json:"content"`
	Count int   `json:"-"`
}

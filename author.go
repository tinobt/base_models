package base_models

// Repository
type Author struct {
	Base
    Name 		string 		`json:"name"`
    Avatar 		string  	`json:"avatar"`
    URL 		string 		`json:"url"`
}
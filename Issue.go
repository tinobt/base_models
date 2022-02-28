package modles

// Repository
type Issue struct {
	Base
    Name 		string 		`json:"name"`
	Description string  	`json:"description"`
    URL 		string  	`json:"url"`
    Tags 		[]string   	`json:"tags"`
    Source 		string 		`json:"source"`
	State bool `json:"boolean"`
	Author
}
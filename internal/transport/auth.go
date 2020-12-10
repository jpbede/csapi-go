package transport

// APICreds represents api credentials used for /login
type APICreds struct {
	Name string `json:"name"`
	Key  string `json:"key"`
}

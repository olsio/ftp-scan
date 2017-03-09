package types

type Grab struct {
	IP             string    `json:"ip"`
	Domain         string    `json:"domain,omitempty"`
	Time           string    `json:"timestamp"`
	Data           *GrabData `json:"data,omitempty"`
	Error          *string   `json:"error,omitempty"`
	ErrorComponent string    `json:"error_component,omitempty"`
}

type GrabData struct {
	Banner string `json:"banner,omitempty"`
}

type Result struct {
	Directories []string `json:"directories"`
}

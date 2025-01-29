package auth

// Shared models between PDP and Server
type AuthZENRequest struct {
	Subject  Subject  `json:"subject"`
	Action   Action   `json:"action"`
	Resource Resource `json:"resource"`
	Context  struct{} `json:"context"`
}

type Subject struct {
	Type string `json:"type"`
	ID   string `json:"id"`
}

type Action struct {
	Name string `json:"name"`
}

type Resource struct {
	Type       string     `json:"type"`
	ID         string     `json:"id"`
	Properties Properties `json:"properties,omitempty"`
}

type Properties struct {
	OwnerID string `json:"ownerID,omitempty"`
}

type AuthZENResponse struct {
	Decision bool `json:"decision"`
}

type BatchEvaluationRequest struct {
	Subject     Subject    `json:"subject"`
	Action      Action     `json:"action"`
	Resource    Resource   `json:"resource"`
	Context     struct{}   `json:"context"`
	Evaluations []Resource `json:"evaluations"`
}

type BatchEvaluationResponse struct {
	Evaluations []AuthZENResponse `json:"evaluations"`
}

type Todo struct {
	ID           string `json:"id"`
	Title        string `json:"title"`
	Completed    bool   `json:"completed"`
	OwnerID      string `json:"ownerID"`
	OwnerName    string `json:"ownerName"`
	OwnerEmail   string `json:"ownerEmail"`
	OwnerPicture string `json:"ownerPicture"`
}

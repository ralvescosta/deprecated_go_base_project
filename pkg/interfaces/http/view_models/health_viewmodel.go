package viewmodels

type IntegrationsViewModel struct {
	Name   string `json:"name"`
	Status string `json:"status"`
}

type HealthViewModel struct {
	Status       string                  `json:"status"`
	Integrations []IntegrationsViewModel `json:"integrations"`
}

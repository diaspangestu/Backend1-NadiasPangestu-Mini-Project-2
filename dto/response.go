package dto

type Response struct {
	Success      bool   `json:"success"`
	MessageTitle string `json:"messageTitle"`
	Message      string `json:"message"`
	ResponseTime string `json:"responseTime"`
}

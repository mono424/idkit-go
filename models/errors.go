package models

type ErrorResponse struct {
	Code      string  `json:"code"`
	Detail    string  `json:"detail"`
	Attribute *string `json:"attribute,omitempty"`
}

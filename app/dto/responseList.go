package dto

// ResponseList ...
type ResponseList struct {
	Total    int64       `json:"total_data"`
	Limit    int64       `json:"limit"`
	Page     int64       `json:"page"`
	LastPage int64       `json:"last_page"`
	Data     interface{} `json:"data"`
}

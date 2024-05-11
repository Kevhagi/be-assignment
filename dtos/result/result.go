package resultdto

type ErrorResult struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

type ErrorResultJSON struct {
	Status  int         `json:"status"`
	Message interface{} `json:"message"`
}

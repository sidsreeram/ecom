package response

type Response struct {
	StatusCode int         `json:"stastuscode"`
	Message    string      `json:"message"`
	Data       interface{} `json:"data"`
	Errors     interface{} `json:"error"`
}
type ProductImage struct {
	ID             int
	ProductItemID  int
	FileName       string
	Data           []byte
}

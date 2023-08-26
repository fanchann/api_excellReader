package types

type (
	WebResponse struct {
		Status  int         `json:"status"`
		Message string      `json:"message"`
		Data    interface{} `json:"data"`
	}

	CustomerResponse struct {
		Id   uint   `json:"customer_id"`
		Name string `json:"customer_name"`
	}
)

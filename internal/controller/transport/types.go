package transport

type response struct {
	ErrMessage   string      `json:"err_message,omitempty"`
	ResponseData interface{} `json:"response,omitempty"`
}

package models

/**
  Estructura auxiliar para almacenar la respuestas
*/
type Response struct {
	Code    int    `json:"code,omitempty"`
	Message string `json:"message,omitempty"`
}

func (r *Response) add(code int, message string) {
	if r != nil {
		r.Code = code
		r.Message = message
	}
}

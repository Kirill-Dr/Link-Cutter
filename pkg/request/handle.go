package request

import (
	"link-cutter/pkg/response"
	"net/http"
)

func HandleBody[T any](w *http.ResponseWriter, req *http.Request) (*T, error) {
	body, err := Decode[T](req.Body)
	if err != nil {
		response.JSON(*w, err.Error(), http.StatusBadRequest)
		return nil, err
	}

	err = IsValid(body)
	if err != nil {
		response.JSON(*w, err.Error(), http.StatusBadRequest)
		return nil, err
	}

	return &body, nil
}

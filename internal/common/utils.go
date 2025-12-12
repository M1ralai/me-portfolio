package common

import (
	"encoding/json"
	"net/http"

	"github.com/go-playground/validator"
)

func ReadJson[T any](r *http.Request, validate *validator.Validate) (T, error) {
	var res T
	err := json.NewDecoder(r.Body).Decode(&res)
	if err != nil {
		return res, err
	}
	if validate != nil {
		err = validate.Struct(res)
		if err != nil {
			return res, err
		}
	}
	return res, nil
}

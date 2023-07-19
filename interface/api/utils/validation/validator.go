package validation

import (
	"fmt"
	"net/http"
	"reflect"
	"regexp"
	"strconv"
	"strings"

	"github.com/go-playground/validator/v10"
)

// ValidateReq
func ValidateReq(req interface{}) interface{} {
	validate := validator.New()
	err := validate.Struct(req)
	if err != nil {
		return errorMessage(err, req)
	}

	return nil
}

// ErrorMessage is to set error message for validator to perform readable error message.
func errorMessage(err error, req interface{}) interface{} {
	var errorValidation = map[string]string{}

	if castedObject, ok := err.(validator.ValidationErrors); ok {
		for _, err := range castedObject {
			fieldName, fieldMsg := getFieldInfo(req, err)

			errorValidation[fieldName] = fieldMsg

		}
	}

	return struct {
		Status          bool              `json:"status"`
		Message         string            `json:"message"`
		Code            int               `json:"code"`
		ErrorValidation map[string]string `json:"error_validation"`
	}{
		Status:          false,
		Message:         http.StatusText(http.StatusBadRequest),
		Code:            http.StatusBadRequest,
		ErrorValidation: errorValidation,
	}
}

// getFieldInfo
func getFieldInfo(data interface{}, err validator.FieldError) (string, string) {

	var fieldMsg string
	var arrFieldTag []string

	structNamespace := err.StructNamespace()

	structFields := strings.Split(structNamespace, ".")
	structFields = structFields[1:]

	r := reflect.ValueOf(data).Elem()
	for _, i := range structFields {
		fieldName, idxArr := handleSliceField(i)
		field, ok := r.Type().FieldByName(fieldName)
		if !ok {
			return "", ""
		}

		arrFieldTag = append(arrFieldTag, field.Tag.Get("json"))

		switch err.Tag() {
		case "required", "required_if":
			fieldMsg = field.Tag.Get("required")
		case "min":
			fieldMsg = field.Tag.Get("min")
		case "max":
			fieldMsg = field.Tag.Get("max")
		}

		if r.FieldByName(fieldName).Kind() == reflect.Struct {
			r = reflect.ValueOf(r.FieldByName(fieldName).Interface())
		} else if r.FieldByName(fieldName).Kind() == reflect.Slice {
			if r.FieldByName(fieldName).Len() > 0 {
				r = reflect.ValueOf(r.FieldByName(fieldName).Index(idxArr).Interface())
			}
		}
	}

	return strings.Join(arrFieldTag, "."), fieldMsg
}

// handleSliceField
func handleSliceField(field string) (string, int) {
	rgx := regexp.MustCompile(`\[(.*?)\]`)
	index := 0

	isArray := rgx.MatchString(field)
	if !isArray {
		return field, index
	}

	rs := rgx.FindStringSubmatch(field)
	if intVar, err := strconv.Atoi(rs[1]); err == nil {
		index = intVar
	}

	field = strings.Trim(field, fmt.Sprintf("[%d]", index))
	return field, index
}

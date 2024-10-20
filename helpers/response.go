package helpers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

const (
	StatusAPISuccess = "SUCCESS"
	StatusAPIError   = "ERROR"
	StatusAPIFailure = "FAILURE"
)

const APIFailureMessage = "Internal Server Error"

type API struct {
	statusCode int
	Code       int    `json:"code,omitempty"`
	Status     string `json:"status"`
	State      string `json:"state,omitempty"`
	Message    string `json:"message,omitempty"`
}

type APISuccess struct {
	*API
	Meta interface{} `json:"meta,omitempty"`
	Data interface{} `json:"data,omitempty"`
}

type PaginationParams struct {
	Path        string
	Page        string
	TotalRows   int32
	TotalPages  int32
	PerPage     int32
	OrderBy     string
	SortBy      string
	CurrentPage int32
}

type APIError struct {
	*API
	Errors error `json:"errors,omitempty"`
}

type APIFailure struct {
	*API
	catransaction error
}

func (a *API) StatusCode() int {
	return a.statusCode
}

func NewHTTPResponse(entity string) *API {
	return &API{
		Status: StatusAPISuccess,
	}
}

func formatState(status string) string {
	status = strings.Title(strings.ToLower(status))
	return status
}

// ============================= HANDLE SUCCESS RESPONSE ===================================
func (a *API) Success(data interface{}, code int, message string) *APISuccess {
	a.statusCode = code
	a.Status = StatusAPISuccess
	a.Message = message
	a.State = formatState(a.Status)
	return &APISuccess{
		API:  a,
		Data: data,
	}
}

func SuccessResponseJSON(w http.ResponseWriter, statusCode int, data *APISuccess) error {
	if statusCode == http.StatusNoContent {
		w.WriteHeader(http.StatusNoContent)
		return nil
	}

	jsonData, err := json.MarshalIndent(data, "", "\t")
	if err != nil {
		return err
	}

	w.Header().Set("Content-Type", "application/json")

	w.WriteHeader(statusCode)

	if _, err := w.Write(jsonData); err != nil {
		return err
	}

	return nil
}

func (a *API) SuccessJSON(e http.ResponseWriter, data interface{}, code int, message string) error {
	a.statusCode = code
	a.Status = StatusAPISuccess
	a.Message = message
	a.State = formatState(a.Status)
	return SuccessResponseJSON(e, a.statusCode, &APISuccess{
		API:  a,
		Data: data,
	})
}

func (a *API) SuccessWithMeta(w http.ResponseWriter, data interface{}, meta interface{}, code int, message string) {
	res := a.Success(data, code, message)
	res.Meta = meta

	SuccessResponseJSON(w, res.statusCode, res)
}

// SuccessWithoutData returns response format for success state without data.
func (a *API) SuccessWithoutData(w http.ResponseWriter, code int, message string) {
	a.statusCode = code
	a.Status = StatusAPISuccess
	a.Message = message
	a.State = formatState(a.Status)
	response := &APISuccess{
		API: a,
	}
	SuccessResponseJSON(w, a.statusCode, response)
}

func (a *API) Error(code int, message string) *APIError {
	a.statusCode = code
	a.Status = StatusAPIError
	a.Message = message
	a.State = formatState(a.Status)
	return &APIError{
		API: a,
	}
}

// ErrorResponseJSON setting response for error condition
func ErrorResponseJSON(w http.ResponseWriter, statusCode int, data *APIError) error {
	// If there is nothing to marshal then set status code and return.
	if statusCode == http.StatusNoContent {
		w.WriteHeader(http.StatusNoContent)
		return nil
	}

	// Encode the data to JSON.
	jsonData, err := json.MarshalIndent(data, "", "\t")
	if err != nil {
		return err
	}

	// Set the content type and headers once we know marshaling has succeeded.
	w.Header().Set("Content-Type", "application/json")

	// Write the status code to the response.
	w.WriteHeader(statusCode)

	// Send the result back to the client.
	if _, err := w.Write(jsonData); err != nil {
		return err
	}

	return nil
}

// Error returns response format for error state.
func (a *API) ErrorJSON(w http.ResponseWriter, code int, message string) {
	a.statusCode = code
	a.Status = StatusAPIError
	a.Message = message
	a.State = formatState(a.Status)
	response := &APIError{
		API: a,
	}

	ErrorResponseJSON(w, a.statusCode, response)
}

// FieldErrors returns response format error.
func (a *API) FieldErrors(w http.ResponseWriter, err error, code int, message string) {
	fe := a.Error(code, message)
	fe.Errors = err

	ErrorResponseJSON(w, fe.statusCode, fe)
}

// ErrorWithStatusCode returns response format error.
func (a *API) ErrorWithStatusCode(w http.ResponseWriter, code int, message string) {
	a.statusCode = code
	a.Code = code
	a.Status = StatusAPIError
	a.Message = strings.Title(message)
	a.State = formatState(a.Status)
	response := &APIError{
		API: a,
	}

	ErrorResponseJSON(w, a.statusCode, response)
}

// ============================= ======================= ===================================

// ============================= HANDLE FAILURE RESPONSE ===================================
// Failure returns response format for failure state.
func (a *API) Failure(err error, code int) *APIFailure {
	a.statusCode = code
	a.Status = StatusAPIFailure
	a.Message = APIFailureMessage
	a.State = formatState(a.Status)
	return &APIFailure{
		API:           a,
		catransaction: nil,
	}
}

// FailureResponseJSON setting response for failure condition
func FailureResponseJSON(w http.ResponseWriter, statusCode int, data *APIFailure) error {
	// If there is nothing to marshal then set status code and return.
	if statusCode == http.StatusNoContent {
		w.WriteHeader(http.StatusNoContent)
		return nil
	}

	// Encode the data to JSON.
	jsonData, err := json.MarshalIndent(data, "", "\t")
	if err != nil {
		return err
	}

	// Set the content type and headers once we know marshaling has succeeded.
	w.Header().Set("Content-Type", "application/json")

	// Write the status code to the response.
	w.WriteHeader(statusCode)

	// Send the result back to the client.
	if _, err := w.Write(jsonData); err != nil {
		return err
	}

	return nil
}

// Failure returns response format for failure state.
func (a *API) FailureJSON(w http.ResponseWriter, err error, code int) {
	a.statusCode = code
	a.Status = StatusAPIFailure
	a.Message = APIFailureMessage
	a.State = formatState(a.Status)
	response := &APIFailure{
		API:           a,
		catransaction: nil,
	}

	FailureResponseJSON(w, a.statusCode, response)
}

// ============================= ======================= ===================================

// Error implements error interface.
func (f *APIFailure) Error() string {
	b, err := json.Marshal(f) // {"", ""}
	if err != nil {
		return err.Error()
	}
	return string(b)
}

// Catransaction returns error catransaction.
// The Catransaction error is needed for logging.
func (f *APIFailure) Catransaction() error {
	return f.catransaction
}

// Error implement error interface.
func (e *APIError) Error() string {
	b, err := json.Marshal(e)
	if err != nil {
		return err.Error()
	}

	return string(b)
}

func JSON(w http.ResponseWriter, statusCode int, data interface{}) {
	w.WriteHeader(statusCode)
	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		fmt.Fprintf(w, "%s", err.Error())
	}
}

func ERROR(w http.ResponseWriter, statusCode int, err error) {
	if err != nil {
		JSON(w, statusCode, struct {
			Error string `json:"error"`
		}{
			Error: err.Error(),
		})
		return
	}
	JSON(w, http.StatusBadRequest, nil)
}

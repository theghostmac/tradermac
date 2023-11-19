package alphavantage

// APIRequest contains necessary information to execute an HTTP request.
type APIRequest struct {
	method string
	path   string
	params map[string]string
}

// NewAPIRequest instantiates a new APIRequest to AlphaVantage using GET as default.
func NewAPIRequest() *APIRequest {
	apiReq := &APIRequest{
		method: "GET",
		params: make(map[string]string),
	}

	return apiReq
}

// SetAPIReqMethod sets a specific API method preferred.
func (ar *APIRequest) SetAPIReqMethod(method string) *APIRequest {
	ar.method = method
	return ar
}

// SetAPIReqPath sets a specific API path preferred.
func (ar *APIRequest) SetAPIReqPath(path string) *APIRequest {
	ar.path = "/" + path
	return ar
}

// SetAPIReqParam adds a new parameter to the APIRequest.
func (ar *APIRequest) SetAPIReqParam(param, value string) *APIRequest {
	ar.params[param] = value
	return ar
}

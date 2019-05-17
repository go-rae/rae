package rae

// DefaultRestAPIError default RestAPIError impl
type DefaultRestAPIError struct {
	ErrCode string `json:"err_code,omitempty"`
	ErrMsg  string `json:"err_msg,omitempty"`
	ErrDesc string `json:"err_desc,omitempty"`
}

//Code get error code for rae
func (ie DefaultRestAPIError) Code() string {
	return ie.ErrCode
}

//Msg get msg for rae
func (ie DefaultRestAPIError) Msg() string {
	return ie.ErrMsg
}

// Desc get description for rae
func (ie DefaultRestAPIError) Desc() string {
	return ie.ErrDesc
}

// Set set error value for rae
func (ie *DefaultRestAPIError) Set(code, msg, desc string) {
	ie.ErrCode = code
	ie.ErrMsg = msg
	ie.ErrDesc = desc
}

package types

type generalResponseType struct {
	ResponseStatus    bool   `json:"response_status"`
	ResponseCode      int64  `json:"response_code"`
	ResponseMessage   string `json:"response_message"`
	ResponseTimestamp string `json:"response_timestamp"`
}

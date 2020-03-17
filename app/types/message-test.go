package types

type GetMessageResponseTest struct {
	GeneralResponse generalResponseType `json:"general_response"`
	Result          []MessageResponse   `json:"result"`
}

type AddMessageResponseTest struct {
	GeneralResponse generalResponseType `json:"general_response"`
	Result          MessageResponse     `json:"result"`
}
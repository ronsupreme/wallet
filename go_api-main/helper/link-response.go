package helper

type WalletResponse struct {
	Header interface{} `json:"header"`
	Body   interface{} `json:"data"`
}
type EmptyObj2 struct {
}

func BuildLinkRequest(header interface{}, body interface{}) WalletResponse {
	res := WalletResponse{
		Header: header,
		Body:   body,
	}
	return res
}

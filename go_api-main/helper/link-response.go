package helper

type WalletResponse struct {
	Header HeaderResp `json:"header"`
	Body   BodyResp `json:"body"`
}
type HeaderResp struct {
	RespCode string `json:"respCode"`
	RespMsg string `json:"respMsg"`
	BankSignature string `json:"bankSignature"`
}
type BodyResp struct{
	BankRef string `json:"bank_ref"`
}

func BuildLinkReponse(header HeaderResp, body BodyResp) WalletResponse {
	res := WalletResponse{
		Header: header,
		Body:   body,
	}
	return res
}

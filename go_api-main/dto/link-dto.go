package dto

type Request struct {
	Header HeaderLinkDTO `json:"header"`
	Body   BodyLinkDTO   `json:"body"`
}

type HeaderLinkDTO struct {
	Idrequest  string `json:"idrequest"`
	Providerno string `json:"providerno"`
	Signature  string `json:"signature"`
}

type BodyLinkDTO struct {
	Accountno    string `json:"Accountno"`
	Wallet_type  string `json:"wallet_type"`
	Uniqueid     string `json:"uniqueid"`
	Phonenumber  string `json:"phonenumber"`
	Customername string `json:"customername"`
}

func BuildLinkRequest(header HeaderLinkDTO, body BodyLinkDTO) Request {
	res := Request{
		Header: header,
		Body:   body,
	}
	return res
}

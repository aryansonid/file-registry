package handlers

type UploadFileResponse struct {
	CID                 string `json:"cid"`
	TransactionResponse string `json:"transactionResponse"`
}

type GetFileResponse struct {
	CID string `json:"cid"`
}

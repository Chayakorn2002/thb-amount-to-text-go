package dto

type ConvertDecimalToBahtTextRequest struct {
	Amount string `query:"amount" validate:"required"`
}

type ConvertDecimalToBahtTextResponse struct {
	AmountText string `json:"amount_text"`
}

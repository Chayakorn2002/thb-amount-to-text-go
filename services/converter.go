package services

import (
	"context"

	"github.com/Chayakorn2002/thb-amount-to-text-go/dto"
	"github.com/Chayakorn2002/thb-amount-to-text-go/utils"
	"github.com/shopspring/decimal"
)

type ConverterService interface {
	ConvertDecimalToBahtText(ctx context.Context, in *dto.ConvertDecimalToBahtTextRequest) (*dto.ConvertDecimalToBahtTextResponse, error)
}

type converterService struct{}

func NewConverterService() ConverterService {
	return &converterService{}
}

func (s *converterService) ConvertDecimalToBahtText(ctx context.Context, in *dto.ConvertDecimalToBahtTextRequest) (*dto.ConvertDecimalToBahtTextResponse, error) {
	decimalAmt, err := decimal.NewFromString(in.Amount)
	if err != nil {
		return nil, err
	}

	amountText, err := utils.DecimalToBahtText(decimalAmt)
	if err != nil {
		return nil, err
	}

	return &dto.ConvertDecimalToBahtTextResponse{
		AmountText: amountText,
	}, nil
}

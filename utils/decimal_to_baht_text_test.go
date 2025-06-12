package utils_test

import (
	"errors"
	"fmt"
	"strconv"
	"testing"

	"github.com/Chayakorn2002/thb-amount-to-text-go/utils"
	"github.com/shopspring/decimal"
)

type testCase struct {
	input         string
	expectedText  string
	expectedError error
}

func TestDecimalToBahtText(t *testing.T) {
	tests := []testCase{
		// Zero & leading
		{strconv.Itoa(0), "ศูนย์บาทถ้วน", nil},
		{decimal.NewFromFloat(0.0).String(), "ศูนย์บาทถ้วน", nil},
		{strconv.Itoa(0000), "ศูนย์บาทถ้วน", nil},
		{strconv.Itoa(0001), "หนึ่งบาทถ้วน", nil},
		{decimal.NewFromFloat(0000.50).String(), "ห้าสิบสตางค์", nil},

		// Satang only
		{decimal.NewFromFloat(0.01).String(), "หนึ่งสตางค์", nil},
		{decimal.NewFromFloat(0.04).String(), "สี่สตางค์", nil},
		{decimal.NewFromFloat(0.05).String(), "ห้าสตางค์", nil},
		{decimal.NewFromFloat(0.10).String(), "สิบสตางค์", nil},
		{decimal.NewFromFloat(0.25).String(), "ยี่สิบห้าสตางค์", nil},
		{decimal.NewFromFloat(0.50).String(), "ห้าสิบสตางค์", nil},
		{decimal.NewFromFloat(0.99).String(), "เก้าสิบเก้าสตางค์", nil},

		// Baht + Satang
		{decimal.NewFromFloat(1.01).String(), "หนึ่งบาทหนึ่งสตางค์", nil},
		{decimal.NewFromFloat(11.10).String(), "สิบเอ็ดบาทสิบสตางค์", nil},
		{decimal.NewFromFloat(99.99).String(), "เก้าสิบเก้าบาทเก้าสิบเก้าสตางค์", nil},
		{decimal.NewFromFloat(123.01).String(), "หนึ่งร้อยยี่สิบสามบาทหนึ่งสตางค์", nil},

		// Integer only
		{strconv.Itoa(1), "หนึ่งบาทถ้วน", nil},
		{strconv.Itoa(10), "สิบบาทถ้วน", nil},
		{strconv.Itoa(11), "สิบเอ็ดบาทถ้วน", nil},
		{strconv.Itoa(21), "ยี่สิบเอ็ดบาทถ้วน", nil},
		{strconv.Itoa(101), "หนึ่งร้อยหนึ่งบาทถ้วน", nil},
		{strconv.Itoa(1_001), "หนึ่งพันหนึ่งบาทถ้วน", nil},
		{strconv.Itoa(2_001), "สองพันหนึ่งบาทถ้วน", nil},

		// Large numbers
		{strconv.Itoa(123_456), "หนึ่งแสนสองหมื่นสามพันสี่ร้อยห้าสิบหกบาทถ้วน", nil},
		{strconv.Itoa(1_000_000), "หนึ่งล้านบาทถ้วน", nil},
		{strconv.Itoa(2_000_000), "สองล้านบาทถ้วน", nil},
		{strconv.Itoa(2_100_000), "สองล้านหนึ่งแสนบาทถ้วน", nil},
		{strconv.Itoa(123_456_789), "หนึ่งร้อยยี่สิบสามล้านสี่แสนห้าหมื่นหกพันเจ็ดร้อยแปดสิบเก้าบาทถ้วน", nil},
		{decimal.NewFromFloat(123_456_789.50).String(), "หนึ่งร้อยยี่สิบสามล้านสี่แสนห้าหมื่นหกพันเจ็ดร้อยแปดสิบเก้าบาทห้าสิบสตางค์", nil},
		{strconv.Itoa(1_000_000_000), "หนึ่งพันล้านบาทถ้วน", nil},
		{strconv.Itoa(1_000_000_000_000), "หนึ่งล้านล้านบาทถ้วน", nil},
		{strconv.Itoa(9_000_000_000_000), "เก้าล้านล้านบาทถ้วน", nil},
		{strconv.Itoa(999_999_999_999_999), "เก้าร้อยเก้าสิบเก้าล้านล้านเก้าแสนเก้าหมื่นเก้าพันเก้าร้อยเก้าสิบเก้าล้านเก้าแสนเก้าหมื่นเก้าพันเก้าร้อยเก้าสิบเก้าบาทถ้วน", nil},
		{strconv.Itoa(1_000_000_000_000_000), "หนึ่งพันล้านล้านบาทถ้วน", nil},

		// Error cases
		{decimal.NewFromFloat(0.001).String(), "", errors.New("only 2 decimal places are supported, got 3")},
		{decimal.NewFromFloat(123.456).String(), "", errors.New("only 2 decimal places are supported, got 3")},
		{decimal.NewFromFloat(-1.00).String(), "", errors.New("negative numbers are not supported")},
		{strconv.Itoa(1_000_000_000_000_001), "", fmt.Errorf("number too large to convert (must not exceed %s)", utils.MaxValue.String())},
	}

	for _, tc := range tests {
		val, err := decimal.NewFromString(tc.input)
		if err != nil {
			t.Errorf("Failed to parse input %q: %v", tc.input, err)
			continue
		}

		output, err := utils.DecimalToBahtText(val)

		if tc.expectedError != nil {
			if err == nil {
				t.Errorf("Expected error for input %q but got none", tc.input)
			} else if err.Error() != tc.expectedError.Error() {
				t.Errorf("Expected error %q for input %q but got %q", tc.expectedError.Error(), tc.input, err.Error())
			}
			continue
		}

		if err != nil {
			t.Errorf("Unexpected error for input %q: %v", tc.input, err)
			continue
		}

		if output != tc.expectedText {
			t.Errorf("DecimalToBahtText(%q) = %q; want %q", tc.input, output, tc.expectedText)
		}
	}
}

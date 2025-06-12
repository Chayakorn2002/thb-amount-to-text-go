package utils

import (
	"fmt"
	"math/big"
	"strings"

	"github.com/shopspring/decimal"
)

var MaxValue = big.NewInt(1_000_000_000_000_000) // หนึ่งพันล้านล้าน
var digitWords = []string{"", "หนึ่ง", "สอง", "สาม", "สี่", "ห้า", "หก", "เจ็ด", "แปด", "เก้า"}
var unitWords = []string{"", "สิบ", "ร้อย", "พัน", "หมื่น", "แสน"}

func numberToThaiText(n *big.Int) string {
	numStr := fmt.Sprintf("%d", n)

	/*
		Pad the string to handle "ล้าน" in Thai numeric spelling behavior.
		numStr = "100000000" // 100,000,000
		padded = "000100000000" // 000,100,000,000
	*/
	padded := numStr
	if len(padded)%6 != 0 {
		padded = strings.Repeat("0", 6-(len(padded)%6)) + padded
	}

	groupCount := len(padded) / 6

	parts := []string{}
	for i := 0; i < groupCount; i++ {
		group := padded[i*6 : (i+1)*6]
		groupValue := convertGroup(strings.TrimLeft(group, "0"))
		if groupValue != "" {
			/*
				Example:
				- 1,000,000 -> "หนึ่งร้อยล้าน" (groupCount = 2)
				- 1,000,000,000,000 -> "หนึ่งล้านล้าน" (groupCount = 3)
			*/
			for j := 0; j < groupCount-i-1; j++ {
				groupValue += "ล้าน"
			}
			parts = append(parts, groupValue)
		}
	}

	return strings.Join(parts, "")
}

func convertGroup(group string) string {
	result := ""
	digits := []rune(group)
	for i := 0; i < len(digits); i++ {
		pos := len(digits) - i - 1
		d := int(digits[i] - '0') // Take initiative of the ASCII Code

		if d == 0 {
			continue
		}

		word := digitWords[d]
		if pos == 0 && d == 1 && len(digits) == 2 {
			word = "เอ็ด"
		} else if pos == 1 && d == 2 {
			word = "ยี่"
		} else if pos == 1 && d == 1 {
			word = ""
		}

		result += word + unitWords[pos]
	}

	return result
}

func DecimalToBahtText(in decimal.Decimal) (string, error) {
	if in.GreaterThan(decimal.NewFromBigInt(MaxValue, 0)) {
		return "", fmt.Errorf("number too large to convert (must not exceed %s)", MaxValue.String())
	}
	if in.IsNegative() {
		return "", fmt.Errorf("negative numbers are not supported")
	}

	fracPart := in.Mod(decimal.NewFromInt(1)) // e.g., 1234.56 -> 0.56
	fracDigits := -fracPart.Exponent()        // Get number of decimal digits (e.g., 0.56 has exponent -2, so fracDigits = 2)
	if fracDigits > 2 {
		return "", fmt.Errorf("only 2 decimal places are supported, got %d", fracDigits)
	}

	if in.IsZero() {
		return "ศูนย์บาทถ้วน", nil
	}

	intPart := in.BigInt()                                   // e.g., 1234.56 -> 1234
	satang := fracPart.Mul(decimal.NewFromInt(100)).BigInt() // e.g., 0.56 -> 56

	text := numberToThaiText(intPart)
	if text != "" {
		text += "บาท"
	}

	if satang.Cmp(big.NewInt(0)) == 0 {
		text += "ถ้วน"
	} else {
		text += numberToThaiText(satang) + "สตางค์"
	}

	return text, nil
}

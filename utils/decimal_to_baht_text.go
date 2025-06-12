package utils

import (
	"fmt"
	"math/big"
	"strings"

	"github.com/shopspring/decimal"
)

var maxDigits = 150
var digitWords = []string{"", "หนึ่ง", "สอง", "สาม", "สี่", "ห้า", "หก", "เจ็ด", "แปด", "เก้า"}
var unitWords = []string{"", "สิบ", "ร้อย", "พัน", "หมื่น", "แสน"}

func numberToThaiText(n *big.Int) string {
	numStr := fmt.Sprintf("%d", n)

	/*
		Pad the string to handle "ล้าน" in the Thai numeric spelling behavior.
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
	if len(in.BigInt().String()) > maxDigits {
		return "", fmt.Errorf("number too large to convert (%d digits > %d allowed)", len(in.String()), maxDigits)
	}
	if in.IsZero() {
		return "ศูนย์บาทถ้วน", nil
	}

	intPart := in.BigInt()
	fracPart := in.Mod(decimal.NewFromInt(1)).Mul(decimal.NewFromInt(100)).BigInt()

	text := numberToThaiText(intPart)
	if text != "" {
		text += "บาท"
	}

	if fracPart.Cmp(big.NewInt(0)) == 0 {
		text += "ถ้วน"
	} else {
		text += numberToThaiText(fracPart) + "สตางค์"
	}

	return text, nil
}

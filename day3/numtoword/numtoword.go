package numtoword

import (
	"errors"
	"strconv"
)

func NumToWord(num string) (string, error) {

	var result string

	if len(num) == 0 {
		return num, nil
	}

	units := map[string]string{"0": "ноль", "1": "один", "2": "два", "3": "три", "4": "четыре", "5": "пять",
		"6": "шесть", "7": "семь", "8": "восемь", "9": "девять", "10": "десять", "11": "одинадцать", "12": "двенадцать",
		"13": "тринадцать", "14": "четырнадцать", "15": "пятнадцать", "16": "шестнадцать", "17": "семнадцать",
		"18": "восемнадцать", "19": "девятнадцать"}

	ten := map[string]string{"2": "двадцать", "3": "тридцать", "4": "сорок", "5": "пятдесят", "6": "шестдесят",
		"7": "семдесят", "8": "восемдесят", "9": "девяносто"}

	hun := map[string]string{"1": "сто", "2": "двести", "3": "триста", "4": "четыреста", "5": "пятсот",
		"6": "шестьсот", "7": "семьсот", "8": "восемьсот", "9": "девятьсот"}

	if len(num) > 0 {
		if num[0] == '-' {
			return "", errors.New("число не входит в диапазон 0-999")
		}
	}

	if _, err := strconv.ParseInt(num, 10, 64); err != nil {
		return "", errors.New("это не число")
	}
	if len(num) == 1 {
		result = units[num]
	}
	if len(num) == 2 {
		if num[:1] == "1" {
			result = units[num]
		} else if num[1:] == "0" {
			result = ten[num[:1]]
		} else {
			result = ten[num[:1]] + " " + units[num[1:]]
		}
	}
	if len(num) == 3 {
		if num[1:2] == "0" && num[2:] == "0" {
			result = hun[num[:1]]
		} else if num[1:2] == "1" {
			result = hun[num[:1]] + " " + units[num[1:]]
		} else if num[1:2] == "0" {
			result = hun[num[:1]] + " " + units[num[2:]]
		} else if num[2:] == "0" {
			result = hun[num[:1]] + " " + ten[num[1:2]]
		} else {
			result = hun[num[:1]] + " " + ten[num[1:2]] + " " + units[num[2:]]
		}
	}
	if len(num) > 3 {
		return "", errors.New("число не входит в диапазон 0-999")
	}
	return result, nil
}

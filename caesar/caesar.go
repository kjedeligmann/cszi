package main

func mod(num int, mod uint) (res int) {
	res = num % int(mod)

	if res < 0 {
		res += int(mod)
	}

	return res
}

const engAlfLen = 26

func caesar(plaintext *string, key int) []rune {
    runes := []rune(*plaintext)
    ciphertext := make([]rune, len(runes))
	copy(ciphertext, runes)

	for i, char := range ciphertext {
		switch {
		case char >= 'a' && char <= 'z':
			ciphertext[i] = 'a' + rune(mod(int(char-'a')+key, engAlfLen))
		case char >= 'A' && char <= 'Z':
			ciphertext[i] = 'A' + rune(mod(int(char-'A')+key, engAlfLen))
		}
	}

	return ciphertext
}

const ukrAlfLen = 33

var ukrU = []rune("АБВГҐДЕЄЖЗИІЇЙКЛМНОПРСТУФХЦЧШЩЬЮЯ")
var ukrL = []rune("абвгґдеєжзиіїйклмнопрстуфхцчшщьюя")

var ukrLetterToIdx = map[rune]struct {
	Idx     int
	IsUpper bool
}{
	'А': {0, true},
	'Б': {1, true},
	'В': {2, true},
	'Г': {3, true},
	'Ґ': {4, true},
	'Д': {5, true},
	'Е': {6, true},
	'Є': {7, true},
	'Ж': {8, true},
	'З': {9, true},
	'И': {10, true},
	'І': {11, true},
	'Ї': {12, true},
	'Й': {13, true},
	'К': {14, true},
	'Л': {15, true},
	'М': {16, true},
	'Н': {17, true},
	'О': {18, true},
	'П': {19, true},
	'Р': {20, true},
	'С': {21, true},
	'Т': {22, true},
	'У': {23, true},
	'Ф': {24, true},
	'Х': {25, true},
	'Ц': {26, true},
	'Ч': {27, true},
	'Ш': {28, true},
	'Щ': {29, true},
	'Ь': {30, true},
	'Ю': {31, true},
	'Я': {32, true},
	'а': {0, false},
	'б': {1, false},
	'в': {2, false},
	'г': {3, false},
	'ґ': {4, false},
	'д': {5, false},
	'е': {6, false},
	'є': {7, false},
	'ж': {8, false},
	'з': {9, false},
	'и': {10, false},
	'і': {11, false},
	'ї': {12, false},
	'й': {13, false},
	'к': {14, false},
	'л': {15, false},
	'м': {16, false},
	'н': {17, false},
	'о': {18, false},
	'п': {19, false},
	'р': {20, false},
	'с': {21, false},
	'т': {22, false},
	'у': {23, false},
	'ф': {24, false},
	'х': {25, false},
	'ц': {26, false},
	'ч': {27, false},
	'ш': {28, false},
	'щ': {29, false},
	'ь': {30, false},
	'ю': {31, false},
	'я': {32, false},
}

func caesarUkr(plaintext *string, key int) []rune {
    runes := []rune(*plaintext)
    ciphertext := make([]rune, len(runes))
	copy(ciphertext, runes)

	for i, rune := range ciphertext {
		if char, in := ukrLetterToIdx[rune]; in {
			enc := mod(char.Idx+key, ukrAlfLen)
			if char.IsUpper {
				ciphertext[i] = ukrU[enc]
			} else {
				ciphertext[i] = ukrL[enc]
			}
		}
	}

	return ciphertext
}

type CaesarInput struct {
	Plaintext string
	Key       int
}

type English CaesarInput
type Ukrainian CaesarInput

func (i *English) encrypt() string {
	return string(caesar(&i.Plaintext, i.Key))
}

func (i *English) decrypt() string {
	return string(caesar(&i.Plaintext, -i.Key))
}

func (i *Ukrainian) encrypt() string {
	return string(caesarUkr(&i.Plaintext, i.Key))
}

func (i *Ukrainian) decrypt() string {
	return string(caesarUkr(&i.Plaintext, -i.Key))
}

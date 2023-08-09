package transform

var pzpzpz = map[byte]byte{
	'a': '=',
	'b': '2',
	'c': 'x',
	'd': 'x',
	'e': '?',
	'f': '1',
	'g': '{',
	'h': 'P',
	'i': '-',
	'j': 'Z',
	'k': 'D',
	'l': 'H',
	'm': '4',
	'n': 'Q',
	'o': '+',
	'p': 'G',
	'q': 'J',
	'r': 'R',
	's': '7',
	't': 'S',
	'u': '_',
	'v': 'T',
	'w': 'W',
	'x': 'd',
	'y': 'L',
	'z': 'z',
	'A': ':',
	'B': 'B',
	'C': '9',
	'D': 'k',
	'E': '}',
	'F': '8',
	'G': 'p',
	'H': 'l',
	'I': '(',
	'J': 'q',
	'K': '/',
	'L': 'l',
	'M': 'M',
	'N': 'N',
	'O': ',',
	'P': 'h',
	'Q': 'n',
	'R': 'r',
	'S': 't',
	'T': 'v',
	'U': '.',
	'V': 'V',
	'W': 'w',
	'X': 'X',
	'Y': '3',
	'Z': 'j',
	' ': '0',
	',': 'O',
	'/': 'K',
	'}': 'E',
	'{': 'g',
	'(': 'I',
	')': '6',
	'.': 'U',
	';': '5',
	':': 'A',
	'?': 'e',
	'=': 'a',
	'+': 'o',
	'-': 'i',
	'_': 'u',
	'0': ' ',
	'1': 'f',
	'2': 'b',
	'3': 'Y',
	'4': 'm',
	'5': ';',
	'6': ')',
	'7': 's',
	'8': 'F',
	'9': 'C',
}

func Transform(plainText string) string {
	encrypted := make([]byte, len(plainText))

	for i := 0; i < len(plainText); i++ {
		char := plainText[i]
		if replacement, ok := pzpzpz[char]; ok {
			encrypted[i] = replacement
		} else {
			encrypted[i] = char
		}
	}

	return string(encrypted)
}

package main

import (
	"fmt"
	"net/http"
)

var utf8GsmChars = ""

func UTF8ToGSM7() {
	fmt.Println("converting")
}

func (a *App) InitExperimental3() {
	exp := a.Router.PathPrefix("/exp2").Subrouter()
	exp.HandleFunc("/utftogsm", a.utftogsm).Methods("GET")
}

func (a *App) utftogsm(w http.ResponseWriter, r *http.Request) {
	// s := "@"
	s := "}"
	fmt.Println(IsValidGSM(s))

	respondWithJSON(w, 200, "sure")

}

func IsValidGSM(text string) (bool, int) {
	var tot int
	// var str []string
	var str string
	for _, r := range text {
		v, ok := baseGSM7[r]
		if ok {
			str += fmt.Sprintf("%x", v)
			tot++
		} else {
			v, ok = extendedGSM7[r]
			if ok {
				str += fmt.Sprintf("%x%x", esc, v)
				tot += 2
			} else {
				return false, tot
			}
		}
	}
	fmt.Println(str)
	return true, tot
}

const esc byte = 0x1B

// https://github.com/nyaruka/courier/blob/master/gsm7/gsm7.go
// base gsm7 characters in our normal table
var baseGSM7 = map[rune]byte{
	'@':  0x00,
	'£':  0x01,
	'$':  0x02,
	'¥':  0x03,
	'è':  0x04,
	'é':  0x05,
	'ù':  0x06,
	'ì':  0x07,
	'ò':  0x08,
	'Ç':  0x09,
	'\n': 0x0A,
	'Ø':  0x0B,
	'ø':  0x0C,
	'\r': 0x0D,
	'Å':  0x0E,
	'å':  0x0F,
	'Δ':  0x10,
	'_':  0x11,
	'Φ':  0x12,
	'Γ':  0x13,
	'Λ':  0x14,
	'Ω':  0x15,
	'Π':  0x16,
	'Ψ':  0x17,
	'Σ':  0x18,
	'Θ':  0x19,
	'Ξ':  0x1A,
	// 'ESC':      0x1B, // Escape control
	'Æ':  0x1C,
	'æ':  0x1D,
	'ß':  0x1E,
	'É':  0x1F,
	' ':  0x20,
	'!':  0x21,
	'"':  0x22,
	'#':  0x23,
	'¤':  0x24,
	'%':  0x25,
	'&':  0x26,
	'\'': 0x27,
	'(':  0x28,
	')':  0x29,
	'*':  0x2A,
	'+':  0x2B,
	',':  0x2C,
	'-':  0x2D,
	'.':  0x2E,
	'/':  0x2F,
	'0':  0x30,
	'1':  0x31,
	'2':  0x32,
	'3':  0x33,
	'4':  0x34,
	'5':  0x35,
	'6':  0x36,
	'7':  0x37,
	'8':  0x38,
	'9':  0x39,
	':':  0x3A,
	';':  0x3B,
	'<':  0x3C,
	'=':  0x3D,
	'>':  0x3E,
	'?':  0x3F,
	'¡':  0x40,
	'A':  0x41,
	'B':  0x42,
	'C':  0x43,
	'D':  0x44,
	'E':  0x45,
	'F':  0x46,
	'G':  0x47,
	'H':  0x48,
	'I':  0x49,
	'J':  0x4A,
	'K':  0x4B,
	'L':  0x4C,
	'M':  0x4D,
	'N':  0x4E,
	'O':  0x4F,
	'P':  0x50,
	'Q':  0x51,
	'R':  0x52,
	'S':  0x53,
	'T':  0x54,
	'U':  0x55,
	'V':  0x56,
	'W':  0x57,
	'X':  0x58,
	'Y':  0x59,
	'Z':  0x5A,
	'Ä':  0x5B,
	'Ö':  0x5C,
	'Ñ':  0x5D,
	'Ü':  0x5E,
	'§':  0x5F,
	'¿':  0x60,
	'a':  0x61,
	'b':  0x62,
	'c':  0x63,
	'd':  0x64,
	'e':  0x65,
	'f':  0x66,
	'g':  0x67,
	'h':  0x68,
	'i':  0x69,
	'j':  0x6A,
	'k':  0x6B,
	'l':  0x6C,
	'm':  0x6D,
	'n':  0x6E,
	'o':  0x6F,
	'p':  0x70,
	'q':  0x71,
	'r':  0x72,
	's':  0x73,
	't':  0x74,
	'u':  0x75,
	'v':  0x76,
	'w':  0x77,
	'x':  0x78,
	'y':  0x79,
	'z':  0x7A,
	'ä':  0x7B,
	'ö':  0x7C,
	'ñ':  0x7D,
	'ü':  0x7E,
	'à':  0x7F,
}

// extended gsm7 characters, these my be preceded by our escape
var extendedGSM7 = map[rune]byte{
	' ':  0x0A,
	'^':  0x14,
	'{':  0x28,
	'}':  0x29,
	'\\': 0x2F,
	'[':  0x3C,
	'~':  0x3D,
	']':  0x3E,
	'|':  0x40,
	'€':  0x65,
}

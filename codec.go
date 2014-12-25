package gopos

import (
	"errors"
	// "fmt"
	"encoding/hex"
	"strconv"
)

var BadDigit = errors.New("Bad digit in BCD decoding")
var Overflow = errors.New("Overflow occurred in BCD decoding")
var UnknownError = errors.New("Unknown Error")

type codec struct {
	bcd bool
}

//FB_AMOUNT FA_AMOUNT
func (c codec) amount(val int) ([]byte, error) {
	if c.bcd {
		return dec2bcd(uint64(val))
	}
	return []byte(strconv.Itoa(val)), nil
}

//FB_BINARY FA_BINARY
func (c codec) binary(val []byte) ([]byte, error) {
	return nil, nil
}

//FB_BITMAP FA_BITMAP
func (c codec) bitmap(b []byte) ([]byte, error) {
	return nil, nil
}

//FB_LLBINARY FA_LLBINARY
func (c codec) llbinary(b []byte) ([]byte, error) {
	return nil, nil
}

//FB_LLLBINARY FA_LLLBINARY
func (c codec) lllbinary(b []byte) ([]byte, error) {
	return nil, nil
}

//FB_LLCHAR FA_LLCHAR
func (c codec) llchar(b []byte) ([]byte, error) {
	return nil, nil
}

//FB_LLLCHAR FA_LLLCHAR
func (c codec) lllchar(b []byte) ([]byte, error) {
	return nil, nil
}

//FB_LLNUM FA_LLNUM
func (c codec) llnumeric(i int) ([]byte, error) {
	return nil, nil
}

//FB_NUMERIC FA_NUMERIC
func (c codec) numeric(i int) ([]byte, error) {
	return nil, nil
}

//F_CHAR
func (c codec) char(l int, v string) ([]byte, error) {
	return nil, nil
}

/**
[+ASCII_Number+]      encodes either a Number or String representation of ASCII_Number
                      a number to the ASCII represenation of the number, Packed_Number
                      decodes ASCII  numerals to a numberA_Codec
[+A_Codec+]           passes through ASCII string checking they conform to [A-Za-z]AN_Codec
                      during encoding, no validity check during decoding. ANP_Codec
[+AN_Codec+]          passes through ASCII string checking they conform to [A-Za-z0-9]ANS_Codec
                      during encoding, no validity check during decoding. Null_Codec
[+ANP_Codec+]         passes through ASCII string checking they conform to [A-Za-z0-9 ] Track2
                      during encoding, no validity check during decoding.
[+ANS_Codec+]         passes through ASCII string checking they conform to [\x20-\x7E]
                      during encoding, no validity check during decoding.
[+Null_Codec+]        passes anything along untouched.
[<tt>Track2</tt>]     rudimentary check that string conforms to Track2
[+MMDDhhmmssCodec+]   encodes Time, Datetime or String to the described date format, checking
                      that it is a valid date. Decodes to a DateTime instance, decoding and
                      encoding perform validity checks!
[+MMDDCodec+]   encodes Time, Datetime or String to the described date format, checking
                      that it is a valid date. Decodes to a DateTime instance, decoding and
                      encoding perform validity checks!
[+YYMMDDhhmmssCodec+] encodes Time, Datetime or String to the described date format, checking
                      that it is a valid date. Decodes to a DateTime instance, decoding and
                      encoding perform validity checks!
[+YYMMCodec+]         encodes Time, Datetime or String to the described date format (exp date),
                      checking that it is a valid date. Decodes to a DateTime instance, decoding
                      and encoding perform validity checks!

*/

func dec2bcd(valInt uint64) ([]byte, error) {
	var rb []byte
	valStr := strconv.Itoa(int(valInt))
	if len(valStr)%2 != 0 {
		valStr = "0" + valStr
	}
	for i := 0; i < len(valStr); i += 2 {
		hv, err := strconv.ParseInt(valStr[i:i+2], 16, 16)
		if err != nil {
			return nil, err
		}
		rb = append(rb, byte(hv))
	}
	return rb, nil
}

func bcd2dec(bcd []byte) (uint64, error) {
	ri, err := strconv.ParseInt(hex.EncodeToString(bcd), 10, 64)
	if err != nil {
		return 0, err
	}
	return uint64(ri), nil
}

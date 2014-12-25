package gopos

import (
	"encoding/hex"
	"fmt"
	"testing"
)

func TestBCDDEC(t *testing.T) {
	tv := []uint64{
		1, 2, 3,
		57849, 90471, 30193,
		791707123659,
		923456781917981214,
		45012199999086618,
	}
	e := func(err error) {
		if err != nil {
			t.Fatal(err)
		}
	}
	for _, v := range tv {
		bcd, err := dec2bcd(v)
		e(err)
		fmt.Println("BCD: ", hex.EncodeToString(bcd))
		dec, err := bcd2dec(bcd)
		e(err)
		fmt.Println("DEC: ", dec)

	}
}

func BenchmarkDEC2BCD(b *testing.B) {
	e := func(err error) {
		if err != nil {
			b.Fatal(err)
		}
	}
	var intV uint64 = 923456781917981214
	for i := 0; i < b.N; i++ {
		_, err := dec2bcd(intV)
		e(err)
	}
}

func BenchmarkBCD2DEC(b *testing.B) {
	e := func(err error) {
		if err != nil {
			b.Fatal(err)
		}
	}
	var intV uint64 = 923456781917981214

	bcd, err := dec2bcd(intV)
	e(err)
	for i := 0; i < b.N; i++ {
		dec, err := bcd2dec(bcd)
		e(err)
		if dec != intV {
      // e(fmt.Errorf("dec don't matching"))
		}
	}
}

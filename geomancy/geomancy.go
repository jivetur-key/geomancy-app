package geomancy

import (
	"crypto/rand"
	"encoding/binary"
	"fmt"
)

/*
  The geomancy reading is encoded into a unsigned 64 bit intger.
*/

type Geomancy struct {
	value uint64
}

/*
Const variables to representing bits. To hold the place of the figures in
the reading Nibble the mothers, Byte the mothers and daughters
Tribble the mother, daughters and resultants, Word the entire reading .
*/
const (
	Mask    = 0xF
	Nibble  = 4
	Byte    = 8
	Tribble = 12
	Word    = 16
)

/*
Key of the figure names in decimal order 0 - 15.
*/
var figures = []string{
	"Populus", "Tristitia", "Albus", "Fortune Major", "Rubeus",
	"Acquisitio", "Conjunctio", "Caput Draconis", "Laetitia",
	"Carcer", "Amissio", "Puella", "Fortuna Minor", "Puer", "Cauda Draconis", "Via",
}

func New() *Geomancy {
	return &Geomancy{}
}

/*
Private sets the value at the index offsetted by four bits.
*/
func (geo *Geomancy) set(index int, value uint64) {
	geo.value |= value << (index * Nibble)
}

/*
Returns the decimal value at the index shifted 4 bits.
*/
func (geo *Geomancy) get(index int) uint64 {
	return geo.value >> (index * Nibble) & Mask
}

/*
Returns the value of index shifted 4 bit sas string in a 4 bit binary format.
*/
func (geo *Geomancy) String(index int) string {
	value := geo.get(index)
	str := fmt.Sprintf("%04b", value)
	return str
}

func (geo *Geomancy) Name(index int) string {
	return figures[geo.get(index)]
}

/*
Simulates a geomancy reading and encodes into a 64 bit unsigned integer.
Uses 16 random even or odd numbers as the first four figures
*/
func (geo *Geomancy) Generate() error {

	/*
	   Generate the four  mother figures each figure is four bits
	   odd random numbers will set the bits
	*/
	for index := range Nibble {
		var value uint64
		for pos := Nibble - 1; pos >= 0; pos-- {
			var randomNumber uint64
			err := binary.Read(rand.Reader,
				binary.LittleEndian, &randomNumber)
			if err != nil {
				return err
			} else if randomNumber%2 != 0 {
				value |= (1 << pos)
			}
		}
		geo.set(index, value)
	}

	/*
	   Generate the four daughter, the daughters are a transpose of
	   the mothers
	*/

	for x1, x2 := Nibble-1, 0; x1 >= 0; x1, x2 = x1-1, x2+1 {
		var value uint64
		for y1, y2 := 0, Nibble-1; y2 >= 0; y1, y2 = y1+1, y2-1 {
			n := geo.get(y1)
			if (n & (1 << x1)) != 0 {
				value |= (1 << y2)
			}
		}
		geo.set(x2+Nibble, value)
	}

	/*
	   Generate the four resultants or nephews, and the right and
	   left witnesses and the judge, which is the outcome.  This is
	   the XOR of two figures.
	*/
	for i, j := 0, 0; j < Word-2; i, j = i+1, j+2 {
		value := geo.get(j) ^ geo.get(j+1)
		geo.set(i+Byte, value)

	}
	/*
	   The optional sentence, is the XOR of the first mother and
	   the judge.
	*/
	sentence := geo.get(0) ^ geo.get(Word-2)
	geo.set(Word-1, sentence)

	return nil
}

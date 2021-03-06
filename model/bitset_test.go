package model

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"math/big"
	"strings"
	"testing"
)

func TestBitset_Decode_Simple(t *testing.T) {
	bitSet := Bitset{
		Int: big.Int{},
		Min: 1,
		Max: 32,
	}

	day := 15

	err := bitSet.Decode(fmt.Sprintf("%d", day))
	assert.NoError(t, err)
	assert.Equal(t, fmt.Sprintf("%d", day), bitSet.String())
}

func TestBitset_Decode_List(t *testing.T) {
	bitSet := Bitset{
		Int: big.Int{},
		Min: 1,
		Max: 32,
	}

	start := 1
	end := 15
	list := make([]string, 0)

	for start <= end {
		list = append(list, fmt.Sprintf("%d", start))
		start += 3
	}

	listString := strings.Join(list, ",")
	err := bitSet.Decode(listString)
	assert.NoError(t, err)
	assert.Equal(t, listString, bitSet.String())
}

func TestBitset_Decode_Range(t *testing.T) {
	bitSet := Bitset{
		Int: big.Int{},
		Min: 1,
		Max: 32,
	}

	start := 1
	end := 15
	days := make([]string, 0)

	for start <= end {
		days = append(days, fmt.Sprintf("%d", start))
		start++
	}

	err := bitSet.Decode("1-15")
	assert.NoError(t, err)
	assert.Equal(t, strings.Join(days, ","), bitSet.String())
}

func TestBitset_Decode_Full(t *testing.T) {
	bitSet := Bitset{
		Int: big.Int{},
		Min: 1,
		Max: 32,
	}

	start := 1
	end := 31
	list := make([]string, 0)

	for start <= end {
		list = append(list, fmt.Sprintf("%d", start))
		start++
	}

	listString := strings.Join(list, ",")
	err := bitSet.Decode("*")
	assert.NoError(t, err)
	assert.Equal(t, listString, bitSet.String())
}

func TestBitset_Decode_Steps(t *testing.T) {
	bitSet := Bitset{
		Int: big.Int{},
		Min: 1,
		Max: 32,
	}

	start := 1
	end := 31
	step := 3
	list := make([]string, 0)

	for start <= end {
		list = append(list, fmt.Sprintf("%d", start))
		start += step
	}
	listString := strings.Join(list, ",")

	err := bitSet.Decode(fmt.Sprintf("*/%d", step))
	assert.NoError(t, err)
	assert.Equal(t, listString, bitSet.String())
}

func TestBitset_String(t *testing.T) {
	bitSet := Bitset{
		Int: big.Int{},
		Min: 1,
		Max: 32,
	}
	assert.Equal(t, "", bitSet.String())

	bitSet.SetBit(&bitSet.Int, 1, 1)
	bitSet.SetBit(&bitSet.Int, 15, 1)
	bitSet.SetBit(&bitSet.Int, 30, 1)
	assert.Equal(t, "1,15,30", bitSet.String())
}

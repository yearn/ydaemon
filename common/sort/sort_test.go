package sort

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

type TStructNestedNested struct {
	String string  `json:"string"`
	Int    int     `json:"int"`
	Uint   uint64  `json:"uint"`
	Float  float64 `json:"float"`
}
type TStructNested struct {
	String string               `json:"string"`
	Int    int                  `json:"int"`
	Uint   uint64               `json:"uint"`
	Float  float64              `json:"float"`
	Nested *TStructNestedNested `json:"nested"`
}
type TStruct struct {
	String string  `json:"string"`
	Int    int     `json:"int"`
	Uint   uint64  `json:"uint"`
	Float  float64 `json:"float"`
	Bool   bool
	Nested TStructNested `json:"nested"`
}

func PrepareTStruct() []TStruct {
	tStruct := []TStruct{}
	for i := 0; i < 10; i++ {
		tStruct = append(tStruct, TStruct{
			String: strconv.Itoa(i),
			Int:    i,
			Uint:   uint64(i),
			Float:  float64(i),
			Bool:   i%2 == 0,
			Nested: TStructNested{
				String: strconv.Itoa(10 - i),
				Int:    10 - i,
				Uint:   uint64(10) - uint64(i),
				Float:  float64(10 - i),
				Nested: &TStructNestedNested{
					String: strconv.Itoa(i),
					Int:    i,
					Uint:   uint64(i),
					Float:  float64(i),
				},
			},
		})
	}
	return tStruct
}

func TestSortBy(t *testing.T) {
	{
		sortedData := PrepareTStruct()

		SortBy("string", "asc", sortedData)
		assert.Equal(t, "0", sortedData[0].String)
		assert.Equal(t, "2", sortedData[2].String)
		assert.Equal(t, "9", sortedData[9].String)

		SortBy("nested.int", "asc", sortedData)
		assert.Equal(t, 9, sortedData[8].Nested.Int)
		assert.Equal(t, 6, sortedData[5].Nested.Int)
		assert.Equal(t, 1, sortedData[0].Nested.Int)

		SortBy("nested.nested.uint", "asc", sortedData)
		assert.Equal(t, uint64(6), sortedData[6].Nested.Nested.Uint)
		assert.Equal(t, uint64(7), sortedData[7].Nested.Nested.Uint)
		assert.Equal(t, uint64(8), sortedData[8].Nested.Nested.Uint)

		SortBy("nested.nested.uint", "asc", sortedData)
		assert.Equal(t, float64(1), sortedData[1].Nested.Nested.Float)
		assert.Equal(t, float64(2), sortedData[2].Nested.Nested.Float)
		assert.Equal(t, float64(4), sortedData[4].Nested.Nested.Float)

		SortBy("bool", "asc", sortedData)
		assert.Equal(t, false, sortedData[0].Bool)
		assert.Equal(t, false, sortedData[1].Bool)
		assert.Equal(t, false, sortedData[2].Bool)
		assert.Equal(t, false, sortedData[3].Bool)
		assert.Equal(t, false, sortedData[4].Bool)
		assert.Equal(t, true, sortedData[5].Bool)
		assert.Equal(t, true, sortedData[6].Bool)
		assert.Equal(t, true, sortedData[7].Bool)
		assert.Equal(t, true, sortedData[8].Bool)
		assert.Equal(t, true, sortedData[9].Bool)
	}

	{
		sortedData := PrepareTStruct()

		SortBy("string", "desc", sortedData)
		assert.Equal(t, "0", sortedData[9].String)
		assert.Equal(t, "2", sortedData[7].String)
		assert.Equal(t, "9", sortedData[0].String)

		SortBy("nested.int", "desc", sortedData)
		assert.Equal(t, 2, sortedData[8].Nested.Int)
		assert.Equal(t, 5, sortedData[5].Nested.Int)
		assert.Equal(t, 10, sortedData[0].Nested.Int)

		SortBy("nested.nested.uint", "desc", sortedData)
		assert.Equal(t, float64(3), sortedData[6].Nested.Nested.Float)
		assert.Equal(t, float64(2), sortedData[7].Nested.Nested.Float)
		assert.Equal(t, float64(1), sortedData[8].Nested.Nested.Float)

		SortBy("bool", "desc", sortedData)
		assert.Equal(t, true, sortedData[0].Bool)
		assert.Equal(t, true, sortedData[1].Bool)
		assert.Equal(t, true, sortedData[2].Bool)
		assert.Equal(t, true, sortedData[3].Bool)
		assert.Equal(t, true, sortedData[4].Bool)
		assert.Equal(t, false, sortedData[5].Bool)
		assert.Equal(t, false, sortedData[6].Bool)
		assert.Equal(t, false, sortedData[7].Bool)
		assert.Equal(t, false, sortedData[8].Bool)
		assert.Equal(t, false, sortedData[9].Bool)
	}

	{
		sortedData := PrepareTStruct()

		SortBy("blabla", "desc", sortedData)
		assert.Equal(t, sortedData[0].String, sortedData[0].String)
		assert.Equal(t, sortedData[7].Int, sortedData[7].Int)

		SortBy("nested.bla", "asc", sortedData)
		assert.Equal(t, sortedData[0].Float, sortedData[0].Float)
		assert.Equal(t, sortedData[7].Nested.Nested.Uint, sortedData[7].Nested.Nested.Uint)
	}
}

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
		tStruct := PrepareTStruct()
		var sortedData []interface{} = make([]interface{}, len(tStruct))
		for i, d := range tStruct {
			sortedData[i] = d
		}

		SortBy(sortedData, "string", "asc")
		assert.Equal(t, "0", sortedData[0].(TStruct).String)
		assert.Equal(t, "2", sortedData[2].(TStruct).String)
		assert.Equal(t, "9", sortedData[9].(TStruct).String)

		SortBy(sortedData, "nested.int", "asc")
		assert.Equal(t, 9, sortedData[8].(TStruct).Nested.Int)
		assert.Equal(t, 6, sortedData[5].(TStruct).Nested.Int)
		assert.Equal(t, 1, sortedData[0].(TStruct).Nested.Int)

		SortBy(sortedData, "nested.nested.uint", "asc")
		assert.Equal(t, uint64(6), sortedData[6].(TStruct).Nested.Nested.Uint)
		assert.Equal(t, uint64(7), sortedData[7].(TStruct).Nested.Nested.Uint)
		assert.Equal(t, uint64(8), sortedData[8].(TStruct).Nested.Nested.Uint)

		SortBy(sortedData, "nested.nested.uint", "asc")
		assert.Equal(t, float64(1), sortedData[1].(TStruct).Nested.Nested.Float)
		assert.Equal(t, float64(2), sortedData[2].(TStruct).Nested.Nested.Float)
		assert.Equal(t, float64(4), sortedData[4].(TStruct).Nested.Nested.Float)

		SortBy(sortedData, "bool", "asc")
		assert.Equal(t, false, sortedData[0].(TStruct).Bool)
		assert.Equal(t, false, sortedData[1].(TStruct).Bool)
		assert.Equal(t, false, sortedData[2].(TStruct).Bool)
		assert.Equal(t, false, sortedData[3].(TStruct).Bool)
		assert.Equal(t, false, sortedData[4].(TStruct).Bool)
		assert.Equal(t, true, sortedData[5].(TStruct).Bool)
		assert.Equal(t, true, sortedData[6].(TStruct).Bool)
		assert.Equal(t, true, sortedData[7].(TStruct).Bool)
		assert.Equal(t, true, sortedData[8].(TStruct).Bool)
		assert.Equal(t, true, sortedData[9].(TStruct).Bool)
	}

	{
		tStruct := PrepareTStruct()
		var sortedData []interface{} = make([]interface{}, len(tStruct))
		for i, d := range tStruct {
			sortedData[i] = d
		}

		SortBy(sortedData, "string", "desc")
		assert.Equal(t, "0", sortedData[9].(TStruct).String)
		assert.Equal(t, "2", sortedData[7].(TStruct).String)
		assert.Equal(t, "9", sortedData[0].(TStruct).String)

		SortBy(sortedData, "nested.int", "desc")
		assert.Equal(t, 2, sortedData[8].(TStruct).Nested.Int)
		assert.Equal(t, 5, sortedData[5].(TStruct).Nested.Int)
		assert.Equal(t, 10, sortedData[0].(TStruct).Nested.Int)

		SortBy(sortedData, "nested.nested.uint", "desc")
		assert.Equal(t, float64(3), sortedData[6].(TStruct).Nested.Nested.Float)
		assert.Equal(t, float64(2), sortedData[7].(TStruct).Nested.Nested.Float)
		assert.Equal(t, float64(1), sortedData[8].(TStruct).Nested.Nested.Float)

		SortBy(sortedData, "bool", "desc")
		assert.Equal(t, true, sortedData[0].(TStruct).Bool)
		assert.Equal(t, true, sortedData[1].(TStruct).Bool)
		assert.Equal(t, true, sortedData[2].(TStruct).Bool)
		assert.Equal(t, true, sortedData[3].(TStruct).Bool)
		assert.Equal(t, true, sortedData[4].(TStruct).Bool)
		assert.Equal(t, false, sortedData[5].(TStruct).Bool)
		assert.Equal(t, false, sortedData[6].(TStruct).Bool)
		assert.Equal(t, false, sortedData[7].(TStruct).Bool)
		assert.Equal(t, false, sortedData[8].(TStruct).Bool)
		assert.Equal(t, false, sortedData[9].(TStruct).Bool)
	}

	{
		tStruct := PrepareTStruct()
		var sortedData []interface{} = make([]interface{}, len(tStruct))
		for i, d := range tStruct {
			sortedData[i] = d
		}

		SortBy(sortedData, "blabla", "desc")
		assert.Equal(t, tStruct[0].String, sortedData[0].(TStruct).String)
		assert.Equal(t, tStruct[7].Int, sortedData[7].(TStruct).Int)

		SortBy(sortedData, "nested.bla", "asc")
		assert.Equal(t, tStruct[0].Float, sortedData[0].(TStruct).Float)
		assert.Equal(t, tStruct[7].Nested.Nested.Uint, sortedData[7].(TStruct).Nested.Nested.Uint)
	}
}

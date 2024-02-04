package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_getCountMapAndTotalCount(t *testing.T) {
	tests := []struct {
		name         string
		phrase       string
		wantCountMap map[rune]float64
		wantTotal    float64
	}{
		{
			name:   "Simple English Phrase",
			phrase: "hello world",
			wantCountMap: map[rune]float64{
				'h': 1,
				'e': 1,
				'l': 3,
				'o': 2,
				'w': 1,
				'r': 1,
				'd': 1,
			},
			wantTotal: 10,
		},
		{
			name:   "Standard English Phrase",
			phrase: "the table shows the cursor values (created_at, id)",
			wantCountMap: map[rune]float64{
				't': 5,
				'h': 3,
				'e': 6,
				'a': 4,
				'b': 1,
				'l': 2,
				's': 4,
				'o': 2,
				'w': 1,
				'c': 2,
				'u': 2,
				'r': 3,
				'v': 1,
				'd': 2,
				'i': 1,
			},
			wantTotal: 39,
		},
		{
			name:   "Simple Russian Phrase",
			phrase: "мама мыла раму",
			wantCountMap: map[rune]float64{
				'м': 4,
				'а': 4,
				'ы': 1,
				'л': 1,
				'р': 1,
				'у': 1,
			},
			wantTotal: 12,
		},
		{
			name:   "Standard Russian Phrase",
			phrase: "например, давайте выведем поле из таблицы, в которой организовано хранение информации о расписании",
			wantCountMap: map[rune]float64{
				'н': 7,
				'а': 10,
				'п': 3,
				'р': 7,
				'и': 11,
				'м': 3,
				'е': 7,
				'д': 2,
				'в': 5,
				'й': 2,
				'т': 3,
				'ы': 2,
				'о': 9,
				'л': 2,
				'з': 2,
				'б': 1,
				'ц': 2,
				'к': 1,
				'г': 1,
				'х': 1,
				'ф': 1,
				'с': 2,
			},
			wantTotal: 84,
		},
		{
			name:   "Simple Chineeze Phrase",
			phrase: "你好我親愛的朋友",
			wantCountMap: map[rune]float64{
				'你': 1,
				'好': 1,
				'我': 1,
				'親': 1,
				'愛': 1,
				'的': 1,
				'朋': 1,
				'友': 1,
			},
			wantTotal: 8,
		},
		{
			name:   "Only digits and special characters",
			phrase: "/7@|_|_|@",
			wantCountMap: map[rune]float64{
			},
			wantTotal: 0,
		},		
		{
			name:   "Empty Phrase",
			phrase: "",
			wantCountMap: map[rune]float64{
			},
			wantTotal: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			countMap, total := getCountMapAndTotalCount(tt.phrase)
			assert.Equal(t, tt.wantCountMap, countMap)
			assert.Equal(t, tt.wantTotal, total)
		})
	}
}

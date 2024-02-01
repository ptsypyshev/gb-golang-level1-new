package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_getFileNameAndExt(t *testing.T) {
	tests := []struct {
		name         string
		separator    string
		filePath     string
		wantFileName string
		wantExt      string
	}{
		{
			name:         "Unix standard file",
			filePath:     "/home/robotomize/1.txt",
			separator:    "/",
			wantFileName: "1",
			wantExt:      "txt",
		},
		{
			name:         "Unix file with dots in filename",
			filePath:     "/home/robotomize/1.txt.txt",
			separator:    "/",
			wantFileName: "1.txt",
			wantExt:      "txt",
		},
		{
			name:         "Unix file without extension",
			filePath:     "/home/robotomize/1",
			separator:    "/",
			wantFileName: "1",
			wantExt:      "",
		},		
		{
			name:         "Unix hidden file without extension",
			filePath:     "/home/robotomize/.env",
			separator:    "/",
			wantFileName: ".env",
			wantExt:      "",
		},
		{
			name:         "Windows standard file",
			filePath:     `c:\Users\pavel\1.txt`,
			separator:    `\`,
			wantFileName: "1",
			wantExt:      "txt",
		},
		{
			name:         "Windows file with dots in filename",
			filePath:     `c:\Users\pavel\1.txt.txt`,
			separator:    `\`,
			wantFileName: "1.txt",
			wantExt:      "txt",
		},
		{
			name:         "Windows file without extension",
			filePath:     `c:\Users\pavel\1`,
			separator:    `\`,
			wantFileName: "1",
			wantExt:      "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fileName, fileExt := getFileNameAndExt(tt.filePath, tt.separator)
			assert.Equal(t, tt.wantFileName, fileName)
			assert.Equal(t, tt.wantExt, fileExt)
		})
	}
}

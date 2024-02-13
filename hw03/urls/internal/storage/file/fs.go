package file

import (
	"context"
	"encoding/json"
	"io/fs"
	"os"

	"github.com/ptsypyshev/gb-golang-level1-new/hw03/urls/internal/app"
	"github.com/ptsypyshev/gb-golang-level1-new/hw03/urls/internal/models"
	"github.com/ptsypyshev/gb-golang-level1-new/hw03/urls/internal/storage/memory"
)

var _ app.Storage = (*FileStor)(nil)

// FileStor is an implementation of app.Storage interface that stores data in JSON file.
type FileStor struct {
	memory.MemStor
	filePath string
}

// New is a constructor for FileStor.
func New(m *memory.MemStor, f string) *FileStor {
	return &FileStor{
		MemStor:  *m,
		filePath: f,
	}
}

// Close implements app.Storage.
func (f *FileStor) Close(ctx context.Context) error {
	return f.Save()
}

// Load reads JSON-file with data to urls map.
func Load(f string) (*memory.MemStor, error) {
	bytes, err := os.ReadFile(f)
	if err != nil {
		return nil, err
	}

	var res *memory.MemStor
	err = json.Unmarshal(bytes, &res)
	if err != nil {
		return nil, err
	}

	if res.Urls == nil {
		res.Urls = make(map[string]models.URL)
	}
	return res, nil
}

// Save saves JSON-file with data from urls map.
func (f *FileStor) Save() error {
	data, err := json.MarshalIndent(f.MemStor, "", "    ")
	if err != nil {
		return err
	}

	err = os.WriteFile(f.filePath, data, fs.ModePerm)
	if err != nil {
		return err
	}

	return nil
}

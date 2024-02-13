package file

import (
	"context"
	"encoding/json"
	"os"

	"github.com/ptsypyshev/gb-golang-level1-new/hw03/urls/internal/app"
	"github.com/ptsypyshev/gb-golang-level1-new/hw03/urls/internal/models"
	"github.com/ptsypyshev/gb-golang-level1-new/hw03/urls/internal/storage/memory"
)

var _ app.Storage = (*FileStor)(nil)

// FileStor is an implementation of app.Storage interface that stores data in JSON file.
type FileStor struct {
	InMem    app.Storage
	filePath string
}

// New is a constructor for FileStor.
func New(m app.Storage, f string) *FileStor {
	return &FileStor{
		InMem:    m,
		filePath: f,
	}
}

// Add implements app.Storage.
func (f *FileStor) Add(args []string) error {
	return f.InMem.Add(args)
}

// Close implements app.Storage.
func (f *FileStor) Close(ctx context.Context) error {
	return f.Save()
}

// List implements app.Storage.
func (f *FileStor) List() ([]models.URL, error) {
	return f.InMem.List()
}

// Remove implements app.Storage.
func (f *FileStor) Remove(url string) error {
	return f.InMem.Remove(url)
}

// Search implements app.Storage.
func (f *FileStor) Search(t string) ([]models.URL, error) {
	return f.InMem.Search(t)
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
	return res, nil
}

// Save saves JSON-file with data from urls map.
func (f *FileStor) Save() error {
	data, err := json.MarshalIndent(f.InMem, "", "    ")
	if err != nil {
		return err
	}

	err = os.WriteFile(f.filePath, data, os.ModeType)
	if err != nil {
		return err
	}

	return nil
}

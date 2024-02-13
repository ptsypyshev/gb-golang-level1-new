package file

import (
	"context"
	"os"
	"testing"
	"time"

	"github.com/ptsypyshev/gb-golang-level1-new/hw03/urls/internal/models"
	"github.com/ptsypyshev/gb-golang-level1-new/hw03/urls/internal/storage/memory"
	"github.com/stretchr/testify/assert"
)

var (
	data_payload = []byte(`{
		"Urls": {
			"ya.ru": {
				"Description": "desc",
				"Date": "2024-02-13T14:02:17.1945437+05:00",
				"Tags": [
					"1",
					"2"
				],
				"Link": "ya.ru"
			},
			"yandex.com": {
				"Description": "some_some_desc",
				"Date": "2024-02-13T14:03:07.0418499+05:00",
				"Tags": [
					"123",
					"123",
					"444"
				],
				"Link": "yandex.com"
			}
		}
	}`)
	empty_payload = []byte("{}")
	bad_payload   = []byte("it; is; not; json!")

	date1, _ = time.Parse(time.RFC3339Nano, "2024-02-13T14:02:17.1945437+05:00")
	date2, _ = time.Parse(time.RFC3339Nano, "2024-02-13T14:03:07.0418499+05:00")

	ya = models.URL{
		Link:        "ya.ru",
		Description: "desc",
		Tags:        []string{"1", "2"},
		Date:        date1,
	}
	yandex = models.URL{
		Link:        "yandex.com",
		Description: "some_some_desc",
		Tags:        []string{"123", "123", "444"},
		Date:        date2,
	}

	goodFilepath     = "data.json"
	emptyFilepath    = "empty.json"
	badFilepath      = "bad.json"
	nonExistFilepath = "nofile.json"

	goodWriteFilePath = "write.json"
	badWriteFilePath  = os.Getenv("SystemDrive") + string(os.PathSeparator) + "locked.json"

	goodCloseFilePath = "close.json"
)

func TestLoad(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		filepath string
		payload  []byte
		want     *memory.MemStor
		wantErr  bool
	}{
		{
			name:     "Load file with content",
			filepath: goodFilepath,
			payload:  data_payload,
			want: &memory.MemStor{
				Urls: map[string]models.URL{
					"ya.ru":      ya,
					"yandex.com": yandex,
				},
			},
			wantErr: false,
		},
		{
			name:     "Load empty file",
			filepath: emptyFilepath,
			payload:  empty_payload,
			want:     memory.New(),
			wantErr:  false,
		},
		{
			name:     "Load file with invalid content",
			filepath: badFilepath,
			payload:  bad_payload,
			want:     nil,
			wantErr:  true,
		},
		{
			name:     "Load from a non-existent file",
			filepath: nonExistFilepath,
			payload:  nil,
			want:     nil,
			wantErr:  true,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if tt.payload != nil {
				err := os.WriteFile(tt.filepath, tt.payload, 0644)
				assert.NoError(t, err)

				defer func() {
					err = os.Remove(tt.filepath)
					assert.NoError(t, err)
				}()
			}

			got, err := Load(tt.filepath)
			if tt.wantErr {
				assert.Error(t, err)
				return
			}
			assert.NoError(t, err)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestFileStor_Save(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name    string
		f       *FileStor
		wantErr bool
	}{
		{
			name: "Save content to good file",
			f: &FileStor{
				MemStor: memory.MemStor{
					Urls: map[string]models.URL{
						"ya.ru":      ya,
						"yandex.com": yandex,
					},
				},
				filePath: goodWriteFilePath,
			},
			wantErr: false,
		},
		{
			name: "Save content to bad file (no access to file)",
			f: &FileStor{
				MemStor: memory.MemStor{
					Urls: map[string]models.URL{
						"ya.ru":      ya,
						"yandex.com": yandex,
					},
				},
				filePath: badWriteFilePath,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			err := tt.f.Save()
			if tt.wantErr {
				assert.Error(t, err)
				return
			}
			assert.NoError(t, err)

			err = os.Remove(tt.f.filePath)
			assert.NoError(t, err)
		})
	}
}

func TestFileStor_Close(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name    string
		f       *FileStor
		ctx context.Context
		wantErr bool
	}{
		{
			name: "Close FileStor",
			f: &FileStor{
				MemStor: memory.MemStor{
					Urls: map[string]models.URL{
						"ya.ru":      ya,
						"yandex.com": yandex,
					},
				},
				filePath: goodCloseFilePath,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			err := tt.f.Close(tt.ctx)

			assert.NoError(t, err)

			err = os.Remove(tt.f.filePath)
			assert.NoError(t, err)
		})
	}
}

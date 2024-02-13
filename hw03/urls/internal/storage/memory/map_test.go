package memory

import (
	"context"
	"testing"

	"github.com/ptsypyshev/gb-golang-level1-new/hw03/urls/internal/models"
	"github.com/ptsypyshev/gb-golang-level1-new/hw03/urls/internal/storage"
	"github.com/stretchr/testify/assert"
)

func TestMemStor_Add(t *testing.T) {
	t.Parallel()
	m := New()
	tests := []struct {
		name    string
		args    []string
		wantErr error
	}{
		{
			name:    "Successful adding",
			args:    []string{"ya.ru", "some_desc", "tag1", "tag2", "tag3"},
			wantErr: nil,
		},
		{
			name:    "Small args slice",
			args:    []string{"ya.ru", "some_desc"},
			wantErr: storage.ErrBadArguments,
		},
		{
			name:    "Nil args slice",
			args:    nil,
			wantErr: storage.ErrBadArguments,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			err := m.Add(tt.args)
			assert.ErrorIs(t, err, tt.wantErr)
			if tt.wantErr != nil {
				return
			}

			assert.NoError(t, err)
			_, ok := m.Urls[tt.args[0]]
			assert.True(t, ok)
		})
	}
}

func TestMemStor_Close(t *testing.T) {
	t.Parallel()
	m := New()
	assert.NoError(t, m.Close(context.Background()))
}

func TestMemStor_List(t *testing.T) {
	t.Parallel()
	mFull := &MemStor{
		Urls: map[string]models.URL{
			"ya.ru": {
				Link:        "ya.ru",
				Description: "desc_ya",
				Tags:        []string{"111", "222", "333"},
			},
			"google.ru": {
				Link:        "google.ru",
				Description: "desc_google",
				Tags:        []string{"444", "555", "666"},
			},
			"bing.com": {
				Link:        "bing.com",
				Description: "desc_bing",
				Tags:        []string{"777", "888", "999"},
			},
		},
	}

	mEmpty := New()

	tests := []struct {
		name    string
		m       *MemStor
		want    []models.URL
		wantErr error
	}{
		{
			name: "Successful Listing",
			m:    mFull,
			want: []models.URL{
				{
					Link:        "ya.ru",
					Description: "desc_ya",
					Tags:        []string{"111", "222", "333"}},
				{
					Link:        "google.ru",
					Description: "desc_google",
					Tags:        []string{"444", "555", "666"}},
				{
					Link:        "bing.com",
					Description: "desc_bing",
					Tags:        []string{"777", "888", "999"},
				},
			},
			wantErr: nil,
		},
		{
			name:    "Empty storage",
			m:       mEmpty,
			want:    nil,
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got, err := tt.m.List()
			assert.ElementsMatch(t, got, tt.want)
			assert.ErrorIs(t, err, tt.wantErr)
		})
	}
}

func TestMemStor_Remove(t *testing.T) {
	t.Parallel()
	mFull := &MemStor{
		Urls: map[string]models.URL{
			"ya.ru": {
				Link:        "ya.ru",
				Description: "desc_ya",
				Tags:        []string{"111", "222", "333"},
			},
			"google.ru": {
				Link:        "google.ru",
				Description: "desc_google",
				Tags:        []string{"444", "555", "666"},
			},
			"bing.com": {
				Link:        "bing.com",
				Description: "desc_bing",
				Tags:        []string{"777", "888", "999"},
			},
		},
	}

	mEmpty := New()

	tests := []struct {
		name    string
		m       *MemStor
		url     string
		wantErr error
	}{
		{
			name:    "Successful Removing",
			m:       mFull,
			url:     "ya.ru",
			wantErr: nil,
		},
		{
			name:    "Element not found",
			m:       mFull,
			url:     "yandex.ru",
			wantErr: storage.ErrNotFound,
		},
		{
			name:    "Removing from empty storage",
			m:       mEmpty,
			url:     "ya.ru",
			wantErr: storage.ErrNotFound,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			err := tt.m.Remove(tt.url)
			assert.ErrorIs(t, err, tt.wantErr)
		})
	}
}

func TestMemStor_Search(t *testing.T) {
	t.Parallel()
	var (
		yandex = models.URL{
			Link:        "yandex.ru",
			Description: "desc_yandex",
			Tags:        []string{"111", "222", "333"},
		}
		google = models.URL{
			Link:        "google.ru",
			Description: "desc_google_yandex",
			Tags:        []string{"444", "555", "666"},
		}
		bing = models.URL{
			Link:        "bing.com",
			Description: "bing",
			Tags:        []string{"yandex", "888", "999"},
		}
		yahoo = models.URL{
			Link:        "yahoo.com",
			Description: "yahoo",
			Tags:        []string{"yahoo", "111", "111"},
		}
	)

	mFull := &MemStor{
		Urls: map[string]models.URL{
			"yandex.ru": yandex,
			"google.ru": google,
			"bing.com":  bing,
			"yahoo.com": yahoo,
		},
	}

	mEmpty := New()

	tests := []struct {
		name      string
		m         *MemStor
		searchStr string
		want      []models.URL
		wantErr   error
	}{
		{
			name:      "Search yandex successful",
			m:         mFull,
			searchStr: "yandex",
			want:      []models.URL{yandex, google, bing},
			wantErr:   nil,
		},
		{
			name:      "Search 111 successful",
			m:         mFull,
			searchStr: "111",
			want:      []models.URL{yahoo, yandex},
			wantErr:   nil,
		},		
		{
			name:      "Search ya failed (short searchStr)",
			m:         mFull,
			searchStr: "ya",
			want:      nil,
			wantErr:   storage.ErrTooShortSearchWord,
		},
		{
			name:      "Search noway failed",
			m:         mFull,
			searchStr: "noway",
			want:      nil,
			wantErr:   storage.ErrSearchNotFound,
		},
		{
			name:      "Search at empty storage",
			m:         mEmpty,
			searchStr: "yandex",
			want:      nil,
			wantErr:   storage.ErrSearchNotFound,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			
			got, err := tt.m.Search(tt.searchStr)
			assert.ErrorIs(t, err, tt.wantErr)
			if tt.wantErr != nil {
				return
			}

			assert.Equal(t, tt.want, got)
		})
	}
}

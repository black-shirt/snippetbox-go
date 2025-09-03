package main

import (
	"testing"
	"time"

	"github.com/black-shirt/snippetbox/internal/models/assert"
)

func TestHumanDate(t *testing.T) {
	tests := []struct {
		name string
		tm   time.Time
		want string
	}{

		{

			name: "UTC",
			tm:   time.Date(2025, 9, 1, 19, 39, 0, 0, time.UTC),
			want: "01 Sep 2025 at 19:39",
		},
		{
			name: "Empty",
			tm:   time.Time{},
			want: "",
		},
		{
			name: "CET",
			tm:   time.Date(2025, 9, 1, 18, 39, 0, 0, time.FixedZone("CET", 1*60*60)),
			want: "01 Sep 2025 at 17:39",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			hd := humanDate(tt.tm)

			assert.Equal(t, hd, tt.want)
		})
	}
}

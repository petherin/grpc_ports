package json

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"portsvc/proto"
)

func TestRead(t *testing.T) {
	tt := []struct {
		name     string
		path     string
		expected map[string]proto.Port
	}{
		{
			name: "1. Happy path, port map returned on channel",
			path: "test.json",
			expected: map[string]proto.Port{
				"AEAJM": {
					Name:        "Ajman",
					City:        "Ajman",
					Country:     "United Arab Emirates",
					Alias:       []string{},
					Regions:     []string{},
					Coordinates: []float32{55.513645, 25.405216},
					Timezone:    "Asia/Dubai",
					Unlocs:      []string{"AEAJM"},
					Code:        "52000"},
			},
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {

			reader, portCh := New(tc.path)
			go reader.Read()

			// This isn't great because it'll hold up the testing if this times out.
			ctx, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)
			defer cancel()

			select {
			case actual := <-portCh:
				assert.Equal(t, tc.expected, actual)
			case <-ctx.Done():
				t.Fatal("timeout")
				break
			}
		})
	}
}

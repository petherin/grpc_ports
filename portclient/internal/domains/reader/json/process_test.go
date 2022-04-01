package json

import (
	"context"
	"log"
	"os"
	"portsvc/proto"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
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
					Province:    "Ajman",
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
			file, err := os.Open(tc.path)
			if err != nil {
				log.Fatal(err)
			}
			defer file.Close()

			portCh := make(chan map[string]proto.Port)
			reader := New(file, portCh)
			go reader.Read()

			// This isn't great because it'll hold up the testing if it runs long enough to time out.
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

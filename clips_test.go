package hyperdeck

import (
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_Client_ClipsGet(t *testing.T) {
	examples := map[string]struct {
		Input string
		Clips Clips
		Error string
	}{
		"With an error": {
			Input: "error.txt",
			Error: "150 invalid state",
		},
		"When we receive no clips": {
			Input: "empty.txt",
			Clips: []Clip{},
		},
		"With a single clip": {
			Input: "single_clip.txt",
			Clips: []Clip{
				{
					ID:       42,
					Name:     "TestClip",
					StartAt:  Timecode{1, 23, 45, 10},
					Duration: Timecode{0, 1, 34, 12},
				},
			},
		},
		"With a single clip with spaces": {
			Input: "single_clip_with_space.txt",
			Clips: []Clip{
				{
					ID:       42,
					Name:     "A Clip With Spaces",
					StartAt:  Timecode{1, 23, 45, 10},
					Duration: Timecode{0, 1, 34, 12},
				},
			},
		},
		"With multiple clips": {
			Input: "multiple_clips.txt",
			Clips: []Clip{
				{
					ID:       42,
					Name:     "TestClip",
					StartAt:  Timecode{1, 23, 45, 10},
					Duration: Timecode{0, 1, 34, 12},
				}, {
					ID:       43,
					Name:     "01-Clip-Test",
					StartAt:  Timecode{2, 34, 56, 1},
					Duration: Timecode{1, 2, 12, 34},
				},
			},
		},
	}

	for name, example := range examples {
		t.Run(name, func(t *testing.T) {
			c := &Client{
				operations: make(chan Operation),
			}

			var received []byte
			input, err := ioutil.ReadFile("fixtures/clips_get/" + example.Input)
			require.NoError(t, err)

			go func() {
				operation := <-c.operations
				received = operation.Payload
				operation.Result <- []byte(input)
			}()

			clips, err := c.ClipsGet()
			if example.Error != "" {
				require.Error(t, err)
				assert.Contains(t, err.Error(), example.Error)
				return
			}

			require.NoError(t, err)

			assert.Equal(t, "clips get\n", string(received))
			assert.Equal(t, example.Clips, clips)
		})
	}
}

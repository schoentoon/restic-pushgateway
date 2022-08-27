package stats

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParse(t *testing.T) {
	f, err := os.Open("testdata/restic.ljson")
	if err != nil {
		t.Fatal(err)
	}
	defer f.Close()

	stats, err := ParseOutput(f)
	if assert.NoError(t, err) {
		assert.Equal(t, &Stats{
			Type:            "summary",
			FilesNew:        0,
			FilesChanged:    11,
			FilesUnmodified: 6926,
			DirsNew:         0,
			DirsChanged:     15,
			DirsUnmodified:  614,
			DataBlobs:       8,
			TreeBlobs:       15,
			DataAdded:       683689,
			FilesProcessed:  6937,
			BytesProcessed:  10189155310,
			Duration:        1.125009677,
			SnapshotID:      "216ebdee",
		}, stats)
	}
}

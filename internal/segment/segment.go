// Package segment contains all the logic for handling each segment of the
// shell prompt.
package segment

// Segmenter is the interface that all segments must implement.
type Segmenter interface {
	// Render renders the segment.
	Render() string
}

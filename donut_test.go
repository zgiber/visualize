package visualize

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestDonut(t *testing.T) {
	d, err := Donut(DonutOptions{
		Width:     600,
		Value:     74,
		MaxValue:  100,
		Color:     Blue,
		BgColor:   Grey,
		TextColor: DarkGrey,
	})

	// without OCR this is just a visual test
	// that is not automated
	require.NoError(t, err)
	os.WriteFile("donut.png", d, 0755)
}

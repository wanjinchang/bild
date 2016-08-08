package bild

import (
	"image"
	"testing"
)

func TestFlipH(t *testing.T) {
	cases := []struct {
		value    image.Image
		expected *image.RGBA
	}{
		{
			value: &image.RGBA{
				Rect:   image.Rect(0, 0, 2, 2),
				Stride: 8,
				Pix: []uint8{
					0x80, 0x80, 0x80, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF,
					0xFF, 0xFF, 0xFF, 0xFF, 0x80, 0x80, 0x80, 0x80,
				},
			},
			expected: &image.RGBA{
				Rect:   image.Rect(0, 0, 2, 2),
				Stride: 8,
				Pix: []uint8{
					0xFF, 0xFF, 0xFF, 0xFF, 0x80, 0x80, 0x80, 0xFF,
					0x80, 0x80, 0x80, 0x80, 0xFF, 0xFF, 0xFF, 0xFF,
				},
			},
		},
		{
			value: &image.RGBA{
				Rect:   image.Rect(0, 0, 3, 2),
				Stride: 12,
				Pix: []uint8{
					0x80, 0x80, 0x80, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF,
					0xFF, 0xFF, 0xFF, 0xFF, 0x80, 0x80, 0x80, 0x80, 0xFF, 0xFF, 0xFF, 0xFF,
				},
			},
			expected: &image.RGBA{
				Rect:   image.Rect(0, 0, 3, 2),
				Stride: 12,
				Pix: []uint8{
					0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0x80, 0x80, 0x80, 0xFF,
					0xFF, 0xFF, 0xFF, 0xFF, 0x80, 0x80, 0x80, 0x80, 0xFF, 0xFF, 0xFF, 0xFF,
				},
			},
		},
		{
			value: &image.RGBA{
				Rect:   image.Rect(0, 0, 2, 3),
				Stride: 8,
				Pix: []uint8{
					0x80, 0x80, 0x80, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF,
					0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF,
					0x80, 0x80, 0x80, 0x80, 0xFF, 0xFF, 0xFF, 0xFF,
				},
			},
			expected: &image.RGBA{
				Rect:   image.Rect(0, 0, 2, 3),
				Stride: 8,
				Pix: []uint8{
					0xFF, 0xFF, 0xFF, 0xFF, 0x80, 0x80, 0x80, 0xFF,
					0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF,
					0xFF, 0xFF, 0xFF, 0xFF, 0x80, 0x80, 0x80, 0x80,
				},
			},
		},
	}

	for _, c := range cases {
		actual := FlipH(c.value)
		if !rgbaImageEqual(actual, c.expected) {
			t.Error(testFailMessage("FlipH", c.expected, actual))
		}
	}
}

func TestFlipV(t *testing.T) {
	cases := []struct {
		value    image.Image
		expected *image.RGBA
	}{
		{
			value: &image.RGBA{
				Rect:   image.Rect(0, 0, 2, 2),
				Stride: 8,
				Pix: []uint8{
					0x80, 0x80, 0x80, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF,
					0xFF, 0xFF, 0xFF, 0xFF, 0x80, 0x80, 0x80, 0x80,
				},
			},
			expected: &image.RGBA{
				Rect:   image.Rect(0, 0, 2, 2),
				Stride: 8,
				Pix: []uint8{
					0xFF, 0xFF, 0xFF, 0xFF, 0x80, 0x80, 0x80, 0x80,
					0x80, 0x80, 0x80, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF,
				},
			},
		},
		{
			value: &image.RGBA{
				Rect:   image.Rect(0, 0, 3, 2),
				Stride: 12,
				Pix: []uint8{
					0x80, 0x80, 0x80, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF,
					0xFF, 0xFF, 0xFF, 0xFF, 0x80, 0x80, 0x80, 0x80, 0xFF, 0xFF, 0xFF, 0xFF,
				},
			},
			expected: &image.RGBA{
				Rect:   image.Rect(0, 0, 3, 2),
				Stride: 12,
				Pix: []uint8{
					0xFF, 0xFF, 0xFF, 0xFF, 0x80, 0x80, 0x80, 0x80, 0xFF, 0xFF, 0xFF, 0xFF,
					0x80, 0x80, 0x80, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF,
				},
			},
		},
		{
			value: &image.RGBA{
				Rect:   image.Rect(0, 0, 2, 3),
				Stride: 8,
				Pix: []uint8{
					0x80, 0x80, 0x80, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF,
					0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF,
					0x80, 0x80, 0x80, 0x80, 0xFF, 0xFF, 0xFF, 0xFF,
				},
			},
			expected: &image.RGBA{
				Rect:   image.Rect(0, 0, 2, 3),
				Stride: 8,
				Pix: []uint8{
					0x80, 0x80, 0x80, 0x80, 0xFF, 0xFF, 0xFF, 0xFF,
					0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF,
					0x80, 0x80, 0x80, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF,
				},
			},
		},
	}

	for _, c := range cases {
		actual := FlipV(c.value)
		if !rgbaImageEqual(actual, c.expected) {
			t.Error(testFailMessage("FlipV", c.expected, actual))
		}
	}
}

package animation

import (
	"strings"
)

// ==========================================
// 1. APPLICATION CODE (Implementation)
// ==========================================

type Animation struct {
	Text   string
	Frames int
	art    []string
}

func NewAnimation(text string, frames int) *Animation {
	return &Animation{
		Text:   text,
		Frames: frames,
		art:    make([]string, frames),
	}
}

func (a *Animation) GetFrame(index int) string {
	return a.art[index%a.Frames]
}

func (a *Animation) Play() string {
	return strings.Join(a.art, "\n\n---\n\n")
}

func (a *Animation) GenerateSpinFrames() {
	for i := 0; i < a.Frames; i++ {
		spaces := strings.Repeat(" ", i)
		line := spaces + a.Text
		width := len(line)

		// Build 10 lines where every line is perfectly padded to match the first line's width
		var lines []string
		for l := 0; l < 10; l++ {
			if l == 0 {
				lines = append(lines, line)
			} else {
				lines = append(lines, strings.Repeat(" ", width))
			}
		}
		a.art[i] = strings.Join(lines, "\n") + "\n"
	}
}

func (a *Animation) GenerateWaveFrames() {
	for i := 0; i < a.Frames; i++ {
		padding := strings.Repeat(" ", (i%4)+1)
		line := padding + a.Text
		width := len(line)

		var lines []string
		for l := 0; l < 10; l++ {
			if l == 0 {
				lines = append(lines, line)
			} else {
				lines = append(lines, strings.Repeat(" ", width))
			}
		}
		a.art[i] = strings.Join(lines, "\n") + "\n"
	}
}

func (a *Animation) GenerateZoomFrames() {
	for i := 0; i < a.Frames; i++ {
		var lines []string
		// Every line gets an identical indentation so they all share the exact same width
		indent := strings.Repeat(" ", (i)%5)
		line := indent + a.Text

		for l := 0; l < 10; l++ {
			lines = append(lines, line)
		}
		a.art[i] = strings.Join(lines, "\n") + "\n"
	}
}

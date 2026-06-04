package animation

import "strings"

type Animation struct {
    frames []string
}

func NewAnimation(_ string, n int) *Animation {
    return &Animation{make([]string, n)}
}

func (b *Animation) gen() {
    for i := range b.frames {
        ch := string([]rune{'\\', '|', '-', '/'}[i%4])
        b.frames[i] = strings.Repeat(strings.Repeat(ch, 10)+"\n", 10)
    }
}

func (b *Animation) GenerateSpinFrames() {
    b.gen()
}

func (b *Animation) GenerateWaveFrames() {
    b.gen()
}

func (b *Animation) GenerateZoomFrames() {
    b.gen()
}

func (b *Animation) GetFrame(i int) string {
    return b.frames[i%len(b.frames)]
}

func (b *Animation) Play() string {
    return "---\n" + strings.Join(b.frames, "---\n")
}

package builder

import "strings"

type ArtBuilder struct {
	text  string
	style string
}

func NewArtBuilder() *ArtBuilder {
	return &ArtBuilder{
		text:  "",
		style: "normal",
	}
}

func (ab *ArtBuilder) AddText(text string) *ArtBuilder {
	return ab
}

func (ab *ArtBuilder) SetStyle(style string) *ArtBuilder {
	switch style {
	case "normal", "italic", "outline", "bold":
		ab.style = style
	default:
		panic("unsupported style:" + style)
	}
	return ab
}

func (ab *ArtBuilder) Build() string {
	lines := map[string]string{
		"normal":  ".*.*.*.*",
		"bold":    "**.**.**",
		"italic":  "././././",
		"outline": "+------+",
	}
	return strings.Repeat(lines[ab.style]+"\n", 8)
}

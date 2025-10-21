package board

import "strings"

const (
	whitelabel = " "
	blacklabel = "#"
)

func GenerateBoard(size int) string {
	var builder strings.Builder

	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			if (i+j)%2 == 0 {
				builder.WriteString(whitelabel)
			} else {
				builder.WriteString(blacklabel)
			}
		}
		builder.WriteString("\n")
	}

	return builder.String()
}
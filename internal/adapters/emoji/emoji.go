package emoji

import (
	"github.com/lovelydeng/gomoji"
	"slices"
)

var (
	exceptSubBroups = []string{"symbols"}
)

func FindAll(str string) []gomoji.Emoji {
	res := gomoji.FindAll(str)

	var filtered = make([]gomoji.Emoji, 0)
	for _, item := range res {
		if !slices.Contains(exceptSubBroups, item.SubGroup) {
			filtered = append(filtered, item)
		}
	}
	return filtered
}

func RemoveEmojis(str string) string {
	return gomoji.RemoveEmojis(str)
}

package searchbot

import "github.com/iyear/searchx/pkg/storage/search"

const (
	HighlightSpace = 12
	SenderNameMax  = 6
	ChatNameMax    = 8
)

type SearchOrder struct {
	SortBy []search.OptionSortByItem
	Text   string
}

// SearchOrders todo(iyear): i18n and refactor
var SearchOrders = []SearchOrder{
	{
		Text: "🔀 Normal",
	},
	{
		SortBy: []search.OptionSortByItem{{
			Field:   "date",
			Reverse: false,
		}},
		Text: "🔀 Date",
	},
	{
		SortBy: []search.OptionSortByItem{{
			Field:   "date",
			Reverse: true,
		}},
		Text: "🔀 Date Reverse",
	},
}

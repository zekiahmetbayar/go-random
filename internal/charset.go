package internal

import "github.com/zekiahmetbayar/go-random/constants"

func GetCharset(numbers, letters, specials bool) []string {
	var G_CHARSET []string
	if numbers {
		G_CHARSET = append(G_CHARSET, constants.NMB_CHARSET...)
	}
	if letters {
		G_CHARSET = append(G_CHARSET, constants.EN_CHARSET...)
	}
	if specials {
		G_CHARSET = append(G_CHARSET, constants.SPEC_CHARSET...)
	}

	return G_CHARSET
}

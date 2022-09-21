package regex

func matchOne(pattern, text string) bool {
	// パターンが空ならどんなテキストでもマッチする
	if len(pattern) == 0 {
		return true
	}

	// パターンが定義されていて、テキストが空ならマッチしない
	if len(text) == 0 {
		return false
	}

	// パターンがワイルドカードでテキストが空でないならマッチする
	if pattern == "." {
		return true
	}

	return pattern == text
}

func match(pattern, text string) bool {
	// パターンが空ならどんなテキストでもマッチさせる
	if len(pattern) == 0 {
		return true
	}

	return matchOne(pattern[0:1], text[0:1]) && match(pattern[1:], text[1:])
}

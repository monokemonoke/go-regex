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

	// パターンが文字列の末尾で, テキストも文字列の末尾ならマッチさせる
	if pattern == "$" && text == "" {
		return true
	}

	return matchOne(pattern[0:1], text[0:1]) && match(pattern[1:], text[1:])
}

func search(pattern, text string) bool {
	// パターンが文字列の先頭を含んでいている場合
	if pattern[0:1] == "^" {
		return match(pattern[1:], text)
	}

	// テキスト全てのスタートポイントに関してパターンの一致を検索
	for i := range text {
		// もしどこかの位置でマッチしたらtrueを返す
		if match(pattern, text[i:]) {
			return true
		}
	}

	// どの位置でもマッチしなければfalseを返す
	return false
}

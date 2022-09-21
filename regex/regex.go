package regex

// 指定された位置の文字を取得する関数. もしindexがおかしな位置だった場合は空文字を返す
func getChar(str string, index int) string {
	if index < 0 || len(str) < index+1 {
		return ""
	}

	return str[index : index+1]
}

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

	// パターンの2文字目に?がある場合
	if len(pattern) > 2 && pattern[1:2] == "?" {
		return matchQuestion(pattern, text)
	}

	return matchOne(getChar(pattern, 0), getChar(text, 0)) && match(pattern[1:], text[1:])
}

func matchQuestion(pattern, text string) bool {
	// ?の前の文字以外は一致するパターン
	notUsed := match(pattern[2:], text)

	// ?の前の文字とも一致するパターン
	used := false
	if len(pattern) >= 2 {
		used = matchOne(getChar(pattern, 0), getChar(text, 0)) && match(pattern[2:], text[1:])
	}

	return notUsed || used
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

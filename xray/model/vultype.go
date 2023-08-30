package model

func vultype() map[string]string {
	return map[string]string{
		"xss":            "reflected-xss",
		"sqldet":         "sql-injection",
		"cmd-injection":  "cmd-injection",
		"path-traversal": "path-traversal",
		"xxe":            "xxe",
		"ssrf":           "ssrf",
		"brute-force":    "crypto-bad-ciphers",
		"redirect":       "unvalidated-redirect",
		"crlf-injection": "header-injection",
		"upload":         "path-traversal",
		"xstream":        "unsafe-json-deserialize",
	}
}
func vulLevel() map[string]string {
	return map[string]string{
		"xss":            "MEDIUM",
		"sqldet":         "HIGH",
		"cmd-injection":  "HIGH",
		"path-traversal": "HIGH",
		"xxe":            "MEDIUM",
		"ssrf":           "ssrf",
		"brute-force":    "LOW",
		"redirect":       "LOW",
		"crlf-injection": "LOW",
		"upload":         "HIGH",
		"xstream":        "HIGH",
	}
}
func GetVultype(input string) string {
	value, ok := vultype()[input]
	if ok {
		return value
	} else {
		return input
	}
}

func GetVulLevel(input string) string {
	value, ok := vulLevel()[input]
	if ok {
		return value
	} else {
		return "HIGH"
	}
}

package model

func Vultype() map[string]string {
	return map[string]string{
		"xss":            "reflected-xss",
		"sqldet":         "sql-injection",
		"cmd-injection":  "cmd-injection",
		"path-traversal": "path-traversal",
		"xxe":            "xxe",
		"ssrf":           "ssrf",
		"brute-force":    "crypto-bad-ciphers",
		"redirect":       "unvalidated-redirect",
	}
}
func VulLevel() map[string]string {
	return map[string]string{
		"xss":            "MEDIUM",
		"sqldet":         "HIGH",
		"cmd-injection":  "HIGH",
		"path-traversal": "HIGH",
		"xxe":            "MEDIUM",
		"ssrf":           "ssrf",
		"brute-force":    "LOW",
		"redirect":       "LOW",
	}
}

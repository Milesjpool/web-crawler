package util

import "net/url"

func AsStrings(urls []url.URL) []string {
	strings := make([]string, len(urls))
	for i := 0; i<len(urls) ;i++  {
		strings[i] = urls[i].String()
	}
	return strings
}

package subpageproviders

import "net/url"

type SubPageProvider interface {
	GetSubPages(pageUrl *url.URL) []url.URL
}
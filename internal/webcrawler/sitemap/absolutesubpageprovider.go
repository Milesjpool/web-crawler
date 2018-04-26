package sitemap

import "net/url"

type AbsoluteSubPageProvider struct {
	SubPageProvider SubPageProvider
}

func (s AbsoluteSubPageProvider) GetSubPages(pageUrl *url.URL) []url.URL {
	subPages := s.SubPageProvider.GetSubPages(pageUrl)

	for i, subPageUrl := range subPages {
		if !subPageUrl.IsAbs(){
			makeAbsolute(pageUrl, &subPages[i])
		}
	}

	return subPages
}

func makeAbsolute(rootUrl *url.URL, path *url.URL) {
	path.Host = rootUrl.Host
	path.Scheme = rootUrl.Scheme
}


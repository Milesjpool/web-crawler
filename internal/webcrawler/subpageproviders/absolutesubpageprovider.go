package subpageproviders

import "net/url"

type AbsoluteSubPageProvider struct {
	SubPageProvider SubPageProvider
}

func (s AbsoluteSubPageProvider) GetSubPages(pageUrl *url.URL) []url.URL {
	subPages := s.SubPageProvider.GetSubPages(pageUrl)

	for i, subPageUrl := range subPages {
		subPages[i].Fragment = ""
		subPages[i].RawQuery = ""
		if !subPageUrl.IsAbs(){
			makeAbsolute(pageUrl, &subPages[i])
		}
	}

	return subPages
}

func makeAbsolute(pageUrl *url.URL, subPageUrl *url.URL) {

	subPageUrl.Host = pageUrl.Host
	subPageUrl.Scheme = pageUrl.Scheme
}


package subpageproviders

import "net/url"

type AbsoluteSubPageProvider struct {
	SubPageProvider SubPageProvider
}

func (s AbsoluteSubPageProvider) GetSubPages(pageUrl *url.URL) []url.URL {
	subPages := s.SubPageProvider.GetSubPages(pageUrl)

	for i, subPageUrl := range subPages {
		if !subPageUrl.IsAbs(){
			subPages[i] = *pageUrl.ResolveReference(&subPageUrl)
		}
		subPages[i].Fragment = ""
		subPages[i].RawQuery = ""
	}

	return subPages
}


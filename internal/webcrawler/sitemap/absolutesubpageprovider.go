package sitemap

import "net/url"

type AbsoluteSubPageProvider struct {
	SubPageProvider SubPageProvider
}

func (s AbsoluteSubPageProvider) GetSubPages(pageUrl *url.URL) []string {
	subPages := s.SubPageProvider.GetSubPages(pageUrl)

	for i, v := range subPages {
		subUrl, _ := url.Parse(v)

		if !subUrl.IsAbs(){
			subPages[i] = formAbsUrl(pageUrl, subUrl)
		}

	}

	return subPages
}

func formAbsUrl(rootUrl *url.URL, path *url.URL) string {
	return rootUrl.Scheme + "://" + rootUrl.Host + path.Path
}


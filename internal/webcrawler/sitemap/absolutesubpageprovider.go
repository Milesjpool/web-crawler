package sitemap

import "net/url"

type AbsoluteSubPageProvider struct {
	SubPageProvider SubPageProvider
}

func (s AbsoluteSubPageProvider) GetSubPages(rawUrl string) []string {
	rootUrl, _ := url.Parse(rawUrl)
	subPages := s.SubPageProvider.GetSubPages(rootUrl.String())

	for i, v := range subPages {
		subUrl, _ := url.Parse(v)

		if !subUrl.IsAbs(){
			subPages[i] = formAbsUrl(rootUrl, subUrl)
		}

	}

	return subPages
}

func formAbsUrl(rootUrl *url.URL, path *url.URL) string {
	return rootUrl.Scheme + "://" + rootUrl.Host + path.Path
}


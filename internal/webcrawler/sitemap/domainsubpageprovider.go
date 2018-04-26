package sitemap

import "net/url"

type DomainSubPageProvider struct {
	SubPageProvider SubPageProvider
}

func (s DomainSubPageProvider) GetSubPages(url *url.URL) []string {
	allSubPages :=s.SubPageProvider.GetSubPages(url)

	filteredSubPages := []string{}

	for _, v := range allSubPages {
		subUrl, _ := url.Parse(v)
		if subUrl.Host == url.Host{
			filteredSubPages = append(filteredSubPages, v)
		}
	}

	return filteredSubPages
}

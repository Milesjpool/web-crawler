package subpageproviders

import "net/url"

type DomainSubPageProvider struct {
	SubPageProvider SubPageProvider
}

func (s DomainSubPageProvider) GetSubPages(pageUrl *url.URL) []url.URL {
	allSubPages :=s.SubPageProvider.GetSubPages(pageUrl)

	var filteredSubPages []url.URL

	for _, subUrl := range allSubPages {
		if subUrl.Host == pageUrl.Host {
			filteredSubPages = append(filteredSubPages, subUrl)
		}
	}

	return filteredSubPages
}

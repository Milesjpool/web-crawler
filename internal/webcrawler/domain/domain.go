package domain

import "net/url"

func ExtractDomain(rawurl string) (host string, err error) {
  u, err := url.Parse(rawurl)
  if err != nil {
    return "", err
  }
  return u.Host, nil
}

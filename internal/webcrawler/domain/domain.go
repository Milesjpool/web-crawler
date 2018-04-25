package domain

import "net/url"

func From(rawUrl string) (host string, err error) {
  u, err := url.Parse(rawUrl)
  if err != nil {
    return "", err
  }
  return u.Host, nil
}

package domain

import (
  "testing"
)

func TestFrom(t *testing.T){

  tables := []struct {
    url string
    expected string
    expectErr bool
  }{
    {"http://mydomain.com/path?query=1", "mydomain.com", false},
    {"https://myother.domain.com/path2", "myother.domain.com", false},
    {"something%%silly", "",  true},
    {"", "",  false},
  }

  for _, table := range tables {
    host, err := From(table.url);
    expected := table.expected

    if host != expected {
      t.Error("Expected: ", expected, ", but was: ", host);
    }
    if !table.expectErr && (err != nil)  {
      t.Error("Root domain extraction failed with unexpected error: ", err)
    }
    if table.expectErr && (err == nil)  {
      t.Error("Root domain extraction did not give expected error.")
    }
  }
}

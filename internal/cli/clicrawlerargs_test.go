package cli

import "testing"

func TestParseCliArguments(t *testing.T) {
	args := []string{"http://my.domain.com/path"}
	crawlerArgs := ParseCliArguments(args)

	expectedDomain := "my.domain.com"

	if crawlerArgs.domain != expectedDomain {
		t.Error("Expected: ", expectedDomain, ", but was: ", crawlerArgs.domain);
	}
}

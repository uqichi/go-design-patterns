package command

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	prefCmd := addPrefixCommand{
		text:   "hoge",
		prefix: "fuga",
	}
	prefCmd.execute()
	fmt.Println(prefCmd.text)

	sufCmd := addSuffixCommand{
		text:   "hoge",
		suffix: "fuga",
	}
	sufCmd.execute()
	fmt.Println(sufCmd.text)
}

package testdata

import "github.com/gogf/gf/os/gres"

func init() {
	if err := gres.Add("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"); err != nil {
		panic("add binary content to resource manager failed: " + err.Error())
	}
}
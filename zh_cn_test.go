package timeago

import "testing"

func TestGetZhCN(t *testing.T) {
	if getZhCN()["weeks2"] != "周" {
		t.Error("getZhCN must return map of strings with translations for zh-cn language")
	}

	if getZhCN()["days"] != "天" {
		t.Error("getZhCN must return map of strings with translations for zh-cn language")
	}
}

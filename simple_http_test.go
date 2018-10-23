package main

import (
	"testing"
)

func TestLoadPage(t *testing.T) {
	page, _ := loadPage("TestPage")
	if string(page.Body) != "This is a sample Page." {
		t.Errorf("Incorrect page loaded: got %s, wanted %s", string(page.Body), "This is a sample Page.")
	}
}

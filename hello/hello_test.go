package hello

import (
	"testing"
	"fmt"
)

func TestGreetsGitHub(t *testing.T) {
	result := Greet()
	fmt.Println(result)
	if result != "Hello GitHub Actions. Dev.to is pretty awesome" {
		t.Errorf("Greet() = %s; want Hello GitHub Actions. Dev.to is awesome", result)
	}
}

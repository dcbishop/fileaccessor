package fileaccessor

import (
	"fmt"
	"testing"
)

func ExampleMemFileAccessor() {
	fa := MemFileAccessor{
		map[string]string{
			"data/manifest.yaml": `---
manifest:
    gamename: ExampleName`,

			"data/settings.yaml": `---
settings:
	width: 800
	height: 600
	windowmode: fullscreen`,
			"data/example.yaml": "---",
		},
	}

	data, err := fa.ReadFile("data/example.yaml")

	if err != nil {
		return
	}

	fmt.Println(string(data))

	// Output: ---
}

func TestMemFileAccessorError(t *testing.T) {
	fa := MemFileAccessor{}
	data, err := fa.ReadFile("data/notafile")
	if err == nil {
		t.Error("Did not return an error for non-existant file")
	}
	if data != nil {
		t.Error("Data was not equal to nil")
	}
}

package fileaccessor

import (
	"fmt"
	"testing"
)

func ExampleMemFileAccessor() {
	fa := Virtual{
		map[string][]byte{
			"data/manifest.yaml": []byte(`---
manifest:
    gamename: ExampleName`),

			"data/settings.yaml": []byte(`---
settings:
	width: 800
	height: 600
	windowmode: fullscreen`),
		},
	}

	fa.WriteFile("data/example.txt", []byte("This is a virtual file."), 0777)

	data, err := fa.ReadFile("data/example.txt")

	if err != nil {
		fmt.Println("Failed to find virtual file.")
		return
	}

	fmt.Println(string(data))

	// Output: This is a virtual file.
}

func TestMemFileAccessorError(t *testing.T) {
	fa := Virtual{}
	data, err := fa.ReadFile("data/notafile")
	if err == nil {
		t.Error("Did not return an error for non-existant file")
	}
	if data != nil {
		t.Error("Data was not equal to nil")
	}
}

package editor

import (
	"fmt"

	goeditor "github.com/confluentinc/go-editor"
)

type Editor interface {
	Edit(filename string) error
}

type editor struct {
	editor *goeditor.BasicEditor
}

func NewEditor() Editor {
	return &editor{
		editor: goeditor.NewEditor(),
	}
}

// Edit opens the given file
// in the $EDITOR
func (editor *editor) Edit(filename string) error {
	err := editor.editor.Launch(filename)
	if err != nil {
		return fmt.Errorf("write: editor exited with an error")
	} else {
		return nil
	}
}

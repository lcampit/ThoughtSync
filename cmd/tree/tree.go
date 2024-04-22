package tree

import (
	"io/fs"
	"path/filepath"
	"strings"

	"github.com/pterm/pterm"
)

var ignoreDirs = []string{".git", ".obsidian"}

func GetTree(path string) error {
	tree := pterm.TreeNode{
		Text:     path,
		Children: []pterm.TreeNode{},
	}

	err := filepath.Walk(path, func(path string, info fs.FileInfo, err error) error {
		for _, skipDir := range ignoreDirs {
			if strings.Contains(path, skipDir) {
				return filepath.SkipDir
			}
		}
		childNode := pterm.TreeNode{
			Text: info.Name(),
		}
		tree.Children = append(tree.Children, childNode)
		return nil
	})
	pterm.DefaultTree.WithRoot(tree).Render()
	return err
}

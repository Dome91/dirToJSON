package dirToJSON

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"strings"
)

type file struct {
	Name     string `json:"name"`
	Path     string `json:"path"`
	IsDir    bool   `json:"is_dir"`
	Children []file `json:"children"`
}

// Parse traverses the directory starting from root, building the corresponding JSON.
// It does not resolve symlinks and instead treats them as files.
func Parse(root string) ([]byte, error) {
	var tree file
	info, err := os.Stat(root)
	if err != nil {
		return nil, err
	}

	tree.Name = root
	if info.IsDir() {
		tree.Path = root + "/"
		tree.IsDir = true
	} else {
		tree.Path = root
		tree.IsDir = false
	}

	err = traverse(&tree, root)
	if err != nil {
		return nil, err
	}

	return json.Marshal(tree)
}

func traverse(tree *file, dir string) error {
	var err error
	infos, err := ioutil.ReadDir(dir)
	if err != nil {
		return err
	}

	for i, info := range infos {
		path := strings.Join([]string{tree.Path, info.Name()}, "")
		tree.Children = append(tree.Children, file{Name: info.Name(), Path: path})
		if info.IsDir() {
			tree.Children[i].Path += "/"
			tree.Children[i].IsDir = true
			err = traverse(&tree.Children[i], path)
			if err != nil {
				return err
			}
		} else {
			tree.Children[i].IsDir = false
		}
	}

	return nil
}

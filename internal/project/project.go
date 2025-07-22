package project

import (
	"os"
	"path/filepath"
)

type DirType int

const (
	Pages DirType = iota
	Styles
	Assets
)

var Dirs = map[DirType]string{
	Pages:  "pages",
	Styles: "styles",
	Assets: "assets",
}

type Project struct {
}

func New() (*Project, error) {
	return &Project{}, nil
}

func (p *Project) Create(name string) error {
	err := os.MkdirAll(name, 0755)
	if err != nil {
		return err
	}

	for _, dirName := range Dirs {
		err := os.MkdirAll(filepath.Join(name, dirName), 0755)
		if err != nil {
			return err
		}
	}

	return nil
}

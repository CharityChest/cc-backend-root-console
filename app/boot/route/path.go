package route

import "strings"

type PathInstance struct {
	prefix []string
	path   string
}

func (p *PathInstance) getPrefix() []string {
	return p.prefix
}

func (p *PathInstance) getRelativePath() string {
	return p.path
}

func (p *PathInstance) getFullPath() string {
	return strings.Join(p.prefix, "/") + "/" + p.path
}

func (p *PathInstance) isEqual(path Path) bool {
	return p.getFullPath() == path.getFullPath()
}

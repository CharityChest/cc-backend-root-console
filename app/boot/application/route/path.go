package route

import "strings"

type pathInstance struct {
	prefix []string
	path   string
}

func (p *pathInstance) getPrefix() []string {
	return p.prefix
}

func (p *pathInstance) getRelativePath() string {
	return p.path
}

func (p *pathInstance) getFullPath() string {
	return "/" + strings.Join(p.prefix, "/") + "/" + p.path
}

func (p *pathInstance) isEqual(path Path) bool {
	return p.getFullPath() == path.getFullPath()
}

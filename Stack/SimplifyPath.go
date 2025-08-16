package stack

import "strings"

func simplifyPath(path string) string {
	simplifiedPath := make([]string, 0)
	dirs := strings.Split(path, "/")

	for _, dir := range dirs {
		if dir == "" || dir == "." {
			continue
		}
		if dir != ".." {
			simplifiedPath = append(simplifiedPath, dir)
		} else if len(simplifiedPath) > 0 {
			simplifiedPath = simplifiedPath[:len(simplifiedPath)-1]
		}
	}

	return "/" + strings.Join(simplifiedPath, "/")
}

package main

func groupIndent(lines []string) [][]string {
	var group []string
	var groups [][]string

	for _, line := range lines {
		if len(line) == 0 {
			continue
		}
		firstChar := []rune(line)[0]
		if firstChar == ' ' || firstChar == '\t' {
			group = append(group, line)
		} else {
			if len(group) > 0 {
				groups = append(groups, group)
			}
			group = nil
			group = append(group, line)
		}
	}

	groups = append(groups, group)

	return groups
}

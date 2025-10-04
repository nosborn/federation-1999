package text

import "strings"

func ListOfObjects(names []string) string {
	var sb strings.Builder
	for i, name := range names {
		if i > 0 {
			if i == len(names)-1 {
				if len(names) > 2 {
					sb.WriteByte(',')
				}
				sb.WriteString(" and ")
			} else {
				sb.WriteString(", ")
			}
		}
		sb.WriteString(name)
	}
	sb.WriteByte('.')
	return sb.String()
}

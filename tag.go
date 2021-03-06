package urlquery

import "strings"

type tag struct {
	name    string
	options []string
}

//according to official standard
func newTag(s string) *tag {
	arr := strings.Split(s, ",")
	t := &tag{
		name:    arr[0],
		options: arr[1:],
	}
	return t
}

func (t *tag) getName() string {
	return t.name
}

//contains
func (t *tag) contains(option string) bool {
	for _, o := range t.options {
		if o == option {
			return true
		}
	}
	return false
}

package template

type TemplateMap map[string]string

func (t TemplateMap) Get(name string) string {
	template := t[name]
	return template
}

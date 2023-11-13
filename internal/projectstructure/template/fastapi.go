package template

type TemplateMap map[string]string

var FastAPITemplates TemplateMap

const main = `print("hello, world!")
`

func init() {
	FastAPITemplates = TemplateMap{
		"main.py": main,
	}
}

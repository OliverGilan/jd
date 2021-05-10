package models

type Area struct {
	Code int
	Name string
	Categories map[int]Category `mapstructure:"categories"`
}

func (a *Area) getNextCategoryCode() int {
	code := 1
	for code < 10 {
		if _, ok := a.Categories[code]; ok {
			code = code + 1
			if code == 10{
				code = -1
			}
		}else{
			break
		}
	}

	return code
}

func (a *Area) isCategoryCodeAvailable(code int) bool{
	if _, ok := a.Categories[code]; ok {
		return false
	}
	return true
}
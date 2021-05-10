package models

type Category struct {
	Code int
	Name string
	Items map[int]string `mapstructure:"items"`
}

func (c *Category) GetNextItemCode() int {
	code := 1
	for code < 100 {
		if _, ok := c.Items[code]; ok {
			code = code + 1
			if code == 100{
				code = -1
			}
		}else{
			break
		}
	}

	return code
}

func (c *Category) IsItemCodeAvailable(code int) bool{
	if _, ok := c.Items[code]; ok {
		return false
	}
	return true
}
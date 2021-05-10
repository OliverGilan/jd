package models

type Config struct {
	DefaultProject int `mapstructure:"default_project"`
	Projects map[int]Project `mapstructure:"projects"`
	Paths map[int]string
}

func (c *Config) GetNextAreaCode() int {
	code := 100
	for code < 1000 {
		if _, ok := c.Projects[code]; ok {
			code = code + 1
			if code == 1000{
				code = -1
			}
		}else{
			break
		}
	}

	return code
}

func (p *Config) IsAreaCodeAvailable(code int) bool{
	if _, ok := p.Projects[code]; ok {
		return false
	}
	return true
}
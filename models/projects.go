package models

type Project struct {
	Code int 
	Name string
	Areas map[int]Area `mapstructure:"areas"`
}

// func (p *Project) GetAreas() []map[int]Area{
// 	return p.Areas
// }

func (p *Project) getNextAreaCode() int {
	code := 10
	for code < 100 {
		if _, ok := p.Areas[code]; ok {
			code = code + 10
			if code == 100{
				code = -1
			}
		}else{
			break
		}
	}

	return code
}

func (p *Project) isAreaCodeAvailable(code int) bool{
	if _, ok := p.Areas[code]; ok {
		return false
	}
	return true
}
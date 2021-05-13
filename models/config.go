package models

import "strings"

type Config struct {
	DefaultProject int `mapstructure:"default_project"`
	Projects map[int]Project `mapstructure:"projects"`
	Paths map[int]string `mapstructure:"paths"`
}

func (c *Config) GetNextProjectCode() int {
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

func (p *Config) IsProjectCodeAvailable(code int) bool{
	if _, ok := p.Projects[code]; ok {
		return false
	}
	return true
}

func (p *Config) GetActiveProject(cwd string) int {
	pathParts := strings.SplitAfter(cwd, "/")
	temp := ""
	var lastMatching int;

	for i := 0; i < len(pathParts); i++{
		temp += pathParts[i]
		for code, path := range p.Paths {
			if path == temp{
				lastMatching = code
				break
			}
		}
	}
	if lastMatching == 0 {
		return p.DefaultProject
	}
	return lastMatching
}
package models

type Project struct {
	Code int 
	Name string
	Areas []map[int]Area `mapstructure:"areas"`
}
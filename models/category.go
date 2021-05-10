package models

type Category struct {
	Code int
	Name string
	Items []map[int]string `mapstructure:"items"`
}
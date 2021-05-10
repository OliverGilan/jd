package models

type Area struct {
	Code int
	Name string
	Categories []map[int]Category `mapstructure:"categories"`
}
package models

type Personalidade struct {
	Nome     string `json:"nome"`
	Historia string `json:"historis"`
}

var Personalidades []Personalidade

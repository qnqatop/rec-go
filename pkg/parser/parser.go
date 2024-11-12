package parser

import "github.com/gocolly/colly/v2"

// общий интерфейс поведения парсеров для университетов
type Parserer interface {
	ParseDirections() error
	Collector(domains []string, withProxy bool) (*colly.Collector, error)
}

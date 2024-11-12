package parser

import (
	"context"
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"
	"time"

	"rec-start/pkg/db"

	"github.com/go-pg/pg/v10/orm"
	"github.com/gocolly/colly/v2"
)

type SFEDUParser struct {
	repo db.BaseRepo
}

func NewSFEDUParser(dbc orm.DB) *SFEDUParser {
	return &SFEDUParser{
		repo: db.NewBaseRepo(dbc),
	}
}

func (s *SFEDUParser) ParseDirections() error {
	domains := []string{"sfedu.ru"}
	c, err := s.Collector(domains, false)
	if err != nil {
		return err
	}

	return c.Visit("https://sfedu.ru/www/stat_pages22.show?p=WTS/N13696/P")

}

func (s *SFEDUParser) Collector(domains []string, withProxy bool) (*colly.Collector, error) {
	c := colly.NewCollector(
		colly.AllowedDomains(domains...),
		colly.UserAgent("Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/70.0.3538.77 Safari/537.36"),
	)

	ctx := context.Background()

	// Устанавливаем ограничение частоты запросов
	err := c.Limit(&colly.LimitRule{
		DomainGlob:  "*",
		Parallelism: 2,
		Delay:       1 * time.Second,
	})
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})

	c.OnResponse(func(r *colly.Response) {
		fmt.Println("Received response from", r.Request.URL)
	})

	c.OnError(func(r *colly.Response, err error) {
		fmt.Println("Request URL:", r.Request.URL, "failed with response:", r, "\nError:", err)
	})

	if withProxy {
		err = c.SetProxy("http://proxyserver:port")
		if err != nil {
			log.Fatal(err)
		}
	}

	var code, name, cost, city string

	c.OnHTML("table:nth-of-type(1) tbody tr", func(e *colly.HTMLElement) {

		totalColumns := len(e.ChildTexts("td"))
		var specializations []string

		// Для каждой строки `tr` выбираем все элементы `td`
		e.ForEach("td", func(i int, el *colly.HTMLElement) {
			text := strings.TrimSpace(el.Text)
			if totalColumns == 2 {

				switch i {
				case 0:
					// Парсим направление и специальные предметы
					lines := strings.Split(text, "\n")
					name = lines[0] // Название направления
					if len(lines) > 1 {
						specializations = lines[1:] // Остальные строки - предметы
					}
				case 1:
					// Разделяем баллы на предметы и их значения
					lines := strings.Split(text, "\n")
					for _, line := range lines {
						line = strings.TrimSpace(line)
						if line != "" {
							specializations = append(specializations, line)
						}
					}
				}
			} else {
				switch i {
				case 0:
					code = text // Код направления, например, "45.03.01"
				case 1:
					// Парсим направление и специальные предметы
					lines := strings.Split(text, "\n")
					name = lines[0] // Название направления
					if len(lines) > 1 {
						specializations = lines[1:] // Остальные строки - предметы
					}
				case 2:
					// Разделяем баллы на предметы и их значения
					lines := strings.Split(text, "\n")
					for _, line := range lines {
						line = strings.TrimSpace(line)
						if line != "" {
							specializations = append(specializations, line)
						}
					}
				case 3:
					cost = text // Стоимость обучения
				case 4:
					city = text // Город
				}
			}
		})

		cc, err := s.repo.CitiesByFilters(ctx, &db.CitySearch{Name: &city}, db.PagerNoLimit)
		if err != nil {
			log.Printf(err.Error())
		}

		var cityID int
		if len(cc) == 0 {
			cModel := db.City{Name: city}
			c, err := s.repo.AddCity(ctx, &cModel)
			if err != nil {
				log.Printf(err.Error())
			}
			cityID = c.ID
		} else {
			cityID = cc[0].ID
		}

		text := strings.Join(specializations, " ")
		re := regexp.MustCompile(`\d+\.[^\d]+?-\s*\d{2}`)
		matches := re.FindAllString(text, -1)
		money, err := strconv.Atoi(strings.ReplaceAll(cost, " ", ""))
		_, err = s.repo.AddDirection(ctx, &db.Direction{
			Code:         code,
			Name:         name,
			Params:       &db.DirectionParams{CountPoint: strings.Join(matches, "\n")},
			Cost:         &money,
			UniversityID: 1,
			CityID:       cityID,
		})
		if err != nil {
			log.Printf(err.Error())
		}
	})

	return c, nil

}

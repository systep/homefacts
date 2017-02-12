package homefacts

import (
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type parser struct {
}

func (p *parser) parseHTML(resp *http.Response) (*Result, error) {
	result := &Result{
		CrimeRate:            &resultProperty{},
		SchoolRating:         &resultProperty{},
		RegisteredOffenders:  &resultProperty{},
		AvgHomePrice:         &resultProperty{},
		Foreclosures:         &resultProperty{},
		NaturalHazards:       &resultProperty{},
		EnvironmentalHazards: &resultProperty{},
	}

	doc, err := goquery.NewDocumentFromResponse(resp)

	if err != nil {
		return nil, err
	}

	result = getResultText(doc, result)
	result = getResultIcons(doc, result)

	return result, nil
}

func getResultText(doc *goquery.Document, result *Result) *Result {
	doc.Find(".prop-summary").First().Find(".col3").Each(func(i int, s *goquery.Selection) {
		switch i {
		case 0:
			result.CrimeRate.Text = strings.Trim(s.Text(), "")
		case 1:
			result.SchoolRating.Text = strings.Trim(s.Text(), "")
		case 2:
			result.RegisteredOffenders.Text = strings.Trim(s.Text(), "")
		case 3:
			result.AvgHomePrice.Text = strings.Trim(s.Text(), "")
		case 4:
			result.Foreclosures.Text = strings.Trim(s.Text(), "")
		case 5:
			result.NaturalHazards.Text = strings.Trim(s.Text(), "")
		case 6:
			result.EnvironmentalHazards.Text = strings.Trim(s.Text(), "")
		}
	})

	return result
}

func getResultIcons(doc *goquery.Document, result *Result) *Result {
	doc.Find(".prop-summary").First().Find(".col1").Find("i").Each(func(i int, s *goquery.Selection) {
		switch i {
		case 0:
			v, _ := s.Attr("class")
			result.CrimeRate.Class = strings.Trim(v, "")
		case 1:
			v, _ := s.Attr("class")
			result.SchoolRating.Class = strings.Trim(v, "")
		case 2:
			v, _ := s.Attr("class")
			result.RegisteredOffenders.Class = strings.Trim(v, "")
		case 3:
			v, _ := s.Attr("class")
			result.AvgHomePrice.Class = strings.Trim(v, "")
		case 4:
			v, _ := s.Attr("class")
			result.Foreclosures.Class = strings.Trim(v, "")
		case 5:
			v, _ := s.Attr("class")
			result.NaturalHazards.Class = strings.Trim(v, "")
		case 6:
			v, _ := s.Attr("class")
			result.EnvironmentalHazards.Class = strings.Trim(v, "")
		}
	})

	return result
}

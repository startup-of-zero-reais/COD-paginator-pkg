package paginator

import (
	"fmt"
	"log"
	"net/url"
	"strconv"
)

type (
	LinksInterface interface {
		getKey() url.Values
		MakeFirstPage()
		MakeLastPage()
		GetPrevPage()
		GetNextPage()
	}

	SimplePaginator struct {
		Page    uint `json:"page,omitempty"`
		PerPage uint `json:"per_page,omitempty"`
	}

	LinksResult struct {
		*MetaResult  `json:"-"`
		CurrentQuery map[string]interface{} `json:"-"`
		NextPage     string                 `json:"next_page,omitempty"`
		PrevPage     string                 `json:"prev_page,omitempty"`
		FirstPage    string                 `json:"first_page,omitempty"`
		LastPage     string                 `json:"last_page,omitempty"`
	}
)

func NewLinksResult(metadata *MetaResult, query map[string]interface{}) LinksInterface {
	links := &LinksResult{
		MetaResult:   metadata,
		CurrentQuery: query,
	}

	links.MakeFirstPage()
	links.MakeLastPage()
	links.GetPrevPage()
	links.GetNextPage()

	return links
}

func (l *LinksResult) getKey() url.Values {
	q := url.Values{}

	for key, value := range l.CurrentQuery {
		switch key {
		case Key:
			switch value.(type) {
			case string:
				q.Add(key, value.(string))
			case bool:
				q.Add(key, strconv.FormatBool(value.(bool)))
			case int:
				q.Add(key, strconv.Itoa(value.(int)))
			}
		}
	}

	return q
}

func (l *LinksResult) MakeFirstPage() {
	simplePager := &SimplePaginator{
		Page:    1,
		PerPage: uint(l.MetaResult.config.ItemsPerPage),
	}

	q := l.getKey()

	q.Add("page", fmt.Sprintf("%d", simplePager.Page))
	q.Add("per_page", fmt.Sprintf("%d", simplePager.PerPage))

	l.FirstPage = fmt.Sprintf(
		"%s?%s",
		l.MetaResult.config.BaseURL,
		q.Encode(),
	)
}

func (l *LinksResult) MakeLastPage() {
	simplePager := &SimplePaginator{
		Page:    l.MetaResult.PageTotal,
		PerPage: uint(l.config.ItemsPerPage),
	}

	log.Println("PAGE TOTAL", l.PageTotal)

	if l.PageTotal <= 1 {
		l.LastPage = ""
		return
	}

	q := l.getKey()

	q.Add("page", fmt.Sprintf("%d", simplePager.Page))
	q.Add("per_page", fmt.Sprintf("%d", simplePager.PerPage))

	l.LastPage = fmt.Sprintf(
		"%s?%s",
		l.config.BaseURL,
		q.Encode(),
	)
}

func (l *LinksResult) GetPrevPage() {
	page := l.MetaResult.Page

	if (l.MetaResult.Page - 1) > 0 {
		page = l.MetaResult.Page - 1
	} else {
		l.PrevPage = ""
		return
	}

	simplePager := &SimplePaginator{
		Page:    page,
		PerPage: uint(l.config.ItemsPerPage),
	}

	q := l.getKey()

	q.Add("page", fmt.Sprintf("%d", simplePager.Page))
	q.Add("per_page", fmt.Sprintf("%d", simplePager.PerPage))

	l.PrevPage = fmt.Sprintf(
		"%s?%s",
		l.config.BaseURL,
		q.Encode(),
	)
}

func (l *LinksResult) GetNextPage() {
	page := l.MetaResult.Page

	if (l.MetaResult.Page + 1) <= l.MetaResult.PageTotal {
		page = l.MetaResult.Page + 1
	} else {
		l.NextPage = ""
		return
	}

	simplePager := &SimplePaginator{
		Page:    page,
		PerPage: uint(l.config.ItemsPerPage),
	}

	q := l.getKey()

	q.Add("page", fmt.Sprintf("%d", simplePager.Page))
	q.Add("per_page", fmt.Sprintf("%d", simplePager.PerPage))

	l.NextPage = fmt.Sprintf(
		"%s?%s",
		l.config.BaseURL,
		q.Encode(),
	)
}

package paginator

import (
	"errors"
	"fmt"
	"log"
	"net/url"
	"reflect"
	"strings"
)

func (p *pager) paginateSingle(item interface{}, paginated interface{}) error {
	pValue := reflect.ValueOf(paginated).Elem()
	iValue := reflect.ValueOf(item)

	for i := 0; i < iValue.NumField(); i++ {
		if paginator, ok := iValue.Type().Field(i).Tag.Lookup("paginator"); ok {
			paginatedField := pValue.Field(i)
			it := iValue.Field(i)
			paginatedField.Set(it)

			err := p.extractTags(paginator, it)
			if err != nil {
				return err
			}

			for k, _ := range p.kv {
				switch k {
				case Self:
					newSelf := fmt.Sprintf("%s?%s=%s", p.Config.BaseURL, p.kv[Key], p.kv[k])

					newSelfValue := reflect.ValueOf(newSelf)
					if strings.Index(paginator, Self) >= 0 {
						paginatedField.Set(newSelfValue)
					}
				}
			}
		}
	}

	return nil
}

func (p *pager) paginateCollection() error {
	paginatedSlice := reflect.MakeSlice(reflect.SliceOf(p.items.Index(0).Type()), 0, p.items.Len())

	for i := 0; i < p.items.Len(); i++ {
		log.Printf("ITEM %+v", p.items.Index(i))
		item := p.items.Index(i)

		paginateEl := reflect.New(item.Type())
		err := p.paginateSingle(item.Interface(), paginateEl.Interface())
		if err != nil {
			return err
		}

		paginatedSlice = reflect.Append(paginatedSlice, paginateEl.Elem())
	}

	p.paginated.Elem().Set(paginatedSlice)

	return nil
}

func (p *pager) extractTags(tag string, f reflect.Value) error {
	// Ex.: key:name _self non_paginate
	tags := strings.Split(tag, ";")

	for _, t := range tags {
		// Ex.: [key name] [_self] [non_paginate]
		separate := strings.Split(t, ":")

		tagKey := separate[0]
		tagValue := ""

		if len(separate) > 1 {
			tagValue = separate[1]
		}

		switch tagKey {
		case Key:
			p.kv[tagKey] = tagValue
			p.kv[Self] = url.QueryEscape(f.Interface().(string))
		case NonPaginate:
			p.kv[tagKey] = true
		}
	}

	return nil
}

func (p *pager) scanMode() (Mode, error) {
	validKinds := func(inputKinds ...reflect.Kind) bool {
		if len(inputKinds) <= 0 {
			return false
		}

		for i, ik := range inputKinds {
			if ik == 0 {
				return false
			}

			if i == 1 && ik.String() != "ptr" {
				return false
			}
		}

		return true
	}

	itemsValueKind := p.items.Kind()
	paginatedValueKind := p.paginated.Kind()

	if !validKinds(itemsValueKind, paginatedValueKind) {
		return "", errors.New("os tipos de items e de paginated sao diferentes")
	}

	if p.items.Type().Kind().String() == "slice" {
		return Collection, nil
	}

	return Single, nil
}

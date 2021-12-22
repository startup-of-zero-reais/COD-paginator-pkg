package paginator_test

import (
	"fmt"
	paginator "github.com/startup-of-zero-reais/COD-paginator-pkg"
	"testing"
)

/*
	TAGS
	key = chave que aparecera na query Ex.: key=value
	_self = chave que recebera o metadado _self Ex.: http://baseURL.com/pagina?`key`=value
	not_paginate = chave que define se o url tera a query de pagination Ex.: ?page=1&per_page=10
*/

type (
	EntityModel struct {
		Name   string `json:"name" paginator:"key:name"`
		Self   string `json:"_link" paginator:"_self"`
		Void   string `json:"-"`
		Hidden string `json:"-" paginator:"-"`
	}
)

func TestNewPaginator(t *testing.T) {

}

func TestPaginate(t *testing.T) {
	t.Run("should paginate single item", func(t *testing.T) {
		input := EntityModel{
			Name: "John",
		}

		output := &EntityModel{
			Name:   "John",
			Self:   "https://baseURL.com?name=John&page=1&per_page=10",
			Void:   "",
			Hidden: "",
		}

		pager := paginator.NewPaginator(&paginator.Config{
			BaseURL:      "https://baseURL.com.br",
			ItemsPerPage: 10,
		})

		var result EntityModel
		res := pager.Paginate(input, &result).Json()

		fmt.Printf("\nOUTPUT: %+v\nRESULT: %+v\n\n", output, res)
	})

	t.Run("should paginate slice of items", func(t *testing.T) {
		inputs := []EntityModel{
			{Name: "John doe"},
			{Name: "Doe"},
		}

		outputs := []EntityModel{
			{Name: "John doe", Self: "https://baseURL.com?key=John&page=1&per_page=10"},
			{Name: "Doe", Self: "https://baseURL.com?key=Doe&page=1&per_page=10"},
		}

		pager := paginator.NewPaginator(&paginator.Config{
			BaseURL:      "https://baseURL.com",
			ItemsPerPage: 10,
		})

		var results []EntityModel
		result := pager.Paginate(inputs, &results).Json()

		fmt.Printf("\n\nOUTPUT: %+v\nRESULT: %+v\n\n", outputs, result)
	})

	t.Run("should show meta and links", func(t *testing.T) {
		inputs := []EntityModel{
			{Name: "John doe"},
			{Name: "Doe"},
		}

		outputs := []EntityModel{
			{Name: "John doe", Self: "https://baseURL.com?key=John&page=1&per_page=10"},
			{Name: "Doe", Self: "https://baseURL.com?key=Doe&page=1&per_page=10"},
		}

		pager := paginator.NewPaginator(&paginator.Config{
			BaseURL:      "https://baseURL.com",
			ItemsPerPage: 10,
		})

		var results []EntityModel
		result := pager.WithMeta(&paginator.Metadata{
			Total: 20,
			Page:  2,
		}).Paginate(inputs, &results).Json()

		fmt.Printf("\n\nOUTPUT: %+v\nRESULT: %+v\n\n", outputs, result)
	})
}

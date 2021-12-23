package paginator

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"reflect"
)

type (
	Mode string

	Logger interface {
		SetOutput(w io.Writer)
		Printf(format string, v ...interface{})
		Print(v ...interface{})
		Println(v ...interface{})
		Fatal(v ...interface{})
		Fatalf(format string, v ...interface{})
		Fatalln(v ...interface{})
		Panic(v ...interface{})
		Panicf(format string, v ...interface{})
		Panicln(v ...interface{})
	}

	Metadata struct {
		config *Config
		Total  uint
		Page   uint
	}

	DataResult struct {
		Data     interface{}         `json:"data"`
		Metadata MetaResultInterface `json:"metadata,omitempty"`
		Links    *LinksResult        `json:"links,omitempty"`
	}

	Paginator interface {
		WithMeta(metadata *Metadata) Paginator
		Paginate(items interface{}, result interface{}) Paginator
		Json() string

		paginateSingle(items interface{}, result interface{}) error
		paginateCollection() error
		extractTags(tag string, field reflect.Value) error
	}

	Config struct {
		*log.Logger
		BaseURL      string
		ItemsPerPage uint32
	}

	pager struct {
		*Config
		*Metadata
		*log.Logger

		items     reflect.Value
		paginated reflect.Value

		kv map[string]interface{}
	}
)

const (
	Collection = Mode("collection")
	Single     = Mode("single")
)

const (
	Key         = "key"
	Self        = "_self"
	NonPaginate = "non_paginate"
	Empty       = "-"
)

func NewPaginator(config *Config) Paginator {
	if config.ItemsPerPage == 0 {
		config.ItemsPerPage = 10
	}

	if config.Logger == nil {
		config.Logger = log.Default()
	}

	return &pager{
		Config: config,
		Logger: config.Logger,
		kv: map[string]interface{}{
			Key:         "id",
			NonPaginate: false,
		},
		Metadata: &Metadata{
			config: config,
		},
	}
}

func (p *pager) WithMeta(metadata *Metadata) Paginator {
	if metadata.config == nil {
		metadata.config = p.Config
	}

	p.Metadata = metadata
	return p
}

func (p *pager) Paginate(items interface{}, paginated interface{}) Paginator {
	p.items = reflect.ValueOf(items)
	p.paginated = reflect.ValueOf(paginated)
	mode, err := p.scanMode()

	if err != nil {
		log.Fatalf(err.Error())
	}

	if mode == Collection {
		err = p.paginateCollection()
		if err != nil {
			log.Fatalf(err.Error())
		}
		return p
	}

	err = p.paginateSingle(items, paginated)
	if err != nil {
		log.Fatalf(err.Error())
	}

	return p
}

func (p *pager) Json() string {
	metadata := NewMetaResult(p.Metadata)

	dataResult := &DataResult{
		Data:     p.paginated.Elem().Interface(),
		Metadata: metadata,
		Links:    NewLinksResult(metadata, p.kv),
	}

	buf := &bytes.Buffer{}
	jsEncode := json.NewEncoder(buf)
	jsEncode.SetEscapeHTML(false)
	_ = jsEncode.Encode(dataResult)

	return buf.String()
}

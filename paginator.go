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
		Paginate(items interface{}, result interface{}) (DataResultInterface, error)

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

func (p *pager) Paginate(items interface{}, paginated interface{}) (DataResultInterface, error) {
	p.items = reflect.ValueOf(items)
	p.paginated = reflect.ValueOf(paginated)
	mode, err := p.scanMode()

	if err != nil {
		return nil, err
	}

	metadata := NewMetaResult(p.Metadata)

	if mode == Collection {
		err = p.paginateCollection()
		if err != nil {
			return nil, err
		}

		return &DataResult{
			Data:     p.paginated.Elem().Interface(),
			Metadata: metadata,
			Links:    NewLinksResult(metadata, p.kv),
		}, nil
	}

	err = p.paginateSingle(items, paginated)
	if err != nil {
		return nil, err
	}

	return &DataResult{
		Data:     p.paginated.Elem().Interface(),
		Metadata: metadata,
		Links:    NewLinksResult(metadata, p.kv),
	}, nil
}

func (d *DataResult) Json() string {
	buf := &bytes.Buffer{}
	jsEncode := json.NewEncoder(buf)
	jsEncode.SetEscapeHTML(false)

	isSlice := reflect.TypeOf(d.Data)

	if isSlice.Kind() == reflect.Slice {
		_ = jsEncode.Encode(d)
	} else {
		_ = jsEncode.Encode(d.Data)
	}
	return buf.String()
}

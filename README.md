# Code craft paginator

Pacote de paginação de dados

Instalação:

```shell
go get -u github.com/startup-of-zero-reais/COD-paginator-pkg
```

Modo de uso:

**Declaração de Entidade a ser paginada**
```go
package main

type EntityModel struct {
    ID     string `json:"id" paginator:"key:id"`
    Self   string `json:"_link" paginator:"_self"`
    Void   string `json:"-"`
    Hidden string `json:"-" paginator:"-"`
}
```

**Uso do pacote**
```go
package main

import "github.com/startup-of-zero-reais/COD-paginator-pkg"

pager := paginator.NewPaginator(&paginator.Config{
    BaseURL:      "https://baseURL.com.br",
    ItemsPerPage: 10,
})
```

**Como recuperar o JSON paginado:**
```go
// ...

var results []EntityModel

// Com metadados
p := pager.WithMeta(&paginator.Metadata{
    Total: 20,
    Page:  2,
}).Paginate(items, &results)

// Sem metadados
p := pager.Paginate(item, &results)

// Json String com o resultado
jsonResult := p.Json()

log.Println(jsonResult)

// ...
```
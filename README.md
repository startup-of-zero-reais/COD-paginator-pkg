# Code craft paginator

### Primeiros passos

- [Instalação](#-Instalação)
- [Modo de uso](#-Modo-de-uso)
- [Exemplo com Echo v4](#-Exemplo-com-Echo-v4)

Pacote de paginação de dados

### Instalação

```shell
go get -u github.com/startup-of-zero-reais/COD-paginator-pkg
```

### Modo de uso

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
// Total de 30 Registros na base de dados
Total: 30,
// Pagina atual da listagem
Page:  2,
// ItemsPerPage é opcional. Caso queira alterar valor definido em paginator.Config{}
ItemsPerPage: 5
}).Paginate(items, &results)

// Sem metadados
p := pager.Paginate(items, &results)

// Json String com o resultado
jsonResult := p.Json()

log.Println(jsonResult)

// Output:
// 
// {
//     "data": [
//       {
//         "id": "31f6bb14-d876-4e20-b1a6-bc873de55c8f",
//         "_link": "https://baseURL.com.br?id=31f6bb14-d876-4e20-b1a6-bc873de55c8f"
//       },
//       {
//         "id": "41f6bb14-d876-4e20-b1a6-bc873de55c80",
//         "_link": "https://baseURL.com.br?id=41f6bb14-d876-4e20-b1a6-bc873de55c80"
//       }
//     ],
//     "metadata": {
//       "total": 30,
//       "page": 2,
//       "page_total": 3
//     },
//     "links": {
//       "first_page": "https://baseURL.com.br?page=1&per_page=10",
//       "last_page": "https://baseURL.com.br?page=3&per_page=10",
//       "next_page": "https://baseURL.com.br?page=3&per_page=10",
//       "prev_page": "https://baseURL.com.br?page=1&per_page=10"
//     }
// }
//
// ...
```

## Exemplo com [Echo v4](https://echo.labstack.com/)

```go
package main

import (
	"github.com/startup-of-zero-reais/COD-paginator-pkg"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"

	"your.package/services"
)

type Users struct {
	ID   string `json:"id,omitempty" paginator:"key:id"`
	Name string `json:"name,omitempty" paginator:"-"`
	Link string `json:"_link,omitempty" paginator:"_self"`
}

func main() {
	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		page := c.QueryParam("page")
		perPage, _ := strconv.Atoi(
			c.QueryParam("per_page"),
		)

		users, totalOfUsers := services.GetUsers(page, perPage)

		pager := paginator.NewPaginator(&paginator.Config{
			BaseURL:      "http://localhost:3000",
			ItemsPerPage: uint32(perPage),
		})

		var paginated []Users
		result, err := pager.WithMeta(&paginator.Metadata{
			Total:        totalOfUsers,
			Page:         page,
			ItemsPerPage: uint32(perPage),
		}).Paginate(users, &paginated)

        if err != nil {
            return c.JSON(http.StatusInternalServerError, map[string]string{
                "message": err.Error(),
            })
        }

        return c.JSON(http.StatusOK, result.Json())
    })
	
    e.Start(":3000")
}
```





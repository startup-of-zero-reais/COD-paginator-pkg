package paginator_test

import (
	paginator "github.com/startup-of-zero-reais/COD-paginator-pkg"
	"github.com/startup-of-zero-reais/COD-paginator-pkg/mocks"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestNewPaginator(t *testing.T) {
	t.Run("should create an paginator", func(t *testing.T) {
		config := &paginator.Config{
			BaseURL:      "https://test.url.com",
			ItemsPerPage: 10,
		}

		pager := paginator.NewPaginator(config)

		require.NotNil(t, pager)
	})
	t.Run("should create an paginator with default items per page", func(t *testing.T) {
		config := &paginator.Config{
			BaseURL: "https://test.url.com",
		}

		pager := paginator.NewPaginator(config)

		require.NotNil(t, pager)
		require.Equal(t, uint32(10), config.ItemsPerPage)
	})
}

func TestPager_WithMeta(t *testing.T) {
	t.Run("should return an paginator with metadata", func(t *testing.T) {
		config := &paginator.Config{
			BaseURL:      "https://test.url.com",
			ItemsPerPage: 10,
		}

		pager := paginator.NewPaginator(config)

		p := pager.WithMeta(&paginator.Metadata{
			Total:        20,
			Page:         1,
			ItemsPerPage: 15,
		})

		require.NotNil(t, pager)
		require.NotNil(t, p)
	})
	t.Run("should change items per_page on result", func(t *testing.T) {
		config := &paginator.Config{
			BaseURL:      "https://test.url.com",
			ItemsPerPage: 10,
		}

		pager := paginator.NewPaginator(config)

		var result []mocks.ItemMock
		p, err := pager.WithMeta(&paginator.Metadata{
			Total:        20,
			Page:         1,
			ItemsPerPage: 15,
		}).Paginate(mocks.NewItemSlice(), &result)

		require.Contains(t, p.Json(), "per_page=15")
		require.Nil(t, err)
	})
}

func TestPager_Paginate(t *testing.T) {
	prePaginateTest := func() (*paginator.Config, []mocks.ItemMock) {
		config := &paginator.Config{
			BaseURL:      "https://test.url.com",
			ItemsPerPage: 10,
		}
		var result []mocks.ItemMock

		return config, result
	}

	t.Run("should paginate items and return data result", func(t *testing.T) {
		config := &paginator.Config{
			BaseURL:      "https://test.url.com",
			ItemsPerPage: 10,
		}

		pager := paginator.NewPaginator(config)

		items := mocks.NewItemSlice()
		var result []mocks.ItemMock
		p, err := pager.Paginate(items, &result)

		require.NotNil(t, p)
		require.Nil(t, err)
		require.Len(t, result, len(items))
		require.Contains(t, p.Json(), "links")
		require.Contains(t, p.Json(), "metadata")
		require.Contains(t, p.Json(), "data")
	})
	t.Run("should insert self link, not metadata and links", func(t *testing.T) {
		config := &paginator.Config{
			BaseURL:      "https://test.url.com",
			ItemsPerPage: 10,
		}

		pager := paginator.NewPaginator(config)

		item := *mocks.NewItem()
		var result mocks.ItemMock
		p, err := pager.Paginate(item, &result)

		require.Nil(t, err)
		require.NotNil(t, p)
		require.NotContains(t, p.Json(), "metadata")
		require.NotContains(t, p.Json(), "links")
	})
	t.Run("should fatal if err", func(t *testing.T) {
		c, result := prePaginateTest()
		pager := paginator.NewPaginator(c)

		var any interface{}
		p, err := pager.Paginate(any, &result)

		require.NotNil(t, err)
		require.Nil(t, p)
	})
}

package paginator

type (
	MetaResultInterface interface {
		buildPageTotal()
	}

	MetaResult struct {
		config    *Config
		Total     uint `json:"total,omitempty"`
		Page      uint `json:"page,omitempty"`
		PageTotal uint `json:"page_total,omitempty"`
	}
)

func NewMetaResult(metadata *Metadata) *MetaResult {
	metaResult := &MetaResult{
		config:    metadata.config,
		Total:     metadata.Total,
		Page:      metadata.Page,
		PageTotal: 1,
	}

	metaResult.buildPageTotal()

	return metaResult
}

func (m *MetaResult) buildPageTotal() {
	pages := m.Total / uint(m.config.ItemsPerPage)
	if (m.Total % uint(m.config.ItemsPerPage)) != 0 {
		pages += 1
	}

	m.PageTotal = pages

	if m.Total <= uint(m.config.ItemsPerPage) || m.PageTotal <= 0 {
		m.PageTotal = 1
	}
}

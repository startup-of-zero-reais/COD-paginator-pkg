package mocks

type (
	ItemMock struct {
		ID      string `json:"id,omitempty" paginator:"key:id"`
		Name    string `json:"name,omitempty"`
		Link    string `json:"link,omitempty" paginator:"_self"`
		Ignored string `json:"ignored,omitempty" paginator:"-"`
	}
)

func NewItem(overrideArgs ...map[string]string) *ItemMock {
	itemMock := &ItemMock{
		ID:      "31f6bb14-d876-4e20-b1a6-bc873de55c8f",
		Name:    "John doe",
		Ignored: "Ignored",
	}

	if len(overrideArgs) > 0 {
		for key, value := range overrideArgs[0] {
			switch key {
			case "id":
				itemMock.ID = value
			case "name":
				itemMock.Name = value
			case "ignored":
				itemMock.Ignored = value
			default:

			}
		}
	}

	return itemMock
}

func NewItemSlice() []ItemMock {
	return []ItemMock{
		*NewItem(map[string]string{"id": "31f6bb14-d876-4e20-b1a6-bc873de55c8f"}),
		*NewItem(map[string]string{"id": "41f6bb14-d876-4e20-b1a6-bc873de55c8f", "name": "Dummy"}),
	}
}

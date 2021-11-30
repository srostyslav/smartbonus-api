package smartbonus

import (
	"errors"
)

// Tag is smartbonus filter, used in smartbonus app catalog
type Tag struct {
	Id      string `json:"id"`
	Name    string `json:"name"`
	GroupId string `json:"group_id,omitempty"`
	IsGroup bool   `json:"is_group"`
}

// Sync list of tags
// warning: ensure that count of elements has to be less than or equal 500 in a request
func syncTags(storeId string, tags []Tag) error {
	if len(tags) > 500 {
		return errors.New("length of tags must be less or equal than 500 elements")
	} else if len(tags) == 0 {
		return errors.New("no element found")
	}

	var result string
	body := map[string]interface{}{"store": storeId, "elements": tags}

	if err := sendPostRequest(rootPath+"v2/sync/tag", body, &result); err != nil {
		return err
	} else if result != "Sync success" {
		return errors.New(result)
	}

	return nil
}

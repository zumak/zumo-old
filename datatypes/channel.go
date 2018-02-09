package datatypes

import "encoding/json"

// Channel is
type Channel struct {
	ID     string
	Name   string
	Labels map[string]string
	Member map[string]struct{}
}
type Members map[string]struct{}

func (m Members) MarshalJSON() ([]byte, error) {
	result := []string{}
	for name := range m {
		result = append(result, name)
	}
	return json.Marshal(result)
}
func (m Members) UnmarshalJSON(data []byte) error {
	if m == nil {
		m = map[string]struct{}{}
	}

	result := []string{}
	err := json.Unmarshal(data, result)
	if err != nil {
		return err
	}

	for _, name := range result {
		m[name] = struct{}{}
	}
	return nil
}

/*
func (m Member) MarshalJSON() {

}
func (m Member) UnmarshalJSON() {
}
*/

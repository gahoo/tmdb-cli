package formatter

import (
	"encoding/json"
	"strings"
)

// FilterData marshals the API response to JSON, generic interface{}, and then extracts ONLY the fields requested by the user.
func FilterData(data interface{}, fieldsStr string) (interface{}, error) {
	if fieldsStr == "" {
		return data, nil
	}

	b, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	var raw interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return nil, err
	}

	fields := strings.Split(fieldsStr, ",")
	for i, f := range fields {
		fields[i] = strings.TrimSpace(f)
	}

	filtered := retainFields(raw, fields)
	if filtered == nil {
		return make(map[string]interface{}), nil
	}
	return filtered, nil
}

func retainFields(node interface{}, fields []string) interface{} {
	switch v := node.(type) {
	case []interface{}:
		var out []interface{}
		for _, item := range v {
			if f := retainFields(item, fields); f != nil {
				out = append(out, f)
			}
		}
		if len(out) == 0 {
			return nil
		}
		return out
	case map[string]interface{}:
		out := make(map[string]interface{})
		for key, val := range v {
			// Check if this exact key is requested OR if any sub-field is requested
			relevantFields := getRelevantFields(key, fields)
			if len(relevantFields) > 0 {
				keepDirectly := false
				var subFields []string
				for _, rf := range relevantFields {
					if rf == "" {
						keepDirectly = true
					} else {
						subFields = append(subFields, rf)
					}
				}

				if keepDirectly {
					out[key] = val
				} else if subVal := retainFields(val, subFields); subVal != nil {
					out[key] = subVal
				}
			}
		}
		if len(out) == 0 {
			return nil
		}
		return out
	default:
		return nil
	}
}

// getRelevantFields returns the sub-parts of fields that match the current key.
// If "title" is requested and key is "title", returns [""].
// If "credits.cast" is requested and key is "credits", returns ["cast"].
func getRelevantFields(key string, fields []string) []string {
	var matches []string
	for _, f := range fields {
		if f == key {
			matches = append(matches, "")
		} else if strings.HasPrefix(f, key+".") {
			matches = append(matches, f[len(key)+1:])
		}
	}
	return matches
}

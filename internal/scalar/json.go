package scalar

import (
	"encoding/json"
	"fmt"
	"io"
)

// JSON is a custom scalar that marshals/unmarshals arbitrary JSON values.
// It is backed by json.RawMessage so GORM can scan JSONB columns directly.
type JSON json.RawMessage

func (j JSON) MarshalGQL(w io.Writer) {
	if len(j) == 0 {
		_, _ = io.WriteString(w, "null")
		return
	}
	_, _ = w.Write(j)
}

func (j *JSON) UnmarshalGQL(v interface{}) error {
	switch val := v.(type) {
	case string:
		*j = JSON(val)
	case []byte:
		*j = JSON(val)
	case map[string]interface{}:
		b, err := json.Marshal(val)
		if err != nil {
			return err
		}
		*j = JSON(b)
	case nil:
		*j = nil
	default:
		return fmt.Errorf("JSON scalar: unsupported type %T", v)
	}
	return nil
}

func (j JSON) RawMessage() json.RawMessage {
	return json.RawMessage(j)
}

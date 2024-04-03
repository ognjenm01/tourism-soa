package model

import "encoding/json"

func ToEnum(status string) SystemStatus {
	switch status {
	case "DRAFT":
		return DRAFT
	case "PUBLISHED":
		return PUBLISHED
	case "CLOSED":
		return CLOSED
	default:
		return -1
	}
}

func ToString(status SystemStatus) string {
	switch status {
	case 0:
		return "DRAFT"
	case 1:
		return "PUBLISHED"
	case 2:
		return "CLOSED"
	default:
		return "YE"
	}
}

func (s *SystemStatus) UnmarshalJSON(bytes []byte) error {
	var status string
	err := json.Unmarshal(bytes, &status)
	if err != nil {
		return err
	}

	*s = ToEnum(status)

	return err
}

func (s SystemStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(ToString(s))
}

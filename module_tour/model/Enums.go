package model

import "encoding/json"

func String2TourDifficulty(difficulty string) TourDifficulty {
	switch difficulty {
	case "EASY":
		return 0
	case "MEDIUM":
		return 1
	case "HARD":
		return 2
	case "EXTREME":
		return 3
	default:
		return -1
	}
}

func TourDifficulty2String(difficulty TourDifficulty) string {
	switch difficulty {
	case 0:
		return "EASY"
	case 1:
		return "MEDIUM"
	case 2:
		return "HARD"
	case 3:
		return "EXTREME"
	default:
		return "UNDEFINED"
	}
}

func (t *TourDifficulty) UnmarshalJSON(bytes []byte) error {
	var difficulty string
	err := json.Unmarshal(bytes, &difficulty)
	if err != nil {
		return err
	}
	*t = String2TourDifficulty(difficulty)
	return err
}

func (t *TourDifficulty) MarshalJSON() ([]byte, error) {
	return json.Marshal(TourDifficulty2String(*t))
}

func string2TourStatus(status string) TourStatus {
	switch status {
	case "DRAFT":
		return 0
	case "PUBLISHED":
		return 1
	case "ARCHIVED":
		return 2
	case "DISABLED":
		return 3
	case "CUSTOM":
		return 4
	case "CAMPAIGN":
		return 5
	default:
		return -1
	}
}

func tourStatus2String(status TourStatus) string {
	switch status {
	case 0:
		return "DRAFT"
	case 1:
		return "PUBLISHED"
	case 2:
		return "ARCHIVED"
	case 3:
		return "DISABLED"
	case 4:
		return "CUSTOM"
	case 5:
		return "CAMPAIGN"
	default:
		return "UNDEFINED"
	}
}

func (t *TourStatus) UnmarshalJSON(bytes []byte) error {
	var status string
	err := json.Unmarshal(bytes, &status)
	if err != nil {
		return err
	}
	*t = string2TourStatus(status)
	return err
}

func (t *TourStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(tourStatus2String(*t))
}

func string2TransportType(transportType string) TransportType {
	switch transportType {
	case "WALK":
		return 0
	case "CAR":
		return 1
	case "BIKE":
		return 2
	case "BOAT":
		return 3
	default:
		return -1
	}
}

func transportType2string(transportType TransportType) string {
	switch transportType {
	case 0:
		return "WALK"
	case 1:
		return "CAR"
	case 2:
		return "BIKE"
	case 3:
		return "BOAT"
	default:
		return "UNDEFINED"
	}
}

func (t *TransportType) MarshalJSON() ([]byte, error) {
	return json.Marshal(transportType2string(*t))
}

func (t *TransportType) UnmarshalJSON(bytes []byte) error {
	var transportType string
	err := json.Unmarshal(bytes, &transportType)
	if err != nil {
		return err
	}
	*t = string2TransportType(transportType)
	return err
}

func string2TourProgressStatus(status string) TourProgressStatus {
	switch status {
	case "IN_PROGRESS":
		return 0
	case "ABANDONED":
		return 1
	case "COMPLETED":
		return 2
	default:
		return -1
	}
}

func tourProgressStatus2String(status TourProgressStatus) string {
	switch status {
	case 0:
		return "IN_PROGRESS"
	case 1:
		return "ABANDONED"
	case 2:
		return "COMPLETED"
	default:
		return "UNDEFINED"
	}
}

func (tp *TourProgressStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(tourProgressStatus2String(*tp))
}

func (tp *TourProgressStatus) UnmarshalJSON(bytes []byte) error {
	var status string
	err := json.Unmarshal(bytes, &status)
	if err != nil {
		return err
	}
	*tp = string2TourProgressStatus(status)
	return err
}

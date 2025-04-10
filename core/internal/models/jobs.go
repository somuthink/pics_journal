package models

import "encoding/json"

// type JobBase struct {
// 	JobID  string
// 	Prompt string
// }

type FluxJob struct {
	JobID  string
	UserID uint

	InputName string
	Prompt    string
	Typo      string
}

type LlmJob struct {
	JobID  string
	UserID uint

	Prompt string
	Typo   string
}

func (l LlmJob) MarshalBinary() ([]byte, error) {
	return json.Marshal(l)
}

func (l LlmJob) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, l)
}

func (f FluxJob) MarshalBinary() ([]byte, error) {
	return json.Marshal(f)
}

func (f FluxJob) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, f)
}

package models

// type JobBase struct {
// 	JobID  string
// 	Prompt string
// }

type FluxJob struct {
	JobID  string
	UserID uint

	Prompt string
	Typo   string
}

type LlmJob struct {
	JobID  string
	UserID uint

	Prompt string
	Typo   string
}

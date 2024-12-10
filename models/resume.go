package models

import (
	"github.com/gofrs/uuid"
)

type Resume struct {
	Name       string      `json:"name"`
	ResumeUUID uuid.UUID   `json:"resume_uuid"`
	UserUUID   uuid.UUID   `json:"user_uuid"`
	ResumeDesc *ResumeDesc `json:"resume_desc"`
}
type ResumeDesc struct {
	JobExps   []*JobExp  `json:"job_exps"`
	Candidate *Candidate `json:"candidate"`
	Skills    []string   `json:"skills"`
}
type Candidate struct {
	EmploymentType int           `json:"employment_type"`
	WorkingHours   int           `json:"working_hours"`
	WishSalaryFrom int           `json:"wish_salary_from"`
	WishSalaryTo   int           `json:"wish_salary_to"`
	Specialization string        `json:"specialization"`
	CandidateBio   *CandidateBio `json:"candidate_bio"`
}
type CandidateBio struct {
	Gender  int    `json:"gender"`
	Age     int    `json:"age"`
	Address string `json:"address"`
}
type JobExp struct {
	DateFrom int      `json:"date_from"`
	DateTo   int      `json:"date_to"`
	JobName  string   `json:"job_name"`
	JobDesc  *JobDesc `json:"job_desc"`
	JobPos   string   `json:"job_pos"`
	JobSpec  string   `json:"job_spec"`
}
type JobDesc struct {
	Location string `json:"location"`
	Website  string `json:"website"`
}

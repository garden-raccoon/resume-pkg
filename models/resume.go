package models

import (
	proto "github.com/garden-raccoon/resume-pkg/protocols/resume"

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

// Proto is
func (res Resume) Proto() *proto.Resume {

	pb := &proto.Resume{
		ResumeUuid: res.ResumeUUID.Bytes(),
		UserUuid:   res.UserUUID.Bytes(),
		Name:       res.Name,
	}
	resumeDesc := apiResumeDescToProto(res.ResumeDesc)
	pb.ResumeDesc = resumeDesc
	return pb
}

func apiResumeDescToProto(resumeDesc *ResumeDesc) *proto.ResumeDesc {
	return &proto.ResumeDesc{
		JobExps:   jobExpsToProto(resumeDesc.JobExps),
		Candidate: apiCandidateToProto(resumeDesc.Candidate),
		Skills:    resumeDesc.Skills,
	}
}

func apiCandidateToProto(candApi *Candidate) *proto.Candidate {
	return &proto.Candidate{
		EmploymentType: int64(candApi.EmploymentType),
		WorkingHours:   int64(candApi.WorkingHours),
		WishSalaryFrom: int64(candApi.WishSalaryFrom),
		WishSalaryTo:   int64(candApi.WishSalaryTo),
		Specialization: candApi.Specialization,
		CandidateBio:   apiCandidateBioToProto(candApi.CandidateBio.Gender, candApi.CandidateBio.Age, candApi.CandidateBio.Address),
	}
}

func apiCandidateBioToProto(gender, age int, addr string) *proto.CandidateBio {
	return &proto.CandidateBio{
		Gender:  int64(gender),
		Age:     int64(age),
		Address: addr,
	}
}

func jobExpsToProto(jobExpsApi []*JobExp) []*proto.JobExp {
	var jobExps []*proto.JobExp
	for i := range jobExpsApi {
		jobExps = append(jobExps, apiJobExpToProto(jobExpsApi[i]))
	}
	return jobExps
}

func apiJobExpToProto(jobExpApi *JobExp) *proto.JobExp {
	jobExp := &proto.JobExp{
		DateFrom: int64(jobExpApi.DateFrom),
		DateTo:   int64(jobExpApi.DateTo),
		JobName:  jobExpApi.JobName,
		JobPos:   jobExpApi.JobPos,
		JobSpec:  jobExpApi.JobSpec,
		JobDesc:  apiJobDescToProto(jobExpApi.JobDesc.Website, jobExpApi.JobDesc.Location),
	}
	return jobExp
}

func apiJobDescToProto(site, loc string) *proto.JobDesc {
	return &proto.JobDesc{
		Location: loc,
		Website:  site,
	}
}

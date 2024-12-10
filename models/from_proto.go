package models

import (
	proto "github.com/garden-raccoon/resume-pkg/protocols/resume"
	"github.com/gofrs/uuid"
)

func ResumesFromProto(pb *proto.Resumes) (rs []*Resume) {
	for _, b := range pb.Resumes {
		rs = append(rs, ResumeFromProto(b))
	}
	return
}

func ResumeFromProto(pb *proto.Resume) *Resume {
	vac := &Resume{
		Name:       pb.Name,
		ResumeUUID: uuid.FromBytesOrNil(pb.ResumeUuid),
		UserUUID:   uuid.FromBytesOrNil(pb.UserUuid),
		ResumeDesc: resumeDescFromProto(pb.ResumeDesc),
	}
	return vac
}

func resumeDescFromProto(pb *proto.ResumeDesc) *ResumeDesc {
	return &ResumeDesc{
		JobExps:   jobExpsFromProto(pb.JobExps),
		Candidate: candidateFromProto(pb.Candidate),
		Skills:    pb.Skills,
	}
}

func candidateFromProto(pb *proto.Candidate) *Candidate {
	return &Candidate{
		EmploymentType: int(pb.EmploymentType),
		WorkingHours:   int(pb.WorkingHours),
		WishSalaryFrom: int(pb.WishSalaryFrom),
		WishSalaryTo:   int(pb.WishSalaryTo),
		Specialization: pb.Specialization,
		CandidateBio:   candidateBioFromProto(int(pb.CandidateBio.Gender), int(pb.CandidateBio.Age), pb.CandidateBio.Address),
	}
}

func candidateBioFromProto(gender, age int, addr string) *CandidateBio {
	return &CandidateBio{
		Gender:  gender,
		Age:     age,
		Address: addr,
	}
}

func jobExpsFromProto(pb []*proto.JobExp) []*JobExp {
	var jobExps []*JobExp
	for i := range pb {
		jobExps = append(jobExps, jobExpFromProto(pb[i]))
	}
	return jobExps
}

func jobExpFromProto(jobExpApi *proto.JobExp) *JobExp {
	jobExp := &JobExp{
		DateFrom: int(jobExpApi.DateFrom),
		DateTo:   int(jobExpApi.DateTo),
		JobName:  jobExpApi.JobName,
		JobPos:   jobExpApi.JobPos,
		JobSpec:  jobExpApi.JobSpec,
		JobDesc:  jobDescFromProto(jobExpApi.JobDesc.Website, jobExpApi.JobDesc.Location),
	}
	return jobExp
}

func jobDescFromProto(site, loc string) *JobDesc {
	return &JobDesc{
		Location: loc,
		Website:  site,
	}
}

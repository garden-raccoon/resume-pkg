package models

import proto "github.com/garden-raccoon/resume-pkg/protocols/resume"

// Proto is
func (res Resume) Proto() *proto.Resume {

	pb := &proto.Resume{
		ResumeUuid: res.ResumeUUID.Bytes(),
		UserUuid:   res.UserUUID.Bytes(),
		Name:       res.Name,
	}
	resumeDesc := resumeDescToProto(res.ResumeDesc)
	pb.ResumeDesc = resumeDesc
	return pb
}

func resumeDescToProto(resumeDesc *ResumeDesc) *proto.ResumeDesc {
	return &proto.ResumeDesc{
		JobExps:   jobExpsToProto(resumeDesc.JobExps),
		Candidate: candidateToProto(resumeDesc.Candidate),
		Skills:    resumeDesc.Skills,
	}
}

func candidateToProto(candApi *Candidate) *proto.Candidate {
	return &proto.Candidate{
		EmploymentType: int64(candApi.EmploymentType),
		WorkingHours:   int64(candApi.WorkingHours),
		WishSalaryFrom: int64(candApi.WishSalaryFrom),
		WishSalaryTo:   int64(candApi.WishSalaryTo),
		Specialization: candApi.Specialization,
		CandidateBio:   candidateBioToProto(candApi.CandidateBio.Gender, candApi.CandidateBio.Age, candApi.CandidateBio.Address),
	}
}

func candidateBioToProto(gender, age int, addr string) *proto.CandidateBio {
	return &proto.CandidateBio{
		Gender:  int64(gender),
		Age:     int64(age),
		Address: addr,
	}
}

func jobExpsToProto(jobExpsApi []*JobExp) []*proto.JobExp {
	var jobExps []*proto.JobExp
	for i := range jobExpsApi {
		jobExps = append(jobExps, jobExpToProto(jobExpsApi[i]))
	}
	return jobExps
}

func jobExpToProto(jobExpApi *JobExp) *proto.JobExp {
	jobExp := &proto.JobExp{
		DateFrom: int64(jobExpApi.DateFrom),
		DateTo:   int64(jobExpApi.DateTo),
		JobName:  jobExpApi.JobName,
		JobPos:   jobExpApi.JobPos,
		JobSpec:  jobExpApi.JobSpec,
		JobDesc:  jobDescToProto(jobExpApi.JobDesc.Website, jobExpApi.JobDesc.Location),
	}
	return jobExp
}

func jobDescToProto(site, loc string) *proto.JobDesc {
	return &proto.JobDesc{
		Location: loc,
		Website:  site,
	}
}

package resume

import (
	"context"
	"fmt"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"

	"github.com/garden-raccoon/resume-pkg/models"

	proto "github.com/garden-raccoon/resume-pkg/protocols/resume"
)

// timeOut is  hardcoded GRPC requests timeout value
const timeOut = 60

type IResumeAPI interface {
	CreateResume(res *models.Resume) error

	GetResumes() ([]*models.Resume, error)
	// Close GRPC Api connection
	Close() error
}

// Api is profile-service GRPC Api
// structure with client Connection
type Api struct {
	addr    string
	timeout time.Duration
	*grpc.ClientConn
	proto.ResumeServiceClient
}

// New create new Battles Api instance
func New(addr string) (IResumeAPI, error) {
	api := &Api{timeout: timeOut * time.Second}

	if err := api.initConn(addr); err != nil {
		return nil, fmt.Errorf("create Battles Api:  %w", err)
	}

	api.ResumeServiceClient = proto.NewResumeServiceClient(api.ClientConn)
	return api, nil
}

func (api *Api) CreateResume(res *models.Resume) (err error) {
	ctx, cancel := context.WithTimeout(context.Background(), api.timeout)
	defer cancel()

	_, err = api.ResumeServiceClient.CreateResume(ctx, res.Proto())
	if err != nil {
		return fmt.Errorf("create meals api request: %w", err)
	}
	return nil
}

// initConn initialize connection to Grpc servers
func (api *Api) initConn(addr string) (err error) {
	var kacp = keepalive.ClientParameters{
		Time:                10 * time.Second, // send pings every 10 seconds if there is no activity
		Timeout:             time.Second,      // wait 1 second for ping ack before considering the connection dead
		PermitWithoutStream: true,             // send pings even without active streams
	}

	api.ClientConn, err = grpc.Dial(addr, grpc.WithInsecure(), grpc.WithKeepaliveParams(kacp))
	return
}
func (api *Api) GetResumes() ([]*models.Resume, error) {
	ctx, cancel := context.WithTimeout(context.Background(), api.timeout)
	defer cancel()

	var resp *proto.Resumes
	empty := new(proto.ResumeDbEmpty)
	resp, err := api.ResumeServiceClient.GetResumes(ctx, empty)
	if err != nil {
		return nil, fmt.Errorf("GetOrders api request: %w", err)
	}

	resumes := models.ResumesFromProto(resp)

	return resumes, nil
}

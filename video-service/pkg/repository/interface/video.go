package interfaces

import "videoservice/pkg/pb"

type VideoRepo interface {
	CreateVideoid(string) error
	FindAllVideo() ([]*pb.VideoID, error)
}

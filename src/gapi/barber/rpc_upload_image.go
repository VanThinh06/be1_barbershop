package gapi

import (
	"barbershop/src/pb/barber"
	"bytes"
	"context"

	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *Server) UploadImageService(ctx context.Context, req *barber.UploadImageRequest) (*barber.UploadImageResponse, error) {
	imgBuffer := bytes.NewBuffer(req.ImageData)

	uploadParams := uploader.UploadParams{
		PublicID:       req.FileName,
		Transformation: "c_fill,h_300,w_300,q_auto:low",
	}

	resp, err := server.ServiceApp.CloudinaryApp.Upload.Upload(ctx, imgBuffer, uploadParams)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "internal")
	}

	rsp := &barber.UploadImageResponse{
		ImageUrl: resp.URL,
		Message:  resp.Error.Message,
	}
	return rsp, nil
}

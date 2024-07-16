package gapi

import (
	"barbershop/src/pb/barber"
	"bytes"
	"context"
	"errors"

	"github.com/cloudinary/cloudinary-go/v2/api/admin"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *Server) UploadImageServiceItem(ctx context.Context, req *barber.UploadImageRequest) (*barber.UploadImageResponse, error) {
	imgBuffer := bytes.NewBuffer(req.ImageData)

	uploadParams := uploader.UploadParams{
		PublicID:       req.FileName,
		AssetFolder:    req.BarberShopId + "/" + req.FolderName,
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
	if resp.Error.Message != "" {
		return rsp, errors.New(resp.Error.Message)
	}
	return rsp, nil
}

func (server *Server) DeleteImage(ctx context.Context, publicID string) error {

	delParams := admin.DeleteAssetsParams{
		PublicIDs: []string{publicID},
	}

	_, err := server.ServiceApp.CloudinaryApp.Admin.DeleteAssets(ctx, delParams)
	if err != nil {
		return err
	}
	return nil
}

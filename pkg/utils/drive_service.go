package utils

import (
	"context"
	"fmt"
	"log"
	"mime/multipart"

	"google.golang.org/api/drive/v3"
	"google.golang.org/api/option"
)

// DriveService mendefinisikan kontrak untuk layanan Google Drive
type DriveService interface {
	UploadToDrive(file *multipart.FileHeader, folderID string) (string, error)
	DeleteFile(fileID string) error
	GetFileURL(fileID string) string
	GetViewURL(fileID string) string
}

// DriveServiceImpl implementasi konkret dari DriveService
type DriveServiceImpl struct {
	service *drive.Service
}

// NewDriveService membuat instance baru DriveService
func NewDriveService(serviceAccountKeyPath string) (DriveService, error) {
	ctx := context.Background()

	// Inisialisasi service menggunakan service account
	srv, err := drive.NewService(ctx,
		option.WithCredentialsFile(serviceAccountKeyPath),
		option.WithScopes(drive.DriveScope),
	)
	if err != nil {
		return nil, fmt.Errorf("gagal membuat Drive service: %v", err)
	}

	return &DriveServiceImpl{service: srv}, nil
}

// UploadToDrive mengupload file ke Google Drive
func (d *DriveServiceImpl) UploadToDrive(file *multipart.FileHeader, folderID string) (string, error) {
	// Buka file
	src, err := file.Open()
	if err != nil {
		return "", fmt.Errorf("gagal membuka file: %v", err)
	}
	defer src.Close()

	// Buat metadata file
	driveFile := &drive.File{
		Name:     file.Filename,
		Parents:  []string{folderID},
		MimeType: file.Header.Get("Content-Type"),
	}

	// Upload file
	res, err := d.service.Files.Create(driveFile).Media(src).Do()
	if err != nil {
		return "", fmt.Errorf("gagal mengupload file: %v", err)
	}

	// Set permission public (opsional)
	if err := d.setPublicPermission(res.Id); err != nil {
		log.Printf("Peringatan: Gagal mengatur permission publik: %v", err)
	}

	return res.Id, nil
}

// DeleteFile menghapus file dari Google Drive
func (d *DriveServiceImpl) DeleteFile(fileID string) error {
	err := d.service.Files.Delete(fileID).Do()
	if err != nil {
		return fmt.Errorf("gagal menghapus file: %v", err)
	}
	return nil
}

// setPublicPermission mengatur akses publik ke file
func (d *DriveServiceImpl) setPublicPermission(fileID string) error {
	permission := &drive.Permission{
		Type: "anyone",
		Role: "reader",
	}
	_, err := d.service.Permissions.Create(fileID, permission).Do()
	return err
}

// GetFileURL menghasilkan URL untuk mengunduh file
func (d *DriveServiceImpl) GetFileURL(fileID string) string {
	return fmt.Sprintf("https://drive.google.com/uc?id=%s&export=download", fileID)
}

// GetViewURL menghasilkan URL untuk melihat file di Google Drive
func (d *DriveServiceImpl) GetViewURL(fileID string) string {
	return fmt.Sprintf("https://drive.google.com/file/d/%s/view", fileID)
}

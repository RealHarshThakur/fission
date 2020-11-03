package storagesvc

import (
	"os"
	"path"

	"github.com/graymeta/stow"
	stowgs "github.com/graymeta/stow/google"
	uuid "github.com/satori/go.uuid"
)

type (
	gcsStorage struct {
		storageType     StorageType
		configJSON      string
		configProjectID string
		ConfigScopes    string
		bucketName      string
		subDir          string
		region          string
	}
)

// NewGCSStorage returns a new s3 storage struct
func NewGCSStorage(args ...string) Storage {

	configJSON := os.Getenv("GOOGLE_APPLICATION_CREDENTIALS_JSON")
	configProjectID := os.Getenv("PROJECT_ID")
	bucketName := os.Getenv("STORAGE_GCS_BUCKET_NAME")
	subDir := os.Getenv("STORAGE_GCS_SUB_DIR")
	region := os.Getenv("STORAGE_GCS_REGION")

	return gcsStorage{
		storageType:     StorageTypeGCS,
		configJSON:      configJSON,
		configProjectID: configProjectID,
		bucketName:      bucketName,
		subDir:          subDir,
		region:          region,
	}
}

func (gcs gcsStorage) getStorageType() StorageType {
	return gcs.storageType
}

func (gcs gcsStorage) getContainerName() string {
	return gcs.bucketName
}

func (gcs gcsStorage) getUploadFileName() string {
	uploadName := uuid.NewV4().String()
	return path.Join(gcs.subDir, uploadName)
}

func (gcs gcsStorage) dial() (stow.Location, error) {
	kind := "gcs"
	config := stow.ConfigMap{
		stowgs.ConfigJSON:      gcs.configJSON,
		stowgs.ConfigProjectId: gcs.configProjectID,
		stowgs.ConfigScopes:    gcs.ConfigScopes,
	}
	stowLoc, err := stow.Dial(kind, config)
	if err != nil {
		return nil, err
	}
	return stowLoc, nil
}

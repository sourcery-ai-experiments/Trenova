// Package minio provides a client for managing media storage using Minio.
package minio

import (
	"context"

	"github.com/minio/minio-go/v7"
	"github.com/pkg/errors"
)

// Client represents a client to interact with the Minio server.
type Client struct {
	client   *minio.Client  // client is the Minio client.
	options  *minio.Options // options are the configurations for the Minio client.
	endpoint string         // endpoint is the Minio server URL.
}

// NewClient initializes a new Minio client.
// It panics if the Minio client cannot be created.
func NewClient(endpoint string, opts *minio.Options) *Client {
	mClient := &Client{
		options:  opts,
		endpoint: endpoint,
	}
	mClient.initialize()
	return mClient
}

// initialize sets up the Minio client. It panics if the client cannot be created.
func (c *Client) initialize() {
	client, err := minio.New(c.endpoint, c.options)
	if err != nil {
		panic(errors.Wrap(err, "failed to create Minio client"))
	}
	c.client = client
}

func (c *Client) Ping(ctx context.Context) error {
	_, err := c.client.ListBuckets(ctx)
	return err
}

// CreateBucket ensures a bucket is available for storing media.
//
// It creates the bucket if it does not already exist.
func (c *Client) CreateBucket(ctx context.Context, bucketName, location string) error {
	exists, err := c.client.BucketExists(ctx, bucketName)
	if err != nil {
		return errors.Wrap(err, "checking if bucket exists failed")
	}
	if !exists {
		if err = c.client.MakeBucket(ctx, bucketName, minio.MakeBucketOptions{Region: location}); err != nil {
			return errors.Wrap(err, "creating bucket failed")
		}
	}
	return nil
}

// UploadMedia uploads a file to the specified bucket and object name.
//
// It returns UploadInfo or an error if the upload fails.
func (c *Client) UploadMedia(ctx context.Context, bucketName, filePath, objectName string) (minio.UploadInfo, error) {
	contentType := "application/octet-stream"
	ui, err := c.client.FPutObject(ctx, bucketName, objectName, filePath, minio.PutObjectOptions{ContentType: contentType})
	if err != nil {
		return minio.UploadInfo{}, errors.Wrap(err, "upload failed")
	}
	return ui, nil
}

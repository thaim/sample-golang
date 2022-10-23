package main

import (
	"bytes"
	"context"
	"io/ioutil"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/s3"
)

// ...
type mockGetObjectAPI func(ctx context.Context, params *s3.GetObjectInput, optFns ...func(*s3.Options)) (*s3.GetObjectOutput, error)

func (m mockGetObjectAPI) GetObject(ctx context.Context, params *s3.GetObjectInput, optFns ...func(*s3.Options)) (*s3.GetObjectOutput, error) {
	return m(ctx, params, optFns...)
}

func TestGetObjectFromS3(t *testing.T) {
	mockClient := func(t *testing.T) S3GetObjectAPI {
		return mockGetObjectAPI(func(ctx context.Context, params *s3.GetObjectInput, optFns ...func(*s3.Options)) (*s3.GetObjectOutput, error) {
			t.Helper()
			if params.Bucket == nil {
				t.Fatal("expect bucket to not be nil")
			}
			if e, a := "fooBucket", *params.Bucket; e != a {
				t.Errorf("expect %v, got %v", e, a)
			}
			if params.Key == nil {
				t.Fatal("expect key to not be nil")
			}
			if e, a := "barKey", *params.Key; e != a {
				t.Errorf("expect %v, got %v", e, a)
			}

			return &s3.GetObjectOutput{
				Body: ioutil.NopCloser(bytes.NewReader([]byte("this is the body foo bar baz"))),
			}, nil
		})
	}

	cases := []struct {
		name string
		client func(t *testing.T) S3GetObjectAPI
		bucket string
		key	string
		expect []byte
	}{
		{
			name: "return content",
			client: mockClient,
			bucket: "fooBucket",
			key:	"barKey",
			expect: []byte("this is the body foo bar baz"),
		},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			ctx := context.TODO()
			content, err := GetObjectFromS3(ctx, tt.client(t), tt.bucket, tt.key)
			if err != nil {
				t.Fatalf("expect no error, got %v", err)
			}
			if e, a := tt.expect, content; bytes.Compare(e, a) != 0 {
				t.Errorf("expect %v, got %v", e, a)
			}
		})
	}
}

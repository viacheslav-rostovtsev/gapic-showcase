// Copyright 2019 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// AUTO-GENERATED CODE. DO NOT EDIT.

package client

import (
	"context"
	"math"
	"time"

	"github.com/golang/protobuf/proto"
	genprotopb "github.com/googleapis/gapic-showcase/server/genproto"
	gax "github.com/googleapis/gax-go/v2"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
	"google.golang.org/api/transport"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
)

// TestingCallOptions contains the retry settings for each method of TestingClient.
type TestingCallOptions struct {
	CreateSession []gax.CallOption
	GetSession    []gax.CallOption
	ListSessions  []gax.CallOption
	DeleteSession []gax.CallOption
	ReportSession []gax.CallOption
	ListTests     []gax.CallOption
	DeleteTest    []gax.CallOption
	VerifyTest    []gax.CallOption
}

func defaultTestingClientOptions() []option.ClientOption {
	return []option.ClientOption{
		option.WithEndpoint("localhost:7469:443"),
		option.WithScopes(DefaultAuthScopes()...),
	}
}

func defaultTestingCallOptions() *TestingCallOptions {
	backoff := gax.Backoff{
		Initial:    100 * time.Millisecond,
		Max:        time.Minute,
		Multiplier: 1.3,
	}

	idempotent := []gax.CallOption{
		gax.WithRetry(func() gax.Retryer {
			return gax.OnCodes([]codes.Code{
				codes.Aborted,
				codes.Internal,
				codes.Unavailable,
				codes.Unknown,
			}, backoff)
		}),
	}

	nonidempotent := []gax.CallOption{
		gax.WithRetry(func() gax.Retryer {
			return gax.OnCodes([]codes.Code{
				codes.Unavailable,
			}, backoff)
		}),
	}

	return &TestingCallOptions{
		GetSession:    idempotent,
		ListSessions:  idempotent,
		ListTests:     idempotent,
		CreateSession: nonidempotent,
		DeleteSession: nonidempotent,
		ReportSession: nonidempotent,
		DeleteTest:    nonidempotent,
		VerifyTest:    nonidempotent,
	}
}

// TestingClient is a client for interacting with  API.
//
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
type TestingClient struct {
	// The connection to the service.
	conn *grpc.ClientConn

	// The gRPC API client.
	testingClient genprotopb.TestingClient

	// The call options for this service.
	CallOptions *TestingCallOptions

	// The x-goog-* metadata to be sent with each request.
	xGoogMetadata metadata.MD
}

// NewTestingClient creates a new testing client.
//
// A service to facilitate running discrete sets of tests
// against Showcase.
func NewTestingClient(ctx context.Context, opts ...option.ClientOption) (*TestingClient, error) {
	conn, err := transport.DialGRPC(ctx, append(defaultTestingClientOptions(), opts...)...)
	if err != nil {
		return nil, err
	}
	c := &TestingClient{
		conn:        conn,
		CallOptions: defaultTestingCallOptions(),

		testingClient: genprotopb.NewTestingClient(conn),
	}
	c.setGoogleClientInfo()

	return c, nil
}

// Connection returns the client's connection to the API service.
func (c *TestingClient) Connection() *grpc.ClientConn {
	return c.conn
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *TestingClient) Close() error {
	return c.conn.Close()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *TestingClient) setGoogleClientInfo(keyval ...string) {
	kv := append([]string{"gl-go", versionGo()}, keyval...)
	kv = append(kv, "gapic", versionClient, "gax", gax.Version, "grpc", grpc.Version)
	c.xGoogMetadata = metadata.Pairs("x-goog-api-client", gax.XGoogHeader(kv...))
}

// CreateSession creates a new testing session.
func (c *TestingClient) CreateSession(ctx context.Context, req *genprotopb.CreateSessionRequest, opts ...gax.CallOption) (*genprotopb.Session, error) {
	ctx = insertMetadata(ctx, c.xGoogMetadata)
	opts = append(c.CallOptions.CreateSession[0:len(c.CallOptions.CreateSession):len(c.CallOptions.CreateSession)], opts...)
	var resp *genprotopb.Session
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.testingClient.CreateSession(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// GetSession gets a testing session.
func (c *TestingClient) GetSession(ctx context.Context, req *genprotopb.GetSessionRequest, opts ...gax.CallOption) (*genprotopb.Session, error) {
	ctx = insertMetadata(ctx, c.xGoogMetadata)
	opts = append(c.CallOptions.GetSession[0:len(c.CallOptions.GetSession):len(c.CallOptions.GetSession)], opts...)
	var resp *genprotopb.Session
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.testingClient.GetSession(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// ListSessions lists the current test sessions.
func (c *TestingClient) ListSessions(ctx context.Context, req *genprotopb.ListSessionsRequest, opts ...gax.CallOption) *SessionIterator {
	ctx = insertMetadata(ctx, c.xGoogMetadata)
	opts = append(c.CallOptions.ListSessions[0:len(c.CallOptions.ListSessions):len(c.CallOptions.ListSessions)], opts...)
	it := &SessionIterator{}
	req = proto.Clone(req).(*genprotopb.ListSessionsRequest)
	it.InternalFetch = func(pageSize int, pageToken string) ([]*genprotopb.Session, string, error) {
		var resp *genprotopb.ListSessionsResponse
		req.PageToken = pageToken
		if pageSize > math.MaxInt32 {
			req.PageSize = math.MaxInt32
		} else {
			req.PageSize = int32(pageSize)
		}
		err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
			var err error
			resp, err = c.testingClient.ListSessions(ctx, req, settings.GRPC...)
			return err
		}, opts...)
		if err != nil {
			return nil, "", err
		}
		return resp.Sessions, resp.NextPageToken, nil
	}
	fetch := func(pageSize int, pageToken string) (string, error) {
		items, nextPageToken, err := it.InternalFetch(pageSize, pageToken)
		if err != nil {
			return "", err
		}
		it.items = append(it.items, items...)
		return nextPageToken, nil
	}
	it.pageInfo, it.nextFunc = iterator.NewPageInfo(fetch, it.bufLen, it.takeBuf)
	it.pageInfo.MaxSize = int(req.PageSize)
	return it
}

// DeleteSession delete a test session.
func (c *TestingClient) DeleteSession(ctx context.Context, req *genprotopb.DeleteSessionRequest, opts ...gax.CallOption) error {
	ctx = insertMetadata(ctx, c.xGoogMetadata)
	opts = append(c.CallOptions.DeleteSession[0:len(c.CallOptions.DeleteSession):len(c.CallOptions.DeleteSession)], opts...)
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		_, err = c.testingClient.DeleteSession(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	return err
}

// ReportSession report on the status of a session.
// This generates a report detailing which tests have been completed,
// and an overall rollup.
func (c *TestingClient) ReportSession(ctx context.Context, req *genprotopb.ReportSessionRequest, opts ...gax.CallOption) (*genprotopb.ReportSessionResponse, error) {
	ctx = insertMetadata(ctx, c.xGoogMetadata)
	opts = append(c.CallOptions.ReportSession[0:len(c.CallOptions.ReportSession):len(c.CallOptions.ReportSession)], opts...)
	var resp *genprotopb.ReportSessionResponse
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.testingClient.ReportSession(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// ListTests list the tests of a sessesion.
func (c *TestingClient) ListTests(ctx context.Context, req *genprotopb.ListTestsRequest, opts ...gax.CallOption) *TestIterator {
	ctx = insertMetadata(ctx, c.xGoogMetadata)
	opts = append(c.CallOptions.ListTests[0:len(c.CallOptions.ListTests):len(c.CallOptions.ListTests)], opts...)
	it := &TestIterator{}
	req = proto.Clone(req).(*genprotopb.ListTestsRequest)
	it.InternalFetch = func(pageSize int, pageToken string) ([]*genprotopb.Test, string, error) {
		var resp *genprotopb.ListTestsResponse
		req.PageToken = pageToken
		if pageSize > math.MaxInt32 {
			req.PageSize = math.MaxInt32
		} else {
			req.PageSize = int32(pageSize)
		}
		err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
			var err error
			resp, err = c.testingClient.ListTests(ctx, req, settings.GRPC...)
			return err
		}, opts...)
		if err != nil {
			return nil, "", err
		}
		return resp.Tests, resp.NextPageToken, nil
	}
	fetch := func(pageSize int, pageToken string) (string, error) {
		items, nextPageToken, err := it.InternalFetch(pageSize, pageToken)
		if err != nil {
			return "", err
		}
		it.items = append(it.items, items...)
		return nextPageToken, nil
	}
	it.pageInfo, it.nextFunc = iterator.NewPageInfo(fetch, it.bufLen, it.takeBuf)
	it.pageInfo.MaxSize = int(req.PageSize)
	return it
}

// DeleteTest explicitly decline to implement a test.
//
// This removes the test from subsequent ListTests calls, and
// attempting to do the test will error.
//
// This method will error if attempting to delete a required test.
func (c *TestingClient) DeleteTest(ctx context.Context, req *genprotopb.DeleteTestRequest, opts ...gax.CallOption) error {
	ctx = insertMetadata(ctx, c.xGoogMetadata)
	opts = append(c.CallOptions.DeleteTest[0:len(c.CallOptions.DeleteTest):len(c.CallOptions.DeleteTest)], opts...)
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		_, err = c.testingClient.DeleteTest(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	return err
}

// VerifyTest register a response to a test.
//
// In cases where a test involves registering a final answer at the
// end of the test, this method provides the means to do so.
func (c *TestingClient) VerifyTest(ctx context.Context, req *genprotopb.VerifyTestRequest, opts ...gax.CallOption) (*genprotopb.VerifyTestResponse, error) {
	ctx = insertMetadata(ctx, c.xGoogMetadata)
	opts = append(c.CallOptions.VerifyTest[0:len(c.CallOptions.VerifyTest):len(c.CallOptions.VerifyTest)], opts...)
	var resp *genprotopb.VerifyTestResponse
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.testingClient.VerifyTest(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// SessionIterator manages a stream of *genprotopb.Session.
type SessionIterator struct {
	items    []*genprotopb.Session
	pageInfo *iterator.PageInfo
	nextFunc func() error

	// InternalFetch is for use by the Google Cloud Libraries only.
	// It is not part of the stable interface of this package.
	//
	// InternalFetch returns results from a single call to the underlying RPC.
	// The number of results is no greater than pageSize.
	// If there are no more results, nextPageToken is empty and err is nil.
	InternalFetch func(pageSize int, pageToken string) (results []*genprotopb.Session, nextPageToken string, err error)
}

// PageInfo supports pagination. See the google.golang.org/api/iterator package for details.
func (it *SessionIterator) PageInfo() *iterator.PageInfo {
	return it.pageInfo
}

// Next returns the next result. Its second return value is iterator.Done if there are no more
// results. Once Next returns Done, all subsequent calls will return Done.
func (it *SessionIterator) Next() (*genprotopb.Session, error) {
	var item *genprotopb.Session
	if err := it.nextFunc(); err != nil {
		return item, err
	}
	item = it.items[0]
	it.items = it.items[1:]
	return item, nil
}

func (it *SessionIterator) bufLen() int {
	return len(it.items)
}

func (it *SessionIterator) takeBuf() interface{} {
	b := it.items
	it.items = nil
	return b
}

// TestIterator manages a stream of *genprotopb.Test.
type TestIterator struct {
	items    []*genprotopb.Test
	pageInfo *iterator.PageInfo
	nextFunc func() error

	// InternalFetch is for use by the Google Cloud Libraries only.
	// It is not part of the stable interface of this package.
	//
	// InternalFetch returns results from a single call to the underlying RPC.
	// The number of results is no greater than pageSize.
	// If there are no more results, nextPageToken is empty and err is nil.
	InternalFetch func(pageSize int, pageToken string) (results []*genprotopb.Test, nextPageToken string, err error)
}

// PageInfo supports pagination. See the google.golang.org/api/iterator package for details.
func (it *TestIterator) PageInfo() *iterator.PageInfo {
	return it.pageInfo
}

// Next returns the next result. Its second return value is iterator.Done if there are no more
// results. Once Next returns Done, all subsequent calls will return Done.
func (it *TestIterator) Next() (*genprotopb.Test, error) {
	var item *genprotopb.Test
	if err := it.nextFunc(); err != nil {
		return item, err
	}
	item = it.items[0]
	it.items = it.items[1:]
	return item, nil
}

func (it *TestIterator) bufLen() int {
	return len(it.items)
}

func (it *TestIterator) takeBuf() interface{} {
	b := it.items
	it.items = nil
	return b
}
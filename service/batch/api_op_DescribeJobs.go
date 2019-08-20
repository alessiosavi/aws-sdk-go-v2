// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package batch

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/batch-2016-08-10/DescribeJobsRequest
type DescribeJobsInput struct {
	_ struct{} `type:"structure"`

	// A list of up to 100 job IDs.
	//
	// Jobs is a required field
	Jobs []string `locationName:"jobs" type:"list" required:"true"`
}

// String returns the string representation
func (s DescribeJobsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeJobsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribeJobsInput"}

	if s.Jobs == nil {
		invalidParams.Add(aws.NewErrParamRequired("Jobs"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DescribeJobsInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/x-amz-json-1.1"), protocol.Metadata{})

	if s.Jobs != nil {
		v := s.Jobs

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "jobs", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddValue(protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v1)})
		}
		ls0.End()

	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/batch-2016-08-10/DescribeJobsResponse
type DescribeJobsOutput struct {
	_ struct{} `type:"structure"`

	// The list of jobs.
	Jobs []JobDetail `locationName:"jobs" type:"list"`
}

// String returns the string representation
func (s DescribeJobsOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DescribeJobsOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.Jobs != nil {
		v := s.Jobs

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "jobs", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	return nil
}

const opDescribeJobs = "DescribeJobs"

// DescribeJobsRequest returns a request value for making API operation for
// AWS Batch.
//
// Describes a list of AWS Batch jobs.
//
//    // Example sending a request using DescribeJobsRequest.
//    req := client.DescribeJobsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/batch-2016-08-10/DescribeJobs
func (c *Client) DescribeJobsRequest(input *DescribeJobsInput) DescribeJobsRequest {
	op := &aws.Operation{
		Name:       opDescribeJobs,
		HTTPMethod: "POST",
		HTTPPath:   "/v1/describejobs",
	}

	if input == nil {
		input = &DescribeJobsInput{}
	}

	req := c.newRequest(op, input, &DescribeJobsOutput{})
	return DescribeJobsRequest{Request: req, Input: input, Copy: c.DescribeJobsRequest}
}

// DescribeJobsRequest is the request type for the
// DescribeJobs API operation.
type DescribeJobsRequest struct {
	*aws.Request
	Input *DescribeJobsInput
	Copy  func(*DescribeJobsInput) DescribeJobsRequest
}

// Send marshals and sends the DescribeJobs API request.
func (r DescribeJobsRequest) Send(ctx context.Context) (*DescribeJobsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeJobsResponse{
		DescribeJobsOutput: r.Request.Data.(*DescribeJobsOutput),
		response:           &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DescribeJobsResponse is the response type for the
// DescribeJobs API operation.
type DescribeJobsResponse struct {
	*DescribeJobsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeJobs request.
func (r *DescribeJobsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
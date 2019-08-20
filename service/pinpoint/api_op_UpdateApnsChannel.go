// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package pinpoint

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/pinpoint-2016-12-01/UpdateApnsChannelRequest
type UpdateApnsChannelInput struct {
	_ struct{} `type:"structure" payload:"APNSChannelRequest"`

	// Specifies the status and settings of the APNs (Apple Push Notification service)
	// channel for an application.
	//
	// APNSChannelRequest is a required field
	APNSChannelRequest *APNSChannelRequest `type:"structure" required:"true"`

	// ApplicationId is a required field
	ApplicationId *string `location:"uri" locationName:"application-id" type:"string" required:"true"`
}

// String returns the string representation
func (s UpdateApnsChannelInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *UpdateApnsChannelInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "UpdateApnsChannelInput"}

	if s.APNSChannelRequest == nil {
		invalidParams.Add(aws.NewErrParamRequired("APNSChannelRequest"))
	}

	if s.ApplicationId == nil {
		invalidParams.Add(aws.NewErrParamRequired("ApplicationId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s UpdateApnsChannelInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/x-amz-json-1.1"), protocol.Metadata{})

	if s.ApplicationId != nil {
		v := *s.ApplicationId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "application-id", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.APNSChannelRequest != nil {
		v := s.APNSChannelRequest

		metadata := protocol.Metadata{}
		e.SetFields(protocol.PayloadTarget, "APNSChannelRequest", v, metadata)
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/pinpoint-2016-12-01/UpdateApnsChannelResponse
type UpdateApnsChannelOutput struct {
	_ struct{} `type:"structure" payload:"APNSChannelResponse"`

	// Provides information about the status and settings of the APNs (Apple Push
	// Notification service) channel for an application.
	//
	// APNSChannelResponse is a required field
	APNSChannelResponse *APNSChannelResponse `type:"structure" required:"true"`
}

// String returns the string representation
func (s UpdateApnsChannelOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s UpdateApnsChannelOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.APNSChannelResponse != nil {
		v := s.APNSChannelResponse

		metadata := protocol.Metadata{}
		e.SetFields(protocol.PayloadTarget, "APNSChannelResponse", v, metadata)
	}
	return nil
}

const opUpdateApnsChannel = "UpdateApnsChannel"

// UpdateApnsChannelRequest returns a request value for making API operation for
// Amazon Pinpoint.
//
// Updates the APNs channel settings for an application.
//
//    // Example sending a request using UpdateApnsChannelRequest.
//    req := client.UpdateApnsChannelRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/pinpoint-2016-12-01/UpdateApnsChannel
func (c *Client) UpdateApnsChannelRequest(input *UpdateApnsChannelInput) UpdateApnsChannelRequest {
	op := &aws.Operation{
		Name:       opUpdateApnsChannel,
		HTTPMethod: "PUT",
		HTTPPath:   "/v1/apps/{application-id}/channels/apns",
	}

	if input == nil {
		input = &UpdateApnsChannelInput{}
	}

	req := c.newRequest(op, input, &UpdateApnsChannelOutput{})
	return UpdateApnsChannelRequest{Request: req, Input: input, Copy: c.UpdateApnsChannelRequest}
}

// UpdateApnsChannelRequest is the request type for the
// UpdateApnsChannel API operation.
type UpdateApnsChannelRequest struct {
	*aws.Request
	Input *UpdateApnsChannelInput
	Copy  func(*UpdateApnsChannelInput) UpdateApnsChannelRequest
}

// Send marshals and sends the UpdateApnsChannel API request.
func (r UpdateApnsChannelRequest) Send(ctx context.Context) (*UpdateApnsChannelResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &UpdateApnsChannelResponse{
		UpdateApnsChannelOutput: r.Request.Data.(*UpdateApnsChannelOutput),
		response:                &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// UpdateApnsChannelResponse is the response type for the
// UpdateApnsChannel API operation.
type UpdateApnsChannelResponse struct {
	*UpdateApnsChannelOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// UpdateApnsChannel request.
func (r *UpdateApnsChannelResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ec2

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/ModifyTrafficMirrorFilterNetworkServicesRequest
type ModifyTrafficMirrorFilterNetworkServicesInput struct {
	_ struct{} `type:"structure"`

	// The network service, for example Amazon DNS, that you want to mirror.
	AddNetworkServices []TrafficMirrorNetworkService `locationName:"AddNetworkService" locationNameList:"item" type:"list"`

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have
	// the required permissions, the error response is DryRunOperation. Otherwise,
	// it is UnauthorizedOperation.
	DryRun *bool `type:"boolean"`

	// The network service, for example Amazon DNS, that you no longer want to mirror.
	RemoveNetworkServices []TrafficMirrorNetworkService `locationName:"RemoveNetworkService" locationNameList:"item" type:"list"`

	// The ID of the Traffic Mirror filter.
	//
	// TrafficMirrorFilterId is a required field
	TrafficMirrorFilterId *string `type:"string" required:"true"`
}

// String returns the string representation
func (s ModifyTrafficMirrorFilterNetworkServicesInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ModifyTrafficMirrorFilterNetworkServicesInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ModifyTrafficMirrorFilterNetworkServicesInput"}

	if s.TrafficMirrorFilterId == nil {
		invalidParams.Add(aws.NewErrParamRequired("TrafficMirrorFilterId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/ModifyTrafficMirrorFilterNetworkServicesResult
type ModifyTrafficMirrorFilterNetworkServicesOutput struct {
	_ struct{} `type:"structure"`

	// The Traffic Mirror filter that the network service is associated with.
	TrafficMirrorFilter *TrafficMirrorFilter `locationName:"trafficMirrorFilter" type:"structure"`
}

// String returns the string representation
func (s ModifyTrafficMirrorFilterNetworkServicesOutput) String() string {
	return awsutil.Prettify(s)
}

const opModifyTrafficMirrorFilterNetworkServices = "ModifyTrafficMirrorFilterNetworkServices"

// ModifyTrafficMirrorFilterNetworkServicesRequest returns a request value for making API operation for
// Amazon Elastic Compute Cloud.
//
// Allows or restricts mirroring network services.
//
// By default, Amazon DNS network services are not eligible for Traffic Mirror.
// Use AddNetworkServices to add network services to a Traffic Mirror filter.
// When a network service is added to the Traffic Mirror filter, all traffic
// related to that network service will be mirrored. When you no longer want
// to mirror network services, use RemoveNetworkServices to remove the network
// services from the Traffic Mirror filter.
//
// FFor information about filter rule properties, see Network Services (https://docs.aws.amazon.com/vpc/latest/mirroring/traffic-mirroring-considerations.html#traffic-mirroring-network-services)
// in the Traffic Mirroring User Guide .
//
//    // Example sending a request using ModifyTrafficMirrorFilterNetworkServicesRequest.
//    req := client.ModifyTrafficMirrorFilterNetworkServicesRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/ModifyTrafficMirrorFilterNetworkServices
func (c *Client) ModifyTrafficMirrorFilterNetworkServicesRequest(input *ModifyTrafficMirrorFilterNetworkServicesInput) ModifyTrafficMirrorFilterNetworkServicesRequest {
	op := &aws.Operation{
		Name:       opModifyTrafficMirrorFilterNetworkServices,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &ModifyTrafficMirrorFilterNetworkServicesInput{}
	}

	req := c.newRequest(op, input, &ModifyTrafficMirrorFilterNetworkServicesOutput{})
	return ModifyTrafficMirrorFilterNetworkServicesRequest{Request: req, Input: input, Copy: c.ModifyTrafficMirrorFilterNetworkServicesRequest}
}

// ModifyTrafficMirrorFilterNetworkServicesRequest is the request type for the
// ModifyTrafficMirrorFilterNetworkServices API operation.
type ModifyTrafficMirrorFilterNetworkServicesRequest struct {
	*aws.Request
	Input *ModifyTrafficMirrorFilterNetworkServicesInput
	Copy  func(*ModifyTrafficMirrorFilterNetworkServicesInput) ModifyTrafficMirrorFilterNetworkServicesRequest
}

// Send marshals and sends the ModifyTrafficMirrorFilterNetworkServices API request.
func (r ModifyTrafficMirrorFilterNetworkServicesRequest) Send(ctx context.Context) (*ModifyTrafficMirrorFilterNetworkServicesResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ModifyTrafficMirrorFilterNetworkServicesResponse{
		ModifyTrafficMirrorFilterNetworkServicesOutput: r.Request.Data.(*ModifyTrafficMirrorFilterNetworkServicesOutput),
		response: &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// ModifyTrafficMirrorFilterNetworkServicesResponse is the response type for the
// ModifyTrafficMirrorFilterNetworkServices API operation.
type ModifyTrafficMirrorFilterNetworkServicesResponse struct {
	*ModifyTrafficMirrorFilterNetworkServicesOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ModifyTrafficMirrorFilterNetworkServices request.
func (r *ModifyTrafficMirrorFilterNetworkServicesResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
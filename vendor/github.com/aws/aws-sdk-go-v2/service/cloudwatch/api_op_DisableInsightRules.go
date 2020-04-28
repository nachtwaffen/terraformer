// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package cloudwatch

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type DisableInsightRulesInput struct {
	_ struct{} `type:"structure"`

	// An array of the rule names to disable. If you need to find out the names
	// of your rules, use DescribeInsightRules (https://docs.aws.amazon.com/AmazonCloudWatch/latest/APIReference/API_DescribeInsightRules.html).
	//
	// RuleNames is a required field
	RuleNames []string `type:"list" required:"true"`
}

// String returns the string representation
func (s DisableInsightRulesInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DisableInsightRulesInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DisableInsightRulesInput"}

	if s.RuleNames == nil {
		invalidParams.Add(aws.NewErrParamRequired("RuleNames"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DisableInsightRulesOutput struct {
	_ struct{} `type:"structure"`

	// An array listing the rules that could not be disabled. You cannot disable
	// built-in rules.
	Failures []PartialFailure `type:"list"`
}

// String returns the string representation
func (s DisableInsightRulesOutput) String() string {
	return awsutil.Prettify(s)
}

const opDisableInsightRules = "DisableInsightRules"

// DisableInsightRulesRequest returns a request value for making API operation for
// Amazon CloudWatch.
//
// Disables the specified Contributor Insights rules. When rules are disabled,
// they do not analyze log groups and do not incur costs.
//
//    // Example sending a request using DisableInsightRulesRequest.
//    req := client.DisableInsightRulesRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/monitoring-2010-08-01/DisableInsightRules
func (c *Client) DisableInsightRulesRequest(input *DisableInsightRulesInput) DisableInsightRulesRequest {
	op := &aws.Operation{
		Name:       opDisableInsightRules,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DisableInsightRulesInput{}
	}

	req := c.newRequest(op, input, &DisableInsightRulesOutput{})
	return DisableInsightRulesRequest{Request: req, Input: input, Copy: c.DisableInsightRulesRequest}
}

// DisableInsightRulesRequest is the request type for the
// DisableInsightRules API operation.
type DisableInsightRulesRequest struct {
	*aws.Request
	Input *DisableInsightRulesInput
	Copy  func(*DisableInsightRulesInput) DisableInsightRulesRequest
}

// Send marshals and sends the DisableInsightRules API request.
func (r DisableInsightRulesRequest) Send(ctx context.Context) (*DisableInsightRulesResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DisableInsightRulesResponse{
		DisableInsightRulesOutput: r.Request.Data.(*DisableInsightRulesOutput),
		response:                  &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DisableInsightRulesResponse is the response type for the
// DisableInsightRules API operation.
type DisableInsightRulesResponse struct {
	*DisableInsightRulesOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DisableInsightRules request.
func (r *DisableInsightRulesResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
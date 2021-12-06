// Code generated by smithy-go-codegen DO NOT EDIT.


package grafana

import (
	"context"
	"github.com/aws/smithy-go/middleware"
	"github.com/grafana/grafana/smithy/build/go/grafana/types"
)

// Query for metrics.
func (c *Client) QueryMetricsV2(ctx context.Context, params *QueryMetricsV2Input, optFns ...func(*Options)) (*QueryMetricsV2Output, error) {
	if params == nil { params = &QueryMetricsV2Input{} }
	
	result, metadata, err := c.invokeOperation(ctx, "QueryMetricsV2", params, optFns, c.addOperationQueryMetricsV2Middlewares)
	if err != nil { return nil, err }
	
	out := result.(*QueryMetricsV2Output)
	out.ResultMetadata = metadata
	return out, nil
}

type QueryMetricsV2Input struct {
	
	// This member is required.
	From *string
	
	// This member is required.
	To *string
	
	Debug *bool
	
	noSmithyDocumentSerde
}

type QueryMetricsV2Output struct {
	
	// This member is required.
	Responses map[string]types.DataResponse
	
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
	
	noSmithyDocumentSerde
}

func (c *Client) addOperationQueryMetricsV2Middlewares(stack *middleware.Stack, options Options) (err error) {
	if err = addSetLoggerMiddleware(stack, options); err != nil {
	return err
	}
	if err = addOpQueryMetricsV2ValidationMiddleware(stack); err != nil {
	return err
	}
	return nil
}
// Code generated by smithy-go-codegen DO NOT EDIT.

package sns

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/sns/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates a topic to which notifications can be published. Users can create at
// most 100,000 standard topics (at most 1,000 FIFO topics). For more information,
// see Creating an Amazon SNS topic
// (https://docs.aws.amazon.com/sns/latest/dg/sns-create-topic.html) in the Amazon
// SNS Developer Guide. This action is idempotent, so if the requester already owns
// a topic with the specified name, that topic's ARN is returned without creating a
// new topic.
func (c *Client) CreateTopic(ctx context.Context, params *CreateTopicInput, optFns ...func(*Options)) (*CreateTopicOutput, error) {
	if params == nil {
		params = &CreateTopicInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateTopic", params, optFns, c.addOperationCreateTopicMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateTopicOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Input for CreateTopic action.
type CreateTopicInput struct {

	// The name of the topic you want to create. Constraints: Topic names must be made
	// up of only uppercase and lowercase ASCII letters, numbers, underscores, and
	// hyphens, and must be between 1 and 256 characters long. For a FIFO
	// (first-in-first-out) topic, the name must end with the .fifo suffix.
	//
	// This member is required.
	Name *string

	// A map of attributes with their corresponding values. The following lists the
	// names, descriptions, and values of the special request parameters that the
	// CreateTopic action uses:
	//
	// * DeliveryPolicy – The policy that defines how Amazon
	// SNS retries failed deliveries to HTTP/S endpoints.
	//
	// * DisplayName – The display
	// name to use for a topic with SMS subscriptions.
	//
	// * FifoTopic – Set to true to
	// create a FIFO topic.
	//
	// * Policy – The policy that defines who can access your
	// topic. By default, only the topic owner can publish or subscribe to the
	// topic.
	//
	// The following attribute applies only to server-side encryption
	// (https://docs.aws.amazon.com/sns/latest/dg/sns-server-side-encryption.html):
	//
	// *
	// KmsMasterKeyId – The ID of an Amazon Web Services managed customer master key
	// (CMK) for Amazon SNS or a custom CMK. For more information, see Key Terms
	// (https://docs.aws.amazon.com/sns/latest/dg/sns-server-side-encryption.html#sse-key-terms).
	// For more examples, see KeyId
	// (https://docs.aws.amazon.com/kms/latest/APIReference/API_DescribeKey.html#API_DescribeKey_RequestParameters)
	// in the Key Management Service API Reference.
	//
	// The following attributes apply
	// only to FIFO topics
	// (https://docs.aws.amazon.com/sns/latest/dg/sns-fifo-topics.html):
	//
	// * FifoTopic –
	// When this is set to true, a FIFO topic is created.
	//
	// * ContentBasedDeduplication
	// – Enables content-based deduplication for FIFO topics.
	//
	// * By default,
	// ContentBasedDeduplication is set to false. If you create a FIFO topic and this
	// attribute is false, you must specify a value for the MessageDeduplicationId
	// parameter for the Publish
	// (https://docs.aws.amazon.com/sns/latest/api/API_Publish.html) action.
	//
	// * When
	// you set ContentBasedDeduplication to true, Amazon SNS uses a SHA-256 hash to
	// generate the MessageDeduplicationId using the body of the message (but not the
	// attributes of the message). (Optional) To override the generated value, you can
	// specify a value for the MessageDeduplicationId parameter for the Publish action.
	Attributes map[string]string

	// The body of the policy document you want to use for this topic. You can only add
	// one policy per topic. The policy must be in JSON string format. Length
	// Constraints: Maximum length of 30,720.
	DataProtectionPolicy *string

	// The list of tags to add to a new topic. To be able to tag a topic on creation,
	// you must have the sns:CreateTopic and sns:TagResource permissions.
	Tags []types.Tag

	noSmithyDocumentSerde
}

// Response from CreateTopic action.
type CreateTopicOutput struct {

	// The Amazon Resource Name (ARN) assigned to the created topic.
	TopicArn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateTopicMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpCreateTopic{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpCreateTopic{}, middleware.After)
	if err != nil {
		return err
	}
	if err = addSetLoggerMiddleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddClientRequestIDMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddComputeContentLengthMiddleware(stack); err != nil {
		return err
	}
	if err = addResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = v4.AddComputePayloadSHA256Middleware(stack); err != nil {
		return err
	}
	if err = addRetryMiddlewares(stack, options); err != nil {
		return err
	}
	if err = addHTTPSignerV4Middleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddRawResponseToMetadata(stack); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecordResponseTiming(stack); err != nil {
		return err
	}
	if err = addClientUserAgent(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = addOpCreateTopicValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateTopic(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = addRequestIDRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = addResponseErrorMiddleware(stack); err != nil {
		return err
	}
	if err = addRequestResponseLogging(stack, options); err != nil {
		return err
	}
	return nil
}

func newServiceMetadataMiddleware_opCreateTopic(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "sns",
		OperationName: "CreateTopic",
	}
}

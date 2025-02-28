package lex

import (
	"github.com/awslabs/goformation/v5/cloudformation/policies"
)

// BotAlias_TextLogDestination AWS CloudFormation Resource (AWS::Lex::BotAlias.TextLogDestination)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lex-botalias-textlogdestination.html
type BotAlias_TextLogDestination struct {

	// CloudWatch AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lex-botalias-textlogdestination.html#cfn-lex-botalias-textlogdestination-cloudwatch
	CloudWatch interface{} `json:"CloudWatch,omitempty"`

	// AWSCloudFormationDeletionPolicy represents a CloudFormation DeletionPolicy
	AWSCloudFormationDeletionPolicy policies.DeletionPolicy `json:"-"`

	// AWSCloudFormationUpdateReplacePolicy represents a CloudFormation UpdateReplacePolicy
	AWSCloudFormationUpdateReplacePolicy policies.UpdateReplacePolicy `json:"-"`

	// AWSCloudFormationDependsOn stores the logical ID of the resources to be created before this resource
	AWSCloudFormationDependsOn []string `json:"-"`

	// AWSCloudFormationMetadata stores structured data associated with this resource
	AWSCloudFormationMetadata map[string]interface{} `json:"-"`

	// AWSCloudFormationCondition stores the logical ID of the condition that must be satisfied for this resource to be created
	AWSCloudFormationCondition string `json:"-"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *BotAlias_TextLogDestination) AWSCloudFormationType() string {
	return "AWS::Lex::BotAlias.TextLogDestination"
}

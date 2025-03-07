package wafv2

import (
	"github.com/awslabs/goformation/v5/cloudformation/policies"
)

// WebACL_ImmunityTimeProperty AWS CloudFormation Resource (AWS::WAFv2::WebACL.ImmunityTimeProperty)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-webacl-immunitytimeproperty.html
type WebACL_ImmunityTimeProperty struct {

	// ImmunityTime AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-webacl-immunitytimeproperty.html#cfn-wafv2-webacl-immunitytimeproperty-immunitytime
	ImmunityTime int `json:"ImmunityTime"`

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
func (r *WebACL_ImmunityTimeProperty) AWSCloudFormationType() string {
	return "AWS::WAFv2::WebACL.ImmunityTimeProperty"
}

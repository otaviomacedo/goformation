package finspace

import (
	"github.com/awslabs/goformation/v5/cloudformation/policies"
)

// Environment_SuperuserParameters AWS CloudFormation Resource (AWS::FinSpace::Environment.SuperuserParameters)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-finspace-environment-superuserparameters.html
type Environment_SuperuserParameters struct {

	// EmailAddress AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-finspace-environment-superuserparameters.html#cfn-finspace-environment-superuserparameters-emailaddress
	EmailAddress string `json:"EmailAddress,omitempty"`

	// FirstName AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-finspace-environment-superuserparameters.html#cfn-finspace-environment-superuserparameters-firstname
	FirstName string `json:"FirstName,omitempty"`

	// LastName AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-finspace-environment-superuserparameters.html#cfn-finspace-environment-superuserparameters-lastname
	LastName string `json:"LastName,omitempty"`

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
func (r *Environment_SuperuserParameters) AWSCloudFormationType() string {
	return "AWS::FinSpace::Environment.SuperuserParameters"
}

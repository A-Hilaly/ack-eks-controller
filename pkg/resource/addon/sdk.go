// Copyright Amazon.com Inc. or its affiliates. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"). You may
// not use this file except in compliance with the License. A copy of the
// License is located at
//
//     http://aws.amazon.com/apache2.0/
//
// or in the "license" file accompanying this file. This file is distributed
// on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
// express or implied. See the License for the specific language governing
// permissions and limitations under the License.

// Code generated by ack-generate. DO NOT EDIT.

package addon

import (
	"context"
	"errors"
	"fmt"
	"reflect"
	"strings"

	ackv1alpha1 "github.com/aws-controllers-k8s/runtime/apis/core/v1alpha1"
	ackcompare "github.com/aws-controllers-k8s/runtime/pkg/compare"
	ackcondition "github.com/aws-controllers-k8s/runtime/pkg/condition"
	ackerr "github.com/aws-controllers-k8s/runtime/pkg/errors"
	ackrequeue "github.com/aws-controllers-k8s/runtime/pkg/requeue"
	ackrtlog "github.com/aws-controllers-k8s/runtime/pkg/runtime/log"
	"github.com/aws/aws-sdk-go/aws"
	svcsdk "github.com/aws/aws-sdk-go/service/eks"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	svcapitypes "github.com/aws-controllers-k8s/eks-controller/apis/v1alpha1"
)

// Hack to avoid import errors during build...
var (
	_ = &metav1.Time{}
	_ = strings.ToLower("")
	_ = &aws.JSONValue{}
	_ = &svcsdk.EKS{}
	_ = &svcapitypes.Addon{}
	_ = ackv1alpha1.AWSAccountID("")
	_ = &ackerr.NotFound
	_ = &ackcondition.NotManagedMessage
	_ = &reflect.Value{}
	_ = fmt.Sprintf("")
	_ = &ackrequeue.NoRequeue{}
)

// sdkFind returns SDK-specific information about a supplied resource
func (rm *resourceManager) sdkFind(
	ctx context.Context,
	r *resource,
) (latest *resource, err error) {
	rlog := ackrtlog.FromContext(ctx)
	exit := rlog.Trace("rm.sdkFind")
	defer func() {
		exit(err)
	}()
	// If any required fields in the input shape are missing, AWS resource is
	// not created yet. Return NotFound here to indicate to callers that the
	// resource isn't yet created.
	if rm.requiredFieldsMissingFromReadOneInput(r) {
		return nil, ackerr.NotFound
	}

	input, err := rm.newDescribeRequestPayload(r)
	if err != nil {
		return nil, err
	}

	var resp *svcsdk.DescribeAddonOutput
	resp, err = rm.sdkapi.DescribeAddonWithContext(ctx, input)
	rm.metrics.RecordAPICall("READ_ONE", "DescribeAddon", err)
	if err != nil {
		if reqErr, ok := ackerr.AWSRequestFailure(err); ok && reqErr.StatusCode() == 404 {
			return nil, ackerr.NotFound
		}
		if awsErr, ok := ackerr.AWSError(err); ok && awsErr.Code() == "ResourceNotFoundException" {
			return nil, ackerr.NotFound
		}
		return nil, err
	}

	// Merge in the information we read from the API call above to the copy of
	// the original Kubernetes object we passed to the function
	ko := r.ko.DeepCopy()

	if ko.Status.ACKResourceMetadata == nil {
		ko.Status.ACKResourceMetadata = &ackv1alpha1.ResourceMetadata{}
	}
	if resp.Addon.AddonArn != nil {
		arn := ackv1alpha1.AWSResourceName(*resp.Addon.AddonArn)
		ko.Status.ACKResourceMetadata.ARN = &arn
	}
	if resp.Addon.AddonName != nil {
		ko.Spec.Name = resp.Addon.AddonName
	} else {
		ko.Spec.Name = nil
	}
	if resp.Addon.AddonVersion != nil {
		ko.Spec.AddonVersion = resp.Addon.AddonVersion
	} else {
		ko.Spec.AddonVersion = nil
	}
	if resp.Addon.ClusterName != nil {
		ko.Spec.ClusterName = resp.Addon.ClusterName
	} else {
		ko.Spec.ClusterName = nil
	}
	if resp.Addon.ConfigurationValues != nil {
		ko.Spec.ConfigurationValues = resp.Addon.ConfigurationValues
	} else {
		ko.Spec.ConfigurationValues = nil
	}
	if resp.Addon.CreatedAt != nil {
		ko.Status.CreatedAt = &metav1.Time{*resp.Addon.CreatedAt}
	} else {
		ko.Status.CreatedAt = nil
	}
	if resp.Addon.Health != nil {
		f6 := &svcapitypes.AddonHealth{}
		if resp.Addon.Health.Issues != nil {
			f6f0 := []*svcapitypes.AddonIssue{}
			for _, f6f0iter := range resp.Addon.Health.Issues {
				f6f0elem := &svcapitypes.AddonIssue{}
				if f6f0iter.Code != nil {
					f6f0elem.Code = f6f0iter.Code
				}
				if f6f0iter.Message != nil {
					f6f0elem.Message = f6f0iter.Message
				}
				if f6f0iter.ResourceIds != nil {
					f6f0elemf2 := []*string{}
					for _, f6f0elemf2iter := range f6f0iter.ResourceIds {
						var f6f0elemf2elem string
						f6f0elemf2elem = *f6f0elemf2iter
						f6f0elemf2 = append(f6f0elemf2, &f6f0elemf2elem)
					}
					f6f0elem.ResourceIDs = f6f0elemf2
				}
				f6f0 = append(f6f0, f6f0elem)
			}
			f6.Issues = f6f0
		}
		ko.Status.Health = f6
	} else {
		ko.Status.Health = nil
	}
	if resp.Addon.MarketplaceInformation != nil {
		f7 := &svcapitypes.MarketplaceInformation{}
		if resp.Addon.MarketplaceInformation.ProductId != nil {
			f7.ProductID = resp.Addon.MarketplaceInformation.ProductId
		}
		if resp.Addon.MarketplaceInformation.ProductUrl != nil {
			f7.ProductURL = resp.Addon.MarketplaceInformation.ProductUrl
		}
		ko.Status.MarketplaceInformation = f7
	} else {
		ko.Status.MarketplaceInformation = nil
	}
	if resp.Addon.ModifiedAt != nil {
		ko.Status.ModifiedAt = &metav1.Time{*resp.Addon.ModifiedAt}
	} else {
		ko.Status.ModifiedAt = nil
	}
	if resp.Addon.Owner != nil {
		ko.Status.Owner = resp.Addon.Owner
	} else {
		ko.Status.Owner = nil
	}
	if resp.Addon.Publisher != nil {
		ko.Status.Publisher = resp.Addon.Publisher
	} else {
		ko.Status.Publisher = nil
	}
	if resp.Addon.ServiceAccountRoleArn != nil {
		ko.Spec.ServiceAccountRoleARN = resp.Addon.ServiceAccountRoleArn
	} else {
		ko.Spec.ServiceAccountRoleARN = nil
	}
	if resp.Addon.Status != nil {
		ko.Status.Status = resp.Addon.Status
	} else {
		ko.Status.Status = nil
	}
	if resp.Addon.Tags != nil {
		f14 := map[string]*string{}
		for f14key, f14valiter := range resp.Addon.Tags {
			var f14val string
			f14val = *f14valiter
			f14[f14key] = &f14val
		}
		ko.Spec.Tags = f14
	} else {
		ko.Spec.Tags = nil
	}

	rm.setStatusDefaults(ko)
	return &resource{ko}, nil
}

// requiredFieldsMissingFromReadOneInput returns true if there are any fields
// for the ReadOne Input shape that are required but not present in the
// resource's Spec or Status
func (rm *resourceManager) requiredFieldsMissingFromReadOneInput(
	r *resource,
) bool {
	return r.ko.Spec.ClusterName == nil || r.ko.Spec.Name == nil

}

// newDescribeRequestPayload returns SDK-specific struct for the HTTP request
// payload of the Describe API call for the resource
func (rm *resourceManager) newDescribeRequestPayload(
	r *resource,
) (*svcsdk.DescribeAddonInput, error) {
	res := &svcsdk.DescribeAddonInput{}

	if r.ko.Spec.Name != nil {
		res.SetAddonName(*r.ko.Spec.Name)
	}
	if r.ko.Spec.ClusterName != nil {
		res.SetClusterName(*r.ko.Spec.ClusterName)
	}

	return res, nil
}

// sdkCreate creates the supplied resource in the backend AWS service API and
// returns a copy of the resource with resource fields (in both Spec and
// Status) filled in with values from the CREATE API operation's Output shape.
func (rm *resourceManager) sdkCreate(
	ctx context.Context,
	desired *resource,
) (created *resource, err error) {
	rlog := ackrtlog.FromContext(ctx)
	exit := rlog.Trace("rm.sdkCreate")
	defer func() {
		exit(err)
	}()
	input, err := rm.newCreateRequestPayload(ctx, desired)
	if err != nil {
		return nil, err
	}

	var resp *svcsdk.CreateAddonOutput
	_ = resp
	resp, err = rm.sdkapi.CreateAddonWithContext(ctx, input)
	rm.metrics.RecordAPICall("CREATE", "CreateAddon", err)
	if err != nil {
		return nil, err
	}
	// Merge in the information we read from the API call above to the copy of
	// the original Kubernetes object we passed to the function
	ko := desired.ko.DeepCopy()

	if ko.Status.ACKResourceMetadata == nil {
		ko.Status.ACKResourceMetadata = &ackv1alpha1.ResourceMetadata{}
	}
	if resp.Addon.AddonArn != nil {
		arn := ackv1alpha1.AWSResourceName(*resp.Addon.AddonArn)
		ko.Status.ACKResourceMetadata.ARN = &arn
	}
	if resp.Addon.AddonName != nil {
		ko.Spec.Name = resp.Addon.AddonName
	} else {
		ko.Spec.Name = nil
	}
	if resp.Addon.AddonVersion != nil {
		ko.Spec.AddonVersion = resp.Addon.AddonVersion
	} else {
		ko.Spec.AddonVersion = nil
	}
	if resp.Addon.ClusterName != nil {
		ko.Spec.ClusterName = resp.Addon.ClusterName
	} else {
		ko.Spec.ClusterName = nil
	}
	if resp.Addon.ConfigurationValues != nil {
		ko.Spec.ConfigurationValues = resp.Addon.ConfigurationValues
	} else {
		ko.Spec.ConfigurationValues = nil
	}
	if resp.Addon.CreatedAt != nil {
		ko.Status.CreatedAt = &metav1.Time{*resp.Addon.CreatedAt}
	} else {
		ko.Status.CreatedAt = nil
	}
	if resp.Addon.Health != nil {
		f6 := &svcapitypes.AddonHealth{}
		if resp.Addon.Health.Issues != nil {
			f6f0 := []*svcapitypes.AddonIssue{}
			for _, f6f0iter := range resp.Addon.Health.Issues {
				f6f0elem := &svcapitypes.AddonIssue{}
				if f6f0iter.Code != nil {
					f6f0elem.Code = f6f0iter.Code
				}
				if f6f0iter.Message != nil {
					f6f0elem.Message = f6f0iter.Message
				}
				if f6f0iter.ResourceIds != nil {
					f6f0elemf2 := []*string{}
					for _, f6f0elemf2iter := range f6f0iter.ResourceIds {
						var f6f0elemf2elem string
						f6f0elemf2elem = *f6f0elemf2iter
						f6f0elemf2 = append(f6f0elemf2, &f6f0elemf2elem)
					}
					f6f0elem.ResourceIDs = f6f0elemf2
				}
				f6f0 = append(f6f0, f6f0elem)
			}
			f6.Issues = f6f0
		}
		ko.Status.Health = f6
	} else {
		ko.Status.Health = nil
	}
	if resp.Addon.MarketplaceInformation != nil {
		f7 := &svcapitypes.MarketplaceInformation{}
		if resp.Addon.MarketplaceInformation.ProductId != nil {
			f7.ProductID = resp.Addon.MarketplaceInformation.ProductId
		}
		if resp.Addon.MarketplaceInformation.ProductUrl != nil {
			f7.ProductURL = resp.Addon.MarketplaceInformation.ProductUrl
		}
		ko.Status.MarketplaceInformation = f7
	} else {
		ko.Status.MarketplaceInformation = nil
	}
	if resp.Addon.ModifiedAt != nil {
		ko.Status.ModifiedAt = &metav1.Time{*resp.Addon.ModifiedAt}
	} else {
		ko.Status.ModifiedAt = nil
	}
	if resp.Addon.Owner != nil {
		ko.Status.Owner = resp.Addon.Owner
	} else {
		ko.Status.Owner = nil
	}
	if resp.Addon.Publisher != nil {
		ko.Status.Publisher = resp.Addon.Publisher
	} else {
		ko.Status.Publisher = nil
	}
	if resp.Addon.ServiceAccountRoleArn != nil {
		ko.Spec.ServiceAccountRoleARN = resp.Addon.ServiceAccountRoleArn
	} else {
		ko.Spec.ServiceAccountRoleARN = nil
	}
	if resp.Addon.Status != nil {
		ko.Status.Status = resp.Addon.Status
	} else {
		ko.Status.Status = nil
	}
	if resp.Addon.Tags != nil {
		f14 := map[string]*string{}
		for f14key, f14valiter := range resp.Addon.Tags {
			var f14val string
			f14val = *f14valiter
			f14[f14key] = &f14val
		}
		ko.Spec.Tags = f14
	} else {
		ko.Spec.Tags = nil
	}

	rm.setStatusDefaults(ko)
	return &resource{ko}, nil
}

// newCreateRequestPayload returns an SDK-specific struct for the HTTP request
// payload of the Create API call for the resource
func (rm *resourceManager) newCreateRequestPayload(
	ctx context.Context,
	r *resource,
) (*svcsdk.CreateAddonInput, error) {
	res := &svcsdk.CreateAddonInput{}

	if r.ko.Spec.Name != nil {
		res.SetAddonName(*r.ko.Spec.Name)
	}
	if r.ko.Spec.AddonVersion != nil {
		res.SetAddonVersion(*r.ko.Spec.AddonVersion)
	}
	if r.ko.Spec.ClientRequestToken != nil {
		res.SetClientRequestToken(*r.ko.Spec.ClientRequestToken)
	}
	if r.ko.Spec.ClusterName != nil {
		res.SetClusterName(*r.ko.Spec.ClusterName)
	}
	if r.ko.Spec.ConfigurationValues != nil {
		res.SetConfigurationValues(*r.ko.Spec.ConfigurationValues)
	}
	if r.ko.Spec.PodIdentityAssociations != nil {
		f5 := []*svcsdk.AddonPodIdentityAssociations{}
		for _, f5iter := range r.ko.Spec.PodIdentityAssociations {
			f5elem := &svcsdk.AddonPodIdentityAssociations{}
			if f5iter.RoleARN != nil {
				f5elem.SetRoleArn(*f5iter.RoleARN)
			}
			if f5iter.ServiceAccount != nil {
				f5elem.SetServiceAccount(*f5iter.ServiceAccount)
			}
			f5 = append(f5, f5elem)
		}
		res.SetPodIdentityAssociations(f5)
	}
	if r.ko.Spec.ResolveConflicts != nil {
		res.SetResolveConflicts(*r.ko.Spec.ResolveConflicts)
	}
	if r.ko.Spec.ServiceAccountRoleARN != nil {
		res.SetServiceAccountRoleArn(*r.ko.Spec.ServiceAccountRoleARN)
	}
	if r.ko.Spec.Tags != nil {
		f8 := map[string]*string{}
		for f8key, f8valiter := range r.ko.Spec.Tags {
			var f8val string
			f8val = *f8valiter
			f8[f8key] = &f8val
		}
		res.SetTags(f8)
	}

	return res, nil
}

// sdkUpdate patches the supplied resource in the backend AWS service API and
// returns a new resource with updated fields.
func (rm *resourceManager) sdkUpdate(
	ctx context.Context,
	desired *resource,
	latest *resource,
	delta *ackcompare.Delta,
) (updated *resource, err error) {
	rlog := ackrtlog.FromContext(ctx)
	exit := rlog.Trace("rm.sdkUpdate")
	defer func() {
		exit(err)
	}()
	if delta.DifferentAt("Spec.Tags") {
		err := syncTags(
			ctx, rm.sdkapi, rm.metrics,
			string(*desired.ko.Status.ACKResourceMetadata.ARN),
			desired.ko.Spec.Tags, latest.ko.Spec.Tags,
		)
		if err != nil {
			return nil, err
		}
	}
	if !delta.DifferentExcept("Spec.Tags") {
		return desired, nil
	}
	input, err := rm.newUpdateRequestPayload(ctx, desired, delta)
	if err != nil {
		return nil, err
	}

	var resp *svcsdk.UpdateAddonOutput
	_ = resp
	resp, err = rm.sdkapi.UpdateAddonWithContext(ctx, input)
	rm.metrics.RecordAPICall("UPDATE", "UpdateAddon", err)
	if err != nil {
		return nil, err
	}
	// Merge in the information we read from the API call above to the copy of
	// the original Kubernetes object we passed to the function
	ko := desired.ko.DeepCopy()

	if resp.Update.CreatedAt != nil {
		ko.Status.CreatedAt = &metav1.Time{*resp.Update.CreatedAt}
	} else {
		ko.Status.CreatedAt = nil
	}
	if resp.Update.Status != nil {
		ko.Status.Status = resp.Update.Status
	} else {
		ko.Status.Status = nil
	}

	rm.setStatusDefaults(ko)
	return &resource{ko}, nil
}

// newUpdateRequestPayload returns an SDK-specific struct for the HTTP request
// payload of the Update API call for the resource
func (rm *resourceManager) newUpdateRequestPayload(
	ctx context.Context,
	r *resource,
	delta *ackcompare.Delta,
) (*svcsdk.UpdateAddonInput, error) {
	res := &svcsdk.UpdateAddonInput{}

	if r.ko.Spec.Name != nil {
		res.SetAddonName(*r.ko.Spec.Name)
	}
	if r.ko.Spec.AddonVersion != nil {
		res.SetAddonVersion(*r.ko.Spec.AddonVersion)
	}
	if r.ko.Spec.ClientRequestToken != nil {
		res.SetClientRequestToken(*r.ko.Spec.ClientRequestToken)
	}
	if r.ko.Spec.ClusterName != nil {
		res.SetClusterName(*r.ko.Spec.ClusterName)
	}
	if r.ko.Spec.ConfigurationValues != nil {
		res.SetConfigurationValues(*r.ko.Spec.ConfigurationValues)
	}
	if r.ko.Spec.PodIdentityAssociations != nil {
		f5 := []*svcsdk.AddonPodIdentityAssociations{}
		for _, f5iter := range r.ko.Spec.PodIdentityAssociations {
			f5elem := &svcsdk.AddonPodIdentityAssociations{}
			if f5iter.RoleARN != nil {
				f5elem.SetRoleArn(*f5iter.RoleARN)
			}
			if f5iter.ServiceAccount != nil {
				f5elem.SetServiceAccount(*f5iter.ServiceAccount)
			}
			f5 = append(f5, f5elem)
		}
		res.SetPodIdentityAssociations(f5)
	}
	if r.ko.Spec.ResolveConflicts != nil {
		res.SetResolveConflicts(*r.ko.Spec.ResolveConflicts)
	}
	if r.ko.Spec.ServiceAccountRoleARN != nil {
		res.SetServiceAccountRoleArn(*r.ko.Spec.ServiceAccountRoleARN)
	}

	return res, nil
}

// sdkDelete deletes the supplied resource in the backend AWS service API
func (rm *resourceManager) sdkDelete(
	ctx context.Context,
	r *resource,
) (latest *resource, err error) {
	rlog := ackrtlog.FromContext(ctx)
	exit := rlog.Trace("rm.sdkDelete")
	defer func() {
		exit(err)
	}()
	input, err := rm.newDeleteRequestPayload(r)
	if err != nil {
		return nil, err
	}
	var resp *svcsdk.DeleteAddonOutput
	_ = resp
	resp, err = rm.sdkapi.DeleteAddonWithContext(ctx, input)
	rm.metrics.RecordAPICall("DELETE", "DeleteAddon", err)
	return nil, err
}

// newDeleteRequestPayload returns an SDK-specific struct for the HTTP request
// payload of the Delete API call for the resource
func (rm *resourceManager) newDeleteRequestPayload(
	r *resource,
) (*svcsdk.DeleteAddonInput, error) {
	res := &svcsdk.DeleteAddonInput{}

	if r.ko.Spec.Name != nil {
		res.SetAddonName(*r.ko.Spec.Name)
	}
	if r.ko.Spec.ClusterName != nil {
		res.SetClusterName(*r.ko.Spec.ClusterName)
	}

	return res, nil
}

// setStatusDefaults sets default properties into supplied custom resource
func (rm *resourceManager) setStatusDefaults(
	ko *svcapitypes.Addon,
) {
	if ko.Status.ACKResourceMetadata == nil {
		ko.Status.ACKResourceMetadata = &ackv1alpha1.ResourceMetadata{}
	}
	if ko.Status.ACKResourceMetadata.Region == nil {
		ko.Status.ACKResourceMetadata.Region = &rm.awsRegion
	}
	if ko.Status.ACKResourceMetadata.OwnerAccountID == nil {
		ko.Status.ACKResourceMetadata.OwnerAccountID = &rm.awsAccountID
	}
	if ko.Status.Conditions == nil {
		ko.Status.Conditions = []*ackv1alpha1.Condition{}
	}
}

// updateConditions returns updated resource, true; if conditions were updated
// else it returns nil, false
func (rm *resourceManager) updateConditions(
	r *resource,
	onSuccess bool,
	err error,
) (*resource, bool) {
	ko := r.ko.DeepCopy()
	rm.setStatusDefaults(ko)

	// Terminal condition
	var terminalCondition *ackv1alpha1.Condition = nil
	var recoverableCondition *ackv1alpha1.Condition = nil
	var syncCondition *ackv1alpha1.Condition = nil
	for _, condition := range ko.Status.Conditions {
		if condition.Type == ackv1alpha1.ConditionTypeTerminal {
			terminalCondition = condition
		}
		if condition.Type == ackv1alpha1.ConditionTypeRecoverable {
			recoverableCondition = condition
		}
		if condition.Type == ackv1alpha1.ConditionTypeResourceSynced {
			syncCondition = condition
		}
	}
	var termError *ackerr.TerminalError
	if rm.terminalAWSError(err) || err == ackerr.SecretTypeNotSupported || err == ackerr.SecretNotFound || errors.As(err, &termError) {
		if terminalCondition == nil {
			terminalCondition = &ackv1alpha1.Condition{
				Type: ackv1alpha1.ConditionTypeTerminal,
			}
			ko.Status.Conditions = append(ko.Status.Conditions, terminalCondition)
		}
		var errorMessage = ""
		if err == ackerr.SecretTypeNotSupported || err == ackerr.SecretNotFound || errors.As(err, &termError) {
			errorMessage = err.Error()
		} else {
			awsErr, _ := ackerr.AWSError(err)
			errorMessage = awsErr.Error()
		}
		terminalCondition.Status = corev1.ConditionTrue
		terminalCondition.Message = &errorMessage
	} else {
		// Clear the terminal condition if no longer present
		if terminalCondition != nil {
			terminalCondition.Status = corev1.ConditionFalse
			terminalCondition.Message = nil
		}
		// Handling Recoverable Conditions
		if err != nil {
			if recoverableCondition == nil {
				// Add a new Condition containing a non-terminal error
				recoverableCondition = &ackv1alpha1.Condition{
					Type: ackv1alpha1.ConditionTypeRecoverable,
				}
				ko.Status.Conditions = append(ko.Status.Conditions, recoverableCondition)
			}
			recoverableCondition.Status = corev1.ConditionTrue
			awsErr, _ := ackerr.AWSError(err)
			errorMessage := err.Error()
			if awsErr != nil {
				errorMessage = awsErr.Error()
			}
			recoverableCondition.Message = &errorMessage
		} else if recoverableCondition != nil {
			recoverableCondition.Status = corev1.ConditionFalse
			recoverableCondition.Message = nil
		}
	}
	// Required to avoid the "declared but not used" error in the default case
	_ = syncCondition
	if terminalCondition != nil || recoverableCondition != nil || syncCondition != nil {
		return &resource{ko}, true // updated
	}
	return nil, false // not updated
}

// terminalAWSError returns awserr, true; if the supplied error is an aws Error type
// and if the exception indicates that it is a Terminal exception
// 'Terminal' exception are specified in generator configuration
func (rm *resourceManager) terminalAWSError(err error) bool {
	if err == nil {
		return false
	}
	awsErr, ok := ackerr.AWSError(err)
	if !ok {
		return false
	}
	switch awsErr.Code() {
	case "ResourceLimitExceeded",
		"ResourceNotFound",
		"ResourceInUse",
		"OptInRequired",
		"InvalidParameterCombination",
		"InvalidParameterValue",
		"InvalidParameterException",
		"InvalidQueryParameter",
		"MalformedQueryString",
		"MissingAction",
		"MissingParameter",
		"ValidationError":
		return true
	default:
		return false
	}
}

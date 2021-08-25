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

package scaling_policy

import (
	"context"
	"strings"

	ackv1alpha1 "github.com/aws-controllers-k8s/runtime/apis/core/v1alpha1"
	ackcompare "github.com/aws-controllers-k8s/runtime/pkg/compare"
	ackerr "github.com/aws-controllers-k8s/runtime/pkg/errors"
	ackrtlog "github.com/aws-controllers-k8s/runtime/pkg/runtime/log"
	"github.com/aws/aws-sdk-go/aws"
	svcsdk "github.com/aws/aws-sdk-go/service/applicationautoscaling"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	svcapitypes "github.com/aws-controllers-k8s/applicationautoscaling-controller/apis/v1alpha1"
)

// Hack to avoid import errors during build...
var (
	_ = &metav1.Time{}
	_ = strings.ToLower("")
	_ = &aws.JSONValue{}
	_ = &svcsdk.ApplicationAutoScaling{}
	_ = &svcapitypes.ScalingPolicy{}
	_ = ackv1alpha1.AWSAccountID("")
	_ = &ackerr.NotFound
)

// sdkFind returns SDK-specific information about a supplied resource
func (rm *resourceManager) sdkFind(
	ctx context.Context,
	r *resource,
) (latest *resource, err error) {
	rlog := ackrtlog.FromContext(ctx)
	exit := rlog.Trace("rm.sdkFind")
	defer exit(err)
	// If any required fields in the input shape are missing, AWS resource is
	// not created yet. Return NotFound here to indicate to callers that the
	// resource isn't yet created.
	if rm.requiredFieldsMissingFromReadManyInput(r) {
		return nil, ackerr.NotFound
	}

	input, err := rm.newListRequestPayload(r)
	if err != nil {
		return nil, err
	}
	var resp *svcsdk.DescribeScalingPoliciesOutput
	resp, err = rm.sdkapi.DescribeScalingPoliciesWithContext(ctx, input)
	rm.metrics.RecordAPICall("READ_MANY", "DescribeScalingPolicies", err)
	if err != nil {
		if awsErr, ok := ackerr.AWSError(err); ok && awsErr.Code() == "UNKNOWN" {
			return nil, ackerr.NotFound
		}
		return nil, err
	}

	// Merge in the information we read from the API call above to the copy of
	// the original Kubernetes object we passed to the function
	ko := r.ko.DeepCopy()

	found := false
	for _, elem := range resp.ScalingPolicies {
		if elem.Alarms != nil {
			f0 := []*svcapitypes.Alarm{}
			for _, f0iter := range elem.Alarms {
				f0elem := &svcapitypes.Alarm{}
				if f0iter.AlarmARN != nil {
					f0elem.AlarmARN = f0iter.AlarmARN
				}
				if f0iter.AlarmName != nil {
					f0elem.AlarmName = f0iter.AlarmName
				}
				f0 = append(f0, f0elem)
			}
			ko.Status.Alarms = f0
		} else {
			ko.Status.Alarms = nil
		}
		if elem.PolicyARN != nil {
			if ko.Status.ACKResourceMetadata == nil {
				ko.Status.ACKResourceMetadata = &ackv1alpha1.ResourceMetadata{}
			}
			tmpARN := ackv1alpha1.AWSResourceName(*elem.PolicyARN)
			ko.Status.ACKResourceMetadata.ARN = &tmpARN
		}
		if elem.PolicyName != nil {
			ko.Spec.PolicyName = elem.PolicyName
		} else {
			ko.Spec.PolicyName = nil
		}
		if elem.PolicyType != nil {
			ko.Spec.PolicyType = elem.PolicyType
		} else {
			ko.Spec.PolicyType = nil
		}
		if elem.ResourceId != nil {
			ko.Spec.ResourceID = elem.ResourceId
		} else {
			ko.Spec.ResourceID = nil
		}
		if elem.ScalableDimension != nil {
			ko.Spec.ScalableDimension = elem.ScalableDimension
		} else {
			ko.Spec.ScalableDimension = nil
		}
		if elem.ServiceNamespace != nil {
			ko.Spec.ServiceNamespace = elem.ServiceNamespace
		} else {
			ko.Spec.ServiceNamespace = nil
		}
		if elem.StepScalingPolicyConfiguration != nil {
			f8 := &svcapitypes.StepScalingPolicyConfiguration{}
			if elem.StepScalingPolicyConfiguration.AdjustmentType != nil {
				f8.AdjustmentType = elem.StepScalingPolicyConfiguration.AdjustmentType
			}
			if elem.StepScalingPolicyConfiguration.Cooldown != nil {
				f8.Cooldown = elem.StepScalingPolicyConfiguration.Cooldown
			}
			if elem.StepScalingPolicyConfiguration.MetricAggregationType != nil {
				f8.MetricAggregationType = elem.StepScalingPolicyConfiguration.MetricAggregationType
			}
			if elem.StepScalingPolicyConfiguration.MinAdjustmentMagnitude != nil {
				f8.MinAdjustmentMagnitude = elem.StepScalingPolicyConfiguration.MinAdjustmentMagnitude
			}
			if elem.StepScalingPolicyConfiguration.StepAdjustments != nil {
				f8f4 := []*svcapitypes.StepAdjustment{}
				for _, f8f4iter := range elem.StepScalingPolicyConfiguration.StepAdjustments {
					f8f4elem := &svcapitypes.StepAdjustment{}
					if f8f4iter.MetricIntervalLowerBound != nil {
						f8f4elem.MetricIntervalLowerBound = f8f4iter.MetricIntervalLowerBound
					}
					if f8f4iter.MetricIntervalUpperBound != nil {
						f8f4elem.MetricIntervalUpperBound = f8f4iter.MetricIntervalUpperBound
					}
					if f8f4iter.ScalingAdjustment != nil {
						f8f4elem.ScalingAdjustment = f8f4iter.ScalingAdjustment
					}
					f8f4 = append(f8f4, f8f4elem)
				}
				f8.StepAdjustments = f8f4
			}
			ko.Spec.StepScalingPolicyConfiguration = f8
		} else {
			ko.Spec.StepScalingPolicyConfiguration = nil
		}
		if elem.TargetTrackingScalingPolicyConfiguration != nil {
			f9 := &svcapitypes.TargetTrackingScalingPolicyConfiguration{}
			if elem.TargetTrackingScalingPolicyConfiguration.CustomizedMetricSpecification != nil {
				f9f0 := &svcapitypes.CustomizedMetricSpecification{}
				if elem.TargetTrackingScalingPolicyConfiguration.CustomizedMetricSpecification.Dimensions != nil {
					f9f0f0 := []*svcapitypes.MetricDimension{}
					for _, f9f0f0iter := range elem.TargetTrackingScalingPolicyConfiguration.CustomizedMetricSpecification.Dimensions {
						f9f0f0elem := &svcapitypes.MetricDimension{}
						if f9f0f0iter.Name != nil {
							f9f0f0elem.Name = f9f0f0iter.Name
						}
						if f9f0f0iter.Value != nil {
							f9f0f0elem.Value = f9f0f0iter.Value
						}
						f9f0f0 = append(f9f0f0, f9f0f0elem)
					}
					f9f0.Dimensions = f9f0f0
				}
				if elem.TargetTrackingScalingPolicyConfiguration.CustomizedMetricSpecification.MetricName != nil {
					f9f0.MetricName = elem.TargetTrackingScalingPolicyConfiguration.CustomizedMetricSpecification.MetricName
				}
				if elem.TargetTrackingScalingPolicyConfiguration.CustomizedMetricSpecification.Namespace != nil {
					f9f0.Namespace = elem.TargetTrackingScalingPolicyConfiguration.CustomizedMetricSpecification.Namespace
				}
				if elem.TargetTrackingScalingPolicyConfiguration.CustomizedMetricSpecification.Statistic != nil {
					f9f0.Statistic = elem.TargetTrackingScalingPolicyConfiguration.CustomizedMetricSpecification.Statistic
				}
				if elem.TargetTrackingScalingPolicyConfiguration.CustomizedMetricSpecification.Unit != nil {
					f9f0.Unit = elem.TargetTrackingScalingPolicyConfiguration.CustomizedMetricSpecification.Unit
				}
				f9.CustomizedMetricSpecification = f9f0
			}
			if elem.TargetTrackingScalingPolicyConfiguration.DisableScaleIn != nil {
				f9.DisableScaleIn = elem.TargetTrackingScalingPolicyConfiguration.DisableScaleIn
			}
			if elem.TargetTrackingScalingPolicyConfiguration.PredefinedMetricSpecification != nil {
				f9f2 := &svcapitypes.PredefinedMetricSpecification{}
				if elem.TargetTrackingScalingPolicyConfiguration.PredefinedMetricSpecification.PredefinedMetricType != nil {
					f9f2.PredefinedMetricType = elem.TargetTrackingScalingPolicyConfiguration.PredefinedMetricSpecification.PredefinedMetricType
				}
				if elem.TargetTrackingScalingPolicyConfiguration.PredefinedMetricSpecification.ResourceLabel != nil {
					f9f2.ResourceLabel = elem.TargetTrackingScalingPolicyConfiguration.PredefinedMetricSpecification.ResourceLabel
				}
				f9.PredefinedMetricSpecification = f9f2
			}
			if elem.TargetTrackingScalingPolicyConfiguration.ScaleInCooldown != nil {
				f9.ScaleInCooldown = elem.TargetTrackingScalingPolicyConfiguration.ScaleInCooldown
			}
			if elem.TargetTrackingScalingPolicyConfiguration.ScaleOutCooldown != nil {
				f9.ScaleOutCooldown = elem.TargetTrackingScalingPolicyConfiguration.ScaleOutCooldown
			}
			if elem.TargetTrackingScalingPolicyConfiguration.TargetValue != nil {
				f9.TargetValue = elem.TargetTrackingScalingPolicyConfiguration.TargetValue
			}
			ko.Spec.TargetTrackingScalingPolicyConfiguration = f9
		} else {
			ko.Spec.TargetTrackingScalingPolicyConfiguration = nil
		}
		found = true
		break
	}
	if !found {
		return nil, ackerr.NotFound
	}

	rm.setStatusDefaults(ko)
	return &resource{ko}, nil
}

// requiredFieldsMissingFromReadManyInput returns true if there are any fields
// for the ReadMany Input shape that are required but not present in the
// resource's Spec or Status
func (rm *resourceManager) requiredFieldsMissingFromReadManyInput(
	r *resource,
) bool {
	return false
}

// newListRequestPayload returns SDK-specific struct for the HTTP request
// payload of the List API call for the resource
func (rm *resourceManager) newListRequestPayload(
	r *resource,
) (*svcsdk.DescribeScalingPoliciesInput, error) {
	res := &svcsdk.DescribeScalingPoliciesInput{}

	if r.ko.Spec.ResourceID != nil {
		res.SetResourceId(*r.ko.Spec.ResourceID)
	}
	if r.ko.Spec.ScalableDimension != nil {
		res.SetScalableDimension(*r.ko.Spec.ScalableDimension)
	}
	if r.ko.Spec.ServiceNamespace != nil {
		res.SetServiceNamespace(*r.ko.Spec.ServiceNamespace)
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
	defer exit(err)
	input, err := rm.newCreateRequestPayload(ctx, desired)
	if err != nil {
		return nil, err
	}

	var resp *svcsdk.PutScalingPolicyOutput
	_ = resp
	resp, err = rm.sdkapi.PutScalingPolicyWithContext(ctx, input)
	rm.metrics.RecordAPICall("CREATE", "PutScalingPolicy", err)
	if err != nil {
		return nil, err
	}
	// Merge in the information we read from the API call above to the copy of
	// the original Kubernetes object we passed to the function
	ko := desired.ko.DeepCopy()

	if resp.Alarms != nil {
		f0 := []*svcapitypes.Alarm{}
		for _, f0iter := range resp.Alarms {
			f0elem := &svcapitypes.Alarm{}
			if f0iter.AlarmARN != nil {
				f0elem.AlarmARN = f0iter.AlarmARN
			}
			if f0iter.AlarmName != nil {
				f0elem.AlarmName = f0iter.AlarmName
			}
			f0 = append(f0, f0elem)
		}
		ko.Status.Alarms = f0
	} else {
		ko.Status.Alarms = nil
	}
	if ko.Status.ACKResourceMetadata == nil {
		ko.Status.ACKResourceMetadata = &ackv1alpha1.ResourceMetadata{}
	}
	if resp.PolicyARN != nil {
		arn := ackv1alpha1.AWSResourceName(*resp.PolicyARN)
		ko.Status.ACKResourceMetadata.ARN = &arn
	}

	rm.setStatusDefaults(ko)
	return &resource{ko}, nil
}

// newCreateRequestPayload returns an SDK-specific struct for the HTTP request
// payload of the Create API call for the resource
func (rm *resourceManager) newCreateRequestPayload(
	ctx context.Context,
	r *resource,
) (*svcsdk.PutScalingPolicyInput, error) {
	res := &svcsdk.PutScalingPolicyInput{}

	if r.ko.Spec.PolicyName != nil {
		res.SetPolicyName(*r.ko.Spec.PolicyName)
	}
	if r.ko.Spec.PolicyType != nil {
		res.SetPolicyType(*r.ko.Spec.PolicyType)
	}
	if r.ko.Spec.ResourceID != nil {
		res.SetResourceId(*r.ko.Spec.ResourceID)
	}
	if r.ko.Spec.ScalableDimension != nil {
		res.SetScalableDimension(*r.ko.Spec.ScalableDimension)
	}
	if r.ko.Spec.ServiceNamespace != nil {
		res.SetServiceNamespace(*r.ko.Spec.ServiceNamespace)
	}
	if r.ko.Spec.StepScalingPolicyConfiguration != nil {
		f5 := &svcsdk.StepScalingPolicyConfiguration{}
		if r.ko.Spec.StepScalingPolicyConfiguration.AdjustmentType != nil {
			f5.SetAdjustmentType(*r.ko.Spec.StepScalingPolicyConfiguration.AdjustmentType)
		}
		if r.ko.Spec.StepScalingPolicyConfiguration.Cooldown != nil {
			f5.SetCooldown(*r.ko.Spec.StepScalingPolicyConfiguration.Cooldown)
		}
		if r.ko.Spec.StepScalingPolicyConfiguration.MetricAggregationType != nil {
			f5.SetMetricAggregationType(*r.ko.Spec.StepScalingPolicyConfiguration.MetricAggregationType)
		}
		if r.ko.Spec.StepScalingPolicyConfiguration.MinAdjustmentMagnitude != nil {
			f5.SetMinAdjustmentMagnitude(*r.ko.Spec.StepScalingPolicyConfiguration.MinAdjustmentMagnitude)
		}
		if r.ko.Spec.StepScalingPolicyConfiguration.StepAdjustments != nil {
			f5f4 := []*svcsdk.StepAdjustment{}
			for _, f5f4iter := range r.ko.Spec.StepScalingPolicyConfiguration.StepAdjustments {
				f5f4elem := &svcsdk.StepAdjustment{}
				if f5f4iter.MetricIntervalLowerBound != nil {
					f5f4elem.SetMetricIntervalLowerBound(*f5f4iter.MetricIntervalLowerBound)
				}
				if f5f4iter.MetricIntervalUpperBound != nil {
					f5f4elem.SetMetricIntervalUpperBound(*f5f4iter.MetricIntervalUpperBound)
				}
				if f5f4iter.ScalingAdjustment != nil {
					f5f4elem.SetScalingAdjustment(*f5f4iter.ScalingAdjustment)
				}
				f5f4 = append(f5f4, f5f4elem)
			}
			f5.SetStepAdjustments(f5f4)
		}
		res.SetStepScalingPolicyConfiguration(f5)
	}
	if r.ko.Spec.TargetTrackingScalingPolicyConfiguration != nil {
		f6 := &svcsdk.TargetTrackingScalingPolicyConfiguration{}
		if r.ko.Spec.TargetTrackingScalingPolicyConfiguration.CustomizedMetricSpecification != nil {
			f6f0 := &svcsdk.CustomizedMetricSpecification{}
			if r.ko.Spec.TargetTrackingScalingPolicyConfiguration.CustomizedMetricSpecification.Dimensions != nil {
				f6f0f0 := []*svcsdk.MetricDimension{}
				for _, f6f0f0iter := range r.ko.Spec.TargetTrackingScalingPolicyConfiguration.CustomizedMetricSpecification.Dimensions {
					f6f0f0elem := &svcsdk.MetricDimension{}
					if f6f0f0iter.Name != nil {
						f6f0f0elem.SetName(*f6f0f0iter.Name)
					}
					if f6f0f0iter.Value != nil {
						f6f0f0elem.SetValue(*f6f0f0iter.Value)
					}
					f6f0f0 = append(f6f0f0, f6f0f0elem)
				}
				f6f0.SetDimensions(f6f0f0)
			}
			if r.ko.Spec.TargetTrackingScalingPolicyConfiguration.CustomizedMetricSpecification.MetricName != nil {
				f6f0.SetMetricName(*r.ko.Spec.TargetTrackingScalingPolicyConfiguration.CustomizedMetricSpecification.MetricName)
			}
			if r.ko.Spec.TargetTrackingScalingPolicyConfiguration.CustomizedMetricSpecification.Namespace != nil {
				f6f0.SetNamespace(*r.ko.Spec.TargetTrackingScalingPolicyConfiguration.CustomizedMetricSpecification.Namespace)
			}
			if r.ko.Spec.TargetTrackingScalingPolicyConfiguration.CustomizedMetricSpecification.Statistic != nil {
				f6f0.SetStatistic(*r.ko.Spec.TargetTrackingScalingPolicyConfiguration.CustomizedMetricSpecification.Statistic)
			}
			if r.ko.Spec.TargetTrackingScalingPolicyConfiguration.CustomizedMetricSpecification.Unit != nil {
				f6f0.SetUnit(*r.ko.Spec.TargetTrackingScalingPolicyConfiguration.CustomizedMetricSpecification.Unit)
			}
			f6.SetCustomizedMetricSpecification(f6f0)
		}
		if r.ko.Spec.TargetTrackingScalingPolicyConfiguration.DisableScaleIn != nil {
			f6.SetDisableScaleIn(*r.ko.Spec.TargetTrackingScalingPolicyConfiguration.DisableScaleIn)
		}
		if r.ko.Spec.TargetTrackingScalingPolicyConfiguration.PredefinedMetricSpecification != nil {
			f6f2 := &svcsdk.PredefinedMetricSpecification{}
			if r.ko.Spec.TargetTrackingScalingPolicyConfiguration.PredefinedMetricSpecification.PredefinedMetricType != nil {
				f6f2.SetPredefinedMetricType(*r.ko.Spec.TargetTrackingScalingPolicyConfiguration.PredefinedMetricSpecification.PredefinedMetricType)
			}
			if r.ko.Spec.TargetTrackingScalingPolicyConfiguration.PredefinedMetricSpecification.ResourceLabel != nil {
				f6f2.SetResourceLabel(*r.ko.Spec.TargetTrackingScalingPolicyConfiguration.PredefinedMetricSpecification.ResourceLabel)
			}
			f6.SetPredefinedMetricSpecification(f6f2)
		}
		if r.ko.Spec.TargetTrackingScalingPolicyConfiguration.ScaleInCooldown != nil {
			f6.SetScaleInCooldown(*r.ko.Spec.TargetTrackingScalingPolicyConfiguration.ScaleInCooldown)
		}
		if r.ko.Spec.TargetTrackingScalingPolicyConfiguration.ScaleOutCooldown != nil {
			f6.SetScaleOutCooldown(*r.ko.Spec.TargetTrackingScalingPolicyConfiguration.ScaleOutCooldown)
		}
		if r.ko.Spec.TargetTrackingScalingPolicyConfiguration.TargetValue != nil {
			f6.SetTargetValue(*r.ko.Spec.TargetTrackingScalingPolicyConfiguration.TargetValue)
		}
		res.SetTargetTrackingScalingPolicyConfiguration(f6)
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
) (*resource, error) {
	// TODO(jaypipes): Figure this out...
	return nil, ackerr.NotImplemented
}

// sdkDelete deletes the supplied resource in the backend AWS service API
func (rm *resourceManager) sdkDelete(
	ctx context.Context,
	r *resource,
) (latest *resource, err error) {
	rlog := ackrtlog.FromContext(ctx)
	exit := rlog.Trace("rm.sdkDelete")
	defer exit(err)
	input, err := rm.newDeleteRequestPayload(r)
	if err != nil {
		return nil, err
	}
	var resp *svcsdk.DeleteScalingPolicyOutput
	_ = resp
	resp, err = rm.sdkapi.DeleteScalingPolicyWithContext(ctx, input)
	rm.metrics.RecordAPICall("DELETE", "DeleteScalingPolicy", err)
	return nil, err
}

// newDeleteRequestPayload returns an SDK-specific struct for the HTTP request
// payload of the Delete API call for the resource
func (rm *resourceManager) newDeleteRequestPayload(
	r *resource,
) (*svcsdk.DeleteScalingPolicyInput, error) {
	res := &svcsdk.DeleteScalingPolicyInput{}

	if r.ko.Spec.PolicyName != nil {
		res.SetPolicyName(*r.ko.Spec.PolicyName)
	}
	if r.ko.Spec.ResourceID != nil {
		res.SetResourceId(*r.ko.Spec.ResourceID)
	}
	if r.ko.Spec.ScalableDimension != nil {
		res.SetScalableDimension(*r.ko.Spec.ScalableDimension)
	}
	if r.ko.Spec.ServiceNamespace != nil {
		res.SetServiceNamespace(*r.ko.Spec.ServiceNamespace)
	}

	return res, nil
}

// setStatusDefaults sets default properties into supplied custom resource
func (rm *resourceManager) setStatusDefaults(
	ko *svcapitypes.ScalingPolicy,
) {
	if ko.Status.ACKResourceMetadata == nil {
		ko.Status.ACKResourceMetadata = &ackv1alpha1.ResourceMetadata{}
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

	if rm.terminalAWSError(err) || err == ackerr.SecretTypeNotSupported || err == ackerr.SecretNotFound {
		if terminalCondition == nil {
			terminalCondition = &ackv1alpha1.Condition{
				Type: ackv1alpha1.ConditionTypeTerminal,
			}
			ko.Status.Conditions = append(ko.Status.Conditions, terminalCondition)
		}
		var errorMessage = ""
		if err == ackerr.SecretTypeNotSupported || err == ackerr.SecretNotFound {
			errorMessage = err.Error()
		} else {
			awsErr, _ := ackerr.AWSError(err)
			errorMessage = awsErr.Message()
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
				errorMessage = awsErr.Message()
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
	// No terminal_errors specified for this resource in generator config
	return false
}

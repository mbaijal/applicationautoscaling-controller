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

package scalable_target

import (
	"errors"

	svcsdk "github.com/aws/aws-sdk-go/service/applicationautoscaling"
)

// addIDToListRequest adds requested-resource VpcId to ListRequest.
// Return error to indicate to callers that the resource is not yet created.
func addIDToListRequest(r *resource, input *svcsdk.DescribeScalableTargetsInput) error {
	if r.ko.Spec.ResourceID == nil || *r.ko.Spec.ResourceID == "" {
		return errors.New("unable to extract resourceID from Kubernetes resource")
	}
	input.SetResourceIds([]*string{r.ko.Spec.ResourceID})
	return nil
}

// customCheckRequiredFieldsMissingMethod returns true if there are any fields
// for the ReadOne Input shape that are required but not present in the
// resource's Spec or Status
// ModelPackage can be Unversioned or Versioned both cannot be nil since ModelPackageName
// is required for Unversioned and  ModelPackageGroupName is required for Versioned
func (rm *resourceManager) customCheckRequiredFieldsMissingMethod(
	r *resource,
) bool {
	return r.ko.Spec.ResourceID == nil

}

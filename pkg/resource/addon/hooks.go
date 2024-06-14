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

package addon

import (
	"context"
	"strings"

	ackrtlog "github.com/aws-controllers-k8s/runtime/pkg/runtime/log"
	svcsdk "github.com/aws/aws-sdk-go/service/eks"

	"github.com/aws-controllers-k8s/eks-controller/apis/v1alpha1"
	"github.com/aws-controllers-k8s/eks-controller/pkg/tags"
)

var syncTags = tags.SyncTags

// setResourceDefaults queries the EKS API for the current state of the
// fields that are not returned by the ReadOne or List APIs.
func (rm *resourceManager) setResourceAdditionalFields(ctx context.Context, r *v1alpha1.Addon, associationARNs []*string) (err error) {
	rlog := ackrtlog.FromContext(ctx)
	exit := rlog.Trace("rm.setResourceAdditionalFields")
	defer func() { exit(err) }()

	pias, err := rm.describeAddonPodIdentityAssociations(ctx, r.Spec.ClusterName, associationARNs)
	if err != nil {
		return err
	}
	r.Spec.PodIdentityAssociations = pias
	return nil
}

func (rm *resourceManager) describeAddonPodIdentityAssociations(
	ctx context.Context,
	clusterName *string,
	associationARNs []*string,
) (pias []*v1alpha1.AddonPodIdentityAssociations, err error) {
	rlog := ackrtlog.FromContext(ctx)
	exit := rlog.Trace("rm.getPodIdentityAssociations")
	defer func() { exit(err) }()

	for _, associationARN := range associationARNs {
		associationID := getAssociationID(*associationARN)
		resp, err := rm.sdkapi.DescribePodIdentityAssociationWithContext(
			ctx,
			&svcsdk.DescribePodIdentityAssociationInput{
				ClusterName:   clusterName,
				AssociationId: &associationID,
			},
		)
		if err != nil {
			return nil, err
		}
		pias = append(pias, &v1alpha1.AddonPodIdentityAssociations{
			RoleARN:        resp.Association.RoleArn,
			ServiceAccount: resp.Association.ServiceAccount,
		})
	}

	return pias, nil
}

func getAssociationID(associationARN string) string {
	parts := strings.Split(associationARN, "/")
	return parts[len(parts)-1]
}

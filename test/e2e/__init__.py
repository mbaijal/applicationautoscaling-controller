# Copyright Amazon.com Inc. or its affiliates. All Rights Reserved.
#
# Licensed under the Apache License, Version 2.0 (the "License"). You may
# not use this file except in compliance with the License. A copy of the
# License is located at
#
# 	 http://aws.amazon.com/apache2.0/
#
# or in the "license" file accompanying this file. This file is distributed
# on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
# express or implied. See the License for the specific language governing
# permissions and limitations under the License.

import pytest
from typing import Dict, Any
from pathlib import Path

from acktest.k8s import resource as k8s
from acktest.resources import load_resource_file

SERVICE_NAME = "applicationautoscaling"
CRD_GROUP = "applicationautoscaling.services.k8s.aws"
ADOPTED_RESOURCE_CRD_GROUP = "services.k8s.aws"
CRD_VERSION = "v1alpha1"

# PyTest marker for the current service
service_marker = pytest.mark.service(arg=SERVICE_NAME)

bootstrap_directory = Path(__file__).parent
resource_directory = Path(__file__).parent / "resources"


def load_autoscaling_resource(
    resource_name: str, additional_replacements: Dict[str, Any] = {}
):
    """Overrides the default `load_resource_file` to access the specific resources
    directory for the current service.
    """
    return load_resource_file(
        resource_directory,
        resource_name,
        additional_replacements=additional_replacements,
    )


def create_applicationautoscaling_resource(
    resource_plural, resource_name, spec_file, replacements, namespace="default"
):
    """
    Wrapper around k8s.load_and_create_resource to create a ApplicationAutoscaling resource
    """

    reference, spec, resource = k8s.load_and_create_resource(
        resource_directory,
        CRD_GROUP,
        CRD_VERSION,
        resource_plural,
        resource_name,
        spec_file,
        replacements,
        namespace,
    )

    return reference, spec, resource

def create_adopted_resource(resource_name, spec_file, replacements, namespace="default"):
    """
    Wrapper around k8s.load_and_create_resource to create a Adopoted resource
    """

    reference, spec, resource = k8s.load_and_create_resource(
        resource_directory,
        ADOPTED_RESOURCE_CRD_GROUP,
        CRD_VERSION,
        "adoptedresources",
        resource_name,
        spec_file,
        replacements,
        namespace,
    )

    return reference, spec, resource
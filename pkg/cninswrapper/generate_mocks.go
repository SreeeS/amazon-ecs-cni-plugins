// Copyright 2017 Amazon.com, Inc. or its affiliates. All Rights Reserved.
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

package cninswrapper

//go:generate go run ../../scripts/mockgen.go github.com/aws/amazon-ecs-cni-plugins/pkg/cninswrapper NS mocks/cninswrapper_mocks.go
//go:generate go run ../../scripts/mockgen.go github.com/containernetworking/plugins/pkg/ns NetNS mocks_netns/netns_mocks.go

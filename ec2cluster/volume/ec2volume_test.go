// Copyright 2019 GRAIL, Inc. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.

package volume

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/aws/aws-sdk-go/service/ec2/ec2iface"
)

type mockEC2Client struct {
	ec2iface.EC2API
	descVolsFn     func() (*ec2.DescribeVolumesOutput, error)
	descVolsModsFn func() (*ec2.DescribeVolumesModificationsOutput, error)
	modVolsFn      func(volId string) (*ec2.ModifyVolumeOutput, error)
}

func (e *mockEC2Client) DescribeVolumesWithContext(ctx aws.Context, input *ec2.DescribeVolumesInput, _ ...request.Option) (*ec2.DescribeVolumesOutput, error) {
	return e.descVolsFn()
}

func (e *mockEC2Client) DescribeVolumesModificationsWithContext(ctx aws.Context, input *ec2.DescribeVolumesModificationsInput, _ ...request.Option) (*ec2.DescribeVolumesModificationsOutput, error) {
	return e.descVolsModsFn()
}

func (e *mockEC2Client) ModifyVolumeWithContext(ctx aws.Context, input *ec2.ModifyVolumeInput, _ ...request.Option) (*ec2.ModifyVolumeOutput, error) {
	return e.modVolsFn(*input.VolumeId)
}

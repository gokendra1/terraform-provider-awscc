// Code generated by generators/resource/main.go; DO NOT EDIT.

package memorydb_test

import (
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-provider-awscc/internal/acctest"
)

func TestAccAWSMemoryDBSubnetGroup_basic(t *testing.T) {
	td := acctest.NewTestData(t, "AWS::MemoryDB::SubnetGroup", "awscc_memorydb_subnet_group", "test")

	td.ResourceTest(t, []resource.TestStep{
		{
			Config:      td.EmptyConfig(),
			ExpectError: regexp.MustCompile("Missing required argument"),
		},
	})
}
// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package customerprofiles_test

import (
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-provider-awscc/internal/acctest"
)

func TestAccAWSCustomerProfilesCalculatedAttributeDefinitionDataSource_basic(t *testing.T) {
	td := acctest.NewTestData(t, "AWS::CustomerProfiles::CalculatedAttributeDefinition", "awscc_customerprofiles_calculated_attribute_definition", "test")

	td.DataSourceTest(t, []resource.TestStep{
		{
			Config:      td.EmptyDataSourceConfig(),
			ExpectError: regexp.MustCompile("Missing required argument"),
		},
	})
}

func TestAccAWSCustomerProfilesCalculatedAttributeDefinitionDataSource_NonExistent(t *testing.T) {
	td := acctest.NewTestData(t, "AWS::CustomerProfiles::CalculatedAttributeDefinition", "awscc_customerprofiles_calculated_attribute_definition", "test")

	td.DataSourceTest(t, []resource.TestStep{
		{
			Config:      td.DataSourceWithNonExistentIDConfig(),
			ExpectError: regexp.MustCompile("Not Found"),
		},
	})
}
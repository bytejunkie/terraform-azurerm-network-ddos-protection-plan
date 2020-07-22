package test

import (
	"math/rand"
	"strings"
	"testing"

	"github.com/gruntwork-io/terratest/modules/terraform"
	test_structure "github.com/gruntwork-io/terratest/modules/test-structure"
	"github.com/stretchr/testify/assert"
)

// Test01 commented out as can only have one per subscription per region.

// func TestTerraformVirtualNetwork(t *testing.T) {
// 	t.Parallel()

// 	// randomString := string(RandomString(4))
// 	resourceGroupNameStrings := []string{"mns", "rg", "01"}

// 	DDOSPlanNameStrings := []string{"mns", "ddos", "01"}
// 	DDOSPlanLocation := "westeurope"

// 	// Root folder where terraform files should be (relative to the test folder)
// 	rootFolder := ".."
// 	// Relative path to terraform module being tested from the root folder
// 	terraformFolderRelativeToRoot := "calling-code"
// 	// Copy the terraform folder to a temp folder
// 	tempTestFolder := test_structure.CopyTerraformFolderToTemp(t, rootFolder, terraformFolderRelativeToRoot)

// 	terraformOptions := &terraform.Options{
// 		// The path to where our Terraform code is located
// 		TerraformDir: tempTestFolder,

// 		// Variables to pass to our Terraform code using -var options
// 		Vars: map[string]interface{}{
// 			"resource_group_name_strings":   resourceGroupNameStrings,
// 			"resource_group_name_separator": "-",

// 			"ddos_plan_name_strings":   DDOSPlanNameStrings,
// 			"ddos_plan_name_separator": "-",
// 			"ddos_plan_location":       DDOSPlanLocation,
// 		},

// 		BackendConfig: map[string]interface{}{
// 			"bucket": "byt-terraform-state-bucket",
// 			"key":    "testing/test01.tfstate",
// 			"region": "eu-west-2",
// 		},
// 		// Variables to pass to our Terraform code using -var-file options
// 		// VarFiles: []string{"varfile.tfvars"},

// 		// Disable colors in Terraform commands so its easier to parse stdout/stderr
// 		NoColor: true,
// 	}

// 	// At the end of the test, run `terraform destroy` to clean up any resources that were created
// 	defer terraform.Destroy(t, terraformOptions)

// 	// This will run `terraform init` and `terraform apply` and fail the test if there are any errors
// 	terraform.InitAndApply(t, terraformOptions)

// 	// Run `terraform output` to get the values of output variables
// 	actualDDOSPlanName := terraform.Output(t, terraformOptions, "ddos_plan_name")
// 	actualDDOSPlanLocation := terraform.Output(t, terraformOptions, "ddos_plan_location")

// 	// Verify we're getting back the outputs we expect
// 	DDOSPlanName := strings.Join(DDOSPlanNameStrings, "-")
// 	assert.Equal(t, DDOSPlanName, actualDDOSPlanName)
// 	assert.Equal(t, DDOSPlanLocation, actualDDOSPlanLocation)
// }

func TestTerraformVirtualNetworkTags(t *testing.T) {
	t.Parallel()

	// randomString := string(RandomString(4))
	resourceGroupNameStrings := []string{"mns", "rg", "02"}

	DDOSPlanNameStrings := []string{"mns", "ddos", "02"}
	DDOSPlanLocation := "westeurope"
	DDOSPlanTags := map[string]string{"environment": "development", "location": "westeurope"}

	// Root folder where terraform files should be (relative to the test folder)
	rootFolder := ".."
	// Relative path to terraform module being tested from the root folder
	terraformFolderRelativeToRoot := "calling-code"
	// Copy the terraform folder to a temp folder
	tempTestFolder := test_structure.CopyTerraformFolderToTemp(t, rootFolder, terraformFolderRelativeToRoot)

	terraformOptions := &terraform.Options{
		// The path to where our Terraform code is located
		TerraformDir: tempTestFolder,

		// Variables to pass to our Terraform code using -var options
		Vars: map[string]interface{}{
			"resource_group_name_strings":   resourceGroupNameStrings,
			"resource_group_name_separator": "-",

			"ddos_plan_name_strings":   DDOSPlanNameStrings,
			"ddos_plan_name_separator": "-",
			"ddos_plan_location":       DDOSPlanLocation,
			"ddos_plan_tags":           DDOSPlanTags,
		},

		BackendConfig: map[string]interface{}{
			"bucket": "byt-terraform-state-bucket",
			"key":    "testing/test01.tfstate",
			"region": "eu-west-2",
		},
		// Variables to pass to our Terraform code using -var-file options
		// VarFiles: []string{"varfile.tfvars"},

		// Disable colors in Terraform commands so its easier to parse stdout/stderr
		NoColor: true,
	}

	// At the end of the test, run `terraform destroy` to clean up any resources that were created
	defer terraform.Destroy(t, terraformOptions)

	// This will run `terraform init` and `terraform apply` and fail the test if there are any errors
	terraform.InitAndApply(t, terraformOptions)

	// Run `terraform output` to get the values of output variables
	actualDDOSPlanName := terraform.Output(t, terraformOptions, "ddos_plan_name")
	actualDDOSPlanLocation := terraform.Output(t, terraformOptions, "ddos_plan_location")
	actualDDOSPlanTags := terraform.OutputMap(t, terraformOptions, "ddos_plan_tags")

	// Verify we're getting back the outputs we expect
	DDOSPlanName := strings.Join(DDOSPlanNameStrings, "-")
	assert.Equal(t, DDOSPlanName, actualDDOSPlanName)
	assert.Equal(t, DDOSPlanLocation, actualDDOSPlanLocation)
	assert.Equal(t, DDOSPlanTags, actualDDOSPlanTags)
}

func RandomString(n int) string {
	var letters = []rune("abcdefghijklmnopqrstuvwxyz")

	s := make([]rune, n)
	for i := range s {
		s[i] = letters[rand.Intn(len(letters))]
	}
	return string(s)
}

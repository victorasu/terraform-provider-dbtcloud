package license_map_test

import (
	"fmt"
	"regexp"
	"strconv"
	"testing"

	"github.com/dbt-labs/terraform-provider-dbtcloud/pkg/framework/acctest_helper"
	"github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

func TestAccDbtCloudLicenseMapResource(t *testing.T) {

	groupName := acctest.RandStringFromCharSet(10, acctest.CharSetAlpha)
	groupName2 := acctest.RandStringFromCharSet(10, acctest.CharSetAlpha)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest_helper.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest_helper.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckDbtCloudLicenseMapDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccDbtCloudLicenseMapResourceBasicConfig("developer", groupName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDbtCloudLicenseMapExists("dbtcloud_license_map.test_license_map"),
					resource.TestCheckResourceAttr(
						"dbtcloud_license_map.test_license_map",
						"license_type",
						"developer",
					),
					resource.TestCheckResourceAttr(
						"dbtcloud_license_map.test_license_map",
						"sso_license_mapping_groups.#",
						"1",
					),
					resource.TestCheckResourceAttr(
						"dbtcloud_license_map.test_license_map",
						"sso_license_mapping_groups.0",
						groupName,
					),
				),
			},
			// MODIFY
			{
				Config: testAccDbtCloudLicenseMapResourceBasicConfig(
					"developer",
					fmt.Sprintf(`%s","%s`, groupName, groupName2),
				),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDbtCloudLicenseMapExists("dbtcloud_license_map.test_license_map"),
					resource.TestCheckResourceAttr(
						"dbtcloud_license_map.test_license_map",
						"license_type",
						"developer",
					),
					resource.TestCheckResourceAttr(
						"dbtcloud_license_map.test_license_map",
						"sso_license_mapping_groups.#",
						"2",
					),
					resource.TestCheckTypeSetElemAttr(
						"dbtcloud_license_map.test_license_map",
						"sso_license_mapping_groups.*",
						groupName,
					),
					resource.TestCheckTypeSetElemAttr(
						"dbtcloud_license_map.test_license_map",
						"sso_license_mapping_groups.*",
						groupName2,
					),
				),
			},
			// IMPORT
			{
				ResourceName:            "dbtcloud_license_map.test_license_map",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{},
			},
		},
	})
}

func testAccDbtCloudLicenseMapResourceBasicConfig(licenseType string, groupName string) string {
	return fmt.Sprintf(`

resource "dbtcloud_license_map" "test_license_map" {
    license_type       = "%s"
    sso_license_mapping_groups = ["%s"]
}
`, licenseType, groupName)
}

func testAccCheckDbtCloudLicenseMapExists(resource string) resource.TestCheckFunc {
	return func(state *terraform.State) error {
		rs, ok := state.RootModule().Resources[resource]
		if !ok {
			return fmt.Errorf("Not found: %s", resource)
		}
		if rs.Primary.ID == "" {
			return fmt.Errorf("No Record ID is set")
		}
		apiClient, err := acctest_helper.SharedClient()
		if err != nil {
			return fmt.Errorf("Issue getting the client")
		}
		licenseMapID, err := strconv.Atoi(rs.Primary.ID)
		if err != nil {
			return fmt.Errorf("Can't get groupID")
		}
		_, err = apiClient.GetLicenseMap(licenseMapID)
		if err != nil {
			return fmt.Errorf("error fetching item with resource %s. %s", resource, err)
		}
		return nil
	}
}

func testAccCheckDbtCloudLicenseMapDestroy(s *terraform.State) error {
	apiClient, err := acctest_helper.SharedClient()
	if err != nil {
		return fmt.Errorf("Issue getting the client")
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "dbtcloud_license_map" {
			continue
		}
		licenseMapID, err := strconv.Atoi(rs.Primary.ID)
		if err != nil {
			return fmt.Errorf("Can't get licenseMapID")
		}
		_, err = apiClient.GetLicenseMap(licenseMapID)
		if err == nil {
			return fmt.Errorf("License Map still exists")
		}
		notFoundErr := "resource-not-found"
		expectedErr := regexp.MustCompile(notFoundErr)
		if !expectedErr.Match([]byte(err.Error())) {
			return fmt.Errorf("expected %s, got %s", notFoundErr, err)
		}
	}

	return nil
}

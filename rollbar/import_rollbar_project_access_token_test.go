package rollbar

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"testing"
)

func TestAccRollbarProjectAccessToken_importBasic(t *testing.T) {
	projectName := fmt.Sprintf("project-tftest-%s", acctest.RandString(10))
	tokenName := fmt.Sprintf("token-tftest-%s", acctest.RandString(10))

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
		},
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckRollbarProjectAccessToken_basic(projectName, tokenName),
			},
			{
				ResourceName:      "rollbar_project_access_token.foobar",
				ImportStateIdFunc: testAccRollbarProjectAccessTokenImportStateIdFunc("rollbar_project_access_token.foobar"),
				ImportState:       true,
				ImportStateCheck: func(s []*terraform.InstanceState) error {
					if len(s) != 1 {
						return fmt.Errorf("expected 1 state: %#v", s)
					}

					rs := s[0]

					// Do a simple check of state since the resource ID is complex
					if rs.Attributes["access_token"] == "" || rs.Attributes["name"] == "" {
						return fmt.Errorf("rollbar_project_access_token import not successful")
					}

					return nil
				},
			},
		},
	})
}

func testAccRollbarProjectAccessTokenImportStateIdFunc(resourceName string) resource.ImportStateIdFunc {
	return func(s *terraform.State) (string, error) {
		rs, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return "", fmt.Errorf("Not found: %s", resourceName)
		}

		return fmt.Sprintf("%s:%s", rs.Primary.Attributes["project_id"],
			rs.Primary.Attributes["access_token"]), nil
	}
}

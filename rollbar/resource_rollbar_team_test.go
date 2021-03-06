package rollbar

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"testing"
)

func TestAccRollbarTeam_Basic(t *testing.T) {
	name := fmt.Sprintf("tftest-%s", acctest.RandString(10))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckRollbarTeam_basic(name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"rollbar_team.foobar", "name", name),
					resource.TestCheckResourceAttr(
						"rollbar_team.foobar", "access_level", "standard"),
				),
			},
		},
	})
}

func testAccCheckRollbarTeam_basic(name string) string {
	return fmt.Sprintf(`
resource "rollbar_team" "foobar" {
	name = "%s"
	access_level = "standard"
}
`, name)
}

package dropbox

import (
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
)

func TestAccDropboxPaperDoc_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccDropboxPaperDocConfig,
				Check:  resource.ComposeTestCheckFunc(),
			},
		},
	})
}

func TestAccDropboxPaperDoc_foldered(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccDropboxPaperDocFolderConfig,
				Check:  resource.ComposeTestCheckFunc(),
			},
		},
	})
}

const testAccDropboxPaperDocConfig = `
resource "dropbox_paper_document" "doc" {
	content_file  = "${file("../token.txt")}"
	import_format = "plain_text"
}
`

// TODO: Input valid parent folder ID
const testAccDropboxPaperDocFolderConfig = `
resource "dropbox_paper_document" "foldered_doc" {
	content_file  = "${file("../token.txt")}"
	parent_folder = ""
	import_format = "plain_text"
}
`

package helpers

import (
	"fmt"
	"strings"

	"github.com/jesseduffield/gocui"
	"github.com/jesseduffield/lazygit/pkg/gui/types"
)

type ArchiveBundleHelper struct {
	c *HelperCommon
}

func NewArchiveBundleHelper(
	c *HelperCommon,
) *ArchiveBundleHelper {
	return &ArchiveBundleHelper{
		c: c,
	}
}

func (self *ArchiveBundleHelper) CreateArchive(refName string) error {
	self.c.LogAction(fmt.Sprintf("create archive called %s", refName))

	// TODO: Add a waiting status thingy
	return self.c.Prompt(types.PromptOpts{
		Title: "Name of folder inside archive (optional)",
		HandleConfirm: func(prefix string) error {
			// If the prefix doesn't have a trailing /,
			// git will apply the prefix to each file instead of
			// putting the files inside a folder inside the archive
			if prefix != "" && !strings.HasSuffix(prefix, "/") {
				prefix += "/"
			}
			self.c.LogAction(fmt.Sprintf("Selected prefix: %s", prefix))

			return self.c.Prompt(types.PromptOpts{
				Title: "Choose an archive name (without extension)",
				HandleConfirm: func(fileName string) error {
					self.c.LogAction(fmt.Sprintf("Selected fileName has no extension: %s", fileName))
					return self.c.Menu(types.CreateMenuOptions{
						Title: "Amend commit attribute",
						Items: []*types.MenuItem{
							{
								Label: ".zip",
								OnPress: func() error {
									return self.runArchiveCommand(refName, fileName, prefix, ".zip")
								},
								Key: '1',
							},
							{
								Label: ".tar.gz",
								OnPress: func() error {
									return self.runArchiveCommand(refName, fileName, prefix, ".tar.gz")
								},
								Key: '2',
							},
							{
								Label: ".tar",
								OnPress: func() error {
									return self.runArchiveCommand(refName, fileName, prefix, ".tar")
								},
								Key: '3',
							},
							{
								Label: ".tgz",
								OnPress: func() error {
									return self.runArchiveCommand(refName, fileName, prefix, ".tgz")
								},
								Key: '4',
							},
						},
					})
				},
			})
		},
	})
}

func (self *ArchiveBundleHelper) runArchiveCommand(refName string, fileName string, prefix string, suffix string) error {
	return self.c.WithWaitingStatus("Archiving...", func(gocui.Task) error {
		return self.c.Git().Archive.Archive(refName, fileName+suffix, prefix)
	})
}

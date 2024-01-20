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
	return self.c.Prompt(types.PromptOpts{
		Title: self.c.Tr.ArchiveChoosePrefixTitle,
		HandleConfirm: func(prefix string) error {
			// If the prefix doesn't have a trailing /,
			// git will apply the prefix to each file instead of
			// putting the files inside a folder inside the archive
			if prefix != "" && !strings.HasSuffix(prefix, "/") {
				prefix += "/"
			}

			return self.c.Prompt(types.PromptOpts{
				Title: self.c.Tr.ArchiveChooseFileName,
				HandleConfirm: func(fileName string) error {
					validArchiveFormats, err := self.c.Git().Archive.GetValidArchiveFormats()
					if err != nil {
						return err
					}

					menuItems := make([]*types.MenuItem, len(validArchiveFormats))

					for i, format := range validArchiveFormats {
						format := format

						menuItems[i] = &types.MenuItem{
							Label: format,
							OnPress: func() error {
								return self.runArchiveCommand(refName, fileName, prefix, format)
							},
						}
					}

					return self.c.Menu(types.CreateMenuOptions{
						Title: self.c.Tr.ArchiveChooseFormatMenuTitle,
						Items: menuItems,
					})
				},
			})
		},
	})
}

func (self *ArchiveBundleHelper) runArchiveCommand(refName string, fileName string, prefix string, suffix string) error {
	self.c.LogAction(fmt.Sprintf("Creating archive for ref: %s", refName))
	return self.c.WithWaitingStatus(self.c.Tr.ArchiveWaitingStatusMessage, func(gocui.Task) error {
		return self.c.Git().Archive.Archive(refName, fileName+suffix, prefix)
	})
}

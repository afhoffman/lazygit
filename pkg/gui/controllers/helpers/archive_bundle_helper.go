package helpers

import (
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
		Title: "Name of folder inside archive (optional)",
		HandleConfirm: func(prefix string) error {
			// If the prefix doesn't have a trailing /,
			// git will apply the prefix to each file instead of
			// putting the files inside a folder inside the archive
			if prefix != "" && !strings.HasSuffix(prefix, "/") {
				prefix += "/"
			}

			return self.c.Prompt(types.PromptOpts{
				Title: "Choose an archive name (without extension)",
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
						Title: "Select archive format",
						Items: menuItems,
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

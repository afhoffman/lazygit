package helpers

import (
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

func (self *ArchiveBundleHelper) CreateBundle(refName string) error {
	return self.c.Prompt(types.PromptOpts{
		Title: self.c.Tr.BundleChooseBundleName,
		HandleConfirm: func(bundleName string) error {
			return self.runBundleCommand(bundleName, refName)
		},
	})
}

func (self *ArchiveBundleHelper) runBundleCommand(bundleName string, refNames ...string) error {
	return self.c.WithWaitingStatus(self.c.Tr.BundleWaitingStatusMessage, func(gocui.Task) error {
		return self.c.Git().Bundle.Bundle(bundleName, refNames...)
	})
}

package bundle

import (
	"github.com/jesseduffield/lazygit/pkg/config"
	. "github.com/jesseduffield/lazygit/pkg/integration/components"
	"github.com/jesseduffield/lazygit/pkg/integration/tests/shared"
)

var Tag = NewIntegrationTest(NewIntegrationTestArgs{
	Description:  "Create an archive of a tag",
	ExtraCmdArgs: []string{},
	Skip:         false,
	SetupConfig:  func(config *config.AppConfig) {},
	SetupRepo: func(shell *Shell) {
		shared.MergeConflictsSetup(shell)
		shell.CreateLightweightTag("tag", "HEAD^")
	},
	Run: func(t *TestDriver, keys config.KeybindingConfig) {
		createdBundleName := "test-bundle.bundle"

		t.Views().Tags().
			Focus().
			Lines(
				Contains("tag").IsSelected(),
			).Press(keys.Branches.Bundle)

		t.ExpectPopup().Menu().Title(Equals("Create archive or bundle")).
			Select(Contains("bundle")).
			Confirm()

		t.ExpectPopup().Prompt().
			Title(Equals("Choose bundle name")).
			Type(createdBundleName).
			Confirm()

		// This could be better tested by checking refs
		// returned from `git ls-remote`, but honestly if we
		// made it this far without throwing an error, it's pretty
		// likely that something outside the scope of lazygit has gone wrong
		//
		// It would be worth making this test more robust when support for
		// archiving a range of tags is added so we can make sure all of
		// the desired tags make it into the bundle.
		t.FileSystem().PathPresent(createdBundleName)
	},
})

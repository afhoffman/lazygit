package archive_and_bundle

import (
	"github.com/jesseduffield/lazygit/pkg/config"
	. "github.com/jesseduffield/lazygit/pkg/integration/components"
	"github.com/jesseduffield/lazygit/pkg/integration/tests/shared"
)

var ArchiveBranch = NewIntegrationTest(NewIntegrationTestArgs{
	Description:  "Create an archive of a branch",
	ExtraCmdArgs: []string{},
	Skip:         false,
	SetupConfig:  func(config *config.AppConfig) {},
	SetupRepo: func(shell *Shell) {
		shared.MergeConflictsSetup(shell)
	},
	Run: func(t *TestDriver, keys config.KeybindingConfig) {
		t.Views().Branches().
			Focus().
			Lines(
				Contains("first-change-branch").IsSelected(),
				Contains("second-change-branch"),
				Contains("original-branch"),
			).
			Press(keys.Branches.ArchiveOrBundle)

		// t.ExpectPopup().CommitMessagePanel().
		// 	Title(Equals("Bing bong")).
		// 	Type("new-tag").
		// 	Confirm()
		//
		// t.Views().Tags().Focus().
		// 	Lines(
		// 		MatchesRegexp(`new-tag`).IsSelected(),
		// 	)
		//
		// t.Git().
		// 	TagNamesAt("HEAD", []string{}).
		// 	TagNamesAt("master", []string{"new-tag"})
	},
})

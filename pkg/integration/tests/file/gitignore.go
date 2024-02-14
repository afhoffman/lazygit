package file

import (
	"github.com/jesseduffield/lazygit/pkg/config"
	. "github.com/jesseduffield/lazygit/pkg/integration/components"
)

var Gitignore = NewIntegrationTest(NewIntegrationTestArgs{
	Description:  "Verify that we can't ignore the .gitignore file, then ignore/exclude other files",
	ExtraCmdArgs: []string{},
	Skip:         false,
	SetupConfig: func(config *config.AppConfig) {
	},
	SetupRepo: func(shell *Shell) {
		shell.CreateFile(".gitignore", "")
		shell.CreateFile("toExclude", "")
		shell.CreateFile("toIgnore", "")
		shell.CreateFile("toRangeIgnore1", "")
		shell.CreateFile("toRangeIgnore2", "")
		shell.CreateFile("toRangeIgnore3", "")
		shell.CreateFile("folder1/toRangeIgnore4", "")
		shell.CreateFile("folder1/toRangeIgnore5", "")
	},
	Run: func(t *TestDriver, keys config.KeybindingConfig) {
		t.Views().Files().
			IsFocused().
			Lines(
				Contains("folder1").IsSelected(),
				Contains(`?? toRangeIgnore4`),
				Contains(`?? toRangeIgnore5`),
				Contains(`?? .gitignore`),
				Contains(`?? toExclude`),
				Contains(`?? toIgnore`),
				Contains(`?? toRangeIgnore1`),
				Contains(`?? toRangeIgnore2`),
				Contains(`?? toRangeIgnore3`),
			).
			NavigateToLine(Contains(".gitignore")).
			Press(keys.Files.IgnoreFile).
			// ensure we can't exclude the .gitignore file
			Tap(func() {
				t.ExpectPopup().Menu().Title(Equals("Ignore or exclude file")).Select(Contains("Add to .git/info/exclude")).Confirm()

				t.ExpectPopup().Alert().Title(Equals("Error")).Content(Equals("Cannot exclude .gitignore")).Confirm()
			}).
			SelectNextItem().
			Press(keys.Files.IgnoreFile).
			// exclude a file
			Tap(func() {
				t.ExpectPopup().Menu().Title(Equals("Ignore or exclude file")).Select(Contains("Add to .git/info/exclude")).Confirm()

				t.FileSystem().FileContent(".gitignore", Equals(""))
				t.FileSystem().FileContent(".git/info/exclude", Contains("toExclude"))
			}).
			Press(keys.Files.IgnoreFile).
			// ignore a file
			Tap(func() {
				t.ExpectPopup().Menu().Title(Equals("Ignore or exclude file")).Select(Contains("Add to .gitignore")).Confirm()

				t.FileSystem().FileContent(".gitignore", Equals("toIgnore\n"))
				t.FileSystem().FileContent(".git/info/exclude", Contains("toExclude"))
			})

		t.Views().Files().
			IsFocused().
			Lines(
				Contains("folder1"),
				Contains(`?? toRangeIgnore4`),
				Contains(`?? toRangeIgnore5`),
				Contains(`?? .gitignore`),
				Contains(`?? toRangeIgnore1`).IsSelected(),
				Contains(`?? toRangeIgnore2`),
				Contains(`?? toRangeIgnore3`),
			).
			NavigateToLine(Contains("folder1")).
			Press(keys.Universal.ToggleRangeSelect).
			NavigateToLine(Contains("toRangeIgnore3")).
			Press(keys.Files.IgnoreFile).
			Tap(func() {
				t.ExpectPopup().Menu().Title(Equals("Ignore or exclude file")).Select(Contains("Add to .gitignore")).Confirm()

				// This selection will include the .gitignore file. Make sure we can't ignore that.
				t.ExpectPopup().Alert().Title(Equals("Error")).Content(Equals("Cannot ignore .gitignore")).Confirm()

				t.FileSystem().FileContent(".gitignore", Equals("toIgnore\n"+
					"folder1\n"+
					"toRangeIgnore1\n"+
					"toRangeIgnore2\n"+
					"toRangeIgnore3\n"))
				t.FileSystem().FileContent(".git/info/exclude", Contains("toExclude"))
			})
	},
})

package git_commands

type BundleCommands struct {
	*GitCommon
}

func NewBundleCommands(gitCommon *GitCommon) *BundleCommands {
	return &BundleCommands{
		GitCommon: gitCommon,
	}
}

func (self *BundleCommands) Bundle(bundleName string, refNames ...string) error {
	cmdArgs := NewGitCmd("bundle").
		Arg("create").
		Arg(bundleName).
		Arg(refNames...)

	return self.cmd.New(cmdArgs.ToArgv()).DontLog().Run()
}

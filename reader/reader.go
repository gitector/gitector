package reader

type InitParams struct {
	Scope            string
	ConfigFilePath   string
	Limit            int
	DirectInput      string
	Directory        string
	UsingDirectInput bool
	Output           string
}

type GitCommit struct {
	Title       string
	Description string
	RawMessage  string
	Signature   Signature
	FilesCount  int
	IsMerge     bool
}

type Signature struct {
	Name  string
	Email string
}

func ReadGitCommits(configParams InitParams) []GitCommit {
	if !configParams.UsingDirectInput {
		return ReadGitCommitsFromDirectory(configParams.Directory, configParams.Scope)
	} else {
		return ReadDirectInput(configParams.Directory)
	}
}

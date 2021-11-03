	"io/ioutil"
	Name               string
	OldName            string
	Index              int
	Addition, Deletion int
	Type               DiffFileType
	IsCreated          bool
	IsDeleted          bool
	IsBin              bool
	IsLFSFile          bool
	IsRenamed          bool
	IsAmbiguous        bool
	IsSubmodule        bool
	Sections           []*DiffSection
	IsIncomplete       bool
	IsProtected        bool
			_, err := io.Copy(ioutil.Discard, reader)
	// FIXME: There are numerous issues with this:
	// - and this doesn't really account for changes in encoding
	var buf bytes.Buffer
		buf.Reset()
				buf.WriteString(l.Content[1:])
				buf.WriteString("\n")
		charsetLabel, err := charset.DetectEncoding(buf.Bytes())
		if charsetLabel != "UTF-8" && err == nil {
			encoding, _ := stdcharset.Lookup(charsetLabel)
			if encoding != nil {
				d := encoding.NewDecoder()
				for _, sec := range f.Sections {
					for _, l := range sec.Lines {
						if l.Type == DiffLineSection {
							continue
						}
						if c, _, err := transform.String(d, l.Content[1:]); err == nil {
							l.Content = l.Content[0:1] + c
						}
			diffLine := &DiffLine{Type: DiffLineAdd, RightIdx: rightLine}
			diffLine := &DiffLine{Type: DiffLineDel, LeftIdx: leftLine}
		if line[1:] == models.LFSMetaFileIdentifier {
		} else if curFileLFSPrefix && strings.HasPrefix(line[1:], models.LFSMetaFileOidPrefix) {
			oid := strings.TrimPrefix(line[1:], models.LFSMetaFileOidPrefix)
				m := &models.LFSMetaObject{Oid: oid}
				count, err := models.Count(m)
// GetDiffRange builds a Diff between two commits of a repository.
// passing the empty string as beforeCommitID returns a diff from the
// parent commit.
func GetDiffRange(repoPath, beforeCommitID, afterCommitID string, maxLines, maxLineCharacters, maxFiles int) (*Diff, error) {
	return GetDiffRangeWithWhitespaceBehavior(repoPath, beforeCommitID, afterCommitID, maxLines, maxLineCharacters, maxFiles, "")
}

func GetDiffRangeWithWhitespaceBehavior(repoPath, beforeCommitID, afterCommitID string, maxLines, maxLineCharacters, maxFiles int, whitespaceBehavior string) (*Diff, error) {
	gitRepo, err := git.OpenRepository(repoPath)
	if err != nil {
		return nil, err
	}
	defer gitRepo.Close()
	// FIXME: graceful: These commands should likely have a timeout
	ctx, cancel := context.WithCancel(git.DefaultContext)
	var cmd *exec.Cmd
		diffArgs := []string{"diff", "--src-prefix=\\a/", "--dst-prefix=\\b/", "-M"}
		cmd = exec.CommandContext(ctx, git.GitExecutable, diffArgs...)
		diffArgs := []string{"diff", "--src-prefix=\\a/", "--dst-prefix=\\b/", "-M"}
		cmd = exec.CommandContext(ctx, git.GitExecutable, diffArgs...)
	shortstatArgs := []string{beforeCommitID + "..." + afterCommitID}
// GetDiffCommit builds a Diff representing the given commitID.
func GetDiffCommit(repoPath, commitID string, maxLines, maxLineCharacters, maxFiles int) (*Diff, error) {
	return GetDiffRangeWithWhitespaceBehavior(repoPath, "", commitID, maxLines, maxLineCharacters, maxFiles, "")
}

func GetDiffCommitWithWhitespaceBehavior(repoPath, commitID string, maxLines, maxLineCharacters, maxFiles int, whitespaceBehavior string) (*Diff, error) {
	return GetDiffRangeWithWhitespaceBehavior(repoPath, "", commitID, maxLines, maxLineCharacters, maxFiles, whitespaceBehavior)
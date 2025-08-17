package command

type Repository string
type Branch string
type Remote string
type CommitMessage string
type Author string
type Email string

type ForceFlag bool

const (
	Force   ForceFlag = true
	NoForce ForceFlag = false
)

type VerboseFlag bool

const (
	Verbose   VerboseFlag = true
	NoVerbose VerboseFlag = false
)

type QuietFlag bool

const (
	Quiet   QuietFlag = true
	NoQuiet QuietFlag = false
)

type DryRunFlag bool

const (
	DryRun   DryRunFlag = true
	NoDryRun DryRunFlag = false
)

type AllFlag bool

const (
	All   AllFlag = true
	NoAll AllFlag = false
)

type InteractiveFlag bool

const (
	Interactive   InteractiveFlag = true
	NoInteractive InteractiveFlag = false
)

type flags struct {
	Repository    Repository
	Branch        Branch
	Remote        Remote
	CommitMessage CommitMessage
	Author        Author
	Email         Email
	Force         ForceFlag
	Verbose       VerboseFlag
	Quiet         QuietFlag
	DryRun        DryRunFlag
	All           AllFlag
	Interactive   InteractiveFlag
}

func (r Repository) Configure(flags *flags)      { flags.Repository = r }
func (b Branch) Configure(flags *flags)          { flags.Branch = b }
func (r Remote) Configure(flags *flags)          { flags.Remote = r }
func (c CommitMessage) Configure(flags *flags)   { flags.CommitMessage = c }
func (a Author) Configure(flags *flags)          { flags.Author = a }
func (e Email) Configure(flags *flags)           { flags.Email = e }
func (f ForceFlag) Configure(flags *flags)       { flags.Force = f }
func (v VerboseFlag) Configure(flags *flags)     { flags.Verbose = v }
func (q QuietFlag) Configure(flags *flags)       { flags.Quiet = q }
func (d DryRunFlag) Configure(flags *flags)      { flags.DryRun = d }
func (a AllFlag) Configure(flags *flags)         { flags.All = a }
func (i InteractiveFlag) Configure(flags *flags) { flags.Interactive = i }

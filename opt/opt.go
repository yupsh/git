package opt

// Custom types for parameters
type Repository string
type Branch string
type Remote string
type CommitMessage string
type Author string
type Email string

// Boolean flag types with constants
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

// Flags represents the configuration options for git commands
type Flags struct {
	Repository    Repository      // Repository path or URL
	Branch        Branch          // Branch name
	Remote        Remote          // Remote name
	CommitMessage CommitMessage   // Commit message
	Author        Author          // Author name
	Email         Email           // Author email
	Force         ForceFlag       // Force operation
	Verbose       VerboseFlag     // Verbose output
	Quiet         QuietFlag       // Quiet output
	DryRun        DryRunFlag      // Dry run mode
	All           AllFlag         // All files/branches
	Interactive   InteractiveFlag // Interactive mode
}

// Configure methods for the opt system
func (r Repository) Configure(flags *Flags)      { flags.Repository = r }
func (b Branch) Configure(flags *Flags)          { flags.Branch = b }
func (r Remote) Configure(flags *Flags)          { flags.Remote = r }
func (c CommitMessage) Configure(flags *Flags)   { flags.CommitMessage = c }
func (a Author) Configure(flags *Flags)          { flags.Author = a }
func (e Email) Configure(flags *Flags)           { flags.Email = e }
func (f ForceFlag) Configure(flags *Flags)       { flags.Force = f }
func (v VerboseFlag) Configure(flags *Flags)     { flags.Verbose = v }
func (q QuietFlag) Configure(flags *Flags)       { flags.Quiet = q }
func (d DryRunFlag) Configure(flags *Flags)      { flags.DryRun = d }
func (a AllFlag) Configure(flags *Flags)         { flags.All = a }
func (i InteractiveFlag) Configure(flags *Flags) { flags.Interactive = i }

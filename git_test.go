package git

import (
	`bytes`
	`context`
	`io`
	`reflect`
	`testing`

	yup `github.com/yupsh/framework`
)

func TestAdd(t *testing.T) {
	type args struct {
		parameters []any
	}
	tests := []struct {
		name string
		args args
		want yup.Command
	}{
		{},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Add(tt.args.parameters...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Add() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAddAll(t *testing.T) {
	type args struct {
		parameters []any
	}
	tests := []struct {
		name string
		args args
		want yup.Command
	}{
		{},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := AddAll(tt.args.parameters...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AddAll() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCheckout(t *testing.T) {
	type args struct {
		branch     string
		parameters []any
	}
	tests := []struct {
		name string
		args args
		want yup.Command
	}{
		{},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Checkout(tt.args.branch, tt.args.parameters...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Checkout() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCheckoutNew(t *testing.T) {
	type args struct {
		branch     string
		parameters []any
	}
	tests := []struct {
		name string
		args args
		want yup.Command
	}{
		{},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CheckoutNew(tt.args.branch, tt.args.parameters...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CheckoutNew() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCommit(t *testing.T) {
	type args struct {
		message    string
		parameters []any
	}
	tests := []struct {
		name string
		args args
		want yup.Command
	}{
		{},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Commit(tt.args.message, tt.args.parameters...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Commit() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCommitAll(t *testing.T) {
	type args struct {
		message    string
		parameters []any
	}
	tests := []struct {
		name string
		args args
		want yup.Command
	}{
		{},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CommitAll(tt.args.message, tt.args.parameters...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CommitAll() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDiff(t *testing.T) {
	type args struct {
		parameters []any
	}
	tests := []struct {
		name string
		args args
		want yup.Command
	}{
		{},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Diff(tt.args.parameters...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Diff() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFetch(t *testing.T) {
	type args struct {
		parameters []any
	}
	tests := []struct {
		name string
		args args
		want yup.Command
	}{
		{},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Fetch(tt.args.parameters...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Fetch() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFetchOrigin(t *testing.T) {
	type args struct {
		parameters []any
	}
	tests := []struct {
		name string
		args args
		want yup.Command
	}{
		{},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FetchOrigin(tt.args.parameters...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FetchOrigin() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGit(t *testing.T) {
	type args struct {
		parameters []any
	}
	tests := []struct {
		name string
		args args
		want yup.Command
	}{
		{},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Git(tt.args.parameters...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Git() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLog(t *testing.T) {
	type args struct {
		parameters []any
	}
	tests := []struct {
		name string
		args args
		want yup.Command
	}{
		{},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Log(tt.args.parameters...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Log() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLogOneline(t *testing.T) {
	type args struct {
		parameters []any
	}
	tests := []struct {
		name string
		args args
		want yup.Command
	}{
		{},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := LogOneline(tt.args.parameters...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LogOneline() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPull(t *testing.T) {
	type args struct {
		parameters []any
	}
	tests := []struct {
		name string
		args args
		want yup.Command
	}{
		{},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Pull(tt.args.parameters...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Pull() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPullOrigin(t *testing.T) {
	type args struct {
		branch     string
		parameters []any
	}
	tests := []struct {
		name string
		args args
		want yup.Command
	}{
		{},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := PullOrigin(tt.args.branch, tt.args.parameters...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PullOrigin() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPush(t *testing.T) {
	type args struct {
		parameters []any
	}
	tests := []struct {
		name string
		args args
		want yup.Command
	}{
		{},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Push(tt.args.parameters...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Push() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPushOrigin(t *testing.T) {
	type args struct {
		branch     string
		parameters []any
	}
	tests := []struct {
		name string
		args args
		want yup.Command
	}{
		{},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := PushOrigin(tt.args.branch, tt.args.parameters...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PushOrigin() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStatus(t *testing.T) {
	type args struct {
		parameters []any
	}
	tests := []struct {
		name string
		args args
		want yup.Command
	}{
		{},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Status(tt.args.parameters...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Status() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_command_Execute(t *testing.T) {
	type args struct {
		ctx   context.Context
		stdin io.Reader
	}
	tests := []struct {
		name       string
		c          command
		args       args
		wantStdout string
		wantStderr string
		wantErr    bool
	}{
		{},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			stdout := &bytes.Buffer{}
			stderr := &bytes.Buffer{}
			err := tt.c.Execute(tt.args.ctx, tt.args.stdin, stdout, stderr)
			if (err != nil) != tt.wantErr {
				t.Errorf("Execute() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotStdout := stdout.String(); gotStdout != tt.wantStdout {
				t.Errorf("Execute() gotStdout = %v, want %v", gotStdout, tt.wantStdout)
			}
			if gotStderr := stderr.String(); gotStderr != tt.wantStderr {
				t.Errorf("Execute() gotStderr = %v, want %v", gotStderr, tt.wantStderr)
			}
		})
	}
}

func Test_command_String(t *testing.T) {
	tests := []struct {
		name string
		c    command
		want string
	}{
		{},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.c.String(); got != tt.want {
				t.Errorf("String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_command_buildExecOptions(t *testing.T) {
	tests := []struct {
		name string
		c    command
		want []any
	}{
		{},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.c.buildExecOptions(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("buildExecOptions() = %v, want %v", got, tt.want)
			}
		})
	}
}

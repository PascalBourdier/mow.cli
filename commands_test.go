package cli

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNormalize(t *testing.T) {
	cases := []struct {
		opts   func(*Cmd)
		input  []string
		output []string
	}{
		{func(cmd *Cmd) {
		}, []string{}, []string{}},
		{func(cmd *Cmd) {
			cmd.Bool(BoolOpt{Name: "a", Value: true, Desc: ""})
		}, []string{"ab"}, []string{"ab"}},
		{func(cmd *Cmd) {
			cmd.Bool(BoolOpt{Name: "a", Value: true, Desc: ""})
		}, []string{"-a"}, []string{"-a", "true"}},
		{func(cmd *Cmd) {
			cmd.Bool(BoolOpt{Name: "a", Value: true, Desc: ""})
		}, []string{"-a=false"}, []string{"-a", "false"}},
		{func(cmd *Cmd) {
			cmd.Bool(BoolOpt{Name: "a", Value: true, Desc: ""})
			cmd.Bool(BoolOpt{Name: "b", Value: true, Desc: ""})
		}, []string{"-ab"}, []string{"-a", "true", "-b", "true"}},
		{func(cmd *Cmd) {
			cmd.Bool(BoolOpt{Name: "a", Value: true, Desc: ""})
			cmd.Bool(BoolOpt{Name: "b", Value: true, Desc: ""})
			cmd.String(StringOpt{Name: "c", Value: "", Desc: ""})
		}, []string{"-abc", "hello"}, []string{"-a", "true", "-b", "true", "-c", "hello"}},
		{func(cmd *Cmd) {
			cmd.String(StringOpt{Name: "s", Value: "", Desc: ""})
		}, []string{"-shello"}, []string{"-s", "hello"}},
		{func(cmd *Cmd) {
			cmd.Bool(BoolOpt{Name: "a", Value: true, Desc: ""})
			cmd.String(StringOpt{Name: "b", Value: "", Desc: ""})
		}, []string{"-ab", "test"}, []string{"-a", "true", "-b", "test"}},
	}
	for _, cas := range cases {
		cmd := &Cmd{optionsIdx: map[string]*opt{}}
		cas.opts(cmd)
		nz, cons, err := cmd.normalize(cas.input)
		require.Nil(t, err, "Unexpected error %v", err)
		t.Logf("%v %v", nz, cons)
		require.Equal(t, cas.output, nz)

	}
}

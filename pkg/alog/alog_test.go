package alog

import (
	"bytes"
	"fmt"
	"os"
	"regexp"
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestNew(t *testing.T) {
	cfg := &Config{}
	got := New(cfg)
	want := os.Stdout

	require.NotNil(t, got)
	require.Equal(t, want, got.writer)
}

func TestNew_WithFile(t *testing.T) {
	cfg := &Config{Out: "logs.out"}

	got := New(cfg)
	require.NotNil(t, got)

	_, err := os.Stat(cfg.Out)
	require.NoError(t, err)

	err = os.Remove(cfg.Out)
	if err != nil {
		t.Fatal(err)
	}
}

func Test_WriteToLog(t *testing.T) {
	const (
		msgInfo = iota
		msgError
	)

	var output bytes.Buffer
	l := New(&Config{})
	l.writer = &output

	l.Start()
	defer l.Stop()

	formatLog := func(str string) string {
		return fmt.Sprintf(`\[(\d{4})-(\d{2})-(\d{2}) (\d{2}):(\d{2}):(\d{2})\] - %s`, str)
	}

	test := func(wg *sync.WaitGroup, msg string, msgType int) {
		defer wg.Done()

		t.Helper()

		switch msgType {
		case msgError:
			l.Fatal(fmt.Errorf(msg))
		case msgInfo:
			l.Info(msg)
		}

		time.Sleep(1 * time.Millisecond)
		got := output.String()

		re := regexp.MustCompile(formatLog(msg))

		require.True(t, re.MatchString(got))
	}

	tests := []struct {
		msg     string
		msgType int
	}{
		{
			msg:     "test info",
			msgType: msgInfo,
		},
		{
			msg: "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore " +
				"et dolore magna aliqua.\nUt enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut " +
				"aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse " +
				"cillum dolore eu fugiat nulla pariatur.\nExcepteur sint occaecat cupidatat non proident, sunt in " +
				"culpa qui officia deserunt mollit anim id est laborum.",
			msgType: msgInfo,
		},
		{
			msg:     "test error",
			msgType: msgError,
		},
	}

	var wg sync.WaitGroup

	wg.Add(len(tests))
	for _, tt := range tests {
		tt := tt
		go test(&wg, tt.msg, tt.msgType)
	}

	wg.Wait()
}

package fetch_test

import "testing"

type testLogger struct {
	t *testing.T
}

func (l *testLogger) Fatal(args ...interface{}) {
	l.t.Error(args...)
	panic("Fatal error occurred")
}

func (l *testLogger) Fatalf(format string, args ...interface{}) {
	l.t.Errorf(format, args...)
	panic("Fatal error occurred")
}

func (l *testLogger) Fatalln(args ...interface{}) {
	l.t.Error(args...)
	panic("Fatal error occurred")
}

func (l *testLogger) Print(args ...interface{}) {
	l.t.Log(args...)
}

func (l *testLogger) Printf(format string, args ...interface{}) {
	l.t.Logf(format, args...)
}

func (l *testLogger) Println(args ...interface{}) {
	l.t.Log(args...)
}

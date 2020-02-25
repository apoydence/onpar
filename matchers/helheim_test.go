// This file was generated by github.com/nelsam/hel.  Do not
// edit this code by hand unless you *really* know what you're
// doing.  Expect any changes made manually to be overwritten
// the next time hel regenerates this file.

package matchers_test

type mockDiffer struct {
	DiffCalled chan bool
	DiffInput  struct {
		Actual, Expected chan interface{}
	}
	DiffOutput struct {
		Ret0 chan string
	}
}

func newMockDiffer() *mockDiffer {
	m := &mockDiffer{}
	m.DiffCalled = make(chan bool, 100)
	m.DiffInput.Actual = make(chan interface{}, 100)
	m.DiffInput.Expected = make(chan interface{}, 100)
	m.DiffOutput.Ret0 = make(chan string, 100)
	return m
}
func (m *mockDiffer) Diff(actual, expected interface{}) string {
	m.DiffCalled <- true
	m.DiffInput.Actual <- actual
	m.DiffInput.Expected <- expected
	return <-m.DiffOutput.Ret0
}

type mockMatcher struct {
	MatchCalled chan bool
	MatchInput  struct {
		Actual chan interface{}
	}
	MatchOutput struct {
		ResultValue chan interface{}
		Err         chan error
	}
}

func newMockMatcher() *mockMatcher {
	m := &mockMatcher{}
	m.MatchCalled = make(chan bool, 100)
	m.MatchInput.Actual = make(chan interface{}, 100)
	m.MatchOutput.ResultValue = make(chan interface{}, 100)
	m.MatchOutput.Err = make(chan error, 100)
	return m
}
func (m *mockMatcher) Match(actual interface{}) (resultValue interface{}, err error) {
	m.MatchCalled <- true
	m.MatchInput.Actual <- actual
	return <-m.MatchOutput.ResultValue, <-m.MatchOutput.Err
}

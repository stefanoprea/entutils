package matcher

import (
  "fmt"
  "errors"
)

type errorMatcher struct {
  err error
}

func errorMatcherSprintf(message string, args ...interface{}) matcher {
  return errorMatcher{
    err: errors.New(
      fmt.Sprintf(message, args...),
    ),
  }
}

func (self errorMatcher) Match(actual interface{}) (bool, error) {
  return false, self.err
}

func (self errorMatcher) FailureMessage(actual interface{}) string {
  return self.err.Error()
}

func (self errorMatcher) NegatedFailureMessage(actual interface{}) string {
  return self.err.Error()
}

func (self errorMatcher) Except(_ ...string) matcher {
  return self
}

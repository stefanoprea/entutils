package matcher

import (
  "github.com/stefanoprea/entutils/field"
  "fmt"
  "github.com/onsi/gomega/types"
)

type matcher interface {
  Match(interface{}) (bool, error)
  FailureMessage(interface{}) string
  NegatedFailureMessage(interface{}) string
  Except(...string) matcher
}

var _ types.GomegaMatcher = matcher(nil)

type matcherImpl struct {
  expected []field.WrappedField
}

func MatchPublicFieldsOf(expected interface{}) matcher {
  return matcherImpl{field.IterPublicFields(expected)}
}

func (self matcherImpl) Match(actual interface{}) (bool, error) {
  for _, expectedField := range self.expected {
    fieldName := expectedField.GetName()
    actualField, found := field.GetFieldWithName(actual, fieldName)
    if !found || !expectedField.Equals(actualField) {
      return false, nil
    }
  }
  return true, nil
}

func (self matcherImpl) FailureMessage(actual interface{}) string {
  return fmt.Sprintf(
    "Expected record %v to equal %v",
    self.expected,
    actual,
  )
}

func (self matcherImpl) NegatedFailureMessage(actual interface{}) string {
  return fmt.Sprintf(
    "Expected record %v to not equal %v",
    self.expected,
    actual,
  )
}

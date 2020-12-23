package matcher

import (
  . "github.com/onsi/gomega"
  "testing"
)

func Test_MatchPublicFieldsOf(t *testing.T) {
  g := NewWithT(t)
  expected, actual := makeConfigsWithSameUserAndURL()
  g.Expect(actual).To(MatchPublicFieldsOf(expected))
}

func Test_Negative_MatchPublicFieldsOf(t *testing.T) {
  g := NewWithT(t)
  expected, actual := makeConfigsWithSameURL()
  g.Expect(actual).To(Not(MatchPublicFieldsOf(expected)))
}

func Test_Except(t *testing.T) {
  g := NewWithT(t)
  expected, actual := makeConfigsWithSameURL()
  g.Expect(actual).To(MatchPublicFieldsOf(expected).Except("URL", "User"))
}

func Test_Negative_Except(t *testing.T) {
  g := NewWithT(t)
  expected, actual := makeConfigsWithSameURL()
  g.Expect(actual).To(Not(MatchPublicFieldsOf(expected).Except("URL")))
}

func Test_Except_bad_fields(t *testing.T) {
  g := NewWithT(t)
  expected, actual := makeConfigsWithSameURL()
  _, err := MatchPublicFieldsOf(actual).Except("homepage").Match(expected)
  g.Expect(err).To(Not(BeNil()))
}

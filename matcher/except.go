package matcher

import (
  "github.com/stefanoprea/entutils/field"
  "github.com/stefanoprea/entutils/stringset"
)

func (self matcherImpl) Except(exceptedFields ...string) matcher {
  exceptedFieldsSet := stringset.FromSlice(exceptedFields)
  dest := []field.WrappedField{}
  for _, field := range self.expected {
    fieldName := field.GetName()
    if !exceptedFieldsSet.Contains(fieldName) {
      dest = append(dest, field)
    } else {
      exceptedFieldsSet.Remove(fieldName)
    }
  }
  if !exceptedFieldsSet.IsEmpty() {
    return errorMatcherSprintf(
      "Except() received bad field names: %v",
      exceptedFieldsSet.ToString(),
      )
  }
  return matcherImpl{dest}
}

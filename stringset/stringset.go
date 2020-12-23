package stringset

import (
  "fmt"
  "strings"
)

type set map[string]bool

func New() set {
  return set{}
}

func FromSlice(slice []string) set {
  self := set{}
  for _, elem := range slice {
    self[elem] = true
  }
  return self
}

func (self set) Contains(elem string) bool {
  _, ok := self[elem]
  return ok
}

func (self set) ToString() string {
  return fmt.Sprintf(
    "stringset.set(%v)",
    strings.Join(self.ToSlice(), ", "),
    )
}

func (self set) ToSlice() []string {
  slice := []string{}
  for elem, _ := range self {
    slice = append(slice, elem)
  }
  return slice
}

func (self set) Len() int {
  return len(self)
}

func (self set) IsEmpty() bool {
  return self.Len() == 0
}

func (self set) Add(elem string) {
  self[elem] = true
}

func (self set) Remove(elem string) {
  delete(self, elem)
}

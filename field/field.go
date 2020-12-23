package field

import (
  "reflect"
)

type WrappedField struct {
	reflect.StructField
  reflect.Value
}

func (self *WrappedField)Equals(other WrappedField) bool {
  return reflect.DeepEqual(self.StructField, other.StructField) &&
    reflect.DeepEqual(self.Value.Interface(), other.Value.Interface())
}

func (self *WrappedField)GetName() string {
  return self.StructField.Name
}

func (self *WrappedField)IsPublic() bool {
  return self.StructField.PkgPath == ""
}

func (self *WrappedField) ApplyToModel(model interface{}) {
	setter := self.setterInterfaceFromModel(model)
	self.callSetter(setter)
}

func (self *WrappedField) setterInterfaceFromModel(model interface{}) reflect.Value {
	reflectedModel := reflect.ValueOf(model)
	setterName := self.setterName()
	method := reflectedModel.MethodByName(setterName)
	return method
}

func (self *WrappedField) setterName() string {
	return "Set" + self.GetName()
}

func (self *WrappedField) callSetter(setter reflect.Value) {
	arg := []reflect.Value{self.Value}
	setter.Call(arg)
}

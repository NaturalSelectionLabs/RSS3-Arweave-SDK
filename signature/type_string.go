// Code generated by "enumer --values --type=Type --linecomment --output type_string.go"; DO NOT EDIT.

package signature

import (
	"fmt"
	"strings"
)

const _TypeName = "arweaveed25519ethereumsolanaaptosmulti_aptostyped_ethereum"

var _TypeIndex = [...]uint8{0, 7, 14, 22, 28, 33, 44, 58}

const _TypeLowerName = "arweaveed25519ethereumsolanaaptosmulti_aptostyped_ethereum"

func (i Type) String() string {
	i -= 1
	if i >= Type(len(_TypeIndex)-1) {
		return fmt.Sprintf("Type(%d)", i+1)
	}
	return _TypeName[_TypeIndex[i]:_TypeIndex[i+1]]
}

func (Type) Values() []string {
	return TypeStrings()
}

// An "invalid array index" compiler error signifies that the constant values have changed.
// Re-run the stringer command to generate them again.
func _TypeNoOp() {
	var x [1]struct{}
	_ = x[TypeArweave-(1)]
	_ = x[TypeED25519-(2)]
	_ = x[TypeEthereum-(3)]
	_ = x[TypeSolana-(4)]
	_ = x[TypeAptos-(5)]
	_ = x[TypeMultiAptos-(6)]
	_ = x[TypeTypedEthereum-(7)]
}

var _TypeValues = []Type{TypeArweave, TypeED25519, TypeEthereum, TypeSolana, TypeAptos, TypeMultiAptos, TypeTypedEthereum}

var _TypeNameToValueMap = map[string]Type{
	_TypeName[0:7]:        TypeArweave,
	_TypeLowerName[0:7]:   TypeArweave,
	_TypeName[7:14]:       TypeED25519,
	_TypeLowerName[7:14]:  TypeED25519,
	_TypeName[14:22]:      TypeEthereum,
	_TypeLowerName[14:22]: TypeEthereum,
	_TypeName[22:28]:      TypeSolana,
	_TypeLowerName[22:28]: TypeSolana,
	_TypeName[28:33]:      TypeAptos,
	_TypeLowerName[28:33]: TypeAptos,
	_TypeName[33:44]:      TypeMultiAptos,
	_TypeLowerName[33:44]: TypeMultiAptos,
	_TypeName[44:58]:      TypeTypedEthereum,
	_TypeLowerName[44:58]: TypeTypedEthereum,
}

var _TypeNames = []string{
	_TypeName[0:7],
	_TypeName[7:14],
	_TypeName[14:22],
	_TypeName[22:28],
	_TypeName[28:33],
	_TypeName[33:44],
	_TypeName[44:58],
}

// TypeString retrieves an enum value from the enum constants string name.
// Throws an error if the param is not part of the enum.
func TypeString(s string) (Type, error) {
	if val, ok := _TypeNameToValueMap[s]; ok {
		return val, nil
	}

	if val, ok := _TypeNameToValueMap[strings.ToLower(s)]; ok {
		return val, nil
	}
	return 0, fmt.Errorf("%s does not belong to Type values", s)
}

// TypeValues returns all values of the enum
func TypeValues() []Type {
	return _TypeValues
}

// TypeStrings returns a slice of all String values of the enum
func TypeStrings() []string {
	strs := make([]string, len(_TypeNames))
	copy(strs, _TypeNames)
	return strs
}

// IsAType returns "true" if the value is listed in the enum definition. "false" otherwise
func (i Type) IsAType() bool {
	for _, v := range _TypeValues {
		if i == v {
			return true
		}
	}
	return false
}

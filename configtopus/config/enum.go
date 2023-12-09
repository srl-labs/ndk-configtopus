/*
Package config is a generated package which contains definitions
of structs which represent a YANG schema. The generated schema can be
compressed by a series of transformations (compression was false
in this case).

This package was generated by /go/pkg/mod/github.com/openconfig/ygot@v0.29.16/genutil/names.go
using the following YANG input files:
  - yang/app.yang

Imported modules were sourced from:
*/
package config

import (
	"github.com/openconfig/ygot/ygot"
)

// E_App_Configtopus_ActionLeafNode is a derived int64 type which is used to represent
// the enumerated node App_Configtopus_ActionLeafNode. An additional value named
// App_Configtopus_ActionLeafNode_UNSET is added to the enumeration which is used as
// the nil value, indicating that the enumeration was not explicitly set by
// the program importing the generated structures.
type E_App_Configtopus_ActionLeafNode int64

// IsYANGGoEnum ensures that App_Configtopus_ActionLeafNode implements the yang.GoEnum
// interface. This ensures that App_Configtopus_ActionLeafNode can be identified as a
// mapped type for a YANG enumeration.
func (E_App_Configtopus_ActionLeafNode) IsYANGGoEnum() {}

// ΛMap returns the value lookup map associated with  App_Configtopus_ActionLeafNode.
func (E_App_Configtopus_ActionLeafNode) ΛMap() map[string]map[int64]ygot.EnumDefinition {
	return ΛEnum
}

// String returns a logging-friendly string for E_App_Configtopus_ActionLeafNode.
func (e E_App_Configtopus_ActionLeafNode) String() string {
	return ygot.EnumLogString(e, int64(e), "E_App_Configtopus_ActionLeafNode")
}

const (
	// App_Configtopus_ActionLeafNode_UNSET corresponds to the value UNSET of App_Configtopus_ActionLeafNode.
	App_Configtopus_ActionLeafNode_UNSET E_App_Configtopus_ActionLeafNode = 0
	// App_Configtopus_ActionLeafNode_disable corresponds to the value disable of App_Configtopus_ActionLeafNode.
	App_Configtopus_ActionLeafNode_disable E_App_Configtopus_ActionLeafNode = 1
	// App_Configtopus_ActionLeafNode_enable corresponds to the value enable of App_Configtopus_ActionLeafNode.
	App_Configtopus_ActionLeafNode_enable E_App_Configtopus_ActionLeafNode = 2
)
package lint

import (
	"regexp"
	"strings"
)

// RuleInfo stores information of a rule.
type RuleInfo struct {
	Name        RuleName   // rule name in the set.
	Description string     // a short description of this rule.
	URI         string     // a link to a document for more details.
	FileTypes   []FileType // types of files that this rule targets to.
	Category    Category   // category of problems this rule produces.

	noPositional struct{} // Prevent positional composite literal instantiation
}

// FileType defines a file type that a rule is targeting to.
type FileType string

const (
	// ProtoFile indicates that the targeted file is a protobuf-definition file.
	ProtoFile FileType = "proto-file"
)

// Category defines the category of the findings produced by a rule.
type Category string

const (
	// Error indicates that in the API, something will cause errors.
	Error Category = "error"
	// Warning indicates that in the API, something can be do better.
	Warning Category = "warning"
	// DefaultCategory denotes the default value of Category
	DefaultCategory Category = Warning
)

// Status defines whether a rule is enabled, disabled or deprecated.
type Status string

const (
	// Enabled indicates that a rule should be enabled.
	Enabled Status = "enabled"
	// Disabled indicates that a rule should be disabled.
	Disabled Status = "disabled"
	// Deprecated indicates that a rule should be deprecated.
	Deprecated Status = "Deprecated"
	// DefaultStatus denotes the default value of Status.
	DefaultStatus Status = Disabled
)

// RuleName is an identifier for a rule. Allowed characters include a-z, A-Z, 0-9, -, _. The
// namespace separator :: is allowed between RuleName segments (for example, my_namespace::my_rule).
type RuleName string

const nameSeparator string = "::"

var ruleNameValidator = regexp.MustCompile("^([a-zA-Z0-9-_]+(::)?)+[a-zA-Z0-9-_]+$")

func NewRuleName(segments ...string) RuleName {
	return RuleName(strings.Join(segments, nameSeparator))
}

func (r RuleName) IsValid() bool {
	return r != "" && ruleNameValidator.Match([]byte(r))
}

func (r RuleName) WithPrefix(prefix RuleName) RuleName {
	return RuleName(string(prefix) + nameSeparator + string(r))
}

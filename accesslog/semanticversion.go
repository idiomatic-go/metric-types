package accesslog

// Envoy uses SemVer (https://semver.org/). Major/minor versions indicate
// expected behaviors and APIs, the patch version field is used only
// for security fixes and can be generally ignored.
type SemanticVersion struct {
	MajorNumber uint32
	MinorNumber uint32
	Patch       uint32
}

// Code generated by "stringer -type=ValidationLevelLegacy"; DO NOT EDIT.

package carbon20

import "strconv"

const _ValidationLevelLegacy_name = "StrictLegacyMediumLegacyNoneLegacy"

var _ValidationLevelLegacy_index = [...]uint8{0, 12, 24, 34}

func (i ValidationLevelLegacy) String() string {
	if i < 0 || i >= ValidationLevelLegacy(len(_ValidationLevelLegacy_index)-1) {
		return "ValidationLevelLegacy(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _ValidationLevelLegacy_name[_ValidationLevelLegacy_index[i]:_ValidationLevelLegacy_index[i+1]]
}

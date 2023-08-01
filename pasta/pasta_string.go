// Code generated by "stringer -type=Pasta"; DO NOT EDIT.

package pasta

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[Elbow-1]
	_ = x[Macaroni-2]
	_ = x[SewerPipe-3]
	_ = x[Spagatti-4]
	_ = x[Rigatoni-5]
}

const _Pasta_name = "ElbowMacaroniSewerPipeSpagattiRigatoni"

var _Pasta_index = [...]uint8{0, 5, 13, 22, 30, 38}

func (i Pasta) String() string {
	i -= 1
	if i < 0 || i >= Pasta(len(_Pasta_index)-1) {
		return "Pasta(" + strconv.FormatInt(int64(i+1), 10) + ")"
	}
	return _Pasta_name[_Pasta_index[i]:_Pasta_index[i+1]]
}
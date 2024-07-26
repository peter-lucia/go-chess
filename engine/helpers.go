package engine

import "errors"

func ConvertUICoordsToEngineCoords(uiPiece string) (int, int, error) {
	if len(uiPiece) != 2 {
		return 0, 0, errors.New("uninterpretable ui coords")
	}
	// cols: a,b,c,d,e,f,g,h
	col := int(uiPiece[0] - 'a')
	// rows: 1,2,3,4,5,6,7,8
	row := int(uiPiece[1] - '1')
	return row, col, nil

}

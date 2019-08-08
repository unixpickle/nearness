package nearness

import (
	"math"
	"strings"

	"github.com/pkg/errors"
)

func ParseBoards(s string) ([]*Board, error) {
	var res []*Board
	for _, substr := range strings.Split(s, ";") {
		b, err := ParseBoard(substr)
		if err != nil {
			return nil, errors.Wrap(err, "parse boards")
		}
		res = append(res, b)
	}
	return res, nil
}

func ParseBoard(s string) (*Board, error) {
	replacer := strings.NewReplacer("(", "", ")", "", " ", "", "\n", "", ",", " ")
	s = replacer.Replace(s)
	fields := strings.Fields(s)
	size := int(math.Sqrt(float64(len(fields))))
	if size*size != len(fields) {
		return nil, errors.New("parse board: invalid number of fields")
	}
	result := &Board{
		Size:      size,
		Positions: make([]Position, len(fields)),
	}
	for i, field := range fields {
		pos, err := parsePosition(field)
		if err != nil {
			return nil, err
		}
		result.Positions[i] = *pos
	}
	return result, nil
}

func parsePosition(pos string) (*Position, error) {
	if len(pos) != 2 {
		return nil, errors.New("parse board: invalid position: " + pos)
	}
	col, err := parseCoordinate(pos[0])
	if err != nil {
		return nil, err
	}
	row, err := parseCoordinate(pos[1])
	if err != nil {
		return nil, err
	}
	return &Position{Row: row, Col: col}, nil
}

func parseCoordinate(coord byte) (int, error) {
	if coord >= 'A' && coord <= 'Z' {
		return int(coord - 'A'), nil
	} else if coord >= '1' && coord <= '4' {
		return int(coord-'1') + 26, nil
	}
	return 0, errors.New("parse board: invalid coordinate: " + string(coord))
}

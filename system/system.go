package model

import "errors"

type System struct {
	slot [3]Slot
}

func (s *System) GetSlot(index int) (Slot, error) {
	if index < 0 || index >= len(s.slot) {
		return Slot{}, errors.New("index out of range")
	}

	return s.slot[index], nil
}

func (s *System) AvailableSlots() int {
	return len(s.slot)
}

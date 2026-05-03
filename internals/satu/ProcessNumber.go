package satu

import (
	"errors"
	"fmt"
)

func ProcessNumber(nums []int) error {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recover from panic", r)
		}
	}()
	if nums == nil {
		err := errors.New("no data provided")
		return err
	}
	if len(nums) == 0 {
		err := errors.New("empty list provided")
		return err
	}
	for _, num := range nums {
		res := num * 2
		fmt.Printf("%d x 2 = %d\n", num, res)
	}
	return nil
}

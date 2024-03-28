package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"

	"github.com/gofrs/uuid"
	"golang.design/x/clipboard"
)

func GetAmountToGenerate() (int, error) {
	var amount int
	fmt.Println("How many UUIDs would you like to generate?")
	fmt.Scanln(&amount)
	if amount <= 1 {
		return -1, errors.New("invalid amount of UUIDs")
	}
	return amount, nil
}

func GenerateUUIDs(amount int) []string {
	uuids := make([]string, amount)
	for i := 0; i < amount; i++ {
		uuid, _ := uuid.NewV4()
		uuids[i] = uuid.String()
	}
	return uuids
}

func PromptForUUIDSelection(uuids []string) (int, error) {
	var selectID int
	fmt.Println("Please select a UUID to use:")
	var selectionNum int
	_, err := fmt.Scanln(&selectionNum)
	if err != nil {
		return -1, errors.New("invalid selection")
	} else if selectionNum < 1 || selectionNum > len(uuids) {
		return PromptForUUIDSelection(uuids)
	}
	selectID = selectionNum
	return selectID, nil
}

func ChooseUUID(selectID int, uuids []string) (string, error) {
	if selectID < 0 || selectID >= len(uuids) {
		return "", fmt.Errorf("invalid UUID ID")
	} else if len(uuids) == 1 {
		return uuids[0], nil
	}
	return uuids[selectID-1], nil
}

func PrintUUIDs(uuids []string) {
	for i, uuid := range uuids {
		fmt.Printf("%d: %s\n", i+1, uuid)
	}
}

func CopyUUIDToClipboard(uuid string) error {
	err := clipboard.Init()
	if err != nil {
		return err
	}
	cliperr := clipboard.Write(clipboard.FmtText, []byte(uuid))
	if cliperr != nil {
		return err
	}
	return nil
}

func entry(gen *bool) {
	amount, amountgenerror := GetAmountToGenerate()
	if amountgenerror != nil {
		fmt.Println(amountgenerror)
		return
	}
	uuids := GenerateUUIDs(amount)
	PrintUUIDs(uuids)
	selectID, err := PromptForUUIDSelection(uuids)
	if err != nil {
		fmt.Println(err)
		return
	}
	CopyUUIDToClipboard(uuids[selectID-1])
	fmt.Println(uuids[selectID-1])
	fmt.Println("UUID copied to clipboard")
	fmt.Println("Press enter to exit or y to generate more UUIDs")
	reader := bufio.NewReader(os.Stdin)
	input, _, errorbufio := reader.ReadRune()
	if errorbufio != nil {
		fmt.Println(err)
		return
	}
	if input == 'y' {
		return
	} else {
		*gen = false
		return
	}
}

func main() {
	gen := true
	for gen {
		entry(&gen)
	}
}

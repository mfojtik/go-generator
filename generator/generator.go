package generator

import (
	"fmt"
	"strings"

	"github.com/mfojtik/go-generator/types"
)

func Template(s string) (string, error) {
	switch s {
	case "password":
		return SimplePassword(8)
	case "uuid":
		return RandomUUID(), nil
	default:
		return types.FromTemplate(s)
	}
}

func StrongPassword(length int) (string, error) {
	return types.FromTemplate(fmt.Sprintf("[\\w]{%d}", length))
}

func SimplePassword(length int) (string, error) {
	return types.FromTemplate(fmt.Sprintf("[\\a]{%d}", length))
}

func RandomUUID() string {
	uuid, _ := types.FromTemplate("[\\a]{8}-[\\a]{4}-4[\\a]{3}-8[\\a]{3}-[\\a]{12}")
	return strings.ToLower(uuid)
}

func RandomCleanUUID() string {
	uuid := RandomUUID()
	return strings.Replace(uuid, "-", "", -1)
}

func RandomNumericUUID() string {
	uuid, _ := types.FromTemplate("[\\d]{8}-[\\d]{4}-4[\\d]{3}-8[\\d]{3}-[\\d]{12}")
	return uuid
}

func RandomCleanNumbericUUID() string {
	uuid := RandomNumericUUID()
	return strings.Replace(uuid, "-", "", -1)
}

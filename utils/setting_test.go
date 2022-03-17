package utils

import (
	"fmt"
	"testing"
)

func TestSetting(t *testing.T) {
	fmt.Println(Db, DbUser, DbHost, DbPort, DbPassWord)
}

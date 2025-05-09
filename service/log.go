package service

import (
	"fmt"
	"github.com/lontten/ldb"
	"github.com/ltask/global"
	"github.com/ltask/model"
	"strings"
)

func SaveLog(log ...string) {
	var s = strings.Join(log, "\n")
	t, err := ldb.First[model.Task](global.DB, ldb.W().PrimaryKey(1))
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(t)
	fmt.Println(s)
}

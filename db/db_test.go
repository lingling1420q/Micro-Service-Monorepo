package db

import (
	config "monaco/config"
	"testing"
)

func TestInitCfg(t *testing.T) {
	dbCfg := config.Config.DB
	mysqlCfg := "Username:Password@tcp(Host:Port)/DbName?charset=utf8"
	InitCfg(&mysqlCfg, dbCfg.Mysql)
	t.Log(mysqlCfg)
}

func TestDB(t *testing.T) {
	connError := DB().GetErrors()
	if len(connError) != 0 {
		t.Error(connError)
	}
}

func TestInsertUser(t *testing.T) {
	user := TBL_USERS{
		Name:  "UnitTestName",
		Pwd:   "UnitTestPwd",
		Phone: "17199911145",
		Email: "UnitTestEmail",
	}
	InsertColume(&user)
}

func TestFindUser(t *testing.T) {
	user := TBL_USERS{
		Name:  "UnitTestName",
		Pwd:   "UnitTestPwd",
		Phone: "17199911145",
		Email: "UnitTestEmail",
	}
	r := FindColume(&user)
	var u TBL_USERS
	r.Scan(&u)
	t.Log(u)
}

func TestDeleteUser(t *testing.T) {
	user := TBL_USERS{
		Name:  "UnitTestName",
		Pwd:   "UnitTestPwd",
		Phone: "17199911145",
		Email: "UnitTestEmail",
	}
	Deletecolume(&user)
}

func TestInsertLog(t *testing.T) {
	vl := TBL_VISIT_LOG{
		Host:    "UnitTestHostName",
		Method:  "UnitTestMethod",
		URL:     "UnitTestURL",
		Header:  "UnitTestHeader",
		Query:   "UnitTestQuery",
		JSONRaw: "UnitTestJSONData",
	}
	InsertColume(&vl)
}

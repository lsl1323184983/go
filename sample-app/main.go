package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"time"
)

type user struct {
	name     string
	password string
}

type userInfo struct {
	Name  string
	Age   int `json:"age"`
	Hobby []string
}

func (u user) checkPassword(password string) bool {
	return u.password == password
}

func findUser(users []user, name string) (v *user, err error) {
	for _, u := range users {
		if u.name == name {
			return &u, nil
		}
	}
	return nil, errors.New("not found")
}

func add(a int, b int) int {
	return a + b
}

func main() {
	res := add(1, 2)
	fmt.Println(res)

	u, err := findUser([]user{{"li", "1024"}}, "li")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(u.name)

	p := u.checkPassword("1024")
	fmt.Println(p)

	a := userInfo{Name: "lsl", Age: 25, Hobby: []string{"Golang", "java"}}
	buf, err := json.Marshal(a)
	if err != nil {
		panic(err)
	}
	fmt.Println(buf)
	fmt.Println(string(buf))

	var b userInfo
	err = json.Unmarshal(buf, &b)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%#v\n", b)

	now := time.Now()
	fmt.Println(now)

	t := time.Date(2023, 5, 12, 12, 0, 0, 0, time.UTC)
	t2 := time.Date(2023, 5, 12, 14, 0, 0, 0, time.UTC)
	fmt.Println(t)
	fmt.Println(t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second())
	fmt.Println(t.Format("2006-01-02 15:04:05"))
	diff := t2.Sub(t)
	fmt.Println(diff)
	fmt.Println(diff.Minutes(), diff.Seconds())
	t3, err := time.Parse("2006-01-02 15:04:05", "2023-05-12 00:00:00")
	if err != nil {
		panic(err)
	}
	fmt.Println(t3 == t)
	fmt.Println(now.Unix())

	f, _ := strconv.ParseFloat("1.234", 64)
	fmt.Println(f)

	fmt.Println(os.Args)

	fmt.Println(os.Getenv("PATH"))
	os.Setenv("AA", "BB")

	buf, err = exec.Command("java", "-version").CombinedOutput()

	if err != nil {
		panic(err)
	}
	fmt.Println(string(buf))
}

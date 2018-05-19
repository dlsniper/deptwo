package deptwo

import (
	"fmt"

	"github.com/sirupsen/logrus"
)

type MyDepTwo string

func (m MyDepTwo) Hello() {
	fmt.Printf("Hello from MyDepTwo, %s", m)
	logrus.Println("called: MyDepTwo.Hello()")
}

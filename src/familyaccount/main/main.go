package main

import (
	"GoStudy/src/familyaccount/utils"
	"fmt"
)

func main() {
	var na string
	var pw string
	name := utils.NewAccount().GetName()
	pwd := utils.NewAccount().GetPwd()
	fmt.Println("请输入正确的账户名:")
	fmt.Scanln(&na)
	fmt.Println("请输入正确的密码:")
	fmt.Scanln(&pw)
	if name == na && pwd == pw {
		utils.NewFamilyAccount().MainMenu()
	} else {
		fmt.Println("登录失败，账户或者密码有误！")
	}

}

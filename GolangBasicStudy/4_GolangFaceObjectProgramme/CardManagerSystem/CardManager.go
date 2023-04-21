package main

import (
	"GolangBasicStudy/4_GolangFaceObjectProgramme/CardManagerSystem/CardPackage"
	"fmt"
	"time"
)

/*
	名片管理系统代码:
		系统逻辑:
			1. 显示主界面, 供客户选择不同功能 ==> show()
			2. 新增名片 ==> add_card()
			3. 显示所有名片 ==> show_all_card()
			4. 根据姓名查询指定名片 ==> query_card()
			5. 根据姓名修改名片信息 ==> change_card_message()
			6. 删除指定名片 ==> delete_card()
			7. 退出系统 ==> return
*/

func main() {
	// 无限循环
	for {
		showHomePage()
		fmt.Println("请选择您要办理的业务序号:")
		var choose string
		fmt.Scan(&choose)
		switch choose {
		case "1":
			CardPackage.Card.AddCard(CardPackage.Card{})
		case "2":
			CardPackage.Card.ShowAllCard(CardPackage.Card{})
		case "3":
			CardPackage.Card.QueryCard(CardPackage.Card{})
		case "4":
			CardPackage.Card.ChangeCardMessage(CardPackage.Card{})
		case "5":
			CardPackage.Card.DeleteCard(CardPackage.Card{})
		case "6":
			fmt.Printf("当前系统名片的数量为: %d\n", len(CardPackage.CardList))
		case "7":
			fmt.Println("即将退出系统, 感谢您的使用!!")
			time.Sleep(time.Second * 1)
			return
		}
	}
}

// 定义函数, 展示主界面
func showHomePage() {
	fmt.Println()
	fmt.Println("==============欢迎使用混沌名片管理系统==============")
	fmt.Println("1. 新增名片")
	fmt.Println("2. 显示所有名片")
	fmt.Println("3. 查询指定名片信息")
	fmt.Println("4. 修改名片信息")
	fmt.Println("5. 删除指定名片")
	fmt.Println("6. 显示当前系统名片数量")
	fmt.Println("7. 退出系统")
	fmt.Println("==================================================")
	fmt.Println()
}

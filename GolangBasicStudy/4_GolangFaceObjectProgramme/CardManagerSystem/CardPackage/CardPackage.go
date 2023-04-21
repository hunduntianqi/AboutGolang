package CardPackage

import "fmt"

/*
	定义名片信息包, 编写名片相关属性与方法代码
		名片属性:
			name: str, phone: str, qq: str, email: str
		系统逻辑:
			1. 显示主界面, 供客户选择不同功能 ==> show()
			2. 新增名片 ==> add_card()
			3. 显示所有名片 ==> show_all_card()
			4. 根据姓名查询指定名片 ==> query_card()
			5. 根据姓名修改名片信息 ==> change_card_message()
			6. 删除指定名片 ==> delete_card()
			7. 退出系统 ==> return
*/

// Card 定义名片结构体类
type Card struct {
	name  string
	phone string
	qq    string
	email string
}

// 定义全局变量切片, 存储名片
var CardList = make([]*Card, 0)

// 定义方法, 新增名片
func (Card) AddCard() {
	// 新建名片对象
	fmt.Println("========新增名片对象========")
	var cardName Card = Card{}
	fmt.Println("请输入名片姓名:")
	var name string
	_, err := fmt.Scan(&name)
	if err != nil {
		return
	}
	cardName.name = name
	fmt.Println("请输入名片手机号:")
	var phone string
	_, err = fmt.Scan(&phone)
	if err != nil {
		return
	}
	cardName.phone = phone
	fmt.Println("请输入名片qq:")
	var qq string
	_, err = fmt.Scan(&qq)
	if err != nil {
		return
	}
	cardName.qq = qq
	fmt.Println("请输入名片邮箱:")
	var email string
	_, err = fmt.Scan(&email)
	if err != nil {
		return
	}
	cardName.email = email
	// 将名片添加到切片中
	CardList = append(CardList, &cardName)
	fmt.Printf("名片 %s 已生成!!", name)
}

// 定义方法, 显示所有名片
func (Card) ShowAllCard() {
	fmt.Println("========显示所有名片信息========")
	// 判断名片列表是否为空
	if len(CardList) == 0 {
		fmt.Println("系统中现在没有名片信息, 请先添加名片, 谢谢!!")
		return
	} else {
		fmt.Println("姓名\t手机号\t\tQQ号码\t\t邮箱地址")
		// 遍历列表
		for _, cardName := range CardList {
			fmt.Printf("%s\t%s\t%s\t%s\n", cardName.name, cardName.phone, cardName.qq, cardName.email)
		}
	}
}

// 定义方法, 完成查询指定名片功能
func (Card) QueryCard() {
	fmt.Println("========查询指定名片========")
	// 判断名片列表是否为空
	if len(CardList) == 0 {
		fmt.Println("系统中无名片信息, 请先添加名片, 谢谢!!")
		return
	} else {
		cardName := ""
		fmt.Println("请输入要查询名片的姓名: ")
		_, err := fmt.Scan(&cardName)
		if err != nil {
			return
		}
		// 遍历切片
		for _, card := range CardList {
			if card.name == cardName {
				fmt.Println("你查询的名片信息如下:")
				fmt.Println("姓名\t手机号\t\tQQ号码\t\t邮箱地址")
				fmt.Printf("%s\t%s\t%s\t%s\n", card.name, card.phone, card.qq, card.email)
				return
			}
		}
		// 循环结束, 无满足条件数据
		fmt.Println("系统中没有您要查询的名片信息, 请再检查一下, 谢谢!!")
	}
}

// 定义方法, 完成修改名片信息功能
func (Card) ChangeCardMessage() {
	fmt.Println("========修改指定名片信息========")
	// 判断名片列表是否为空
	if len(CardList) == 0 {
		fmt.Println("系统中无名片信息, 请先添加名片, 谢谢!!")
		return
	} else {
		cardName := ""
		fmt.Println("请输入要修改的名片姓名: ")
		_, err := fmt.Scan(&cardName)
		if err != nil {
			return
		}
		// 遍历列表, 查询名片
		for _, card := range CardList {
			if card.name == cardName {
				fmt.Println("要查找的名片地址为: ", card)
				// 找到对应名片, 修改名片信息
				fmt.Println("请问是否修改名片姓名(Yes / No):")
				var choose string     // 接收用户输入的选择
				var newMessage string // 接受用户修改后的信息
				fmt.Scan(&choose)
				if (choose == "Yes") || (choose == "yes") {
					fmt.Println("请输入修改后的姓名:")
					fmt.Scan(&newMessage)
					card.name = newMessage
				}
				fmt.Println("请问是否修改名片手机号(Yes / No):")
				fmt.Scan(&choose)
				if (choose == "Yes") || (choose == "yes") {
					fmt.Println("请输入修改后名片的手机号:")
					fmt.Scan(&newMessage)
					card.phone = newMessage
				}
				fmt.Println("请问是否修改名片QQ号码(Yes / No):")
				fmt.Scan(&choose)
				if (choose == "Yes") || (choose == "yes") {
					fmt.Println("请输入修改后名片的QQ号码:")
					fmt.Scan(&newMessage)
					card.qq = newMessage
				}
				fmt.Println("请问是否修改名片邮箱地址(Yes / No):")
				fmt.Scan(&choose)
				if (choose == "Yes") || (choose == "yes") {
					fmt.Println("请输入修改后名片的邮箱地址:")
					fmt.Scan(&newMessage)
					card.email = newMessage
				}
				fmt.Println("名片信息已修改成功!!")
				fmt.Println("修改后的名片地址为: ", card)
				return
			}
		}
		// 循环结束, 无满足条件数据
		fmt.Println("系统中没有您要查询的名片信息, 请再检查一下, 谢谢!!")
	}
}

// 定义方法, 完成删除名片操作
func (Card) DeleteCard() {
	fmt.Println("========删除指定名片========")
	// 判断名片列表是否为空
	if len(CardList) == 0 {
		fmt.Println("系统中无名片信息, 请先添加名片, 谢谢!!")
		return
	} else {
		cardName := ""
		fmt.Println("请输入要删除名片的姓名: ")
		_, err := fmt.Scan(&cardName)
		if err != nil {
			return
		}
		// 遍历切片
		for index, card := range CardList {
			if card.name == cardName {
				CardList = append(CardList[:index], CardList[index+1:]...)
				return
			}
		}
		// 循环结束, 无满足条件数据
		fmt.Println("系统中没有您要查询的名片信息, 请再检查一下, 谢谢!!")
	}
}

package main

import "fmt"

type Role struct {
	Name       string
	Profession string
	Attack     int
	Skill      []string
	Money      int
	Goods      []map[string]int
	Wear       map[string]string
}

func main() {
	// 实例化结构体Role
	r := Role{Money: 0, Skill: []string{"普通攻击"}, Attack: 10}
	// 定义变量
	var name, profession string
	// 输出操作提示
	fmt.Printf("欢迎来到西西村！请输入并创建你的名字：\n")
	// 存储用户输入数据
	fmt.Scanln(&name)
	// 设置人物名称
	r.Name = name
	// 设置人物职业和技能
	for {
		// 输出操作提示
		fmt.Printf("请选择你的职业：1-剑士，2-法师，3-弓箭手：\n")
		// 存储用户输入数据
		fmt.Scanln(&profession)
		if profession == "1" {
			r.Profession = "剑士"
			r.Skill = append(r.Skill, "基础剑术")
			break
		} else if profession == "2" {
			r.Profession = "法师"
			r.Skill = append(r.Skill, "基础法术")
			break
		} else if profession == "3" {
			r.Profession = "弓箭手"
			r.Skill = append(r.Skill, "基础箭法")
			break
		} else {
			fmt.Printf("输入有误，请重新选择：\n")
		}
	}
	// 输出游戏角色的基本信息
	fmt.Printf("人物名称：%v\n", r.Name)
	fmt.Printf("人物职业：%v\n", r.Profession)
	fmt.Printf("人物技能：%v\n", r.Skill)
	fmt.Printf("人物装备：%v\n", r.Wear)

	var skill, ways, wear string
	if r.Profession == "剑士" {
		skill = "剑刃冲击"
	} else if r.Profession == "法师" {
		skill = "火球术"
	} else if r.Profession == "弓箭手" {
		skill = "心神凝聚"
	}
	// 与西西村村长进行对话，学习技能
	fmt.Printf("你好，%v，我是西西村村长，最近野兽森林里出现了山贼，捉走了采药的村民。这里有一本%v，希望你能营救村民。\n", name, skill)
	fmt.Printf("系统提示：习得%v\n", skill)
	r.Skill = append(r.Skill, skill)
	// 与药铺老板进行对话，获得金钱和物品
	fmt.Printf("你好，%v，我是药铺老板，我的员工在采药时被山贼捉去，这里有10瓶金疮药和100两银子，希望你能解救我的员工。\n", name)
	r.Goods = append(r.Goods, map[string]int{"金疮药": 10})
	r.Money += 100
	fmt.Printf("系统提示：获得10瓶金疮药和100两银子\n")

	// 让玩家选择路线，用于推进游戏剧情发展
	for {
		// 输出操作提示
		fmt.Printf("选择你的路线：1-兵器铺，2-野兽森林，3-退出\n")
		// 存储用户输入数据
		fmt.Scanln(&ways)
		if ways == "1" {
			fmt.Printf("本店出售兵器，总有一款适合你。\n")
			fmt.Printf("木杖—60两 +18攻击：选择1\n")
			fmt.Printf("木剑—60两 +18攻击：选择2\n")
			fmt.Printf("木弓—60两 +18攻击：选择3\n")
			fmt.Printf("布衣—30两 +10防御：选择4\n")
			// 存储用户输入数据
			fmt.Scanln(&wear)
			// 判断玩家输入数据
			r.Wear = map[string]string{}
			if wear == "1" && r.Profession == "法师" {
				// 设置游戏角色的攻击力、金钱和装备
				r.Attack += 18
				r.Money -= 60
				r.Wear["arms"] = "木杖"
			} else if wear == "2" && r.Profession == "剑士" {
				// 设置游戏角色的攻击力、金钱和装备
				r.Attack += 18
				r.Money -= 60
				r.Wear["arms"] = "木剑"
			} else if wear == "3" && r.Profession == "弓箭手" {
				// 设置游戏角色的攻击力、金钱和装备
				r.Attack += 18
				r.Money -= 60
				r.Wear["arms"] = "木弓"
			} else if wear == "4" {
				// 将衣服的防御力乘以0.2转化为攻击力
				r.Attack += 10 * 0.2
				r.Money -= 30
				r.Wear["clothes"] = "布衣"
			} else {
				fmt.Printf("选择的装备与你的职业不相符\n")
			}
		} else if ways == "2" {
			// 待完善......
		} else if ways == "3" {
			break
		}
	}
	// 输出游戏角色的基本信息
	fmt.Printf("人物名称：%v\n", r.Name)
	fmt.Printf("人物职业：%v\n", r.Profession)
	fmt.Printf("人物技能：%v\n", r.Skill)
	fmt.Printf("人物攻击力：%v\n", r.Attack)
	fmt.Printf("人物金钱：%v\n", r.Money)
	fmt.Printf("人物装备：%v\n", r.Wear)
	fmt.Printf("人物物品：%v\n", r.Goods)
}

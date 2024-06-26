package main

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"time"
)

func main() {
	var ctx = context.Background()
	// 连接Redis数据库
	client := redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	// 测试连接Redis
	ping, _ := client.Ping(ctx).Result()
	fmt.Printf("测试连接Redis：%v\n", ping)
	defer client.Close()

	// 设置字符串类型的数据
	// 参数ctx是内置包context创建上下文对象
	// 参数key和value是键值对，数据类型为字符串
	// 参数expiration是有效期，数据类型为time.Duration
	strSet, _ := client.Set(ctx, "name", "Tom", time.Hour).Result()
	fmt.Printf("设置字符串类型的数据：%v\n", strSet)
	// 获取字符串类型的数据
	strGet, _ := client.Get(ctx, "name").Result()
	fmt.Printf("获取字符串类型的数据：%v\n", strGet)
	// 删除字符串类型的数据
	// 参数keys是不固定参数，参数类型为字符串，代表字符串值
	strDel, _ := client.Del(ctx, "name").Result()
	fmt.Printf("删除字符串类型的数据：%v\n", strDel)

	// 设置哈希类型的数据
	// 参数ctx是内置包context创建上下文对象
	// 参数key是键，数据类型为字符串类型
	// 参数values是不固定参数，参数类型为空接口，代表哈希数值
	hashHset, _ := client.HSet(ctx, "Tom", "age", 10).Result()
	fmt.Printf("设置哈希类型的数据：%v\n", hashHset)
	// 获取哈希类型的数据
	// 参数field是值，数据类型为字符串类型
	hashHGet, _ := client.HGet(ctx, "Tom", "age").Result()
	fmt.Printf("获取哈希类型的数据：%v\n", hashHGet)
	// 删除哈希类型的数据
	// 参数fields是不固定参数，数据类型为字符串类型，代表哈希数值
	hashHDel, _ := client.HDel(ctx, "Tom", "age").Result()
	fmt.Printf("删除哈希类型的数据：%v\n", hashHDel)

	// 在列表中添加一个或多个值
	// 参数ctx是内置包context创建上下文对象
	// 参数key是键，数据类型为字符串类型
	// 参数values是不固定参数，参数类型为空接口，代表列表元素
	litRPush, _ := client.RPush(ctx, "Tom", "English", "Chinese").Result()
	fmt.Printf("在列表中添加一个或多个值：%v\n", litRPush)
	// 获取列表指定范围内的元素
	// 参数start和stop是列表索引，数据类型为整型
	litLRange, _ := client.LRange(ctx, "Tom", 0, 2).Result()
	fmt.Printf("获取列表指定范围内的元素：%v\n", litLRange)
	// 移出并获取列表的第一个元素
	// 参数timeout设置超时，数据类型为time.Duration
	// 参数keys是不固定参数，参数类型为字符串，代表列表元素
	litBLPop, _ := client.BLPop(ctx, time.Second, "Tom").Result()
	fmt.Printf("移出并获取列表的第一个元素：%v\n", litBLPop)

	// 向集合添加一个或多个成员
	// 参数ctx是内置包context创建上下文对象
	// 参数key是键，数据类型为字符串类型
	// 参数members是不固定参数，参数类型为空接口，代表集合成员值
	SetSadd, _ := client.SAdd(ctx, "Tim", 20, "170CM").Result()
	fmt.Printf("向集合添加一个或多个成员：%v\n", SetSadd)
	// 获取集合中的所有成员
	SetSMembers, _ := client.SMembers(ctx, "Tim").Result()
	fmt.Printf("向集合添加一个或多个成员：%v\n", SetSMembers)
	// 移除并返回集合中的一个随机元素
	SetSPop, _ := client.SPop(ctx, "Tim").Result()
	fmt.Printf("移除并返回集合中的一个随机元素：%v\n", SetSPop)

	// 有序集合添加或更新一个或多个成员和成员的分数
	// 参数ctx是内置包context创建上下文对象
	// 参数key是键，数据类型为字符串类型
	// 参数members是不固定参数，数据类型是结构体Z实例化对象，包含集合成员和分数
	z1 := redis.Z{Member: "170CM", Score: 5}
	z2 := redis.Z{Member: 10, Score: 10}
	ZsetZAdd, _ := client.ZAdd(ctx, "Tim", &z1, &z2).Result()
	fmt.Printf("移除并返回集合中的一个随机元素：%v\n", ZsetZAdd)
	// 通过索引区间返回有序集合指定区间内的成员
	// 参数start和stop是有序集合索引区间，数据类型为整型
	ZsetZRange, _ := client.ZRange(ctx, "Tim", 0, 2).Result()
	fmt.Printf("通过索引区间返回有序集合指定区间内的成员：%v\n", ZsetZRange)
	// 移除有序集合中的一个或多个成员
	ZsetZRem, _ := client.ZRem(ctx, "Tim", z1).Result()
	fmt.Printf("移除有序集合中的一个或多个成员：%v\n", ZsetZRem)

	// 新增流类型数据
	// 参数ctx是内置包context创建上下文对象
	// 参数XAddArgs是指针类型的结构体XAddArgs实例化对象，代表流类型的数据结构
	// 实例化结构体XAddArgs只需设置成员Stream和Values
	x1 := redis.XAddArgs{
		Stream: "Lily",
		Values: map[string]interface{}{"age": 10, "height": "160CM"},
	}
	// 结构体XAddArgs实例化对象以指针方式作为参数
	streXAdd, _ := client.XAdd(ctx, &x1).Result()
	fmt.Printf("新增流类型数据：%v\n", streXAdd)
	// 获取流类型所有数据
	// 参数stream代表流数据名称，即结构体XAddArgs的成员Stream
	// 参数start和stop是最小值和最大值，以“-”和“+”表示
	streXRange, _ := client.XRange(ctx, "Lily", "-", "+").Result()
	fmt.Printf("获取流类型所有数据：%v\n", streXRange)
	// 遍历变量streXRange，遍历对象为结构体XMessage，结构体成员ID是流数据ID
	for _, v := range streXRange {
		fmt.Printf("获取流类型所有数据的ID：%v\n", v.ID)
		// 通过流数据ID删除数据
		streXDel, _ := client.XDel(ctx, "Lily", v.ID).Result()
		fmt.Printf("ID：%v已删除，数据量：%v\n", v.ID, streXDel)
	}

	// 新增字符串类型数据
	client.Set(ctx, "Tim", "ABCDEFGHIJKLMN", 0)
	// 将字符串类型数据转为二进制数据，然后修改二级制的位数
	// 参数key代表redis的键
	// 参数offset是二级制的位数偏移量，0代表从左边第一位
	// 参数value只有0和1，因为二级制只有0和1
	bitSetBit, _ := client.SetBit(ctx, "Tim", 0, 1).Result()
	fmt.Printf("位图类型数据：%v\n", bitSetBit)
	// 获取位图类型数据某个偏移量的值
	bitGetBit, _ := client.GetBit(ctx, "Tim", 0).Result()
	fmt.Printf("获取位图类型数据某个偏移量的值：%v\n", bitGetBit)
	// 删除位图数据，即删除字符串数据
	bitDel, _ := client.Del(ctx, "Tim").Result()
	fmt.Printf("删除位图数据：%v\n", bitDel)
}

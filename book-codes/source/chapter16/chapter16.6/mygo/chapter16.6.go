package main

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"time"
)

func main() {
	uri := "mongodb://localhost:27017"
	client, err := mongo.NewClient(options.Client().ApplyURI(uri))
	if err != nil {
		fmt.Printf("连接对象创建失败：%v\n", err)
	}
	// 定义上下文对象ctx，它来自内置包context，这是管理上下文
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	// 使用连接对象连接数据库
	err = client.Connect(ctx)
	if err != nil {
		fmt.Printf("数据库连接失败：%v\n", err)
	}
	// 关闭连接
	defer client.Disconnect(ctx)
	// 通过ping测试是否连接成功
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		fmt.Printf("ping测试是否连接成功：%v\n", err)
	}
	// 获取当前已有的数据库
	databases, err := client.ListDatabaseNames(ctx, bson.M{})
	if err != nil {
		fmt.Printf("当前已有数据库获取失败：%v\n", err)
	}
	fmt.Printf("获取当前已有的数据库：%v\n", databases)

	// 创建数据库DB
	DBDatabase := client.Database("DB")
	// 在数据库DB创建集合user
	user := DBDatabase.Collection("user")

	// 新增集合user一行数据
	userInsert, _ := user.InsertOne(ctx, bson.D{
		{"name", "Tim"},
		{"age", 20},
	})
	fmt.Printf("新增集合user一行数据：%v\n", userInsert)

	// 新增集合user多行数据
	userInserts, _ := user.InsertMany(ctx, []interface{}{
		bson.D{{"name", "Tom"}, {"age", 20}},
		bson.D{{"name", "Lily"}, {"age", 30}},
	})
	fmt.Printf("新增集合user多行数据：%v\n", userInserts)

	// 更新集合user更新一行数据
	// 将name=Tom的数据改为Raboy
	userUpdate, _ := user.UpdateOne(
		ctx,
		bson.M{"name": "Tom"},
		bson.D{
			{"$set", bson.D{{"name", "Raboy"}}},
		},
	)
	fmt.Printf("更新集合user更新一行数据：%v\n", userUpdate)

	// 更新集合user多行数据
	// 将age=20的所有数据改为25
	userUpdates, _ := user.UpdateMany(
		ctx,
		bson.M{"age": 20},
		bson.D{
			{"$set", bson.D{{"age", 25}}},
		},
	)
	fmt.Printf("更新集合user多行数据：%v\n", userUpdates)

	// 替换集合user某行数据的所有数据
	userReplace, _ := user.ReplaceOne(
		ctx,
		bson.M{"name": "Lily"},
		bson.M{
			"name": "Lucy",
			"age":  29,
		},
	)
	fmt.Printf("替换集合user某行数据的所有数据：%v\n", userReplace)

	// 读取集合user的所有数据
	userFinds, _ := user.Find(ctx, bson.M{})
	defer userFinds.Close(ctx)
	// 遍历输出每行数据
	for userFinds.Next(ctx) {
		var datas bson.M
		// 每行数据加载到变量datas
		userFinds.Decode(&datas)
		fmt.Printf("读取集合user当前数据：%v\n", datas)
	}

	// 读取集合user的某行数据
	userFind, _ := user.Find(ctx, bson.M{"age": 25})
	defer userFind.Close(ctx)
	var data []bson.M
	// 数据加载到变量data
	userFind.All(ctx, &data)
	fmt.Printf("读取集合user的某行数据：%v\n", data)

	// 删除集合user某行数据
	userDelete, _ := user.DeleteOne(ctx, bson.M{"name": "Tom"})
	fmt.Printf("删除集合user某行数据：%v\n", userDelete)

	// 删除集合user多行数据
	userDeletes, _ := user.DeleteMany(ctx, bson.M{"age": 25})
	fmt.Printf("删除集合user多行数据：%v\n", userDeletes)

	// 删除集合user和所有数据
	err = user.Drop(ctx)
	if err != nil {
		fmt.Printf("删除集合user和所有数据失败：%v\n", err)
	} else {
		fmt.Printf("删除集合user和所有数据成功\n")
	}
}

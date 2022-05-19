// db_test.go
package db

import (
	"errors"
	"fmt"
	"testing"

	"github.com/ZackWongAither/firstmock/db/mocks"
	gomock "github.com/golang/mock/gomock"
)

func TestGetFromDB(t *testing.T) {
	// 创建gomock控制器，用来记录后续的操作信息
	ctrl := gomock.NewController(t)
	// 断言期望的方法都被执行
	// Go1.14+的单测中不再需要手动调用该方法
	// defer ctrl.Finish()
	// 调用mockgen生成代码中的NewMockDB方法
	// 这里mocks是我们生成代码时指定的package名称
	m := mocks.NewMockDB(ctrl)
	// 打桩（stub）
	// 当传入Get函数的参数为liwenzhou.com时返回1和nil
	m.
		EXPECT().
		Get(gomock.Eq("liwenzhou.com")). // 参数
		Return(1, nil).                  // 返回值
		Times(1)                         // 调用次数

	// 调用GetFromDB函数时传入上面的mock对象m
	if v := GetFromDB(m, "liwenzhou.com"); v != 1 {
		t.Fatal()
	}

}

func TestGetFromDB_InOrder(t *testing.T) {
	ctrl := gomock.NewController(t)
	m := mocks.NewMockDB(ctrl)

	// 指定顺序
	gomock.InOrder(
		m.EXPECT().Get("1"),
		m.EXPECT().Get("2"),
		m.EXPECT().Get("3"),
	)

	// 按顺序调用
	GetFromDB(m, "1")
	GetFromDB(m, "2")
	GetFromDB(m, "3")
}

func TestGetFromDB_Times(t *testing.T) {
	ctrl := gomock.NewController(t)
	m := mocks.NewMockDB(ctrl)

	m.
		EXPECT().
		Get(gomock.Eq("liwenzhou.com")). // 参数
		Return(1, nil).                  // 返回值
		Times(1)                         // 设置Get方法期望被调用次数为1

	// 调用GetFromDB函数时传入上面的mock对象m
	if v := GetFromDB(m, "liwenzhou.com"); v != 1 {
		t.Fatal()
	}
	// 再次调用上方mock的Get方法时不满足调用次数为1的期望
	// if v := GetFromDB(m, "liwenzhou.com"); v != 1 {
	// 	t.Fatal()
	// }
}

func TestGetFromDB_Return(t *testing.T) {
	ctrl := gomock.NewController(t)
	m := mocks.NewMockDB(ctrl)

	// gomock中跟返回值相关的用法有以下几个：

	// Return()：返回指定值
	// Do(func)：执行操作，忽略返回值
	// DoAndReturn(func)：执行并返回指定值

	m.EXPECT().Get(gomock.Any()).Return(20, nil)

	if v := GetFromDB(m, "1"); v != 20 {
		t.Fatal()
	}

	m.EXPECT().Get(gomock.Any()).Do(func(key string) {
		fmt.Printf("input keyd is %v\n", key)
	})

	v := GetFromDB(m, "1")
	fmt.Println(v)

	m.EXPECT().Get(gomock.Any()).DoAndReturn(func(key string) (int, error) {
		fmt.Printf("input key is %v\n", key)
		return 0, errors.New("..")
	})
	if v := GetFromDB(m, "1"); v != -1 {
		t.Fatal()
	}

}

package user

import (
	"errors"
	"log"
	"testing"

	gomock "github.com/golang/mock/gomock"
)

// 静态设置返回值
func TestReturn(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repo := NewMockUserRepository(ctrl)
	// 期望FindOne(1)返回张三用户
	repo.EXPECT().FindOne(1).Return(&User{Name: "张三"}, nil)
	// 期望FindOne(2)返回李四用户
	repo.EXPECT().FindOne(2).Return(&User{Name: "李四"}, nil)
	// 期望给FindOne(3)返回找不到用户的错误
	repo.EXPECT().FindOne(3).Return(nil, errors.New("user not found"))
	// 验证一下结果
	log.Println(repo.FindOne(1)) // 这是张三
	log.Println(repo.FindOne(2)) // 这是李四
	log.Println(repo.FindOne(3)) // user not found
	//log.Println(repo.FindOne(4)) //没有设置4的返回值，却执行了调用，测试不通过
}

// 动态设置返回值
func TestReturnDynamic(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	repo := NewMockUserRepository(ctrl)
	// 常用方法之一：DoAndReturn()，动态设置返回值
	repo.EXPECT().FindOne(gomock.Any()).DoAndReturn(func(i int) (*User, error) {
		if i == 0 {
			return nil, errors.New("user not found")
		}
		if i < 100 {
			return &User{
				Name: "小于100",
			}, nil
		} else {
			return &User{
				Name: "大于等于100",
			}, nil
		}
	})
	log.Println(repo.FindOne(120))
	//log.Println(repo.FindOne(66))
	//log.Println(repo.FindOne(0))
}

func TestTimes(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repo := NewMockUserRepository(ctrl)
	// 默认期望调用一次
	repo.EXPECT().FindOne(1).Return(&User{Name: "张三"}, nil)
	// 期望调用2次
	repo.EXPECT().FindOne(2).Return(&User{Name: "李四"}, nil).Times(2)
	// 调用多少次可以,包括0次
	repo.EXPECT().FindOne(3).Return(nil, errors.New("user not found")).AnyTimes()

	// 验证一下结果
	log.Println(repo.FindOne(1)) // 这是张三
	log.Println(repo.FindOne(2)) // 这是李四
	log.Println(repo.FindOne(2)) // FindOne(2) 需调用两次,注释本行代码将导致测试不通过
	log.Println(repo.FindOne(3)) // user not found, 不限调用次数，注释掉本行也能通过测试
}

func TestOrder(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	repo := NewMockUserRepository(ctrl)
	o1 := repo.EXPECT().FindOne(1).Return(&User{Name: "张三"}, nil)
	o2 := repo.EXPECT().FindOne(2).Return(&User{Name: "李四"}, nil)
	o3 := repo.EXPECT().FindOne(3).Return(nil, errors.New("user not found"))
	gomock.InOrder(o1, o2, o3) //设置调用顺序
	// 按顺序调用，验证一下结果
	log.Println(repo.FindOne(1)) // 这是张三
	log.Println(repo.FindOne(2)) // 这是李四
	log.Println(repo.FindOne(3)) // user not found

	// 如果我们调整了调用顺序，将导致测试不通过：
	// log.Println(repo.FindOne(2)) // 这是李四
	// log.Println(repo.FindOne(1)) // 这是张三
	// log.Println(repo.FindOne(3)) // user not found
}

package example_test

import (
	"fmt"

	"github.com/zzerjae/sejong"
)

func ExampleT() {
	sejong.Locale = "en-GB"

	message, err := sejong.T("message.welcome", "nickname", "John")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(message)

	message, err = sejong.T("message.farewell", "nickname", "John", "time", "5")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(message)

	sejong.Locale = "ko"

	message, err = sejong.T("message.welcome", "nickname", "길동")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(message)

	ko, err := sejong.New("ko")
	if err != nil {
		fmt.Println(err)
		return
	}
	gb, err := sejong.New("en-GB")
	if err != nil {
		fmt.Println(err)
		return
	}

	message, err = ko.T("message.welcome", "nickname", "길동")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(message)

	message, err = gb.T("message.welcome", "nickname", "John")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(message)

	message, err = gb.T("message.friend", "count", "0")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(message)
	message, err = gb.T("message.friend", "count", "1")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(message)
	message, err = gb.T("message.friend", "count", "2")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(message)

	message, err = ko.T("message.friend", "count", "0")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(message)
	message, err = ko.T("message.friend", "count", "1")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(message)
	message, err = ko.T("message.friend", "count", "2")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(message)
	// Output:
	// Hello, John!
	// It's 5 o'clock. Good bye, John!
	// 안녕, 길동!
	// 안녕, 길동!
	// Hello, John!
	// I have no friend.
	// I have a friend.
	// I have 2 friends.
	// 저는 친구가 없어요.
	// 저는 1명의 친구가 있어요.
	// 저는 2명의 친구가 있어요.
}

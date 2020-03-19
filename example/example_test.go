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
	// Output:
	// Hello, John!
	// It's 5 o'clock. Good bye, John!
	// 안녕, 길동!
}

func ExampleMultipleTranslator() {
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

	message, err := ko.T("message.welcome", "nickname", "길동")
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
	// Output:
	// 안녕, 길동!
	// Hello, John!
}

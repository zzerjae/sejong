package example_test

import (
	"fmt"

	"github.com/zzerjae/sejong"
)

func ExampleT() {
	sejong.Locale = "en-GB"

	message := sejong.T("message.welcome", "nickname", "John")
	fmt.Println(message)

	message = sejong.T("message.farewell", "nickname", "John", "time", "5")
	fmt.Println(message)

	sejong.Locale = "ko"

	message = sejong.T("message.welcome", "nickname", "길동")
	fmt.Println(message)
	// Output:
	// Hello, John!
	// It's 5 o'clock. Good bye, John!
	// 안녕, 길동!
}

func ExampleMultipleTranslator(){
	ko := sejong.New("ko")
	gb := sejong.New("en-GB")

	message := ko.T("message.welcome", "nickname", "길동")
	fmt.Println(message)

	message = gb.T("message.welcome", "nickname", "John")
	fmt.Println(message)
	// Output:
	// 안녕, 길동!
	// Hello, John!
}
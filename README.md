# sejong

simple translate library like ruby's i18n module

## Example

Suppose you need some greetings messages in English and Korean respectively.

You can write en-GB.yml

```yml
# en-GB.yml
en-GB:
  message:
    welcome: 'Hello, %{nickname}!'
    farewell: 'It\'s %{time} o\'clock. Good bye, %{nickname}!'
```

and ko.yml.

```yml
# ko.yml
ko:
  message:
    welcome: '안녕, %{nickname}!'
    farewell: '%{time}시 입니다. 잘가요, %{nickname}!'
```

then you can use `sejong.T` for translate.

```go
package main

import (
	"fmt"
	"time"

	"github.com/zzerjae/sejong"
)

func main() {
	sejong.Locale = "en-GB"

	message := sejong.T("message.welcome", "nickname", "John")
	fmt.Println(message) // Hello, John!

	message = sejong.T("message.farewell", "nickname", "John", "time", time.Now().Hour())
	fmt.Println(message) // It's 5 o'clock. Good bye, John!

	sejong.Locale = "ko"

	message = sejong.T("message.welcome", "nickname", "길동")
	fmt.Println(message) // 안녕, 길동!
}
```

## Configuration

### Location of locale files

Default location is the directory of `main.go`, but you can set the env `SEJONG_LOCALE_DIRECTORY` for yours.

### Working with multiple locales

You can also create many different translators for use in your application. Each will have its own unique locale source.

Example:

```go
package main

import (
	"github.com/zzerjae/sejong"
)

func main() {
	ko := sejong.New("ko")
	gb := sejong.New("en-GB")

	message := ko.T("welcome_message", "nickname", "길동")
	fmt.Println(message) // 안녕, 길동!

	message = gb.T("welcome_message", "nickname", "John")
	fmt.Println(message) // Hello, John!
}

```

## QNA

### Why is it called "sejong"?

https://en.wikipedia.org/wiki/Sejong_the_Great

## Author

@zzerjae

## License

MIT

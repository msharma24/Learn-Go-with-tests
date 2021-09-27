package main

import "fmt"

const spanish = "Spanish"
const french = "French"
const englishHelloPrefix = "Hello, "
const frenchHelloPrefix = "Bonjour, "
const spanishHelloPrefix = "Hola, "

func greetingPrefix(language string) (prefix string ){
  switch language {
  case french:
    prefix = frenchHelloPrefix

  case spanish:
    prefix = spanishHelloPrefix

  default:
    prefix = englishHelloPrefix
  }

  return

}

func Hello(name string, language string) string {

  if name == "" {
    name = "World"

  }

  return greetingPrefix(language) + name

}

func main() {
  fmt.Println("World", "")
}

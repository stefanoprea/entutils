package matcher

type Config struct {
  URL string
  User string
  secret string
}

func makeConfigsWithSameUserAndURL() (Config, Config) {
  expected := Config{
    URL: "http://www.example.com/index.html",
    User: "John Doe",
    secret: "25",
  }
  actual := Config{
    URL: "http://www.example.com/index.html",
    User: "John Doe",
    secret: "124",
  }
  return expected, actual
}

func makeConfigsWithSameURL() (Config, Config) {
  expected := Config{
    URL: "http://www.example.com/index.html",
    User: "John Doe",
    secret: "25",
  }
  actual := Config{
    URL: "http://www.example.com/index.html",
    User: "Larry",
    secret: "124",
  }
  return expected, actual
}

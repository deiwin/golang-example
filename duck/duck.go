package duck

type Duck interface {
	Quack() string
	Walk()
	Swim()
}

func FeedBreadcrumbsTo(duck Duck) string {
	return duck.Quack()
}

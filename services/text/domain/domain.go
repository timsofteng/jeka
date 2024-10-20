package text

type RandomText string

type Repo interface {
	Add(text string) error
	Fetch() (text string, err error)
	Count() (count uint, err error)
}

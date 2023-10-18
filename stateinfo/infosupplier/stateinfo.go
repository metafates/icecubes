package infosupplier

import (
	"github.com/metafates/soda/title"
)

type impl struct {
	title    func() title.Title
	subtitle func() string
	status   func() string
}

func (i impl) Title() title.Title {
	return i.title()
}

func (i impl) Subtitle() string {
	return i.subtitle()
}

func (i impl) Status() string {
	return i.status()
}

type Supplier interface {
	Title() title.Title
	Subtitle() string
	Status() string
}

type Option func(supplier *impl)

func WithTitle(t title.Title) Option {
	return func(supplier *impl) {
		supplier.title = func() title.Title {
			return t
		}
	}
}

func WithTitleFunc(title func() title.Title) Option {
	return func(supplier *impl) {
		supplier.title = title
	}
}

func WithSubtitle(subtitle string) Option {
	return func(supplier *impl) {
		supplier.subtitle = func() string {
			return subtitle
		}
	}
}

func WithSubtitleFunc(subtitle func() string) Option {
	return func(supplier *impl) {
		supplier.subtitle = subtitle
	}
}

func WithStatus(status string) Option {
	return func(supplier *impl) {
		supplier.status = func() string {
			return status
		}
	}
}

func WithStatusFunc(status func() string) Option {
	return func(supplier *impl) {
		supplier.status = status
	}
}

func New(options ...Option) Supplier {
	supplier := impl{
		title: func() title.Title {
			return title.New("Title")
		},
		subtitle: func() string {
			return ""
		},
		status: func() string {
			return ""
		},
	}

	for _, option := range options {
		option(&supplier)
	}

	return supplier
}

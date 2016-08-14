package main

import (
	"gopkg.in/qml.v1"
	"fmt"
	"regexp"
)

type YoutuberUI struct {
	Url string
	Rpattern string
	TextAuthor string
	TextComment string
	ImageLinkAuthor string
	LinkAuthor string
}

func (e *YoutuberUI) Start(url string, pattern string) {

	fmt.Println("Getting URL data: ", url)
	fmt.Println("Getting Pattern data: ", pattern)
	CommentsDB = make(map[string]ParsedComment)

	r, err := regexp.Compile(pattern)
	if err != nil {
		fmt.Println(err)
	}


	params := GetCommentsParams{Url: url, Pattern: r}
	GetComments(params)
	RandomComment := GetRandComment()

	e.TextAuthor = RandomComment.AuthorDisplayName
	e.TextComment = RandomComment.TextDisplay
	e.ImageLinkAuthor = RandomComment.AuthorProfileImageURL
	e.LinkAuthor = RandomComment.AuthorChannelURL

	qml.Changed(e, &e.TextAuthor) // нужно, чтобы qml узнал о обновлении переменной
	qml.Changed(e, &e.TextComment) // нужно, чтобы qml узнал о обновлении переменной
	qml.Changed(e, &e.ImageLinkAuthor) // нужно, чтобы qml узнал о обновлении переменной
	qml.Changed(e, &e.LinkAuthor) // нужно, чтобы qml узнал о обновлении переменной
}


func startGui() {
	err := qml.Run(run)
	CheckErr(err)
}

func run() error {
	engine := qml.NewEngine()
	youtuberUI := &YoutuberUI{}
	engine.Context().SetVar("youtuberUI", youtuberUI)
	controls, err := engine.LoadFile("qrc:///assets/main.qml")
	if err != nil {
		return err
	}
	window := controls.CreateWindow(nil)
	window.Show()
	window.Wait()
	return nil
}

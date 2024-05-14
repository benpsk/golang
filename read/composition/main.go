package main

import "fmt"

type author struct {
	firstName string
	lastName  string
	bio       string
}

func (a author) fullName() string {
	return fmt.Sprintf("%s %s", a.firstName, a.lastName)
}

type blogPost struct {
	title   string
	content string
	author
}

func (bp blogPost) details() {
	fmt.Println("Title: ", bp.title)
	fmt.Println("Content: ", bp.content)
	fmt.Println("Author: ", bp.fullName())
	fmt.Println("Bio: ", bp.bio)
}

type website struct {
	blogPosts []blogPost
}

func (w website) contents() {
	fmt.Println("Contents of Website")
	for _, v := range w.blogPosts {
		v.details()
		fmt.Println()
	}
}

func main() {
	author1 := author{
		"Ben",
		"Gunn",
		"Go Developer",
	}
	blogPost1 := blogPost{
		title:   "Inheritance in Go",
		content: "Go supports composition instead of inheritance",
		author:  author1,
	}
	blogPost2 := blogPost{
		title:   "Struct instead of Classes in Go",
		content: "Go does not support classes but methods can be added to structs",
		author:  author1,
	}
	blogPost3 := blogPost{
		title:   "Concurrency",
		content: "Go is a concurrent language and not a parallel one",
		author:  author1,
	}
	w := website{
		blogPosts: []blogPost{blogPost1, blogPost2, blogPost3},
	}
	w.contents()
}

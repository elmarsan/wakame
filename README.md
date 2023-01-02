# Wakame

Minimalist html parser, it allows you to generate tree structure based in html nodes.

[![Go Report Card](https://goreportcard.com/badge/github.com/elmarsan/wakame)](https://goreportcard.com/report/github.com/elmarsan/wakame)
![Coverage](https://img.shields.io/badge/Coverage-98.5%25-brightgreen)
[![License](https://img.shields.io/badge/License-Apache_2.0-blue.svg)](https://opensource.org/licenses/Apache-2.0)
[![Go Reference](https://pkg.go.dev/badge/github.com/elmarsan/wakame.svg)](https://pkg.go.dev/github.com/elmarsan/wakame)

## Usage

- `go get github.com/elmarsan/wakame`


```go

html := `
<div>
	<p>Hello world</p>
	<p>Goodbye world</p>
	<p color="blue">Wakame prototype</p>
	<div id="my-img">
		<img src="https://draxe.com/wp-content/uploads/2018/10/WakameHeader.jpg" class="img-class" alt="Wakame photo" /> 
	</div>
</div>
`

root := wakame.ParseHTML(html)

// find node by tag
paragrahps := root.FindAll("p", map[string]interface{}{})

fmt.Println(paragrahps[0].Content) // Hello World

// find node by attribute
blueText := root.FindAll("p", map[string]interface{}{
	"color": "blue",
})
fmt.Printf("I'm blue with innerText: %s and attributes %#v\n", blueText[0].Content, blueText[0].Attributes)

// you can see node attributes and father
img := root.FindAll("img", map[string]interface{}{})
fmt.Printf("Img src: %s, my father is a: %s", img[0].Attributes["src"], img[0].Parent.Tag)

```

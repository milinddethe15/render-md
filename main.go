package main

import (
	"syscall/js"

	"github.com/gomarkdown/markdown"
	"github.com/gomarkdown/markdown/html"
	"github.com/gomarkdown/markdown/parser"
)

func convertMarkdownWrapper(this js.Value, p []js.Value) interface{} {
	input := p[0].String()

	// Enable Markdown extensions (tables, checkboxes, etc.)
	extensions := parser.CommonExtensions | parser.Tables | parser.Autolink | parser.Strikethrough | parser.NoEmptyLineBeforeBlock

	mdParser := parser.NewWithExtensions(extensions)

	// Enable HTML rendering options
	htmlFlags := html.CommonFlags | html.HrefTargetBlank
	renderer := html.NewRenderer(html.RendererOptions{Flags: htmlFlags})

	// Convert Markdown to HTML
	htmlOutput := markdown.ToHTML([]byte(input), mdParser, renderer)

	return js.ValueOf(string(htmlOutput))
}

func main() {
	js.Global().Set("convertMarkdown", js.FuncOf(convertMarkdownWrapper))
	select {} // Keeps Go running in WebAssembly
}

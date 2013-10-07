package bham

var Filters = []FilterHandler{
	FilterHandler{
		Trigger: ":javascript",
		Open:    `<script type="text/javascript">`,
		Close:   "</script>",
		Handler: Transformer(func(s string) string { return s }),
	},
	FilterHandler{
		Trigger: ":css",
		Open:    `<style>`,
		Close:   "</style>",
		Handler: Transformer(func(s string) string { return s }),
	},
}

type FilterHandler struct {
	Trigger     string
	Open, Close string
	Handler     Transformer
}

type Transformer func(string) string

# Scromble

A small game project, mainly aimed at gaining experience using Go, HTMX and Templ.

This uses a go based webserver - main.go - which serves pages generated using Templ HTML templates. The generated pages make use of HTMX to give the pages reactivity.

Please make sure to run "go generate" after syncing or changing any of the templ files so that the updated go files are generated from the templ files.

When built with "go build" the scromble.exe app will be created. You can view it's output by going to http://localhost:3000/.

## Useful Links

- [Templ](https://templ.guide/)
- [HTMX](https://htmx.org/docs/)
- [Example: todo](https://github.com/stackus/todos/tree/main)
- [Pure CSS](https://purecss.io/)

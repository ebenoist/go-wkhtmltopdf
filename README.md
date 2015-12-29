wkhtmltopdf-go
---

### PoC YMMV

Direct C bindings for wkhtmltopdf as a PoC. Built based on the belief that shelling out is slow and dangerous.

Turns out this is still slow, but likely less dangerous.

### Usage
```Bash
$ go-wkhtmltopdf -html "<b><i>foo</i></b>" > new.pdf
```


### Requirements
You're going to need to install wkhtmltopdf and likely via brew cask.

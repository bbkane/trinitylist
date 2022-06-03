# Dev Notes

Let's make the list first, then pop it in a layout

Notes app: https://github.com/fynelabs/notes

Demo: https://github.com/fyne-io/fyne/tree/master/cmd/fyne_demo

List binding: https://developer.fyne.io/binding/list

## Input stretching

Now I'm trying to make the text input stretch acroos the bottom part

try https://developer.fyne.io/container/box

also see https://www.youtube.com/watch?v=LWn1403gY9E

Ooh I can use a form... https://developer.fyne.io/explore/layouts

## Attaching Enter to input

We can try https://developer.fyne.io/api/v2.1/widget/entry.html#func-entry-typedkey

That didn't work, but adding OnSubmit did. Odd that it's not documented.

Done in 479d5bc

## Adding Color to list entries

This looks harder, so let's try it first...

Trying: https://github.com/fyne-io/fyne/issues/619

Per https://gophers.slack.com/archives/CB4QUBXGQ/p1654222862590979 , I can use https://blogvali.com/button-color-background-image-fyne-gui-golang-tutorial-44/

## Editing existing list entries



# Issues to create

Link to demo code in https://fyne.io/ , https://developer.fyne.io/started/#run-the-demo

Add binding to struct


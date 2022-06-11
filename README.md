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

Done in ccc80e0

## Editing existing list entries

This is gonna use bindstruct... I'm going to have to

https://github.com/search?q=binding.BindStruct&type=code

## List of Structs questions

I need to write a list of questions about how to bind structs...

Examples of using BindStruct with lists

Should I regen the list and repaint all widgets each time? Perf?

# Issues to create

Add binding to struct

# Steal ideas from Orgzly

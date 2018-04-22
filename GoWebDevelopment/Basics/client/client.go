package main

import(
  "fmt"
  "honnef.co/go/js/dom"
)

var d = dom.GetWindow().Document().(dom.HTMLDocument)

fun run() {
  messageInput := d.GetElementByID("messageInput").(*dom.HTMLInputElement)

  alertButtonJS := d.GetElementByID("alertMessageJSGlobal").(*dom.HTMLButtonElement)
  alertButtonJS.AddEventListener("click", false, func(event dom.Event) {
    DisplayAlertMessageJSGlobal(messageInput.Value)
  })

  alertButtonDOM := d.GetElementByID("alertMessageDOM").(*dom.HTMLButtonElement)
  alertButtonDOM.AddEventListener("click", false, func(event dom.Event) {
    go lowerCaseTextTransformer()
  })
}

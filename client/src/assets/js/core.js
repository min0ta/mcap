// DYNAMIC PAGE LOADING, NOT TESTED YET!
//#region 
let parser = new DOMParser()
let doc = document
async function loadHref(href) {
    let pageText = await (await fetch(href)).text()
    let html = parser.parseFromString(pageText, "text/html")
    
    doc.querySelectorAll("script").forEach(v => v.remove())
    doc.querySelector("main").replaceWith(html.querySelector("main"))

    html.querySelectorAll("script").forEach(async v => {
        let script = doc.createElement("script")
        script.append(`{${await (await fetch(v.src)).text()}}`)
        doc.body.appendChild(script)
    })
    history.pushState({}, "", href)
}

// UNCOMMENT THIS TO ENABLE DYNAMIC PAGE LOADING
/*
let a = doc.querySelectorAll("a").forEach(v => {
    if (v.type === "") {
        v.addEventListener("click", (e) => {
            e.preventDefault()
            loadHref(v.href)
        })
    }
})
*/

//#endregion

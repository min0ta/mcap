// DYNAMIC PAGE LOADING, NOT TESTED YET!
//#region 
let parser = new DOMParser()
let doc = document
async function loadHref(href) {
    history.pushState({}, "", href)
    let pageText = await (await fetch(href)).text()
    let html = parser.parseFromString(pageText, "text/html")
    
    doc.querySelectorAll("link").forEach(v => {
        if (!v.attributes.immutable) {
            v.remove()
        }
    })     

    doc.querySelectorAll("script").forEach(v => {
        if (!v.attributes.immutable) {
            v.remove()
        }
    })
    doc.querySelectorAll("style").forEach(async v => {
        if (v.attributes.generated)
            v.remove()
    })
    doc.querySelector("main").replaceWith(html.querySelector("main"))

    html.querySelectorAll("script").forEach(async v => {
        if (!v.attributes.immutable) {
            let script = doc.createElement("script")
            script.append(`{${await (await fetch(v.src)).text()}}`)
            doc.body.appendChild(script)
        }
    })
    html.querySelectorAll("link").forEach(async v => {
        if (!v.attributes.immutable) {
            let css = doc.createElement("style")
            css.setAttribute("generated", "")
            css.append(await (await fetch(v.href)).text())
            doc.body.appendChild(css)
        }
    })
}

function createSmoothAnchor(a) {
    if (a.type === "") {
        a.addEventListener("click", (e) => {
            e.preventDefault()
            loadHref(a.href)
        })
    }
}

let a = doc.querySelectorAll("a").forEach(v => {
    createSmoothAnchor(v)
})


//#endregion

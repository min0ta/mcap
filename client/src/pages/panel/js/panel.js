const state = {
    isServerOnline: false,
    isUILocked: false,
    serverName: null
}

/*** @type {HTMLHeadingElement}*/
const serverHeader = gcel("server-name")
/*** @type {HTMLDivElement}*/
const serverStateElement = gcel("server-state")
/*** @type {HTMLButtonElement}*/
const serverToggleButton = gcel("server-toggle")

/** @type {HTMLSpanElement} */
const errorOutput = gid("js-errout")
serverHeader.textContent = "<server name>"

/** * @param {{serverState: boolean, oldServerStateClass: string, newServerStateClass: string, barIconSrc: string, barText: string, buttonText: string, buttonOldClass: string, buttonNewClass: string}} param0  */
function setServerState({serverState, oldServerStateClass, newServerStateClass, barIconSrc, barText, buttonText, buttonOldClass, buttonNewClass}) {
    state.isServerOnline = serverState
    serverStateElement.classList.replace(oldServerStateClass, newServerStateClass);
    // this is kinda hacky but i was too lazy to select muiltiple elements just to change their src
    /*** @type {HTMLImageElement}*/
    (serverStateElement.children[0]).src = barIconSrc
    serverStateElement.children[1].textContent = barText
    
    serverToggleButton.textContent = buttonText
    serverToggleButton.classList.replace(buttonOldClass, buttonNewClass)
}
/** @param {boolean} online */
function updateServerState(online) {
    online ?
    setServerState({
        serverState: true, 
        oldServerStateClass: "server-state-disabled", 
        newServerStateClass: "server-state-enabled", 
        barIconSrc:"./assets/online.svg", 
        barText:"Включен", 
        buttonText:"Остановить", 
        buttonOldClass: "server-toggle-enabled", 
        buttonNewClass: "server-toggle-disabled"
    }) :
    setServerState({
        serverState: false, 
        oldServerStateClass: "server-state-enabled", 
        newServerStateClass: "server-state-disabled", 
        barIconSrc: "./assets/offline.svg", 
        barText:"Выключен", 
        buttonText: "Запустить", 
        buttonOldClass: "server-toggle-disabled", 
        buttonNewClass: "server-toggle-enabled"
    }) 
}

function lockUI() {
    state.isUILocked = true
    serverToggleButton.disabled = true
}
function unlockUI() {
    state.isUILocked = false
    serverToggleButton.disabled = false
}
function parseParams() {
    let get = window.location.search
    if (get === '') {
        return {}
    }
    const map = {}
    get = get.substring(1)
    let queries = get.split("&")
    for (let i = 0; i<queries.length; i++) {
        let kv = queries[i].split("=")
        map[kv[0]] = kv[1]
    }
    return map
}

function renderError(displayableError, ...err) {
    logError(err.join(" "))
    errorOutput.textContent = displayableError
}

function clearError() {
    errorOutput.textContent = ""
}

async function lockRenderErrUnlock(callback) {
    lockUI()
    try {
        await callback()
    } catch (e) {
        renderError(e)
        throw e
    }
    unlockUI()
}

async function main() {
    lockUI()
    const params = parseParams()
    const server = params["server"]
    if (server == null) {
        renderError("Сервер не указан!", "no server provided", JSON.stringify(params))
        return
    }
    try {
        const info = await api.getServerState(server)
        updateServerState(info.online)
        serverHeader.textContent = info.name
    } catch (e) {
        renderError(e, "getServerState() error! server name =",server,"error:", e)
        throw e
    }  
    unlockUI()
    quitOnClick()

    serverToggleButton.addEventListener("click",async () => {
        if (state.isUILocked) {
            return
        }
        clearError()
        if (state.isServerOnline) {
            try {
                await lockRenderErrUnlock(async () => await api.stopServer(server))
                updateServerState(false)
            } catch (e) {
                console.log(e)
            }
        }
        try {
            lockRenderErrUnlock(async () => await api.startServer(server))
            updateServerState(true)
        } catch (e) {
            console.log(e)
        }
        window.location.reload()
    })
    if (!state.isServerOnline) {
        return
    }
    const ws = createConn(server)
    ws.onmessage = (ev) => {
        console.log(ev.data)
    }
    const interval = setInterval(async () => {
        ws.send(JSON.stringify({type:"keep"}))
    }, 7000)
    ws.onerror = (e) => clearInterval(interval)
    ws.onclose = (e) => clearInterval(interval)
}
main()
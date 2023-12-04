const state = {
    isServerOnline: false,
    isUILocked: false,
}

/*** @type {HTMLHeadingElement}*/
const serverName = gcel("server-name")
/*** @type {HTMLDivElement}*/
const serverStateElement = gcel("server-state")
/*** @type {HTMLButtonElement}*/
const serverToggleButton = gcel("server-toggle")
serverName.textContent = "<server name>"

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
/** @param {boolean} state */
function setServerState(state) {
    state ?
    setServerState({
        serverState: false, 
        oldServerStateClass: "server-state-enabled", 
        newServerStateClass: "server-state-disabled", 
        barIconSrc: "./assets/offline.svg", 
        barText:"Выключен", 
        buttonText: "Запустить", 
        buttonOldClass: "server-toggle-disabled", 
        buttonNewClass: "server-toggle-enabled"
    }) :
    setServerState({
        serverState: true, 
        oldServerStateClass: "server-state-disabled", 
        newServerStateClass: "server-state-enabled", 
        barIconSrc:"./assets/online.svg", 
        barText:"Включен", 
        buttonText:"Остановить", 
        buttonOldClass: "server-toggle-enabled", 
        buttonNewClass: "server-toggle-disabled"
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
/** @param {string} className @returns {Element} */
function gcel(className) {
    return document.getElementsByClassName(className)[0]
}

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

function toggleServerState() {
    if (state.isServerOnline) {
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
        return 
    }
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

function updateServerState() {
    toggleServerState()
}
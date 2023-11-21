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

function __toggleServerState() {
    if (state.isServerOnline) {
        state.isServerOnline = false
        serverStateElement.classList.replace("server-state-enabled", "server-state-disabled");
        // this is kinda hacky but i was too lazy to select muiltiple elements just to change their src
        /*** @type {HTMLImageElement}*/
        (serverStateElement.children[0]).src = "./assets/offline.svg"
        serverStateElement.children[1].textContent = "Выключен"
        
        serverToggleButton.textContent = "Запустить"
        serverToggleButton.classList.replace("server-toggle-disabled", "server-toggle-enabled")
        return
    }
    state.isServerOnline = true
    serverStateElement.classList.replace("server-state-disabled", "server-state-enabled");
    /*** @type {HTMLImageElement}*/
    serverStateElement.children[0].src = "./assets/online.svg"
    serverStateElement.children[1].textContent = "Включен"

    serverToggleButton.textContent = "Остановить"
    serverToggleButton.classList.replace("server-toggle-enabled", "server-toggle-disabled")

}

function lockUI() {
    state.isUILocked = true
    serverToggleButton.disabled = true;
}
function unlockUI() {
    state.isUILocked = false
    serverToggleButton.disabled = false
}

function updateServerState() {
    __toggleServerState()
}
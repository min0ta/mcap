/**
 * @param {string} className 
 * @returns {Element}
 */
function gcel(className) {
    return document.getElementsByClassName(className)[0]
}

const serverName = gcel("server-name")

serverName.textContent = "<server name>"

const serverStateElement = gcel("server-state")

let __serverState = false
function __toggleServerState() {
    if (__serverState) {
        __serverState = false
        serverStateElement.classList.replace("server-state-enabled", "server-state-disabled")
        serverStateElement.children[0].src = "./assets/offline.svg"
    serverStateElement.children[1].textContent = "Выключен"
        return
    }
    __serverState = true
    serverStateElement.classList.replace("server-state-disabled", "server-state-enabled")
    serverStateElement.children[0].src = "./assets/online.svg"
    serverStateElement.children[1].textContent = "Включен"
}

function updateServerState() {
    __toggleServerState()
}
/** @param {string} className @returns {Element} */
function gcel(className) {
    return document.getElementsByClassName(className)[0]
}
/*** @param {string} idName @returns {Element}*/
function gid(idName) {
    return document.getElementById(idName)
}

/**
 * @param {string} err 
 */
function logError(err) {
    let log = sessionStorage.getItem("errlog") ?? ""
    if (log.length > 100) {
        log = ""
    }
    sessionStorage.setItem("errlog",`${log}\n${new Date()}: ${err}`)
}
class ServerApi {
    /** @type {string} */
    rootPath
    constructor(rootPath) {
        this.rootPath = rootPath
    }
    #errorEnum = {
        ErrorBadQuery:0,
        ErrorBadLoginOrPassword:1,
        ErrorUnauthorized: 2,
        ErrorCannotAccessRcon: 3,
        ErrorCannotStartMcServer: 4
    }
    #assert(predicate, errorCode) {
        if (+errorCode === this.#errorEnum.ErrorUnauthorized) {
            window.location.href = "/login/login"
        }   
        if (!predicate) {
            throw this.#parseError(errorCode)
        }
    }
    #parsedErrorArray = ["Неверный запрос!", "Неверный логин или пароль!", "Вы неавторизованы!", "Невозможно получить доступ к RCON!", "Невозможно запустить сервер!"]
    #get(path) {
        return fetch(`${this.rootPath}/${path}`)
    }
    /** @param {string} path @param {Object} body */
    #post(path, body = {}) {
        const b = JSON.stringify(body)
        return fetch(`${this.rootPath}/${path}`, {
            method: "POST",
            body: b
        })
    }
    #parseError(error) {
        return this.#parsedErrorArray[error] ?? "Неизвестная ошибка!"
    }
    /** @param {string} username @param {string} password @throws {Error}*/
    async login(username, password) {
        const q = await (await this.#post("login", {username, password})).json()
        this.#assert(q.err == null, q.err)
    }
    /**@returns {Promise<{name: string, address: string, port: string}[]>}*/
    async getServerList() {
        const q = await (await this.#get("servers")).json()
        this.#assert(q.err == null, q.err)
        return q.list
    }
    /**@param {string} server @returns {Promise<{online: boolean, name: string}>} */
    async getServerState(server) {
        const q = await (await this.#post("server", {server})).json()
        this.#assert(q.err == null, q.err)
        return q
    }
    /**@param {string} server  */
    async startServer(server) {
        const q = await (await this.#post("start", {server})).json()
        this.#assert(q.err == null, q.err)
        return q.success
    }
    /**@param {string} server  */
    async stopServer(server) {
        const q = await (await this.#post("stop", {server})).json()
        this.#assert(q.err == null, q.err)
        return q.success
    }
    async unauth() {
        const q = await (await this.#post("unauth")).json()
        this.#assert(q.err == null, q.err)
        window.location.href = "/login/login"
    }
}
function quitCallback(e) {
    if (e.ctrlKey && (e.code === "KeyX")) {
        const shouldLeave = confirm("вы точно хотите выйти из аккаунта?")
        if (shouldLeave) {
            try {
                api.unauth()
            } catch (e) {
                alert("Не удалось выйти с ошибкой ", e)
                throw e
            }
            window.location.href = "/login/login"
        }
    }
}
function quitOnClick() {
    window.addEventListener("keydown", quitCallback)
}
function onloadhref() {
    window.removeEventListener("keydown", quitCallback)
}
const api = new ServerApi("http://localhost/api")

function createConn(server) {
    const ws = new WebSocket("ws://localhost:8080/logs")
    ws.onopen = () => {
        ws.send(JSON.stringify({server}))
    }
    return ws
}
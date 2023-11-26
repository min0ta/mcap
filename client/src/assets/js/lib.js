/** @param {string} className @returns {Element} */
function gcel(className) {
    return document.getElementsByClassName(className)[0]
}
/*** @param {string} idName @returns {Element}*/
function gid(idName) {
    return document.getElementById(idName)
}

// I WAS DRUNK WHILE DOUNG THIS SHIT SO PLESASE FROGIVE ME
class ServerApi {
    /** @type {string} */
    rootPath
    constructor(rootPath) {
        this.rootPath = rootPath
    }
    #errorEnum = {
        ErrorBadQuery:0,
        ErrorBadLoginOrPassword:1
    }
    #get(path) {
        return fetch(`${this.rootPath}/${path}`)
    }
    /** @param {string} path @param {Object} body */
    #post(path, body) {
        const b = JSON.stringify(body)
        console.log(b)
        return fetch(`${this.rootPath}/${path}`, {
            method: "POST",
            body: b
        })
    }
    #parseError(error) {
        if (+error === this.#errorEnum.ErrorBadQuery) {
            return new Error("Неверный запрос!")
        }
        if (+error === this.#errorEnum.ErrorBadLoginOrPassword) {
            return new Error("Неверное имя пользователя или пароль!")
        }
        return new Error("Неизвестная ошибка!")
    }
    /** @param {string} username @param {string} password @throws {Error}*/
    async login(username, password) {
        const q = await (await this.#post("login", {username, password})).json()
        if (q.err) {
            throw this.#parseError(q.err)
        }
    }
}

const api = new ServerApi("http://localhost/api")
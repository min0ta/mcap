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
        ErrorBadLoginOrPassword:1,
        ErrorUnauthorized: 2,
        ErrorCannotAccessRcon: 3,
        ErrorCannotStartMcServer: 4
    }
    #parsedErrorArray = ["Неверный запрос!", "Неверный логин или пароль!", "Вы неавторизованы!", "Невозможно получить доступ к RCON!", "Невозможно запустить сервер!"]
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
        return this.#parsedErrorArray[error] ?? "Неизвестная ошибка!"
    }
    /** @param {string} username @param {string} password @throws {Error}*/
    async login(username, password) {
        const q = await (await this.#post("login", {username, password})).json()
        if (q.err) {
            throw this.#parseError(q.err)
        }
    }
    /**@returns {Promise<{name: string, address: string, port: string}[]>}*/
    async getServerList() {
        const q = await (await this.#get("servers")).json()
        if (q.err) {
            throw this.#parseError(q.err)
        }
        return q.list
    }
}

const api = new ServerApi("http://localhost/api")
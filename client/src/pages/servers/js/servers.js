/*** @type {HTMLUListElement} */
const serversArray = gid("js-serversArray")

// a lot of magic...
function serverRecord(serverName, href, version, online, playersCount) {
    const template = parser.parseFromString(`
    <li class="${online ? "online" : "offline"}">
        <a href="${href}">
            <div>
                <h2>${serverName}</h2>
                <h4>${version}</h4>
            </div>
            <div class="players">${playersCount}</div>
        </a>
    </li>
    `, "text/html").body.children[0]
    /*** @type {HTMLLIElement}*/
    createSmoothAnchor(template.children[0])
    // template.children[0].children[0].children[0].textContent = serverName
    // template.children[0].children[0].children[1].textContent = version
    // template.classList.add(online ? "online" : "offline")
    // template.children[0].children[1].textContent = playersCount
    return template
}
/**
 * @param {string} serverName 
 * @param {string} href
 * @param {string} version
 * @param {boolean} online
 * @param {number} playersCount 
 */
function push(serverName, href, version, online, playersCount) {
    serversArray.append(
        serverRecord(
            serverName,
            href,
            version,
            online,
            playersCount
        )
    )
}
// really hacky thing - should remove it l8r
(async () => {
    try {
        const list = await api.getServerList()
        for (let i = 0; i<list.length;i++) {
            const elem = list[i]
            push(elem.name, `/panel/panel?server=${elem.name}`, "null", true, 0)
        }
    } catch(e) {
        console.log(e)
        serversArray.append("Невозможно получить список серверов с ошибкой", e.toString())
    }
})()
quitOnClick()

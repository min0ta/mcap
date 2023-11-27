/*** @type {HTMLUListElement} */
const serversArray = gid("js-serversArray")



// a lot of magic...
function serverRecord(serverName, href, version, online) {
    /*** @type {HTMLLIElement}*/
    const template = parser.parseFromString(`<li><a href="#"><h2></h2><h4></h4></a></li>`, "text/html").body.children[0]
    template.children[0].href = href
    template.children[0].children[0].textContent = serverName
    template.children[0].children[1].textContent = version
    template.classList.add(online ? "online" : "offline")
    return template
}
/**
 * @param {string} serverName 
 * @param {string} href
 * @param {string} version
 * @param {boolean} online
 */
function push(serverName, href, version, online) {
    serversArray.append(
        serverRecord(
            serverName,
            href,
            version,
            online
        )
    )
}

push("skywars", "/", "1.8.8", false)
push("bedwars", "/", "1.15.2", false)
push("survival", "/", "1.20.2", true)


/*** @type {HTMLUListElement} */
const serversArray = gid("js-serversArray")



// a lot of magic...
function serverRecord(serverName, href, version, online, playersCount) {
    /*** @type {HTMLLIElement}*/
    const template = parser.parseFromString(`
    <li>
        <a href="#">
            <div>
                <h2></h2>
                <h4></h4>
            </div>
            <div class="players"></div>
        </a>
    </li>
    `, "text/html").body.children[0]
    template.children[0].href = href
    template.children[0].children[0].children[0].textContent = serverName
    template.children[0].children[0].children[1].textContent = version
    template.classList.add(online ? "online" : "offline")
    template.children[0].children[1].textContent = playersCount
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

push("skywars", "/", "1.8.8", true, 10)
push("bedwars", "/", "1.15.2", false, 0)
push("survival", "/", "1.20.2", true, 3)


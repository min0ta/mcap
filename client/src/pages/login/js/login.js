/*** @type {HTMLButtonElement}*/
const passwordLookup = gid("js-lookupBtn")

const state = {
    isUiLocked: false
}

/*** @type {HTMLInputElement} */
const passwordInput = gid("js-password")
passwordLookup.addEventListener("click", e => {
    e.preventDefault()
    if (passwordInput.type === "text") {
        passwordInput.type = "password" 
        state.isPasswordLookedUp = false
        passwordLookup.classList.replace("eyex", "eye")
    } else {
        passwordInput.type = "text"
        state.isPasswordLookedUp = true
        passwordLookup.classList.replace("eye", "eyex")
    } 
})

/*** @type {HTMLButtonElement}*/
const submitButton = gcel("login")

/*** @type {HTMLButtonElement}*/
const usernameInput = gid("js-username")
/** @type {HTMLParagraphElement} */
const errorOutput = gid("js-errorOutput")
submitButton.addEventListener("click", async e => {
    if (state.isUiLocked) {
        return
    }
    e.preventDefault()
    const password = passwordInput.value
    const username = usernameInput.value
    lockUI()
    try {
        await api.login(username, password)
    } catch (e) {
        errorOutput.textContent = e.message
    }
    unlockUI()
})

const loadingEl = gid("js-loading")

function lockUI() {
    submitButton.classList.replace("login", "hidden")
    loadingEl.classList.replace("hidden", "loading")
    state.isUiLocked = true
}


function unlockUI() {
    submitButton.classList.replace("hidden", "login")
    loadingEl.classList.replace("loading", "hidden")
    state.isUiLocked = false
}
const navBar = document.getElementsByClassName("nav-link px-2")

function toggle() {
    const url = window.location.href
    for (var i = 0; i < navBar.length; i++) {
        if (navBar[i].href === url) {
            var current = document.getElementsByClassName("bg-dark text-white active")
            if (current.length > 0) {
                current[0].className = current[0].className.replace(" bg-dark text-white active", " link-dark")
            }
            navBar[i].className = navBar[i].className.replace(" link-dark", " bg-dark text-white active")
        }
    }
}

window.onload = toggle;
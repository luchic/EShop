let pagenum = 1
let page_id = 0

function Test_2() {
    alert("test")
}

function ShowDescription(item) {
    if (getComputedStyle(item).getPropertyValue('--valid') == 0) {
        return
    }
    popup_text = document.getElementById("popup_text").innerHTML = item.id
    const overlay = document.getElementById("overlay")
    overlay.style.display = "block"
}

function CloseDescription() {
    const overlay = document.getElementById("overlay")
    overlay.style.display = "none"
}

function SetItemInvalid(id) {
    document.getElementById(id).style.setProperty('--valid') = 0
    document.getElementById(id).style.setProperty('background-color') = rgba(0, 0, 0, 0)
}

function SetItemValid(id) {
    document.getElementById(id).style.setProperty('--valid') = 1
    document.getElementById(id).style.setProperty('background-color') = rgb(204,153,102)
}

function SetItems() {
    
}

function TurnLeftPage() {
    pagenum -= 2
    page_id -= 1
}

function TurnRightPage() {
    pagenum += 2
    page_id += 1
}
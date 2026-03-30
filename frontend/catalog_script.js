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
    document.getElementById(id).style.setProperty('opacity') = 0
}

function SetItemValid(id) {
    document.getElementById(id).style.setProperty('--valid') = 1
    document.getElementById(id).style.setProperty('opacity') = 1
}

function SetItems() {
    if (page_id == 0) {
        document.getElementById('LeftArrow').style.setProperty('opacity', 0)
    }
    else {
        document.getElementById('LeftArrow').style.setProperty('opacity', 1)
    }
    document.getElementById("pagenum_1").innerHTML = pagenum
    document.getElementById("pagenum_2").innerHTML = pagenum + 1
}

function TurnLeftPage() {
    if (page_id == 0) {
        return
    }
    pagenum -= 2
    page_id -= 1
    console.log(pagenum)
    SetItems()
}

function TurnRightPage() {
    pagenum += 2
    page_id += 1
    console.log(pagenum)
    SetItems()
}
function validation() {
    let x = document.getElementById("num").value
    if (!isNaN(x) && x >= 0 && x <= 10)
        document.getElementById('text').innerHTML = "you're right!"
    else
        document.getElementById('text').innerHTML = "you're wrong!"
}
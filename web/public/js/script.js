var div = document.getElementById("addTask");
var select = document.getElementById("milestoneSelect")
var clone = select.cloneNode(true);

function addTaskSelect() {
    div.innerHTML = div.innerHTML + clone;
}

var selectDiv = document.getElementById("milestoneSelectDiv");
var addTaskDiv = document.getElementById("addTask");
var clone = selectDiv.cloneNode(true);

function addTaskSelect() {
    addTaskDiv.appendChild(clone);
}

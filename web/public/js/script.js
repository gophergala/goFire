var selectDiv = document.getElementById("milestoneSelectDiv")
var clone = selectDiv.cloneNode(true);

function addTaskSelect() {
    document.getElementById("addTask").appendChild(clone);
}

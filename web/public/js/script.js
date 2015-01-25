var div = document.getElementById("addTask");
var selectDiv = document.getElementById("milestoneSelectDiv")
var clone = selectDiv.cloneNode(true);

console.log(clone);
function addTaskSelect() {
    $("addTask").append(clone);
}

document.addEventListener("DOMContentLoaded", function () {
var editor = CodeMirror.fromTextArea(document.getElementById("submissioncode"), {
    lineNumbers: true,
    mode:"text/x-c++src",
    value: ""
});
});
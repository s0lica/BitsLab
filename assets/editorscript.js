document.addEventListener("DOMContentLoaded", function () {
  var editor = CodeMirror.fromTextArea(document.getElementById("task_description"), {
      lineNumbers: true,
      mode: {
        name:"markdown" ,
        highlighFormatting: true,
        extraModes:{
          name:"stex",
          delimiters: [["$$", "$$"]],
        }
      },
      value: ""
  });
  var editor2 = CodeMirror.fromTextArea(document.getElementById("submissioncode"), {
      lineNumbers: true,
      mode:"markdown",
      value: ""
  });
});

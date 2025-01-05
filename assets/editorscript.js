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
});

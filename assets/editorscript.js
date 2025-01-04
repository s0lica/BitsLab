const textarea = document.getElementById("expandingTextarea");
const lineNumbers = document.getElementById("lineNumbers");

function updateLineNumbers() {
  const lines = textarea.value.split("\n").length;
  lineNumbers.innerHTML = Array.from({ length: lines }, (_, i) => i + 1).join("<br>");
}

function handleKeyDown(event) {
  if (event.key === "Enter") {
    textarea.style.height = "auto"; 
    textarea.style.height = `${textarea.scrollHeight}px`; 
  }
}

updateLineNumbers();
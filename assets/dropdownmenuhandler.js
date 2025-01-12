function handleselectchange(){
    const inputselect=document.getElementById('inputchoice');
    const outputselect=document.getElementById('outputchoice');
    const inputtext=document.getElementById('inputtext');
    const inputupload=document.getElementById('inputupload');
    const outputtext=document.getElementById('outputtext');
    const outputupload=document.getElementById('outputupload');
    if(inputselect.value=="type"){
        inputtext.classList.remove('hidden');
        inputtext.setAttribute("required", "");
        inputupload.classList.add('hidden');
        inputupload.removeAttribute("required");
    }
    else{
        inputtext.classList.add('hidden');
        inputtext.removeAttribute("required");
        inputupload.classList.remove('hidden');
        inputupload.setAttribute("required","");
    }
    if (outputselect.value=="type"){
        outputtext.classList.remove('hidden');
        outputtext.setAttribute("required","");
        outputupload.classList.add('hidden');
        outputupload.removeAttribute("required");
    }
    else{
        outputtext.classList.add('hidden');
        outputtext.removeAttribute("required");
        outputupload.classList.remove('hidden');
        outputupload.setAttribute("required","");
    }
}
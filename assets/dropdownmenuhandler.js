function handleselectchange(){
    const inputselect=document.getElementById('inputchoice');
    const outputselect=document.getElementById('outputchoice');
    const inputtext=document.getElementById('inputtext');
    const inputupload=document.getElementById('inputupload');
    const outputtext=document.getElementById('outputtext');
    const outputupload=document.getElementById('outputupload');
    if(inputselect.value=="type"){
        inputtext.classList.remove('hidden');
        inputupload.classList.add('hidden');
    }
    else{
        inputtext.classList.add('hidden');
        inputupload.classList.remove('hidden');
    }
    if (outputselect.value=="type"){
        outputtext.classList.remove('hidden');
        outputupload.classList.add('hidden');
    }
    else{
        outputtext.classList.add('hidden');
        outputupload.classList.remove('hidden');
    }
}

import test from "./test.js"
(function(){
    let btn =        document.getElementById("btn")
    btn.addEventListener("click",()=>{
        alert("hello")
        console.log("test",test);
    })

})()
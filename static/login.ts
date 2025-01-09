let username : string = (document.getElementById('username') as HTMLInputElement).value;
let password : string = (document.getElementById('password') as HTMLInputElement).value;


fetch("/Login",{
    method :  "POST",
    headers:{
        "content-type" : "application/json"
    },
    body : JSON.stringify({
        username : username,
        password : password
    })
})
.then (Response => Response.json())
.then(data => {
    if(data.success){
        alert("Login Success");
        window.location.href = "/";
    }
    else{
        alert("Login Failed");
    }
})
.then (error => {
    console.log("Error : ",error);
})
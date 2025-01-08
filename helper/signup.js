let Username = document.getElementById('username');
let Email = document.getElementById('email');
let Password = document.getElementById('password');
let PhoneNo = document.getElementById('phoneNo');
let ErrorMessage = document.getElementById('errorMessage');

// send data to the server
document.getElementById('signupForm').addEventListener('submit', function(event) {
    event.preventDefault();
    signup();
});
function signup() {
    fetch('/Signup', {
        method: 'POST', //sending post request to /Signup
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify({
            username: Username.value,
            email: Email.value,
            password: Password.value,
            phoneNo: PhoneNo.value
        })
    })
    .then(response => response.json())
    .then(data => {
        if (data.success) {
            // Handle success case
            document.getElementById('successMessage').innerText = "Signup Successful";
            document.getElementById('successMessage').style.display = 'block';
        } else {
            ErrorMessage.textContent = "Signup Failed: " + data.message;
            ErrorMessage.style.display = 'block';
            alert("Signup Failed: " + data.message);
        }
    })
    .catch(error => {
        console.error('Error:', error);
        sendErrorToLoggingService(error);
        alert("Signup Failed: " + error.message);
    });
}

function sendErrorToLoggingService(error) {
    fetch('/logError', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify({
            error: error.message,
            stack: error.stack
        })
    })
    .then(response => response.json())
    .then(data => {
        if (!data.success) {
            console.error('Failed to log error:', data.message);
        }
    })
    .catch(logError => {
        console.error('Error logging the error:', logError);
    });
}
        alert("Signup Failed: " + error.message);


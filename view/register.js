document.getElementById("registerForm").addEventListener("submit", function (event) {
    event.preventDefault();
    const id = document.getElementById("userId").value;
    const password = document.getElementById("password").value;
    const data = {
        id: id,
        password: password
    };
    fetch("/app/user/register", {
        method: "POST",
        headers: {
            "Content-Type": "application/json"
        },
        body: JSON.stringify(data)
    });
});
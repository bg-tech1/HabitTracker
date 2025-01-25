document.getElementById("loginForm").addEventListener("submit", function (event) {
    event.preventDefault();
    const userId = document.getElementById("userId").value;
    const password = document.getElementById("password").value;
    const errorContainer = document.getElementById('errorContainer');
    const data = {
        id: userId,
        password: password
    };
    console.log(data);
    fetch("/app/user/login", {
        method: "POST",
        headers: {
            "Content-Type": "application/json"
        },
        body: JSON.stringify(data)
    })
        .then(response => response.json())
        .then(data => {
            if (data.redirect) {
                window.location.href = data.redirect;
            } else {
                errorContainer.textContent = data.error || 'ユーザー名もしくはパスワードが違います';
                errorContainer.style.display = 'block';
            }
        })
        .then(data => {
            window.location.href = "/view/dashboard.html";
        })
        .catch(error => {
            console.error('Error:', error);
            errorContainer.textContent = 'ログイン処理中にエラーが発生しました。';
            errorContainer.style.display = 'block';
        });
});

document.getElementById("register").addEventListener("click", function (event) {
    event.preventDefault();
    window.location.href = "/view/register.html";
});
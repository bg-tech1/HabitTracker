document.addEventListener('DOMContentLoaded', () => {

    const habitInput = document.getElementById('habit-input');
    const addHabitBtn = document.getElementById('add-habit-btn');
    const habitList = document.getElementById('habit-list');
    const habitModal = document.getElementById('habit-modal');
    const openModalBtn = document.getElementById('open-modal');
    const closeModalBtn = document.getElementById('close-modal');
    const logModal = document.getElementById('log-modal');
    const logDateInput = document.getElementById('log-date-input');
    const logDateBtn = document.getElementById('log-date-btn');
    const closeLogModalBtn = document.getElementById('close-log-modal');

    let selectedHabit = null; // To track the habit being logged

    fetch("/app/habit/dashboard")
        .then(response => response.json())
        .then(data => {
            console.log(data);
            data.habit.forEach(habit => {
                console.log(habit);
                const habitCard = document.createElement('div');
                habitCard.className = 'habit-card';
                const habitName = habit.habit_name;
                const sessionID = getCookie("session_id");
                console.log("sessionIDは", sessionID);
                habitCard.innerHTML = `<h3>${habitName}</h3>
            <div class="d-flex justify-content-between align-items-center">
                <span class="badge bg-secondary">未記録</span>
                <div>
                    <button class="btn btn-primary btn-sm me-2 log-date-btn">日時を記録</button>
                </div>
            </div>`;
                habitCard.querySelector('.log-date-btn').addEventListener('click', () => {
                    selectedHabit = habitName; // Store the habit name
                    logModal.style.display = 'flex';
                });
                habitList.appendChild(habitCard);
            });
        })
        .catch(error => {
            console.error('Error:', error);
        });

    openModalBtn.addEventListener('click', () => {
        habitModal.style.display = 'flex';
    });

    closeModalBtn.addEventListener('click', () => {
        habitModal.style.display = 'none';
    });

    closeLogModalBtn.addEventListener('click', () => {
        logModal.style.display = 'none';
    });

    addHabitBtn.addEventListener('click', () => {
        const habitName = habitInput.value.trim();

        if (habitName === '') {
            alert('習慣を入力してください');
            return;
        }

        fetch("/app/habit/create", {
            method: "POST",
            headers: {
                "Content-Type": "application/json"
            },
            body: JSON.stringify({ habit_name: habitName })
        })
            .then(response => response.json())
            .then(data => {
                console.log(data);
                const habitCard = document.createElement('div');
                habitCard.className = 'habit-card';

                habitCard.innerHTML = `
            <h3>${habitName}</h3>
            <div class="d-flex justify-content-between align-items-center">
                <span class="badge bg-secondary">未記録</span>
                <div>
                    <button class="btn btn-primary btn-sm me-2 log-date-btn">日時を記録</button>
                </div>
            </div>`;

                habitCard.querySelector('.log-date-btn').addEventListener('click', () => {
                    selectedHabit = habitName; // Store the habit name
                    logModal.style.display = 'flex';
                });

                habitList.appendChild(habitCard);
                habitInput.value = '';
                habitModal.style.display = 'none';
            })
            .catch(error => {
                console.error('Error:', error);
            });

    });

    logDateBtn.addEventListener('click', async () => {
        const selectedDate = logDateInput.value;

        if (!selectedDate || !selectedHabit) {
            alert('日付を入力してください');
            return;
        }

        try {
            const response = await fetch('/api/habits/log', {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json'
                },
                body: JSON.stringify({ habit: selectedHabit, date: selectedDate })
            });

            if (!response.ok) {
                throw new Error('Failed to log the habit execution date');
            }

            alert('記録が成功しました');
            logModal.style.display = 'none';
            logDateInput.value = '';
        } catch (error) {
            alert('エラーが発生しました: ' + error.message);
        }
    });
});
function getCookie(name) {
    const value = `; ${document.cookie}`;
    console.log("cookieは", value);
    const parts = value.split(`; ${name}=`);
    if (parts.length === 2) return parts.pop().split(';').shift();
};
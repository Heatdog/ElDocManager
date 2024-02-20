import React, { Component } from 'react'

export class Login extends Component {

    handleSubmit() {
        let loginForm = {
            login: this.state.login,
            password: this.state.password
        }
        var data = JSON.stringify(loginForm);
        let xhr = new XMLHttpRequest();
        xhr.onload = () => {
            if (xhr.status === 200){
                const token = JSON.parse(xhr.responseText);
                alert(token.token);
            }
        };
        alert(data);
        xhr.open("POST", "http://localhost:8080/api/login", true);
        xhr.setRequestHeader("Content-Type", "application/json");
        xhr.send(data);
    }

    fetchSubmit() {
        const loginForm = document.getElementById("loginForm");
        loginForm.addEventListener('submit', async (e) => {
            e.preventDefault();
            const formData = new FormData(loginForm);
            
            let loginData = {
                login: formData.get("login"),
                password: formData.get("password")
            }
            var data = JSON.stringify(loginData);
            let respone = await fetch("http://localhost:8080/api/login", {
                method: "POST",
                body: data,
            });
    
            let result = await respone.json();
            alert(result.token);
        });
    }

  render() {
    return (
        <form id="loginForm" className='login'>
            <h3>Форма входа</h3>
            <p>
                <label>Логин</label><br />
                <input type="text" name='login'  />
            </p>
            <p>
                <label>Пароль:</label><br />
                <input type="text" name='password' />
            </p>
            <input type="submit" value="Отправить" onClick={this.fetchSubmit}/>
        </form>
    );
  }
}

export default Login
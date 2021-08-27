# xxs-cors-example
xxs-cors-example

# How to run a attacker server
- ```bash
    # Open new terminal 
    cd attacker_server
    go run main.go
    ```

# How to run a victim server
- ```bash
    # Open new terminal
    cd victim_server
    go run main.go
    ```
# !!Caution
- Input your kakao-app-js-key to victim_server/victim-community-kakao.html
- victim_server/victim-community-kakao.html
    - ```html
        ... 
        document.addEventListener("DOMContentLoaded", function () {
        loginStatusElem = document.querySelector("#login-status");

        // 카카오 javascript key 입력
        Kakao.init("your-kakao-app-javascript-key");
        Kakao.isInitialized();
        ... 
        ```

# Attack Scenario 1
- Run both servers
- Open a chrome browser
- Access 'http://localhost:8888/victim/community'
- Check attacker server shell log.
- if the shell log shows Authorization and Cookie, the attack is successful.

# Attack Scenario 2
- Run both servers
- Open a chrome browser
- Access 'http://localhost:8888/victim/community/kakao'
- Login kakao
- Click the button named '☆★☆★누르면 놀라운 일이 벌어짐☆★☆★'
- Check attacker server shell log.
- if the shell log shows Authorization and Cookie, the attack is successful.

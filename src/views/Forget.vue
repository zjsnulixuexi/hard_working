<template>
    <div class="all">
        <br><br><br><br><br><br><br>
        <div class="register-container">
            <form @submit.prevent="submitForm">
                <p>找回密码</p>
                <div class="form-group">
                    <label for="username">UserName:</label>
                    <input type="text" id="username" v-model="registerForm.username" required>
                </div>


                <div class="form-group captcha-group">
                    <label for="captcha">Captcha:</label>

                    <input type="text" id="captcha" v-model="registerForm.captcha" required>
                    <button type="button" @click="getCaptcha"
                        style="width: 120px;height: 40px;border-radius: 4px;">获取验证码</button>
                </div>
                <button type="submit" class="submit-btn">提交</button>
            </form>
        </div>
    </div>
</template>

<script scoped>
import axios from 'axios'
export default {
    name: 'RegisterForm',
    data() {
        return {
            registerForm: {
                username: '',
                captcha: ''
            },
            generate_code: '',
            password: '还没给呢，急什么'
            //这个密码，需要在mounted的时候，给传进来
        };
    },
    methods: {
        submitForm() {

            const data = {
                "username": this.registerForm.username,
                "operate": "Password_recovery",
                "operate_type": "Operate_users"
            };
            axios.post('/api/backend', data)
                .then(response => {
                    // this.password=response.data.password
                    console.log(this.registerForm);
                    console.log(response.data);
                    if (response.data.password != "用户不存在") {
                        this.password = response.data.password
                        alert('您的密码为:' + this.password)
                        this.$router.push("/");
                    }
                    else {
                        alert('用户名或验证码不对，请重新填写')
                    }
                })
                .catch(error => {
                    console.error(error);
                });
        },
        getCaptcha() {
            let code = ''
            for (let i = 0; i < 4; i++) {
                // 生成0到9之间的随机数
                const randomNum = Math.floor(Math.random() * 10);
                // 将随机数添加到验证码字符串中
                code += randomNum;
            }
            // 更新this.code为最后一个生成的随机数
            this.code = code.slice(-1); // 取最后一个字符
            this.generate_code = code
            console.log(this.generate_code);
        }
    }
};
</script>

<style scoped>
html body {
    margin: 0;
    border: 0;
}

.all {
    background-color: rgb(144, 157, 157);
    width: 100vw;
    height: 100vh;
}

.register-container {
    max-width: 500px;
    margin: auto;
    padding: 20px;
    border: 1px solid #ccc;
    border-radius: 5px;
    background-color: white
}

p {
    margin: 0;
    text-align: center;
    font-size: 30px;
    font-weight: bold;
}

.form-group {
    margin-bottom: 15px;
}

label {
    display: block;
    margin-bottom: 5px;
}

input[type="text"],
input[type="password"],
input[type="email"],
select {
    width: 100%;
    padding: 10px;
    border: 1px solid #ccc;
    border-radius: 4px;
}

.captcha-group {
    display: flex;
    align-items: center;
}

.captcha-group button {
    margin-left: 10px;
}

.submit-btn {
    padding: 10px 20px;
    background-color: #4CAF50;
    color: white;
    border: none;
    border-radius: 4px;
    cursor: pointer;
    margin: 0 auto;
    display: block;
}

.submit-btn:hover {
    background-color: #45a049;
}
</style>
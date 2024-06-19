<template>
    <div class="all">
        <br><br><br><br><br><br><br>
        <div class="register-container">
            <form @submit.prevent="submitForm">
                <p>注册</p>
                <div class="form-group">
                    <label for="username">Username:</label>
                    <input type="text" id="username" v-model="registerForm.username" required>
                </div>
                <div class="form-group">
                    <label for="password">Password:</label>
                    <input type="password" id="password" v-model="registerForm.password" required>
                </div>
                <div class="form-group">
                    <label for="school">School:</label>
                    <input type="text" id="school" v-model="registerForm.school" required>
                </div>
                <div class="form-group">
                    <label for="sex">Sex:</label>
                    <select id="sex" v-model="registerForm.sex">
                        <option value="">Select Sex</option>
                        <option value="男">Male</option>
                        <option value="女">Female</option>
                    </select>
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
                password: '',
                username: '',
                school: '',
                sex: '',
                captcha: ''
            },
            generate_code: ''
        };
    },

    methods: {
        submitForm() {
            if (this.generate_code == this.registerForm.captcha) {
                console.log('注册成功');
                console.log(this.registerForm);

                const data = {
                    "password": this.registerForm.password,
                    "username": this.registerForm.username,
                    "school": this.registerForm.school,
                    "sex": this.registerForm.sex,
                    "operate": "User_registration",
                    "operate_type": "Operate_users"
                };
                axios.post('/api/backend', data)
                    .then(response => {
                        console.log(response.data);
                        if (response.data.userid != -1) {
                            alert('注册成功')
                            this.$router.push("/");
                        }
                        else {
                            alert('不能注册已有的账户名，请重新注册')
                        }
                    })
                    .catch(error => {
                        console.error(error);
                    });

            }
            else {
                alert('验证码不对，请重新填写');
            }
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
    user-select: none;
    background-color: white;

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
<template>
    <!-- 背景图片 -->
    <div class="background-image"></div>

    <div class="login-box">
        <!-- Login标题 -->
        <h4 class="login-title" style="font-size: 30px;
    color:#fefefe;
    text-shadow:0px 1px 0px #c0c0c0,
       0px 2px 0px #b0b0b0,
       0px 3px 0px #a0a0a0,
       0px 4px 0px #909090,
       0px 5px 10px rgba(0, 0, 0, .9);">Login</h4>

        <!-- 用户名输入框 -->
        <div class="form-group">
            <img src="../assets/账户.png" alt="" class='small_img'>
            <input type="text" id="username" v-model="formData.username" placeholder="Enter your username">
        </div>

        <!-- 密码输入框 -->
        <div class="form-group">
            <img src="../assets/密码.png" alt="" class='small_img'>
            <input type="password" id="password" v-model="formData.password" @keyup.enter="login"
                placeholder="Enter your password">
        </div>

        <!-- 登录按钮 -->
        <button @click="login" class="login-btn">Login</button>

        <!-- 注册链接 -->
        <div class="bottom">
            <div class="register" @click="to_register" style="cursor: pointer;">
                <p>Register?</p>
            </div>
            <div class="forget" @click="to_forget" style="cursor: pointer;">
                <p>Forget?</p>
            </div>
        </div>
    </div>
</template>

<script scoped>
import axios from 'axios'
export default {

    data() {
        return {
            formData: {
                username: '',
                password: '',
            },
        };
    },
    methods: {
        login() {
            const data = {
                "password": this.formData.password,
                "username": this.formData.username,
                "operate": "User_login",
                "operate_type": "Operate_users"
            };
            if (this.formData.password && this.formData.username) {
                axios.post('/api/backend', data)
                    .then(response => {
                        if (response.data.power == 2) {
                            console.log(1);
                            this.$cookies.set("userid", response.data.userid);
                            this.$router.push("/home");
                            console.log('login of response.data', response.data);

                        }
                        else if (response.data.power == -1) {
                            alert('您的账号或密码有误，请重新输入')

                        }
                        else {
                            console.log(2);
                            console.log("Login of response.data", response.data);
                            this.$cookies.set("userid", response.data.userid);
                            this.$router.push("/manager");
                        }
                    })
                    .catch(error => {
                        console.error(error);
                    });
            }
            else {
                alert('请输入完整的账号或密码')
            }
        },
        to_register() {
            this.$router.push("/register");
        },
        to_forget() {
            this.$router.push("/forget");
        }
    }
};
</script>

<style scoped>
.background-image {
    position: absolute;
    top: 0;
    left: 0;
    width: 100%;
    height: 100%;
    background-image: url('../src/assets/背景图.jpg');
    /* 使用您的背景图片路径 */
    background-size: cover;
    background-position: center;
}

.login-box {
    position: absolute;
    top: 50%;
    left: 50%;
    transform: translate(-50%, -50%);
    width: 400px;
    padding: 20px;
    border-radius: 5px;
    background-color: rgba(255, 255, 255, 0.5);
    /* 登录框背景颜色及透明度 */
}

.small_img {
    width: 37px;
    /* 调整图像的宽度 */
    vertical-align: middle;
    padding-right: 5px;
    opacity: 0.2;
    border-radius: 2px;

}

.form-group {
    margin-bottom: 20px;
}

input[type="text"],
input[type="password"] {
    width: calc(100% - 70px);
    /* 减去图像的宽度和两侧的空白间距 */
    padding: 10px;
    border: 1px solid #ccc;
    border-radius: 5px;
    background-color: rgba(255, 255, 255, 0.7);
    vertical-align: middle;
}

.login-title {
    font-family: 'Arial Black', sans-serif;
    /* 设置艺术字体 */
    text-align: center;
}

.form-group {
    margin-bottom: 20px;
}

label {
    display: block;
    margin-bottom: 5px;
}


.login-btn {
    width: 100%;
    padding: 10px;
    border: none;
    border-radius: 5px;
    background-color: #ff7474;
    color: #fff;
    cursor: pointer;
    border-radius: 10px;
}

.login-btn:hover {
    background-color: #f79825;
}

p {
    text-align: center;
    margin: 0;
}


.bottom {
    margin-top: 20px;
}

.register {
    display: inline-block;
    margin-left: 65px;
    margin-right: 65px;
}

.forget {
    display: inline-block;
    margin-left: 65px;
    margin-right: 65px;
}
</style>
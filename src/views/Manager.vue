<template>
    <div style="user-select: none;">
        <p style="margin: 0;font-size: 45px;font-weight: 300px;text-align: center;">欢迎来到管理员界面</p>
    </div>
    <hr>
    <div style="user-select: none;">
        <select v-model="selectedTable" @change="changeTable" style="width: 200px; cursor: pointer;
        height: 50px;font-size: 20px;font-weight:bold; margin-bottom: 20px;background-color: #ccc;">
            <option value="table1">用户信息</option>
            <option value="table2">可视化操作</option>
        </select>
        <button @click="toggleManageMode(true)"
            style="margin-left: 1050px;font-size: 20px;border-radius: 4px;margin-right: 4px;cursor: pointer;">管理</button>
        <button @click="showModal1()" v-if="selectedTable === 'table2'"
            style="font-size: 20px;border-radius: 4px;margin-right: 4px;cursor: pointer;">添加</button>
        <button @click="toggleManageMode(false)"
            style="font-size: 20px;border-radius: 4px;margin-right: 4px;cursor: pointer;">完成</button>
        <div v-if="selectedTable === 'table1'">
            <table>
                <thead>
                    <tr>
                        <th>Username</th>
                        <th>Password</th>
                        <th>School</th>
                        <th>Sex</th>
                        <th v-if="manageMode">操作</th>
                        <th v-if="manageMode">操作</th>
                    </tr>
                </thead>
                <tbody>
                    <tr v-for="(item, index) in table1Data" :key="item.userid">
                        <td>{{ item.username }}</td>
                        <td>{{ item.password }}</td>
                        <td>{{ item.school }}</td>
                        <td>{{ item.sex }}</td>

                        <td v-if="manageMode" @click="add('table1', index)" class="add">
                            权限设置
                        </td>
                        <td v-if="manageMode" @click="removeItem('table1', index)" class="delete">
                            删除
                        </td>
                    </tr>
                </tbody>
            </table>
        </div>
        <div v-else>
            <table>
                <thead>
                    <tr>
                        <th>Code name</th>
                        <th>Code Instruction</th>
                        <th v-if="manageMode">操作</th>
                        <th v-if="manageMode">操作</th>
                    </tr>
                </thead>
                <tbody>
                    <tr v-for="(item, index) in table2Data" :key="item.userid">
                        <td>{{ item.operationname }}</td>
                        <td>{{ item.operationintroduce }}</td>
                        <td v-if="manageMode" @click="add('table2', index)" class="add">
                            状态管理
                        </td>
                        <td v-if="manageMode" @click="removeItem('table2', index)" class="delete">
                            删除
                        </td>
                    </tr>
                </tbody>
            </table>
        </div>
    </div>
    <div v-if="showModal" class="modal">
        <div class="modal-content">
            <span class="close" @click="showModal = false">&times;</span>
            <p style="font-size: 30px; text-align: center;">{{ deadline }}</p>
            <form @submit.prevent="handleAddVisualization(deadline, index)">
                <label for="name" style="font-size: 20px;">名称:</label>
                <input type="text" id="name" v-model="new_name" required style="width: 100%;height: 20px;">
                <br>
                <br>
                <label for="description" style="font-size: 20px;">简介:</label>
                <input type="text" id="description" v-model="new_introduction" style="width: 100%;height: 20px;">
                <br>
                <br>
                <input type="file" id="file" ref="fileInput" v-if="deadline == '添加可视化代码'" @change="Codefile">
                <br>
                <button type="submit" style="padding: auto;">提交</button>
            </form>
        </div>
    </div>
</template>

<script scoped>
import axios from 'axios'
export default {
    data() {
        return {
            selectedTable: 'table1',
            title1: '',
            showModal: false,
            new_name: '',
            new_introduction: '',
            table1Data: [
                // 更多数据...
            ],
            table2Data: [
            ],
            manageMode: false,
            targetUser: '',
            index: '',
            codefile: '',
        };
    },
    mounted() {
        let userid = Number(this.$cookies.get("userid"));
        const data = {
            "userid": userid,
            "operate": "User_query_all",
            "operate_type": "Operate_users"
        };
        axios.post('/api/backend', data)
            .then(response => {
                console.log('Manager of response.data:', response.data);
                this.table1Data = response.data.data
                this.targetUser = this.table1Data.find(user => user.userid === userid)
                if (this.targetUser) {
                    // 获取该用户的power值
                    let targetPower = this.targetUser.power;
                    // 过滤数组，删除那些power值小于等于targetPower的对象
                    this.table1Data = this.table1Data.filter(user => user.power > targetPower);
                }
            })
            .catch(error => {
                console.error(error);
            });

        const data2 = {
            "operate": "Visualoperation_query_all",
            "operate_type": "Operate_visualoperations"
        };
        axios.post('/api/backend', data2)
            .then(response => {
                this.table2Data = response.data.data
                console.log(this.table2Data);
            })
            .catch(error => {
                console.error(error);
            });
    },
    methods: {
        changeTable() {
            // 可以在这里添加切换表格时需要执行的逻辑
        },
        toggleManageMode(mode) {
            this.manageMode = mode;
        },
        async removeItem(table, index) {
            if (table == 'table1') {
                if (confirm('确定要删除用户 ' + this.table1Data[index].username + ' 吗')) {
                    const data = {
                        "userid": this.table1Data[index].userid,
                        "operate": "Manage_user_delete",
                        "operate_type": "Operate_users"
                    };
                    console.log(this.table1Data[index].userid);
                    try {
                        const response = await axios.post('/api/backend', data);
                        await this.update();
                        console.log('response.data', response.data);
                    } catch (error) {
                        console.error(error);
                    }
                }
            } else if (table == 'table2') {
                if (confirm('确定要删除 ' + this.table2Data[index].operationname + ' 吗')) {
                    const data = {
                        "operationid": this.table2Data[index].operationid,
                        "operate": "Visualoperation_delete",
                        "operate_type": "Operate_visualoperations"
                    };
                    axios.post('/api/backend', data)
                        .then(response => {
                            console.log(response.data);
                            this.update();
                        })
                        .catch(error => {
                            console.error(error);
                        });
                }
            }
        },
        async add(table, index) {
            if (table == 'table1') {
                console.log(this.table1Data[index]);
                const data = {
                    "userid": this.table1Data[index].userid,
                    "power": this.table1Data[index].power,
                    "managepower": this.targetUser.power,
                    "operate": "User_power_update",
                    "operate_type": "Operate_users",
                };
                try {
                    const response = await axios.post('/api/backend', data);
                    await this.update();
                    alert(response.data.result)
                } catch (error) {
                    console.error(error);
                }
            }
            else {
                console.log(this.table2Data[index]);
                this.showModal = true
                this.deadline = '重命名可视化'
                this.index = index


            }
        },
        update() {
            let userid = Number(this.$cookies.get("userid"));
            const data = {
                "userid": userid,
                "operate": "User_query_all",
                "operate_type": "Operate_users"
            };
            axios.post('/api/backend', data)
                .then(response => {
                    console.log('Manager of response.data:', response.data);
                    this.table1Data = response.data.data
                    this.targetUser = this.table1Data.find(user => user.userid === userid)
                    if (this.targetUser) {
                        // 获取该用户的power值
                        let targetPower = this.targetUser.power;
                        // 过滤数组，删除那些power值小于等于targetPower的对象
                        this.table1Data = this.table1Data.filter(user => user.power > targetPower);
                        console.log('this.table1Data', this.table1Data);
                    }
                })
                .catch(error => {
                    console.error(error);
                });
            const data2 = {
                "operate": "Visualoperation_query_all",
                "operate_type": "Operate_visualoperations"
            };
            axios.post('/api/backend', data2)
                .then(response => {
                    this.table2Data = response.data.data
                    console.log(this.table2Data);
                })
                .catch(error => {
                    console.error(error);
                });
        },
        handleAddVisualization(type, index) {
            if (type == '添加可视化代码') {
                const data = {
                    "operationname": this.new_name,
                    "operationintroduce": this.new_introduction,
                    "file": this.codefile,
                    "operate": "Visualoperation_create",
                    "operate_type": "Operate_visualoperations"
                };
                console.log(data);
                axios.post('/api/backend', data)
                    .then(response => {
                        console.log(response.data);
                        this.update();
                    })
                    .catch(error => {
                        console.error(error);
                    });
                // 关闭模态框
                this.showModal = false;
            }
            else if (type == "重命名可视化") {
                const data2 = {
                    "operationid": this.table2Data[index].operationid,
                    "operationname": this.new_name,
                    "operationintroduce": this.new_introduction,
                    "operate": "Visualoperation_update",
                    "operate_type": "Operate_visualoperations"
                };
                console.log(data2);
                axios.post('/api/backend', data2)
                    .then(response => {
                        this.update();
                    })
                    .catch(error => {
                        console.error(error);
                    });
                this.showModal = false
            }
        },
        Codefile(e) {
            const file = e.target.files[0];
            // 获取文件名
            const fileName = file.name;
            // 获取文件扩展名
            const fileExtension = fileName.slice(((fileName.lastIndexOf(".") - 1) >>> 0) + 2);

            // 检查文件扩展名是否为 'py'
            if (fileExtension.toLowerCase() === 'py') {
                const reader = new FileReader();
                reader.onload = (e) => {
                    // 如果是Python文件，读取文件内容
                    this.codefile = e.target.result;
                    // 这里可以添加其他处理逻辑
                };
                reader.readAsText(file);
            } else {
                // 如果不是Python文件，可以提示用户或执行其他操作
                alert('请上传Python文件（.py）.');
                // 清除input file的值，以便用户可以重新选择
                e.target.value = '';
            }
        },
        showModal1() {
            this.showModal = true
            this.deadline = '添加可视化代码'
            this.index = -1
        }
    },
};
</script>

<style scoped>
/* 基础样式 */
body {
    font-family: 'Arial', sans-serif;
    background-color: #f5f5f5;
    margin: 0;
    padding: 20px;
    box-sizing: border-box;
}

/* 标题样式 */
.title {
    text-align: center;
    color: #333;
    margin-bottom: 20px;
}

/* 选择框和按钮样式 */
select,
button {
    padding: 10px;
    margin: 5px;
    border: none;
    border-radius: 4px;
    cursor: pointer;
    font-size: 16px;
    transition: background-color 0.3s ease;
}

select {
    background-color: #fff;
    color: #333;
}

button {
    background-color: #5cb85c;
    color: white;
    margin-left: 10px;
    /* 调整按钮间距 */
}

button:hover {
    background-color: #4cae4c;
}

/* 表格样式 */
table {
    width: 100%;
    border-collapse: collapse;
    margin-top: 20px;
}

th,
td {
    border: 1px solid #ddd;
    padding: 15px;
    text-align: left;
}

th {
    background-color: #4CAF50;
    color: white;
}

td {
    background-color: #fff;
}

.add {
    cursor: pointer;
    background-color: #325438;
    color: white;
    border: none;
    padding: 5px 10px;
    border-radius: 3px;
}

.add:hover {
    background-color: #7cca2e;
}

/* 删除按钮样式 */
.delete {
    cursor: pointer;
    background-color: #f44336;
    color: white;
    border: none;
    padding: 5px 10px;
    border-radius: 3px;
}

.delete:hover {
    background-color: #e53935;
}

/* 响应式设计 */
@media (max-width: 768px) {
    .title {
        font-size: 30px;
    }

    select,
    button {
        width: calc(50% - 10px);
        margin: 5px;
    }

    button {
        margin-left: 0;
    }
}

.modal {
    display: block;
    position: fixed;
    z-index: 1;
    left: 0;
    top: 0;
    width: 100%;
    height: 100%;
    overflow: auto;
    background-color: rgb(0, 0, 0);
    background-color: rgba(0, 0, 0, 0.4);
}

.modal-content {
    background-color: #fefefe;
    margin: 15% auto;
    padding: 20px;
    border: 1px solid #888;
    width: 80%;
    max-width: 500px;
}

.close {
    color: #aaa;
    float: right;
    font-size: 28px;
    font-weight: bold;
    cursor: pointer;
}

.close:hover,
.close:focus {
    color: black;
    text-decoration: none;
    cursor: pointer;
}
</style>
<template>
    <div class="overlay" v-if="showOverlay">
        <div class="overlay-content">
            <p>请导入XLSX文件</p>
            <input type="file" accept=".xlsx" @change="checkFileType" />
        </div>
    </div>
    <div class="top">
        <p>{{ title }}</p>
    </div>
    <div class="all">

        <div class="left single">
            <div class="title">
                <p>文件名</p>
            </div>
            <div class="contents">
                <div v-for="(file, index) in files" :key="index" class="file-item" @dblclick="show_content(file)"
                    @mouseover="showDeleteIcon(index)" @mouseout="hideDeleteIcon(index)">
                    {{ file.resultname }} <!-- 假设每个文件对象都有一个name属性 -->
                    <span v-show="deleteIcons[index]" class="delete-icon" @click.stop="deleteFile(index)">删除</span>
                </div>
            </div>
        </div>

        <div class="middle single">
            <div class="title">
                <p>
                    可视化操作
                </p>
            </div>
            <div class="contents">
                <div v-for="item in algorithm" :key="item.title" class="block" @dblclick="operate(item)">
                    {{ item.operationname }}
                </div>
            </div>
        </div>

        <div class="right single">
            <img v-if="show" src="../assets/1.png" alt="">
            <div v-if="!show">
                <div class="r_top">
                    {{ r_top_title }}
                </div>
                <div class="r_content">
                    <!-- <div ref="chart" style="width: 600px;height:400px;"></div> -->
                    <pre style="font-size: 20px;padding-left: 40px; ;">{{ r_content }}</pre> <!-- 显示原始统计数据 -->
                    <!-- {{ r_content }} -->
                </div>
            </div>
        </div>
    </div>
</template>

<script scoped>
import * as XLSX from 'xlsx';
import axios from 'axios';
import * as echarts from 'echarts';
export default {
    data() {
        return {
            show: 1,
            file: {},
            title: '',
            r_top_title: '',
            r_content: '',
            showOverlay: false,
            algorithm: [],
            deleteIcons: [],
            files: [
                // 假设的文件数据结构
                { name: '文件1.txt', content: '文件1', resultid: 1 },
                { name: '文件2.txt', content: '文件2', resultid: 2 },
                { name: '文件3.txt', content: '文件3', resultid: 3 },
                // 这个content，在处理存储文件阶段，就得存储进来，现在的是固定的content
            ],
            tableHeader: [],
            tableData: [],
            json_head: '',
            json_data: '',

        }
    },
    mounted() {
        this.file = this.$cookies.get("fastpaper-document-info");
        console.log(this.file);
        this.title = this.file.filename

        let path = this.file.path
        if (!path) {
            this.showOverlay = true;
        }

        const data = {
            "operate": "Visualoperation_query_all",
            "operate_type": "Operate_visualoperations"
        };
        axios.post('/api/backend', data)
            .then(response => {
                console.log(response.data);
                this.algorithm = response.data.data
            })
            .catch(error => {
                console.error(error);
            });

        const data2 = {
            "operate": "result_query_all",
            "operate_type": "Operate_results",
            "userid": this.file.userid,
            "documentid": this.file.documentid
        };
        axios.post('/api/backend', data2)
            .then(response => {
                console.log('response.data', response.data.data);
                this.files = response.data.data
                console.log();
            })
            .catch(error => {
                console.error(error);
            });
    },
    methods: {
        operate(document) {
            // 根据action执行不同的逻辑
            console.log(document.operationname);
            console.log(document.operationid);

            const data = {
                "operate": "result_create",
                "operate_type": "Operate_results",
                "userid": this.file.userid,
                "documentid": this.file.documentid,
                "kind": document.operationid,
                "resultname": document.operationname,
            };
            console.log(data);
            axios.post('/api/backend', data)
                .then(response => {
                    console.log(response.data);
                    this.updatefiles();
                    if (response.data.resultid == -1) {
                        alert('文件' + response.data.Result)
                    }
                })
                .catch(error => {
                    console.error(error);
                });
        },
        show_content(content) {
            this.show = 0
            console.log(content);
            this.r_top_title = content.resultname
            this.r_content = content.content
            this.initChart(this.r_content);
        },
        initChart(content) {
            // 解析内容并生成图表数据
            const chartData = this.parseRContent(content);
            // 检查图表容器是否存在
            if (this.$refs.chart) {
                const myChart = echarts.init(this.$refs.chart);
                const option = {
                    title: { text: '属性统计图表' },
                    tooltip: {},
                    xAxis: { data: chartData.categories },
                    yAxis: {},
                    series: chartData.series,
                };
                myChart.setOption(option);
            }
        },

        parseRContent(rContent) {
            // 根据rContent的格式解析数据
            // 这里需要根据实际内容格式编写解析逻辑
            // 示例：
            const lines = rContent.split('\n');
            let categories = [];
            let data = [];
            lines.forEach(line => {
                if (line.startsWith('属性:')) {
                    categories.push(line.split(':')[1].trim());
                } else if (line.startsWith('  最大值:')) {
                    data.push(parseFloat(line.split(':')[1].trim()));
                }
            });
            return {
                categories,
                series: [{
                    type: 'bar',
                    data,
                    name: '最大值',
                }],
            };
        },
        checkFileType(e) {
            const file = e.target.files[0];
            console.log(file);
            if (!file) {
                console.log('No file selected.');
                return;
            }
            const reader = new FileReader();
            reader.onload = (e) => {
                // 转换ArrayBuffer为blob再转为DataURL
                const data = e.target.result;
                const arrayBufferView = new Uint8Array(data);
                const blob = new Blob([arrayBufferView], { type: 'application/vnd.openxmlformats-officedocument.spreadsheetml.sheet' });
                const url = URL.createObjectURL(blob);

                // 使用fetch API获取blob的内容，然后使用XLSX库解析
                fetch(url)
                    .then(response => response.blob())
                    .then(blob => {
                        const reader = new FileReader();
                        reader.onload = (e) => {
                            const data = e.target.result;
                            const workbook = XLSX.read(data, { type: 'binary' });
                            const firstSheetName = workbook.SheetNames[0];
                            const worksheet = workbook.Sheets[firstSheetName];
                            const json = XLSX.utils.sheet_to_json(worksheet, { header: 1 });

                            this.tableHeader = json[0];
                            this.tableData = json.slice(1); // 跳过标题行

                            this.json_head = JSON.stringify(this.tableHeader)
                            this.json_data = JSON.stringify(this.tableData);
                            //此时输出类型为字符串

                            this.update()
                        };
                        reader.readAsBinaryString(blob);
                    })
                    .catch(error => console.error('Error parsing file:', error));
            };
            reader.readAsArrayBuffer(file); // 正确使用readAsArrayBuffer
            this.showOverlay = false;

        },
        async update() {
            const data = {
                "userID": this.file.userid,
                "documentID": this.file.documentid,
                "filename": this.file.filename,
                "path": this.json_data,
                "operate": "Document_txt",
                "operate_type": "Operate_documents"
            };
            console.log('data type', typeof (data.path));

            console.log('data', data.path);
            try {
                const response = await axios.post('/api/backend', data);
                this.file.path = this.json_data
                this.$cookies.set("fastpaper-document-info", this.file);
                await this.updatefile();
            } catch (error) {
                console.error(error);
            }
        },
        updatefile() {
            this.file = this.$cookies.get("fastpaper-document-info");
            this.title = this.file.filename
            console.log(this.file.path);
        },
        showDeleteIcon(index) {
            this.deleteIcons[index] = true;
        },
        // 隐藏删除图标
        hideDeleteIcon(index) {
            this.deleteIcons[index] = false;
        },
        // 删除文件并输出ID
        deleteFile(index) {
            console.log(this.files[index]);
            console.log(this.files[index].resultid);

            const data = {
                "operate": "result_delete",
                "operate_type": "Operate_results",
                "userid": this.file.userid,
                "documentid": this.file.documentid,
                "resultid": this.files[index].resultid
            };
            axios.post('/api/backend', data)
                .then(response => {
                    console.log(response.data);
                    this.updatefiles()
                })
                .catch(error => {
                    console.error(error);
                });
        },
        updatefiles() {
            const data2 = {
                "operate": "result_query_all",
                "operate_type": "Operate_results",
                "userid": this.file.userid,
                "documentid": this.file.documentid
            };
            axios.post('/api/backend', data2)
                .then(response => {
                    console.log('response.data', response.data.data);
                    this.files = response.data.data
                    console.log();
                })
                .catch(error => {
                    console.error(error);
                });
        }
    }
}
</script>

<style>
html,
body {
    height: 100%;
}

#app {
    padding: 0;
    margin: 0;
}

.top {
    width: 100%;
    height: 5vh;
    user-select: none;
    background-color: beige;
}

.top p {
    margin: 0;
    /*设置字体*/
    font-size: 30px;
    line-height: 5vh;
    margin-left: 20px;
    /*设置字体大小*/
    font-weight: bolder;
    /*设置字体粗细*/
    -webkit-text-stroke: 1px #626161;
    /*文字描边*/
    -webkit-text-fill-color: transparent;
    /*设置文字的填充颜色*/
}

.all {

    width: calc(100vw - 4px);
    height: 94vh;
    padding: 0;
    margin: 0;
    vertical-align: top;
    border: 2px solid #ccc;
    border-top: 4px solid #ccc;
    user-select: none;
}

.single {
    display: inline-block;
    height: 100%;
}

.left {

    width: 15vw;
    vertical-align: top;
    border-right: 4px solid #ccc;
    background-color: rgb(189, 245, 245);
}

.middle {

    width: 23vw;
    vertical-align: top;
    border-right: 4px solid #ccc;
    background-color: rgb(217, 251, 251);
}

.contents {
    text-align: center;
    cursor: pointer
}

.file-item {
    border-bottom: 2px #ccc solid;
    margin-top: 6px;
    margin-bottom: 6px;
    font-size: 20px;
}

.block {
    border-bottom: 2px #ccc solid;
    margin-bottom: 15px;
    margin-top: 15px;
    font-size: 30px;
}

.right {

    width: calc(62vw - 12px);
    vertical-align: top;
    background-color: rgb(230, 248, 248);
}

img {
    width: 30vw;
    height: 15vh;
    position: absolute;
    top: 44%;
    left: 52%;
}

.title {
    width: 100%;
    height: 60px;

    border-bottom: 4px solid #ccc;

}

.title p {
    text-align: center;
    margin: 0;
    padding: 0;
    line-height: 60px;
    font-size: 26px;
    font-weight: bold;

}

.overlay {
    position: fixed;
    top: 0;
    left: 0;
    right: 0;
    bottom: 0;
    background-color: rgba(0, 0, 0, 0.5);
    display: flex;
    justify-content: center;
    align-items: center;
    z-index: 1000;
    /* 确保浮窗在最上层 */
}

.overlay-content {
    background-color: white;
    padding: 20px;
    border-radius: 5px;
    text-align: center;
    position: relative;
    z-index: 1001;
    /* 确保浮窗内容在背景之上 */
}

.file-item {
    position: relative;
    /* 为删除图标定位 */
    padding-right: 30px;
    /* 留出空间给删除图标 */
}

.delete-icon {
    position: absolute;
    /* 绝对定位 */
    right: 0;
    /* 靠右显示 */
    cursor: pointer;
    /* 鼠标悬浮时显示手型图标 */
    color: red;
    font-weight: bolder;
    /* 设置颜色 */
}

/* 鼠标悬浮时改变删除图标颜色 */
.file-item:hover .delete-icon {
    color: darkred;
}

.r_top {
    text-align: center;
    width: 100%;
    height: 40px;
    line-height: 40px;
    font-size: 30px;
    border-bottom: 2px solid #ccc;
}
</style>
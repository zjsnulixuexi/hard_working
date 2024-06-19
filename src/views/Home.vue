<template>
    <div class="main">
        <div class="global-bar-top">
            <div class="global-bar-top-left">
                <div class="title">Hard-working</div>

            </div>
            <div class="global-bar-top-right">
                <button @click="createNewDocument" class="button"><span>新建</span></button>
                <button @click="toggleRenameButton" class="button"><span>重命名</span></button>
                <button @click="toggleDeleteButton" class="button"><span>删除</span></button>
                <input class="search" placeholder="搜索" @input="fun" style="margin-left: 10px;" />
            </div>
        </div>

        <div class="button_div">
            <button class="button" :class="{ 'active': activeButton === 'all' }" @click="showAll"
                style="background-color: rgba(216, 247, 157, 0.8);width: 100px;"><span>全部文件</span></button>
            <!-- <button class="button" :class="{ 'active': activeButton === 'recent' }" @click="showRecent">最近文件</button> -->
            <button class="button" :class="{ 'active': activeButton === 'collection' }" @click="showCollection"
                style="background-color: rgba(216, 247, 157, 0.8);width: 100px;"><span>收藏文件</span></button>
        </div>
        <div class="global-context">
            <!-- 这里是全部文件 -->
            <div v-if="show" v-for="document in documents_searched" class="context-document-item">
                <div class="ideaFlow" @click="push_edit_page(document)"
                    @contextmenu.prevent="showContextMenu($event, document)"></div>
                <div :path="document.path" :class="textColor" style="font-weight: bold;
font-size: 1.4em; padding:5px 0">{{ document.filename }}</div>
                <div class="update-time" style="font-weight: bold;
font-size: 1.1em;color: rgb(77, 18, 18); padding: 2px 0;">{{ document.updateTime }}</div>
                <button v-if="document.showDeleteButton" class="delete-button"
                    @click.stop="deleteDocument($event, document)">X</button>
                <button v-if="document.isRenaming" class="rename-button"
                    @click.stop="startRenaming($event, document)">R</button>
                <button @click.stop="toggleCollection($event, document)"
                    :class="{ 'collection-button': true, 'collected': document.isCollected }" class="button_star"
                    style="margin-right: 130px;">
                    {{ document.isCollected ? '取消收藏' : '收藏' }}
                </button>
            </div>

            <!-- 这是收藏文件 -->
            <div v-else v-for="document in collectedDocuments" class="context-document-item">
                <div class="ideaFlow" @click="push_edit_page(document)"
                    @contextmenu.prevent="showContextMenu($event, document)"></div>
                <div :path="document.path" :class="textColor" style="font-weight: bold;
        font-size: 1.4em; padding:5px 0">{{ document.filename }}</div>
                <div class="update-time" style="font-weight: bold;
        font-size: 1.1em;color: rgb(77, 18, 18); padding: 2px 0;">{{ document.updateTime }}</div>
                <button @click="toggleCollection($event, document)"
                    :class="{ 'collection-button': true, 'collected': document.isCollected }" class="button_star">
                    {{ document.isCollected ? '取消收藏' : '收藏' }}
                </button>
                <button v-if="document.showDeleteButton" class="delete-button"
                    @click="deleteDocument($event, document)">X</button>
                <button v-if="document.isRenaming" class="rename-button"
                    @click="startRenaming($event, document)">R</button>
            </div>
            <!-- 右键菜单 -->
            <transition name="fade" appear>
                <div v-if="isContextMenuVisible" class="context-menu"
                    :style="{ top: contextMenuPosition.y + 'px', left: contextMenuPosition.x + 'px' }">
                    <ul>
                        <li @click="createNewDocument">新建</li>
                        <li @click="startRenaming($event, selectedDocument)">重命名</li>
                        <li @click="deleteDocument($event, selectedDocument)">删除</li>
                    </ul>
                </div>
            </transition>

        </div>
    </div>
</template>

<script scoped>
import { resolveDirective } from 'vue';
import global from "../components/global.vue"
import axios from 'axios'
export default {
    data() {
        return {
            show: true,
            activeButton: 'all',
            documents: [],
            searchbox_text: "",
            update_time: Date.now(),
            isContextMenuVisible: false,
            contextMenuPosition: { x: 0, y: 0 },
            selectedDocument: null,
            userid: '',
        }
    },
    mounted() {
        this.userid = Number(this.$cookies.get("userid"));
        console.log('userid', this.userid);
        const data = {
            "userID": this.userid,
            "operate": "Document_query_all",
            "operate_type": "Operate_documents"
        };
        axios.post('/api/backend', data)
            .then(response => {
                this.documents = response.data.data
            })
            .catch(error => {
                console.error(error);
            });
    },
    methods: {
        geizhi() {
            this.documents = response
        },
        showAll() {
            this.show = true;
            this.activeButton = 'all';
        },
        showRecent() {
            this.activeButton = 'recent';
        },
        showCollection() {
            this.show = false;
            this.activeButton = 'collection';
        },

        push_edit_page(document_info) {
            this.$cookies.set("fastpaper-document-info", document_info);
            this.$router.push("/run");
            const data = {
                "userID": this.userid,
                "documentID": document_info.documentID,
                "operate": "Document_update_time",
                "operate_type": "Operate_documents"
            };
            axios.post('/api/backend', data)
                .then(response => {
                    console.log(response.data);
                })
                .catch(error => {
                    console.error(error);
                });
        },

        //新建path
        async createNewDocument() {
            const newpath = prompt("请输入新文档的标题", "New Document");
            // 如果用户点击了取消按钮或者没有输入标题，则不进行任何操作
            if (newpath === null || newpath === "") {
                return;
            }
            console.log(1);
            // 检查是否存在重复的名称
            if (this.documents != null) {
                const isDuplicate = this.documents.some(doc => doc.path === newpath);
                if (isDuplicate) {
                    alert("文档标题重复，请重新输入一个名称");
                    return;
                }

                //这里确保当新建时，不会出现红色×和绿色R
                this.documents.forEach(document => {
                    document.showDeleteButton = false;
                    document.isRenaming = false;
                });
            }

            const data = {
                "userID": this.userid,
                "operate": "Document_create",
                "filename": newpath,
                "operate_type": "Operate_documents"
            };
            try {
                console.log(1);
                const response = await axios.post('/api/backend', data);
                console.log('新建成功');
                await this.updateDocuments();
            } catch (error) {
                console.error(error);
            }
        }
        ,

        // 出现红色X
        toggleDeleteButton() {
            this.documents.forEach(document => {
                document.showDeleteButton = !document.showDeleteButton;
            });

            //确保删除时不会出现绿色R（重命名）
            this.documents.forEach(document => {
                document.isRenaming = false;
            });
        },

        //删除文档
        async deleteDocument(event, document) {
            event.stopPropagation(); // 阻止事件冒泡
            this.isContextMenuVisible = false;
            // 显示确认对话框
            const isConfirmed = window.confirm('您确定要删除此文档吗？');

            if (isConfirmed) {
                const index = this.documents.indexOf(document);
                if (index > -1) {
                    this.documents.splice(index, 1);
                }

                const data = {
                    "userID": this.userid,
                    "operate": "Document_delete",
                    "documentID": document.documentid,
                    "operate_type": "Operate_documents"
                };

                try {
                    await axios.post('/api/backend', data);
                    console.log('删除成功');
                    await this.updateDocuments();
                } catch (error) {
                    console.error(error);
                }
            }
        },

        //出现绿色R
        toggleRenameButton() {
            this.documents.forEach(document => {
                document.isRenaming = !document.isRenaming;
            });

            //这里确保重命名不会出现红色×（删除）
            this.documents.forEach(document => {
                document.showDeleteButton = false;
            });
        },
        //重命名文档
        async startRenaming(event, document) {
            event.stopPropagation();
            this.isContextMenuVisible = false;
            const newpath = prompt("请输入新文档的标题", document.path);
            if (newpath === null || newpath === "") {
                return;
            }
            const data = {
                "userID": this.userid,
                "documentID": document.documentid,
                "filename": newpath,
                "path": '',
                "operate": "Document_update",
                "operate_type": "Operate_documents"
            };

            try {
                const response = await axios.post('/api/backend', data);
                console.log('重命名成功');
                await this.updateDocuments(); // 更新文档列表
            } catch (error) {
                console.error(error);
            }
        },


        //收藏按钮
        toggleCollection(event, document) {
            event.stopPropagation();
            document.isCollected = !document.isCollected;
            const data = {
                "userID": this.userid,
                "documentID": document.documentID,
                "isCollected": document.isCollected ? 1 : 0,
                "operate": "Document_update_is_collected",
                "operate_type": "Operate_documents"
            };
            axios.post('/api/backend', data)
                .then(response => {
                    console.log('收藏成功');
                })
                .catch(error => {
                    console.error(error);
                });
        },

        //搜索框
        fun(e) { this.searchbox_text = e.target.value },
        updateDocuments() {
            const data = {
                "userID": this.userid,
                "operate": "Document_query_all",
                "operate_type": "Operate_documents"
            };
            axios.post('/api/backend', data)
                .then(response => {
                    console.log(response.data);
                    this.documents = response.data.data
                })
                .catch(error => {
                    console.error(error);
                });
        },
        showContextMenu(event, doc) {
            this.selectedDocument = doc;
            this.isContextMenuVisible = true;
            this.contextMenuPosition = { x: event.clientX, y: event.clientY };
            document.addEventListener('click', this.hideContextMenu);
            window.addEventListener('click', () => {
                this.contextMenu.visible = false;
            }, { once: true });
        },
        hideContextMenu() {
            this.isContextMenuVisible = false;
            document.removeEventListener('click', this.hideContextMenu);
        },
    },

    computed: {
        documents_searched() {
            if (this.searchbox_text.length > 0) {
                return this.documents.filter(item => item.filename.toLowerCase().indexOf(this.searchbox_text.toLowerCase()) >= 0);
            } else {
                return this.documents;
            }
            // return [{path:"MUMPS: A General Purpose Distributed Memory Sparse Solver"}]            
        },
        collectedDocuments() {
            return this.documents.filter(document => document.isCollected);
        },
        textColor() {
            return this.searchbox_text.length > 0 ? 'path-red' : 'path-black';
        }

    }
}
</script>


<style scoped lang="scss">
.main {
    width: 100%;
    height: 100vh;
    background: radial-gradient(circle, #f6c89f, #ff7e5f);
}

.path-red {
    color: red;
}

.active {
    font-weight: bold;
}

.path-black {
    color: black;
}


.global-bar-top {
    position: fixed;
    height: 50px;
    width: 100vw;
    line-height: 50px;
    padding: 0 20px;
    border-bottom: 1px solid rgb(209, 209, 209);
}

.global-bar-top-left {
    position: absolute;
}

.global-bar-top-right {
    position: absolute;
    right: 60px;
}

.button {
    display: inline-block;
    border-radius: 15px;
    /* 调整按钮的边框圆角 */
    background-color: #ffcd6a;
    /* 修改背景颜色 */
    border: none;
    color: #000000;
    /* 修改文字颜色 */
    text-align: center;
    font-size: 16px;
    /* 调整按钮的字体大小 */
    font-weight: 400;
    padding: 5px;
    /* 调整按钮的内边距 */
    width: 80px;
    /* 调整按钮的宽度 */
    transition: all 0.5s;
    cursor: pointer;
    margin: 5px;
    vertical-align: middle;
}

.button span {
    cursor: pointer;
    display: inline-block;
    position: relative;
    transition: 0.5s;
}

.button span::after {
    content: ">";
    position: absolute;
    opacity: 0;
    top: 0;
    right: -15px;
    /* 调整按钮箭头的位置 */
    transition: 0.5s;
}

.button:hover span {
    padding-right: 15px;
    /* 调整鼠标悬停时的填充 */
}

.button:hover span::after {
    opacity: 1;
    right: 0;
}

.delete-button {
    position: absolute;
    bottom: 270px;
    width: 30px;
    height: 25px;
    font-size: 18px;
    right: 32px;
    margin-top: 20px;
    color: red;
    cursor: pointer;
    margin-left: 5px;
    border: 0;
    border-radius: 6px;
    background-color: rgba(229, 204, 143);

    &:hover {
        color: darkred; // 鼠标悬停时的颜色
        background-color: rgb(255, 163, 58);
    }
}

.rename-button {
    position: absolute;
    bottom: 270px;
    right: 32px;
    width: 30px;
    font-size: 18px;
    height: 25px;
    margin-top: 20px;
    color: rgb(0, 0, 0);
    border: 0;
    cursor: pointer;
    margin-left: 5px;
    border-radius: 6px;
    background-color: rgba(229, 204, 143);

    &:hover {
        color: rgb(0, 82, 0); // 鼠标悬停时的颜色
        background-color: #ffe14b;
    }
}





.global-context {
    position: fixed;
    top: 90px;
    // background-color: antiquewhite;
    width: 100vw;
    height: calc(100vh - 50px);
    overflow: auto;
    display: flex;
    flex-wrap: wrap;
    justify-content: center;
}

.button_div {
    width: 100%;
    height: 40px;
    margin: 0;
    padding: 0;
    position: absolute;
    top: 55px;

}

.context-document-item {
    padding: 20px 20px 40px 20px;
    border-radius: 20px;
    height: 280px;
    position: relative;
}



.context-document-item .ideaFlow {
    border-style: solid;
    width: 300px;
    height: 180px;
    border-radius: 20px;
    border-width: 2px;
    border-color: rgb(76, 247, 252);
    background-image: url('../assets/img.png');
    background-position: center center;
    background-size: 100%;

    &:hover {
        cursor: pointer;
    }
}


.context-document-item div {
    width: 280px;
    font-weight: 300;
    height: 20px;
    // background-color: aqua;
    font-size: 17px;
    font-family: monospace, "Microsoft YaHei";
    overflow: hidden;
    text-align: center;
    text-overflow: ellipsis;
    white-space: nowrap;
    margin: auto;
    user-select: none;

}

.collected {
    background-color: yellow !important; // 已收藏状态下的背景色，使用!important覆盖默认样式
}

.search {
    width: 200px;
    height: 20px;
    padding: 5px;
    border: 2px solid #eca29b;
    /* 橙色按钮的颜色 */
    border-radius: 5px;
    transition: border-color 0.3s, box-shadow 0.3s;
    margin-left: 10px;
    background-color: #f7d1c5;
    outline: none;
}

.search:focus {
    border-color: #d77f6e;
    /* 橙色加深 */
    box-shadow: 0 0 5px rgba(215, 127, 110, 0.5);
    background-color: #f9d6cc;
    /* 淡橙色填充 */
}

.search::placeholder {
    color: #eca29b;
    /* 橙色按钮的颜色 */
}

.title {
    font-size: 26px;
    font-style: italic;
    cursor: default;
    color: #fff134;
    transition: color 0.5s, text-shadow 0.5s;
}

.title:hover {
    color: #ffc61a;
    text-shadow: 0 0 10px #ed4c4c, 0 0 20px #ed4c4c, 0 0 30px #ed4c4c;
}

.button_star {
    display: inline-block;
    border-radius: 15px;
    background-image: -webkit-linear-gradient(top, #fff7d4, #ffe67a);
    background-image: -moz-linear-gradient(top, #fff7d4, #ffe67a);
    background-image: -ms-linear-gradient(top, #fff7d4, #ffe67a);
    background-image: -o-linear-gradient(top, #fff7d4, #ffe67a);
    background-image: linear-gradient(to bottom, #fff7d4, #ffe67a);
    border: none;
    color: #000000;
    text-align: center;
    font-size: 16px;
    font-weight: 400;
    padding: 5px;
    width: 80px;
    transition: all 0.5s;
    cursor: pointer;
    margin: 5px;
    vertical-align: middle;
}

/* 悬停样式 */
.button_star:hover {
    background: #ffe67a;
    background-image: -webkit-linear-gradient(top, #ffe67a, #fff7d4);
    background-image: -moz-linear-gradient(top, #ffe67a, #fff7d4);
    background-image: -ms-linear-gradient(top, #ffe67a, #fff7d4);
    background-image: -o-linear-gradient(top, #ffe67a, #fff7d4);
    background-image: linear-gradient(to bottom, #ffe67a, #fff7d4);
    text-decoration: none;
}

.context-menu {
    position: fixed;
    background-color: #afaf1c;
    border: 2px solid #36a525;
    border-radius: 5px;
    box-shadow: 2px 2px 5px rgba(0, 0, 0, 0.2);
    z-index: 1000;
}

.context-menu ul {
    list-style: none;
    padding: 0;
    margin: 0;
}

.context-menu li {
    padding: 10px;
    cursor: pointer;
}

.context-menu li:hover {
    background-color: rgb(135, 197, 53);
}

.fade-enter-active,
.fade-leave-active {
    transition: opacity 0.5s;
}

.fade-enter,
.fade-leave-to {
    opacity: 0;
}
</style>
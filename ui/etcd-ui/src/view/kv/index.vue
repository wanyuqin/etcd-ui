<template>
    <el-row>
        <el-col :span="1">
            <el-button type="primary" size="small" @click="openPutKeyForm">添 加</el-button>

        </el-col>
        <el-col :span="1">
            <el-button type="info" size="small" :icon="Refresh" @click="refreshKeyTree" />
        </el-col>
    </el-row>
    <el-row :gutter="20">
        <el-col :span="12">
            <el-input v-model="query" placeholder="输入关键字" @input="onQueryChanged" disabled="true" />
            <el-tree-v2 ref="treeRef" :filter-method="filterKey" class="el-tree" :data="treeData" :props="props" :height="600" @node-click="getKeyByName">
            </el-tree-v2>
        </el-col>

        <template v-if="keyDetailVisible">
            <!-- 右侧内容  -->
            <el-col :span="12">
                <el-row :gutter="50">
                    <el-col :span="12">
                        <el-row>
                            <el-col :span="24">
                                <el-input v-model="kv.key" placeholder="" style="width: 100%"></el-input>
                            </el-col>
                        </el-row>
                    </el-col>
                    <el-col :span="12">
                        <el-row>
                            <el-col :span="6">
                                <el-button size="small" type="primary" @click="setTTL">设置TTL</el-button>
                            </el-col>
                            <el-col :span="6">
                                <el-button size="small" type="info" @click="reloadKey">重新加载</el-button>
                            </el-col>
                            <el-col :span="6">
                                <el-button size="small" type="success" @click="watchConfirm">监 听</el-button>
                            </el-col>
                            <el-col :span="6">
                                <el-button size="small" type="danger" @click="deleteConfirm">删 除</el-button>
                            </el-col>
                        </el-row>
                    </el-col>

                </el-row>

                <el-row :gutter="10">
                    <el-col :span="8">
                        Version
                        <el-input v-model="kv.version" disabled />
                    </el-col>
                    <el-col :span="8">
                        CreteReversion
                        <el-input v-model="kv.create_revision" disabled />
                    </el-col>
                    <el-col :span="8">
                        ModReversion
                        <el-input v-model="kv.mod_revision" disabled />
                    </el-col>
                </el-row>
                <div style="margin: 1px 0"></div>

                <!-- 格式化选择框 -->
                <!-- <el-select v-model="value" class="m-2" placeholder="格式化" size="default" @change="formatJson">
                    <el-option v-for="item in selectorOption" :key="item.key" :label="item.label" :value="item.value" />
                </el-select> -->

                <!-- 内容区域 -->
                <el-input class="content-textarea" v-model="kv.value" placeholder="Please input" show-word-limit type="textarea" input-style="height:500px" />
                <div style="margin: 10px 0"></div>
                <el-button size="small" type="success" @click="saveKey">保 存</el-button>
            </el-col>
            <!-- 右侧内容结束 -->
        </template>
    </el-row>

    <!-- 删除确认框 -->
    <el-dialog v-model="centerDialogVisible" title="警告" width="30%" center>
        <span>
            确定要删除 {{ kv.key }} 吗？
        </span>
        <template #footer>
            <span class="dialog-footer">
                <el-button @click="centerDialogVisible = false">取消</el-button>
                <el-button type="danger" @click="doDeleteKey">
                    确定
                </el-button>
            </span>
        </template>
    </el-dialog>
    <!-- 结束 -->

    <!-- 监听确认框 -->
    <el-dialog v-model="watchDialogVisible" title="确认" width="30%" center>
        <span>
            确定要监听 {{ kv.key }} 吗？
        </span>
        <template #footer>
            <span class="dialog-footer">
                <el-button @click="watchDialogVisible = false">取消</el-button>
                <el-button type="danger" @click="doWatchKey">
                    确定
                </el-button>
            </span>
        </template>
    </el-dialog>

    <!-- 添加对话框 -->
    <el-dialog v-model="dialogFormVisible" title="添加KEY-VALUE">
        <el-form :model="newKv">
            <el-form-item label="键" :label-width="formLabelWidth">
                <el-input v-model="newKv.key" autocomplete="off" />
            </el-form-item>
            <el-form-item label="值" :label-width="formLabelWidth">
                <el-input v-model="newKv.value" autocomplete="off" type="textarea" input-style="height:100px" />
            </el-form-item>
            <el-form-item label="过期时间">
                <el-input-number v-model="newKv.ttl" :min="0" />
            </el-form-item>

        </el-form>
        <template #footer>
            <span class="dialog-footer">
                <el-button @click="dialogFormVisible = false">取消</el-button>
                <el-button type="primary" @click="doPutKey">
                    确认
                </el-button>
            </span>
        </template>
    </el-dialog>

</template>
<script  setup>
import { onMounted, ref } from "vue";
import { ElLoading, ElMessage } from "element-plus";
import { Refresh } from "@element-plus/icons-vue";
import { listKv, getKey, putKey, deleteKey, watchKey } from "@/api/kv.js";
import { ElTreeV2 } from "element-plus";
import hljs from "highlight.js";
import "highlight.js/styles/dark.css";

const props = ref({
    value: "id",
    label: "label",
    children: "children",
});
const oldValue = ref("");
const treeData = ref([]);
const centerDialogVisible = ref(false);
const watchDialogVisible = ref(false);
const dialogFormVisible = ref(false);
const kv = ref({});
const keyDetailVisible = ref(false);
const treeRef = ref();
const selectorOption = ref([
    {
        value: "PlainText",
        label: "Plain Text",
    },
    {
        value: "Json",
        label: "Json",
    },
]);
const query = ref("");
const newKv = ref({
    key: "",
    value: "",
    ttl: -1,
});

const getKeyTree = () => {
    listKv().then((response) => {
        console.log(response);
        treeData.value = response.data.data;
    });
};

const refreshKeyTree = () => {
    getKeyTree();
    keyDetailVisible.value = false;
};

const getKeyByName = (node) => {
    getKey(node.id).then((response) => {
        kv.value = response.data.data;
        keyDetailVisible.value = true;
    });
};

const deleteConfirm = () => {
    centerDialogVisible.value = true;
};

const watchConfirm = () => {
    watchDialogVisible.value = true;
};

const openPutKeyForm = () => {
    dialogFormVisible.value = true;
};

const doDeleteKey = () => {
    centerDialogVisible.value = false;
    const loading = ElLoading.service({
        lock: true,
        text: "Loading",
        background: "rgba(0, 0, 0, 0.7)",
    });

    deleteKey(kv.value.key).then((response) => {
        loading.close();
        getKeyTree();
        kv.value = {};
    });
};

const doWatchKey = () => {
    watchKey(kv.value.key).then((response) => {
        watchDialogVisible.value = false;
    });
};

const formatJson = (param) => {
    if (param == "PlainText") {
        kv.value.value = oldValue.value;
    }
    if (param == "Json") {
        var jsonObj = JSON.parse(kv.value.value);
        var jsonString = JSON.stringify(jsonObj, null, 4);
        kv.value.value = jsonString;
    }
};

const doPutKey = () => {
    dialogFormVisible.value = false;
    const loading = ElLoading.service({
        lock: true,
        text: "Loading",
        background: "rgba(0, 0, 0, 0.7)",
    });
    putKey(newKv.value).then((response) => {
        loading.close();
        getKeyTree();
    });
};

const setTTL = () => {
    dialogFormVisible.value = true;
    newKv.value = kv.value;
};

const saveKey = () => {
    putKey(kv.value).then((response) => {
        if (response.data.code == "0000") {
            ElMessage({
                message: "保存成功",
                type: "success",
            });
            reloadKey();
        }
    });
};

const reloadKey = () => {
    if (kv.value.key != "") {
        getKey(kv.value.key).then((response) => {
            kv.value = response.data.data;
        });
    }
};

const onQueryChanged = (query) => {
    return treeRef.value.filter(query);
};

const filterKey = (query, node) => {
    return node.label.includes(query);
};

onMounted(() => {
    getKeyTree();
});
</script>
    

<style>
.custom-tree-node {
    flex: 1;
    display: flex;
    align-items: center;
    justify-content: space-between;
    font-size: 14px;
    padding-right: 8px;
}

.el-row {
    margin-bottom: 20px;
}
.el-row:last-child {
    margin-bottom: 0;
}
.el-col {
    border-radius: 4px;
}
</style>
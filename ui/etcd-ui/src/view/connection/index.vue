<template>
    <el-button type="primary" @click="openAddDiaglog">添 加</el-button>
    <el-table :data="listConnectionData" style="width: 100%">
        <el-table-column label="名称" prop="name" />
        <el-table-column label="类型" prop="type" />
        <el-table-column label="端点" prop="endpoint" />
        <el-table-column label="证书" prop="certificate_name" />
        <el-table-column label="Active" prop="active">
            <template #default="scope">
                <el-tag :type="scope.row.active === 1 ? 'success' : 'info'" disable-transitions>Active</el-tag>
            </template>
        </el-table-column>
        <el-table-column label="创建时间" prop="created_at" />
        <el-table-column align="right">
            <template #header>
                <el-input v-model="listQuery.name" size="small" placeholder="Type to search" />
            </template>
            <template #default="scope">
                <el-button size="small" @click="activeConnection(scope.row)">连 接</el-button>
                <el-button size="small" type="danger" @click="handleDelete(scope.$index, scope.row)">删 除</el-button>
            </template>
        </el-table-column>
    </el-table>

    <!-- 添加表单 -->
    <el-dialog v-model="addConnectionDialogVisible" title="新建连接">

        <el-tabs v-model="activeName" class="demo-tabs">
            <el-tab-pane label="DEFAULT" name="default">
                <el-form :model="defaultConnectionType" :label-position="labelPosition" label-width="100px">
                    <el-form-item label="名称">
                        <el-input v-model="defaultConnectionType.name" placeholder="自定义连接名称"></el-input>
                    </el-form-item>
                    <el-form-item label="Endpoint">
                        <el-input v-model="defaultConnectionType.endpoint" placeholder="Etcd连接地址"></el-input>
                    </el-form-item>

                    <el-form-item label="开启 TLS">
                        <el-switch v-model="switchValue" @change="certificateSwitch" />
                    </el-form-item>
                    <template v-if="certificateSelector">
                        <!-- certificate 选择框 -->
                        <el-form-item label="证书">
                            <el-select v-model="defaultConnectionType.certificate_id" placeholder="请选择证书">
                                <el-option v-for="item in certificateList" :label="item.name" :value="item.id" :key="item" />
                            </el-select>
                        </el-form-item>
                    </template>
                </el-form>
            </el-tab-pane>
            <el-tab-pane label="KEY AUTH" name="keyAuth">Config</el-tab-pane>
            <el-tab-pane label="JWT AUTH" name="jwtAuth">Role</el-tab-pane>
            <el-tab-pane label="BASE AUTH" name="baseAuth">Task</el-tab-pane>
        </el-tabs>

        <!-- <el-form :model="form">

        </el-form> -->
        <template #footer>
            <span class="dialog-footer">
                <el-button @click="addConnectionDialogVisible = false">取 消</el-button>
                <el-button type="primary" @click="doAddConnection">
                    确 定
                </el-button>
            </span>
        </template>
    </el-dialog>
    <!-- 添加表单 结束 -->

</template>

<script setup>
import { onMounted, ref } from "vue";
import {
    listConnection,
    createConnection,
    delConnection,
    updateConnection,
} from "@/api/connection";

import { listCertificates } from "@/api/certificate.js";

import { ElMessage, ElMessageBox } from "element-plus";
import { transformConnectionType } from "@/utils/connection.js";
const queryName = ref();
const listQuery = ref({
    page: 1,
    pageSize: 10,
    name: "",
});
const labelPosition = ref("right");
const defaultConnectionType = ref({
    name: "",
    endpoint: "",
    type: 1,
    tls: 1,
    certificate_id: 0,
});
const activeName = ref("default");
const addConnectionDialogVisible = ref(false);
const listConnectionData = ref();
const certificateSelector = ref(false);
const certificateList = ref([]);
const switchValue = ref(false);
const certificatePageiInfo = ref({
    page: 1,
    pageSize: 10,
    name: "",
});

const getConnectionList = () => {
    listConnection(listQuery.value).then((response) => {
        response.data.data.records.forEach((element) => {
            console.log(element);
            element.type = transformConnectionType(element.type);
        });

        listConnectionData.value = response.data.data.records;
    });
};
const doAddConnection = () => {
    createConnection(defaultConnectionType.value).then((response) => {
        if (response.data.code == "0000") {
            getConnectionList();
            addConnectionDialogVisible.value = false;
            ElMessage({
                message: "添加成功",
                type: "success",
            });
        }
    });
};

const handleDelete = (index, row) => {
    ElMessageBox.confirm("确认是否删除?", "Warning", {
        confirmButtonText: "OK",
        cancelButtonText: "Cancel",
        type: "warning",
    })
        .then(() => {
            delConnection(row.id).then((response) => {
                if (response.data.code == "0000") {
                    getConnectionList();
                    ElMessage({
                        type: "success",
                        message: "删除成功！",
                    });
                }
            });
        })
        .catch(() => {
            // ElMessage({
            //     type: "info",
            //     message: "Delete canceled",
            // });
        });
};

const activeConnection = (row) => {
    row.active = 1;
    const updateConn = {
        id: row.id,
        active: 1,
    };
    updateConnection(updateConn).then((response) => {
        getConnectionList();
    });
};

const openAddDiaglog = () => {
    addConnectionDialogVisible.value = true;
    defaultConnectionType.value.name = "";
    defaultConnectionType.value.endpoint = "";
};

const activeFilter = (value, row) => {
    return (row.tag = value);
};

const getCertificates = () => {
    listCertificates(certificatePageiInfo.value).then((response) => {
        certificateList.value = response.data.data.records;
    });
};

const certificateSwitch = (param) => {
    if (param) {
        getCertificates();
        certificateSelector.value = true;
        defaultConnectionType.value.tls = 2;
    } else {
        defaultConnectionType.value.tls = 1;
        defaultConnectionType.value.certificate_id = 0;
        certificateSelector.value = false;
    }
};

onMounted(() => {
    console.log("aaa");
    getConnectionList();
});
</script>

<style>
</style>
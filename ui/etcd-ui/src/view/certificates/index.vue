<template>
    <el-button type="primary" @click="openAddDiaglog">添 加</el-button>
    <el-table :data="listCertificateData" style="width: 100%">
        <el-table-column label="名称" prop="name" />
        <el-table-column label="创建时间" prop="created_at" />
        <el-table-column align="right">
            <template #header>
                <el-input v-model="listQuery.name" size="small" placeholder="Type to search" />
            </template>
            <template #default="scope">
                <el-button size="small" @click="openCertificateDraw(scope.row)">查 看</el-button>
                <el-button size="small" type="danger" @click="handleDelete(scope.$index, scope.row)">删 除</el-button>
            </template>
        </el-table-column>
    </el-table>

    <!-- 添加表单 -->
    <el-dialog v-model="addCertificateDialogVisible" title="新建证书">

        <el-form :model="certificate" :label-position="labelPosition" label-width="100px">
            <el-form-item label="名称">
                <el-input v-model="certificate.name" placeholder="自定义名称"></el-input>
            </el-form-item>
            <el-form-item label="CA">
                <el-input type="textarea" v-model="certificate.ca" placeholder="" input-style="height:150px"></el-input>
                <el-upload :show-file-list="false" class="upload-demo" :on-change="caUploadOnChange" :auto-upload="false">
                    <el-button type="primary">上传CA</el-button>
                </el-upload>
            </el-form-item>
            <el-form-item label="Cert">
                <el-input type="textarea" v-model="certificate.cert" placeholder="" input-style="height:150px"></el-input>
                <div></div>
                <el-upload class="upload-demo" :show-file-list="false" :on-change="certUploadOnChange" :auto-upload="false">
                    <el-button type="primary">上传Cert</el-button>
                </el-upload>
            </el-form-item>
            <el-form-item label="Key">
                <el-input type="textarea" v-model="certificate.key" placeholder="" input-style="height:150px"></el-input>
                <el-upload class="upload-demo" :show-file-list="false" :on-change="keyUploadOnChange" :auto-upload="false">
                    <el-button type="primary">上传Key</el-button>
                </el-upload>
            </el-form-item>
        </el-form>

        <!-- <el-form :model="form">

        </el-form> -->
        <template #footer>
            <span class="dialog-footer">
                <el-button @click="addCertificateDialogVisible = false">取 消</el-button>
                <el-button type="primary" @click="doAddCertificate">
                    确 定
                </el-button>
            </span>
        </template>
    </el-dialog>
    <!-- 添加表单 结束 -->

    <!-- 侧边详情抽屉 -->
    <el-drawer :direction="drawerDirection" v-model="certificateDraw" :before-close="drawClose">

        <template #header>
            <h4>证 书 详 情</h4>
        </template>
        <template #default>
            <el-form :model="certificate" :label-position="labelPosition" label-width="100px">
                <el-form-item label="名称">
                    <el-input v-model="certificate.name" placeholder="自定义名称"></el-input>
                </el-form-item>
                <el-form-item label="CA">
                    <el-input type="textarea" v-model="certificate.ca" placeholder="" input-style="height:150px"></el-input>
                    <el-upload class="upload-demo" :show-file-list="false" :on-change="caUploadOnChange" :auto-upload="false">
                        <el-button type="primary">上传CA</el-button>
                    </el-upload>
                </el-form-item>
                <el-form-item label="Cert">
                    <el-input type="textarea" v-model="certificate.cert" placeholder="" input-style="height:150px"></el-input>
                    <el-upload class="upload-demo" :show-file-list="false" :on-change="certUploadOnChange" :auto-upload="false">
                        <el-button type="primary">上传Cert</el-button>
                    </el-upload>
                </el-form-item>
                <el-form-item label="Key">
                    <el-input type="textarea" v-model="certificate.key" placeholder="" input-style="height:150px"></el-input>
                    <el-upload class="upload-demo" :show-file-list="false" :on-change="keyUploadOnChange" :auto-upload="false">
                        <el-button type="primary">上传Key</el-button>
                    </el-upload>
                </el-form-item>
            </el-form>
        </template>
        <template #footer>
            <div style="flex: auto">
                <el-button type="primary" @click="updateCertificate">更 新</el-button>
            </div>
        </template>

    </el-drawer>
    <!-- 侧边详情抽屉结束 -->
</template>

<script setup>
import { onMounted, ref } from "vue";
import {
    listCertificates,
    createCertificates,
    delCertificates,
    updateCertificates,
} from "@/api/certificate";
import { ElMessage, ElMessageBox } from "element-plus";
import { transformConnectionType } from "@/utils/connection.js";

const queryName = ref();
const listQuery = ref({
    page: 1,
    pageSize: 10,
    name: "",
});
const labelPosition = ref("right");
const certificate = ref({
    ca: "",
    cert: "",
    key: "",
    name: "",
});
const activeName = ref("default");
const addCertificateDialogVisible = ref(false);
const listCertificateData = ref();
const drawerDirection = ref("rtl");
const certificateDraw = ref(false);

const updateCertificate = () => {
    updateCertificates(certificate.value).then((response) => {
        if (response.data.code == "0000") {
            ElMessage({
                message: "更新成功",
                type: "success",
            });
            drawClose();
        }
    });
};

const drawClose = () => {
    console.log("123");
    certificateDraw.value = false;
    getCertificatesList();
};

const getCertificatesList = () => {
    listCertificates(listQuery.value).then((response) => {
        listCertificateData.value = response.data.data.records;
    });
};

const doAddCertificate = () => {
    createCertificates(certificate.value).then((response) => {
        if (response.data.code == "0000") {
            getCertificatesList();
            addCertificateDialogVisible.value = false;
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
            delCertificates(row.id).then((response) => {
                if (response.data.code == "0000") {
                    getCertificatesList();
                    ElMessage({
                        type: "success",
                        message: "删除成功！",
                    });
                }
            });
        })
        .catch(() => {});
};

const activeConnection = (row) => {
    row.active = 1;
    updateConnection(row).then((response) => {
        getCertificatesList();
    });
};

const openAddDiaglog = () => {
    addCertificateDialogVisible.value = true;
    cleanCertificate();
};

const activeFilter = (value, row) => {
    return (row.tag = value);
};

const cleanCertificate = () => {
    certificate.value.name = "";
    certificate.value.ca = "";
    certificate.value.cert = "";
    certificate.value.key = "";
};

const openCertificateDraw = (row) => {
    console.log(row);
    certificate.value.id = row.id;
    certificate.value.name = row.name;
    certificate.value.ca = row.ca;
    certificate.value.cert = row.cert;
    certificate.value.key = row.key;
    certificateDraw.value = true;
};

const caUploadOnChange = (file) => {
    console.log(file);
    var reader = new FileReader();
    reader.onload = function () {
        if (reader.result) {
            //打印文件内容

            certificate.value.ca = reader.result;
        }
    };
    reader.readAsText(file.raw);
};

const certUploadOnChange = (file) => {
    console.log(file);
    var reader = new FileReader();
    reader.onload = function () {
        if (reader.result) {
            certificate.value.cert = reader.result;
        }
    };
    reader.readAsText(file.raw);
};

const keyUploadOnChange = (file) => {
    console.log(file);
    var reader = new FileReader();
    reader.onload = function () {
        if (reader.result) {
            certificate.value.key = reader.result;
        }
    };
    reader.readAsText(file.raw);
};

onMounted(() => {
    console.log("aaa");
    getCertificatesList();
});
</script>

<style>
.upload-demo {
    margin-top: 10px;
}
</style>
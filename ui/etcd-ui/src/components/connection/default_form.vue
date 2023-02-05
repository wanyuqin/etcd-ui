<template>
    <el-form :model="connection" :label-position="labelPosition" label-width="100px">
        <el-form-item label="名称">
            <el-input v-model="connection.name" placeholder="自定义连接名称"></el-input>
        </el-form-item>
        <el-form-item label="Endpoint">
            <el-input v-model="connection.endpoint" placeholder="Etcd连接地址"></el-input>
        </el-form-item>
    </el-form>
</template>

<script setup>
import { onMounted, ref } from "vue";

const labelPosition = ref("right")

const connection = ref({
    name: "",
    endpoint: "",
    type: 1,
});
</script>

<style>
</style>
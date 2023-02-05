<template>
    <el-image style="width: 150px; height: 50px" src="../../etcd-horizontal-bw.png" :fit="imageFit" />
    <el-menu active-text-color="#ffd04b" background-color="#223749" class="el-menu-vertical" default-active="2" text-color="#fff" :router="true">
        <el-sub-menu index="1">
            <template #title>
                <span>集群</span>
            </template>
            <el-menu-item index="/kv">键/值</el-menu-item>
            <el-menu-item index="/cluster-info">集群信息</el-menu-item>
        </el-sub-menu>

        <el-sub-menu index="2">
            <template #title>
                <span>设置</span>
            </template>
            <el-menu-item index="/connections">连接信息</el-menu-item>
            <el-menu-item index="/certificates">证书</el-menu-item>
        </el-sub-menu>
    </el-menu>
</template>

<script setup>
import {ref} from "vue"
const imageFit = ref("")
</script>

<style>
/* .el-menu-vertical {
    height: 90%;
} */

.el-menu{
  border: 0;
}
</style>
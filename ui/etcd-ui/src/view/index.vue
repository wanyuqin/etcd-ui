<script setup>
import Menu from "@/components/menu/menu.vue";
Menu;
</script>
<template>
    <el-container class="container">
        <el-aside class="aside" width="200px">
            <Menu></Menu>
        </el-aside>
        <el-container>
            <!-- <el-header class="header">
            </el-header> -->
            <el-main class="main">
                <router-view></router-view>
            </el-main>
        </el-container>
    </el-container>
</template>


<style>
.container {
    height: 100%;
}
.header {
    color: #f5f7fa;
    height: 50px;
}

.aside {
    width: "20px";
    color: #f5f7fa;
    background-color: #223749;
    height: 100%;
}

.main {
    height: 100%;
}
</style>
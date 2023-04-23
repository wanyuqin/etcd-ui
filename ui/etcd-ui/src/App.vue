<script setup>
import indexView from "./view/index.vue";
import { ElNotification } from "element-plus";
import { onMounted } from "vue";

const socket = new WebSocket("ws://localhost:8081/v1/connection");

const ping = () => {
    setInterval(() => {
        if (socket.readyState == WebSocket.OPEN) {
            socket.send("ping");
        }
    }, 3000);
};

const getMessage = () => {
    socket.onmessage = function (e) {
        if (e.data == "pong") {
            console.log("❤️");
            return;
        }
        const jsonobj = JSON.parse(e.data);
        console.log(jsonobj);
        ElNotification({
            title: jsonobj.event_type,
            message: jsonobj.key,
        });
    };
};

const onOpen = () => {
    socket.onopen = function (e) {
        ElNotification({
            title: "Socket",
            message: "已连接",
        });
    };
};

onMounted(() => {
    onOpen();
    getMessage();
    ping();
});
</script>

<template>
    <indexView></indexView>
</template>

<style>
#app {
    width: 100%;
    height: 100%;
}
</style>

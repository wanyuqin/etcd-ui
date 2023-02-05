<script setup>
import indexView from "./view/index.vue";
import { ElNotification } from "element-plus";
import { onMounted } from "vue";

const connection = new WebSocket("ws://localhost:8081/v1/connection")

onMounted(()=>{
  connection.onmessage = function(e){
    const jsonobj=  JSON.parse(e.data)
    console.log(jsonobj)
    ElNotification({
    title: jsonobj.event_type,
    message: jsonobj.key,
  })
  }
})

</script>

<template>
  <indexView></indexView>
</template>



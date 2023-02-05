<template>
    <template v-for="cluster in clusters" :key="cluster">
        <div class="tag-group my-2 flex flex-wrap gap-1 items-center">
            <span class="m-1"> {{ cluster.endpoint }}
            </span>
            <el-tag v-if="cluster.is_leader" type="danger" effect="dark" round>Leader</el-tag>
            <el-tag v-else type="info" effect="dark" round>Leader</el-tag>
            <el-tag v-if="cluster.is_learner" type="success" effect="dark" round>Learner</el-tag>
            <el-tag v-else type="info" effect="dark" round>Learner</el-tag>
        </div>
        <el-descriptions :column="3" border>
            <el-descriptions-item label="ID" label-align="right" align="center" label-class-name="my-label" class-name="my-content" width="150px">{{ cluster.id }}</el-descriptions-item>
            <el-descriptions-item label="端点" label-align="right" align="center" label-class-name="my-label" class-name="my-content" width="150px">{{ cluster.endpoint }}</el-descriptions-item>
            <el-descriptions-item label="版本" label-align="right" align="center">{{ cluster.version }}</el-descriptions-item>
            <el-descriptions-item label="DB size" label-align="right" align="center">{{ cluster.db_size }}</el-descriptions-item>
            <el-descriptions-item label="DB In Use" label-align="right" align="center">{{ cluster.db_size_in_use }}</el-descriptions-item>
            <el-descriptions-item label="Reversion" label-align="right" align="center">{{cluster.reversion}}</el-descriptions-item>
            <el-descriptions-item label="Raft Index" label-align="right" align="center">{{cluster.raft_index}}</el-descriptions-item>
            <el-descriptions-item label="Raft Term" label-align="right" align="center">{{ cluster.raft_term }}</el-descriptions-item>
            <el-descriptions-item label="Raft Applied Index" width="5" label-align="right" align="center">{{cluster.raft_applied_index}}</el-descriptions-item>
            <el-descriptions-item label="错误信息" label-align="right" align="center">{{ cluster.error }}</el-descriptions-item>
        </el-descriptions>
    </template>
</template>


<script setup>
import { clusterStatus } from "@/api/cluster";
import { onMounted, ref } from "vue";

const clusters = ref([]);

const getClusterStatus = () => {
    clusterStatus().then((response) => {
        clusters.value = response.data.data;
        console.log(clusters.value);
    });
};

onMounted(() => {
    getClusterStatus();
});
</script>

  <style scoped>
.my-label {
    background: var(--el-color-success-light-9);
}
.my-content {
    background: var(--el-color-danger-light-9);
}

.gap-1 {
    grid-gap: 0.25rem;
    gap: 0.25rem;
}
.items-center {
    align-items: center;
}

.flex-wrap {
    flex-wrap: wrap;
}

.flex {
    display: flex;
}

.my-2 {
    margin-top: 0.5rem;
    margin-bottom: 0.5rem;
}

.m-1 {
    margin: 0.25rem;
    font-weight: 200;
}
</style>
  
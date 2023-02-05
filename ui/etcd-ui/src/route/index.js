import {
    createRouter,
    createWebHashHistory
} from "vue-router";

const kv = () => import("@/view/kv/index.vue")
const clusterInfo = () => import("@/view/cluster/index.vue")
const connections = () => import("@/view/connection/index.vue")
const certificates = () => import("@/view/certificates/index.vue")


const routes = [{
        path: "/kv",
        name: "kv",
        component: kv
    },
    {
        path: "/cluster-info",
        name: "clusterInfo",
        component: clusterInfo
    }, {
        path: "/connections",
        name: "connections",
        component: connections
    }, {
        path: "/certificates",
        name: "certificates",
        component: certificates
    }
]


export const router = createRouter({
    history: createWebHashHistory(),
    routes: routes
})
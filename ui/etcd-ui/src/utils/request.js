import axios from "axios";
import { ElMessage} from 'element-plus'

// 创建axios 实例
const service = axios.create({
    baseURL: "http://127.0.0.1:8081",
    timeout: 5000
})

service.interceptors.response.use(
    response=>{
        const res = response.data

        if (res.code!=="0000"){
            ElMessage.error(res.message)
            return
        }else{
            return response
        }
    
    }
)

export default service
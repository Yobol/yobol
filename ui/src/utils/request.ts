import axios from 'axios'

const service = axios.create({
    baseURL: import.meta.env.BASE_URL,
    timeout: 30000, // 30s
})

export default service
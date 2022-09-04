import service from '@/utils/request'

// @Param data {username: "string", password: "string", remember: "bool"}
export const login = (data: any) => {
    return service({
        url: '/api/v1/auth/login',
        method: 'post',
        data: data
    })
}
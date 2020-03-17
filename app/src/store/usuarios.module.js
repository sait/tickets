import instance from '../utils/api'

export const state = {
    usuarios: [],
    usuario: null,
    msg: null,
    errorMsg: null,
    authData: null
}

export const mutations = {
    saveUsuarios(state, usuarios){
        state.usuarios = usuarios
    },
    saveUsuario(state, usuario){
        state.usuario = usuario
    },
    saveMsg(state, msg) {
        state.msg = msg
    },
    saveErrorMsg(state, errorMsg){
        state.errorMsg = errorMsg
    },
    saveAuthData(state, authData){
        state.authData = authData
    }
}

export const actions = {
    getUsuarios(context) {
        instance.get(`/usuarios`).then(response => {
            context.commit('saveUsuarios', response.data)
        }).catch(e => {
            context.commit('saveErrorMsg', e.reponse)
        })
    },
    getUsuariosByEmail(context, email) {
        instance.get(`/usuarios?email=${email}`).then(response => {
            context.commit('saveUsuario', response.data)
        }).catch(e => {
            context.commit('saveErrorMsg', e.reponse)
        })
    },
    createUsuario(context, data) {
        instance.post(`/usuarios`, data).then(response => {
            context.commit('saveMsg', response.data)
        }).catch(e => {
            context.commit('saveErrorMsg', e.reponse)
        })
    },
    modifyUsuario(context, id, data) {
        instance.put(`/usuarios/${id}`, data).then(response => {
            context.commit('saveMsg', response.data)
        }).catch(e => {
            context.commit('saveErrorMsg', e.reponse)
        })
    },
    iniciarSesion(context, data) {
        instance.get(`/auth/usuarios`, { params: data }).then(response => {
            context.commit('saveAuthData', response)
        }).catch(e => {
            context.commit('saveErrorMsg', e.reponse)
        })
    },
    cerrarSesion(context) {
        instance.get(`/clearCookie`).then(response => {
            context.commit('saveMsg', response.data)
        }).catch(e => {
            context.commit('saveErrorMsg', e.reponse)
        })
    }
}

export const getters = {
    getUsuarios: state => {
        return state.tickets
    },
    getUsuario: state => {
        return  state.ticket
    },
    getMsg: state => {
        return state.msg
    },
    getErrorMsg: state => {
        return state.errorMsg
    },
    getAuthData: state => {
        return state.authData
    }
}

export default {
    state,
    mutations,
    actions,
    getters
}
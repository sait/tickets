import instance from '../utils/api'

export const state = {
    agentes: [],
    agente: null,
    msg: null,
    errorMsg: null,
    authData: null
}

export const mutations = {
    saveAgentes(state, agentes){
        state.agentes = agentes
    },
    saveAgente(state, agente){
        state.agente = agente
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
    getAgentes(context) {
        instance.get(`/agentes`).then(response => {
            context.commit('saveAgentes', response.data)
        }).catch(e => {
            context.commit('saveErrorMsg', e.reponse)
        })
    },
    inicarSesion(context, data) {
        instance.get(`/auth/agentes`, { params: data }).then(response => {
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
    },
    createAgente(context, data) {
        instance.post(`/agentes`, data).then(response => {
            context.commit('saveMsg', response.data)
        }).catch(e => {
            context.commit('saveErrorMsg', e.reponse)
        })
    },
    modifyAgente(context, id, data) {
        return instance.put(`/agentes/${id}`, data).then(response => {
            context.commit('saveMsg', response.data)
        }).catch(e => {
            context.commit('saveErrorMsg', e.reponse)
        })
    }
}

export const getters = {
    getAgentes: state => {
        return state.agentes
    },
    getAgente: state => {
        return  state.agente
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
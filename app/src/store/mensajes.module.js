import instance from '../utils/api'

export const state = {
    mensajes: [],
    mensaje: null,
    msg: null,
    errorMsg: null
}

export const mutations = {
    saveMensajes(state, mensajes){
        state.mensajes = mensajes
    },
    saveMsg(state, msg) {
        state.msg = msg
    },
    saveErrorMsg(state, errorMsg){
        state.errorMsg = errorMsg
    }
}

export const actions = {
    getAllMensajesByTicket(context, id) {
        instance.get(`/tickets/${id}/mensajes`).then(response => {
            context.commit('saveMensajes', response.data)
        }).catch(e => {
            context.commit('saveErrorMsg', e.reponse)
        })
    },
    createMensaje(context, id, data) {
        instance.post(`/tickets/${id}/mensajes`, data).then(response => {
            context.commit('saveMsg', response.data)
        }).catch(e => {
            context.commit('saveErrorMsg', e.reponse)
        })
    }
}

export const getters = {
    getMensajes: state => {
        return state.mensajes
    },
    getMsg: state => {
        return state.msg
    },
    getErrorMsg: state => {
        return state.errorMsg
    }
}

export default {
    state,
    mutations,
    actions,
    getters
}
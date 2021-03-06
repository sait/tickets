import instance from '../utils/api'

export const state = {
    tickets: [],
    ticket: null,
    msg: null,
    errorMsg: null
}

export const mutations = {
    saveTickets(state, tickets){
        state.tickets = tickets
    },
    saveTicket(state, ticket){
        state.ticket = ticket
    },
    saveMsg(state, msg) {
        state.msg = msg
    },
    saveErrorMsg(state, errorMsg){
        state.errorMsg = errorMsg
    }
}

export const actions = {
    fetchTickets(context) {
        instance().get(`/tickets`).then(response => {
            context.commit('saveTickets', response.data)
        }).catch(e => {
            context.commit('saveErrorMsg', e.reponse)
        })
    },
    getTicketByID(context, id) {
        instance.get(`/tickets/${id}`).then(response => {
            context.commit('saveTicket', response.data)
        }).catch(e => {
            context.commit('saveErrorMsg', e.reponse)
        })
    },
    createTicket(context, data) {
        instance.post(`/tickets`, data).then(response => {
            context.commit('saveMsg', response.data)
        }).catch(e => {
            context.commit('saveErrorMsg', e.reponse)
        })
    },
    changeEstadoTicket(context, id, data) {
        instance.patch(`/tickets/${id}`, data).then(response => {
            context.commit('saveMsg', response.data)
        }).catch(e => {
            context.commit('saveErrorMsg', e.reponse)
        })
    },
    getTicketsByUsuario(context, id) {
        instance.get(`/usuarios/${id}/tickets`).then(response => {
            context.commit('saveTickets', response.data)
        }).catch(e => {
            context.commit('saveErrorMsg', e.reponse)
        })
    },
    changeAgenteTicket(context, id, data) {
        instance.put(`/tickets/${id}`, data).then(response => {
            context.commit('saveMsg', response.data)
        }).catch(e => {
            context.commit('saveErrorMsg', e.reponse)
        })
    }
}

export const getters = {
    getTickets: state => {
        return state.tickets
    },
    getTicket: state => {
        return  state.ticket
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
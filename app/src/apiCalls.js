import axios from 'axios'

const instance = axios.create({
    baseURL: 'http://localhost:3000',
    headers: { 'Token': localStorage.getItem("Token") }
})

export default {
    tickets: {
        getTickets() {
            return instance.get(`/tickets`)
        },
        getTicketByID(id) {
            return instance.get(`/tickets/${id}`)
        },
        createTicket(data) {
            return instance.post(`/tickets`, data)
        },
        changeEstadoTicket(id, data) {
            return instance.patch(`/tickets/${id}`, data)
        },
        getTicketsByUsuario(id) {
            return instance.get(`/usuarios/${id}/tickets`)
        },
        changeAgenteTicket(id, data) {
            return instance.put(`/tickets/${id}`, data)
        }
    },
    agentes: {
        getAgentes() {
            return instance.get(`/agentes`)
        },
        createAgente(data) {
            return instance.post(`/agentes`, data)
        },
        iniciarSesion(data) {
            return instance.get(`/auth/agentes`, { params: data })
        },
        cerrarSesion() {
            return instance.get(`/clearCookie`)
        },
        modifyAgente(id, data) {
            return instance.put(`/agentes/${id}`, data)
        },
    },
    usuarios: {
        getUsuarios() {
            return instance.get(`/usuarios`)
        },
        getUsuarioByEmail(email) {
            return instance.get(`/usuarios?email=${email}`)
        },
        createUsuario(data) {
            return instance.post(`/usuarios`, data)
        },
        modifyUsuario(id, data) {
            return instance.put(`/usuarios/${id}`, data)
        },
        iniciarSesion(data) {
            return instance.get(`/auth/usuarios`, { params: data })
        },
        cerrarSesion() {
            return instance.get(`/clearCookie`)
        }
    },
    mensajes: {
        createMensaje(id, data) {
            return instance.post(`/tickets/${id}/mensajes`, data)
        },
        getAllMensajesByTicket(id) {
            return instance.get(`/tickets/${id}/mensajes`)
        }
    }
}
import axios from 'axios'

const baseURL = 'http://localhost:3000'

export default {
    tickets: {
        getTickets() {
            axios.defaults.headers.common['Token'] = localStorage.getItem("Token")
            return axios.get(`${baseURL}/tickets`)
        },
        getTicketByID(id) {
            axios.defaults.headers.common['Token'] = localStorage.getItem("Token")
            console.log(axios.defaults.headers.common)
            return axios.get(`${baseURL}/tickets/${id}`)
        },
        createTicket(data) {
            axios.defaults.headers.common['Token'] = localStorage.getItem("Token")
            return axios.post(`${baseURL}/tickets`, data)
        },
        changeEstadoTicket(id, data) {
            axios.defaults.headers.common['Token'] = localStorage.getItem("Token")
            return axios.patch(`${baseURL}/tickets/${id}`, data)
        },
        getTicketsByUsuario(id) {
            axios.defaults.headers.common['Token'] = localStorage.getItem("Token")
            return axios.get(`${baseURL}/usuarios/${id}/tickets`)
        },
        changeAgenteTicket(id, data) {
            axios.defaults.headers.common['Token'] = localStorage.getItem("Token")
            return axios.put(`${baseURL}/tickets/${id}`, data)
        }
    },
    agentes: {
        getAgentes() {
            axios.defaults.headers.common['Token'] = localStorage.getItem("Token")
            return axios.get(`${baseURL}/agentes`)
        },
        createAgente(data) {
            axios.defaults.headers.common['Token'] = localStorage.getItem("Token")
            return axios.post(`${baseURL}/agentes`, data)
        },
        iniciarSesion(data) {
            axios.defaults.headers.common['Token'] = localStorage.getItem("Token")
            return axios.get(`${baseURL}/auth/agentes`, {params: data})
        },
        cerrarSesion() {
            axios.defaults.headers.common['Token'] = localStorage.getItem("Token")
            return axios.get(`${baseURL}/clearCookie`)
        },
        modifyAgente(id, data) {
            axios.defaults.headers.common['Token'] = localStorage.getItem("Token")
            return axios.put(`${baseURL}/agentes/${id}`, data)
        },
    },
    usuarios: {
        getUsuarios() {
            axios.defaults.headers.common['Token'] = localStorage.getItem("Token")
            return axios.get(`${baseURL}/usuarios`)
        },
        getUsuarioByEmail(email) {
            axios.defaults.headers.common['Token'] = localStorage.getItem("Token")
            return axios.get(`${baseURL}/usuarios?email=${email}`)
        },
        createUsuario(data) {
            axios.defaults.headers.common['Token'] = localStorage.getItem("Token")
            return axios.post(`${baseURL}/usuarios`, data)
        },
        modifyUsuario(id, data) {
            axios.defaults.headers.common['Token'] = localStorage.getItem("Token")
            return axios.put(`${baseURL}/usuarios/${id}`, data)
        },
        iniciarSesion(data) {
            axios.defaults.headers.common['Token'] = localStorage.getItem("Token")
            return axios.get(`${baseURL}/auth/usuarios`, {params: data})
        },
        cerrarSesion() {
            axios.defaults.headers.common['Token'] = localStorage.getItem("Token")
            return axios.get(`${baseURL}/clearCookie`)
        }
    },
    mensajes: {
        createMensaje(id, data) {
            axios.defaults.headers.common['Token'] = localStorage.getItem("Token")
            return axios.post(`${baseURL}/tickets/${id}/mensajes`, data)
        },
        getAllMensajesByTicket(id) {
            axios.defaults.headers.common['Token'] = localStorage.getItem("Token")
            return axios.get(`${baseURL}/tickets/${id}/mensajes`)
        }
    }
}
import store from '../store/index.js'
const exRegEmail = /^(([^<>()[\]\\.,;:\s@"]+(\.[^<>()[\]\\.,;:\s@"]+)*)|(".+"))@((\[[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}])|(([a-zA-Z\-0-9]+\.)+[a-zA-Z]{2,24}))$/
const exRegNumbers = /^\d+$/

export default {
    validador: {
        isEmail(email) {
            return exRegEmail.test(email)
        },
        isNotNull(text) {
            return text.length > 0
        },
        isNumeroTelefono(tel) {
            if (tel.length > 6 && tel.length < 11) {
                return exRegNumbers.test(tel)
            }
            return false
        },
        isPassword(pass) {
            return pass.length > 7
        }
    },
    sesion: {
        cerrar() {
            localStorage.removeItem("Token");
            store.commit("setEmail", null);
            store.commit("setTipoUsuario", null);
            store.commit("setID", null)
            localStorage.removeItem('client')
        }
    }
}
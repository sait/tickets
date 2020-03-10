
export const state = {
    email: null,
    tipoUsuario: null,
    id: null
  }
  
export const mutations = {
    setEmail(state, email) {
      state.email = email
    },
    setTipoUsuario(state, tipo) {
      state.tipoUsuario = tipo
    },
    setID(state, id) {
      state.id = id
    }
  }
export const getters = {
    getEmail: state => {
      return state.email
    },
    getTipoUsuario: state => {
      return state.tipoUsuario
    },
    getID: state => {
      return state.id
    }
  }

  export default {
    state,
    mutations,
    getters
  }
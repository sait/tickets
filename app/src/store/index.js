import Vue from 'vue'
import Vuex from 'vuex'

Vue.use(Vuex)

export default new Vuex.Store({
  state: {
    email: null,
    tipoUsuario: null
  },
  mutations: {
    setEmail(state, email) {
      state.email = email
    },
    setTipoUsuario(state, tipo) {
      state.tipoUsuario = tipo
    }
  },
  getters: {
    getEmail: state => {
      return state.email
    },
    getTipoUsuario: state => {
      return state.tipoUsuario
    }
  },
  actions: {
  },
  modules: {
  }
})

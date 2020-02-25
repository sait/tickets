import Vue from 'vue'
import Vuex from 'vuex'
import  createPersistedState  from  'vuex-persistedstate'

Vue.use(Vuex)

export default new Vuex.Store({
  state: {
    email: null,
    tipoUsuario: null,
    id: null
  },
  mutations: {
    setEmail(state, email) {
      state.email = email
    },
    setTipoUsuario(state, tipo) {
      state.tipoUsuario = tipo
    },
    setID(state, id) {
      state.id = id
    }
  },
  getters: {
    getEmail: state => {
      return state.email
    },
    getTipoUsuario: state => {
      return state.tipoUsuario
    },
    getID: state => {
      return state.id
    }
  },
  actions: {
  },
  modules: {
  },
  plugins: [createPersistedState()]
})

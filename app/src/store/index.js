import Vue from "vue";
import Vuex from "vuex";
import createPersistedState from 'vuex-persistedstate'

import logged from "./logged.module";
import tickets from "./tickets.module";
import agentes from "./agentes.module";
import usuarios from "./usuarios.module";
import mensajes from "./mensajes.module";

Vue.use(Vuex);

export default new Vuex.Store({
  modules: {
    logged,
    tickets,
    agentes,
    usuarios,
    mensajes
  },
  plugins: [createPersistedState()]
});
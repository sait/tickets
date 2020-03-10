import Vue from "vue";
import Vuex from "vuex";
import createPersistedState from 'vuex-persistedstate'

import logged from "./logged.module";

Vue.use(Vuex);

export default new Vuex.Store({
  modules: {
    logged
  },
  plugins: [createPersistedState()]
});
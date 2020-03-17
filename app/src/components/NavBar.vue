<template>
  <div>
    <b-navbar toggleable="lg" type="dark" variant="primary">
      <b-navbar-brand to="/inicio" v-if="false || tipoUsuario == 'Agente'">{{ $t('navInicio') }}</b-navbar-brand>
      <b-navbar-brand
        :to="`/usuarios/inicio`"
        v-if="false || tipoUsuario == 'Usuario'"
      >{{ $t('navInicio') }}</b-navbar-brand>
      <b-navbar-brand
        to="/"
        v-if="false || tipoUsuario != 'Usuario' && tipoUsuario != 'Agente'"
      >{{ $t('navInicio') }}</b-navbar-brand>
      <b-navbar-toggle target="nav-collapse"></b-navbar-toggle>

      <b-collapse id="nav-collapse" is-nav>
        <b-navbar-nav>
          <b-nav-item
            to="/historial"
            v-if="false || tipoUsuario == 'Agente'"
          >{{ $t('navHistorial') }}</b-nav-item>
          <b-nav-item
            to="/agentes"
            v-if="false || tipoUsuario == 'Agente'"
          >{{ $t('navAgentes') }}</b-nav-item>
          <b-nav-item
            to="/usuarios"
            v-if="false || tipoUsuario == 'Agente'"
          >{{ $t('navUsuarios') }}</b-nav-item>
        </b-navbar-nav>

        <!-- Right aligned nav items -->
        <b-navbar-nav class="ml-auto" v-if="$store.getters.getEmail">
          <b-nav-item-dropdown right>
            <!-- Using 'button-content' slot -->
            <template v-slot:button-content>{{$store.getters.getEmail}}</template>
            <b-dropdown-item @click="cerrarSesion">{{ $t('navCerrarSesion') }}</b-dropdown-item>
          </b-nav-item-dropdown>
        </b-navbar-nav>
      </b-collapse>
    </b-navbar>
    <div class="locale-changer">
    <select v-model="$i18n.locale">
      <option v-for="(lang, i) in langs" :key="`Lang${i}`" :value="lang">{{ lang }}</option>
    </select>
  </div>
  </div>
</template>

<script>
import apiCalls from '../apiCalls';
export default {
  data () {
    return { langs: ['en', 'es'] }
  },
  methods: {
    cerrarSesion() {
      localStorage.removeItem("Token");
      this.$store.commit("setEmail", null);
      this.$store.commit("setTipoUsuario", null);
      this.$store.commit("setID", null)
      this.$router.push("/");
      localStorage.removeItem('client')
      //Reparar esto, solo es temporal
      apiCalls.agentes.cerrarSesion().catch(() =>{})
      apiCalls.usuarios.cerrarSesion().catch(() =>{})
    }
  },
  mounted() {
    console.log("Hola")
    this.$store.getters.getTipoUsuario
  },
  computed: {
    tipoUsuario () {
      return this.$store.getters.getTipoUsuario
    }
  }
};
</script>
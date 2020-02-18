<template>
  <div>
    <b-navbar toggleable="lg" type="dark" variant="primary">
      <b-navbar-brand to="/inicio" v-if="false || $store.getters.getTipoUsuario == 'Agente'">Inicio</b-navbar-brand>
      <b-navbar-brand
        to="/inicio/2"
        v-if="false || $store.getters.getTipoUsuario == 'Usuario'"
      >Inicio</b-navbar-brand>
      <b-navbar-brand
        to="/"
        v-if="false || $store.getters.getTipoUsuario != 'Usuario' && $store.getters.getTipoUsuario != 'Agente'"
      >Inicio</b-navbar-brand>
      <b-navbar-toggle target="nav-collapse"></b-navbar-toggle>

      <b-collapse id="nav-collapse" is-nav>
        <b-navbar-nav>
          <b-nav-item
            to="/historial"
            v-if="false || $store.getters.getTipoUsuario == 'Agente'"
          >Historial</b-nav-item>
          <b-nav-item
            to="/agentes"
            v-if="false || $store.getters.getTipoUsuario == 'Agente'"
          >Agentes</b-nav-item>
          <b-nav-item
            to="/usuarios"
            v-if="false || $store.getters.getTipoUsuario == 'Agente'"
          >Usuarios</b-nav-item>
        </b-navbar-nav>

        <!-- Right aligned nav items -->
        <b-navbar-nav class="ml-auto" v-if="$store.getters.getEmail">
          <b-nav-item-dropdown right>
            <!-- Using 'button-content' slot -->
            <template v-slot:button-content>{{$store.getters.getEmail}}</template>
            <b-dropdown-item @click="cerrarSesion">Cerrar sesión</b-dropdown-item>
          </b-nav-item-dropdown>
        </b-navbar-nav>
      </b-collapse>
    </b-navbar>
  </div>
</template>

<script>
import apiCalls from '../apiCalls';
export default {
  methods: {
    cerrarSesion() {
      localStorage.removeItem("Token");
      this.$store.commit("setEmail", null);
      this.$store.commit("setTipoUsuario", null);
      this.$router.push("/");
      localStorage.removeItem('client')
      //Reparar esto, solo es temporal
      apiCalls.agentes.cerrarSesion().catch(() =>{})
      apiCalls.usuarios.cerrarSesion().catch(() =>{})
    }
  }
};
</script>
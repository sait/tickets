<template>
  <div class="container col-sm-6">
    <form-registro @registrarse="registrarUsuario" />
  </div>
</template>

<script>
import FormRegistro from "@/components/FormRegistro.vue";
import apiCalls from "../apiCalls";
export default {
  components: {
    FormRegistro
  },
  methods: {
    registrarUsuario(usuario) {
        console.log(usuario)
      apiCalls.usuarios
        .createUsuario(usuario)
        .then(response => {
          if (response.status == 201) {
            this.iniciarSesion(usuario);
          } else {
            console.log(response);
          }
        })
        .catch(e => {
          console.log(e.response);
        });
    },
    iniciarSesion(usuario) {
      apiCalls.usuarios
        .iniciarSesion(usuario)
        .then(response => {
          localStorage.Token = response.headers.token;
          if (localStorage.getItem("Token")) {
            this.$store.commit("setID", response.data.id);
            this.$router.push(`/usuarios/inicio`);
          }
        })
        .catch(e => {
          console.log(e)
        });
    }
  }
};
</script>
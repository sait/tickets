<template>
  <div class="container col-sm-5">
    <b-card>
      <div v-if="isBadCredentials">
        <p id="alert-text">*Revise su correo electrónico y su contraseña</p>
      </div>
      <form-inicio-sesion @iniciar_sesion="onIniciarSesion"></form-inicio-sesion>
    </b-card>
  </div>
</template>

<script>
import FormInicioSesion from "@/components/FormInicioSesion.vue";
import apiCalls from "../apiCalls";
export default {
  data() {
    return {
      isBadCredentials: false
    };
  },
  components: {
    FormInicioSesion
  },
  methods: {
    onIniciarSesion(agente) {
      apiCalls.agentes
        .iniciarSesion(agente)
        .then(response => {
          localStorage.Token = response.headers.token;
          if (localStorage.getItem("Token")) {
            this.$store.commit("setID", response.data.id)
            this.$router.push("/inicio");
          }
        })
        .catch(() => {
          this.isBadCredentials = true;
        });
    }
  }
};
</script>
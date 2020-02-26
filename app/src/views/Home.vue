<template>
  <div class="container col-sm-6">
    <b-card>
      <div class="row">
        <div class="col-sm-5">
          <div v-if="isBadCredentials">
        <p id="alert-text">*Revise su correo electrónico y su contraseña</p>
      </div>
          <form-inicio-sesion @iniciar_sesion="isActive" />
        </div>
        <div class="col-sm-7">
          <p>
            ¿Aún no estás registrado?
            <router-link to="/registro">Cree una cuenta</router-link>
          </p>
          <p>
            <strong>Soy un agente</strong> -
            <router-link to="/login/agentes">Acceda aquí</router-link>
          </p>
          <b-button variant="primary" to="/tickets">Abrir nuevo ticket</b-button>
        </div>
      </div>
    </b-card>
  </div>
</template>

<script>
import FormInicioSesion from "@/components/FormInicioSesion.vue";
import apiCalls from "../apiCalls";
export default {
  components: {
    FormInicioSesion
  },
  data() {
    return {
      usuario: {
        email: "",
        password: ""
      },
      isBadCredentials: false
    };
  },
  methods: {
    onIniciarSesion(usuario) {
        apiCalls.usuarios
        .iniciarSesion(usuario)
        .then(response => {
          localStorage.Token = response.headers.token;
          if (localStorage.getItem("Token")) {
            this.$store.commit("setID", response.data.id)
            this.$router.push(`/usuarios/inicio`);
          }
        })
        .catch((e) => {
          console.log(e.response)
          this.isBadCredentials = true;
        });
    },
    isActive(usuario) {
      apiCalls.usuarios.getUsuarioByEmail(usuario.email).then(response => {
        if(response.data.estado == 1) {
          this.onIniciarSesion(usuario)
        } else {
          this.isBadCredentials = true
        }
      }).catch(() => {
        this.isBadCredentials = true
      })
    }
  }
};
</script>
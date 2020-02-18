<template>
  <div>
    <b-form-group label="Correo electrónico:">
      <b-form-input v-model="agente.email" type="email" placeholder="ej. example@gmail.com"></b-form-input>
    </b-form-group>
    <b-form-group label="Contraseña">
      <b-form-input v-model="agente.password" type="password" placeholder="********"></b-form-input>
    </b-form-group>
    <b-button variant="primary" @click="retornarDatos">Iniciar sesión</b-button>
  </div>
</template>

<script>
export default {
  data() {
    return {
      agente: {
        email: "",
        password: ""
      }
    };
  },
  methods: {
    retornarDatos(){
      this.$emit('iniciar_sesion', this.agente)
    }
  }
}
</script>
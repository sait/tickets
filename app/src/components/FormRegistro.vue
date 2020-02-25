<template>
  <div>
    <b-form-group label="Nombre completo:*">
      <b-form-input v-model="usuario.nombre" :state="clickOnRegistrar ? validarNombre : null"></b-form-input>
    </b-form-group>

    <div class="row">
      <div class="col-sm-7">
        <b-form-group label="Número de teléfono:*">
          <b-form-input v-model="usuario.telefono" :state="clickOnRegistrar ? validarTelefono : null"></b-form-input>
        </b-form-group>
      </div>
      <div class="col-sm-5">
        <b-form-group label="EXT:">
          <b-form-input v-model="usuario.extension" placeholder="Extensión"></b-form-input>
        </b-form-group>
      </div>
    </div>
    <b-form-group label="Correo electrónico:*">
      <b-form-input v-model="usuario.email" type="email" placeholder="ej. example@gmail.com" :state="clickOnRegistrar ? validarEmail : null"></b-form-input>
    </b-form-group>
    <b-form-group label="Contraseña:*">
      <b-form-input v-model="usuario.password" type="password" placeholder="********" :state="clickOnRegistrar ? validarPassword : null"></b-form-input>
    </b-form-group>
    <b-button variant="primary" @click="retornarDatos">Registrarse</b-button>
  </div>
</template>

<script>
import utilerias from "../utils/utilerias.js";
export default {
  data() {
    return {
      usuario: {
        email: "",
        password: "",
        nombre: "",
        telefono: "",
        extension: ""
      },
      clickOnRegistrar: false
    };
  },
  methods: {
    retornarDatos() {
      this.clickOnRegistrar = true
      if(this.validarCampos()){
        this.$emit("registrarse", this.usuario);
      }
    },
    validarCampos() {
      if(this.validarNombre && this.validarEmail && this.validarPassword && this.validarTelefono){
        return true
      }
      return false
    }
  },
  computed: {
    validarNombre() {
      return utilerias.validador.isNotNull(this.usuario.nombre)
    },
    validarTelefono() {
      return utilerias.validador.isNumeroTelefono(this.usuario.telefono);
    },
    validarEmail() {
      return utilerias.validador.isEmail(this.usuario.email);
    },
    validarPassword() {
      return utilerias.validador.isPassword(this.usuario.password);
    }
  }
};
</script>
<template>
  <div>
    <b-modal
      id="modal-nuevo-usuario"
      title="Agregar nuevo usuario"
      centered
      @hidden="cleanModal()"
      no-close-on-esc
      no-close-on-backdrop
      hide-header-close
    >
      <h6>Nombre*</h6>
      <b-input-group>
        <b-input placeholder="ej. Félix González Morales" v-model="nombre" :state="clickOnAgregar ? validarNombre : null" autofocus />
      </b-input-group>
      <br />
      <h6>Correo Electronico*</h6>
      <b-input-group>
        <b-input placeholder="ej. example@gmail.com" v-model="email" :state="clickOnAgregar ? validarEmail : null"/>
      </b-input-group>
      <br />
      <h6>Contraseña*</h6>
      <b-form-group description="Debe contener al menos 8 dígitos">
        <b-input-group>
          <b-input type="password" placeholder="********" v-model="password" :state="clickOnAgregar ? validarPassword : null"/>
        </b-input-group>
      </b-form-group>
      <br />
      <div class="row">
        <div class="col-sm-7">
          <h6>Teléfono*</h6>
          <b-input-group>
            <b-input placeholder="ej. 6532635658" v-model="telefono" :state="clickOnAgregar ? validarTelefono : null"/>
          </b-input-group>
        </div>
        <div class="col-sm-5">
          <h6>EXT</h6>
          <b-input-group>
            <b-input placeholder="ej. 12" v-model="extension" />
          </b-input-group>
        </div>
      </div>
      <template v-slot:modal-footer>
        <b-button size="sm" variant="danger" @click="closeModal()">Cancelar</b-button>
        <b-button size="sm" variant="success" @click="addUsuario()">Agregar</b-button>
      </template>
    </b-modal>
  </div>
</template>

<script>
import utilerias from "../utils/utilerias.js";
import apiCalls from "../apiCalls";
export default {
  data() {
    return {
      nombre: "",
      email: "",
      password: "",
      telefono: "",
      extension: "",
      clickOnAgregar: false
    };
  },
  methods: {
    addUsuario() {
      this.clickOnAgregar = true
      if(this.validarCampos()){
        const datosUsuario = {
        nombre: this.nombre,
        email: this.email,
        password: this.password,
        telefono: this.telefono,
        extension: this.extension
      };
      apiCalls.usuarios
        .createUsuario(datosUsuario)
        .then(response => {
          if (response.status == 201) {
            this.closeModal();
          } else {
            console.log(response);
          }
        })
        .catch(e => {
          console.log(e);
        });
      }
    },
    closeModal() {
      this.$bvModal.hide("modal-nuevo-usuario");
    },
    cleanModal() {
      this.nombre = "";
      this.email = "";
      this.password = "";
      this.telefono = "";
      this.extension = "";
      this.clickOnAgregar = false;
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
      return utilerias.validador.isNotNull(this.nombre);
    },
    validarEmail() {
      return utilerias.validador.isEmail(this.email);
    },
    validarPassword() {
      return utilerias.validador.isPassword(this.password);
    },
    validarTelefono() {
      return utilerias.validador.isNumeroTelefono(this.telefono)
    }
  }
};
</script>
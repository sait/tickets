<template>
  <div>
    <div class="col-md-7 mx-auto">
      <div class="container text-left">
        <b-card
          header-text-variant="white"
          header-tag="header"
          header-bg-variant="primary"
          footer-tag="footer"
          footer-bg-variant="light"
          footer-border-variant="light"
          title="Mensaje:"
        >
          <template v-slot:header>
            <h4 class="mb-0">Responder</h4>
          </template>
          <b-form-textarea
            v-model="mensaje.descripcion"
            rows="4"
            :state="clickOnEnviar ? validarMensaje : null"
          ></b-form-textarea>
          <template v-slot:footer>
            <h4 class="mb-0"></h4>
            <div class="text-right">
              <b-button variant="outline-primary" @click="enviarMensaje">Enviar</b-button>
            </div>
          </template>
        </b-card>
      </div>
    </div>
  </div>
</template>

<script>
import utilerias from "../utils/utilerias.js";
import apiCalls from "../apiCalls";
export default {
  data() {
    return {
      mensaje: {
        emisor: 1,
        descripcion: "",
        ticketID: 0
      },
      clickOnEnviar: false
    };
  },
  methods: {
    //Falta implementar response y catch errors
    enviarMensaje() {
      this.clickOnEnviar = true;
      if (this.validarMensaje) {
        apiCalls.mensajes
          .createMensaje(this.mensaje.ticket_id, this.mensaje)
          .then(response => {
            console.log(response);
          })
          .catch(e => {
            console.log(e.response);
          });
        this.clickOnEnviar = false;
      }
    }
  },
  computed: {
    validarMensaje() {
      return utilerias.validador.isNotNull(this.mensaje.descripcion);
    }
  },
  mounted() {
    this.mensaje.ticket_id = Number.parseInt(this.$route.params.id);
    console.log(this.mensaje.ticket_id);
  }
};
</script>
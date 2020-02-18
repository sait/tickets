<template>
  <div>
    <div v-for="mensaje in mensajes" :key="mensaje.id" class="col-md-7 mx-auto">
      <div class="container text-left">
        <b-card
          header-text-variant="white"
          header-tag="header"
          header-bg-variant="dark"
          footer-tag="footer"
          footer-bg-variant="light"
          footer-border-variant="light"
          title="Mensaje:"
        >
          <b-card-text>{{mensaje.descripcion}}</b-card-text>
          <template v-slot:footer>
            <h4 class="mb-0">{{mensaje.ticket_id}}</h4>
          </template>
        </b-card>
      </div>
      <br />
    </div>
  </div>
</template>

<script>
import apiCalls from "../apiCalls";
export default {
  data() {
    return {
      ticketID: null,
      mensajes: []
    };
  },
  mounted() {
    this.ticketID = this.$route.params.id;
    this.obtenerMensajes();
  },
  methods: {
    obtenerMensajes() {
      this.mensajes = [];
      if (this.ticketID) {
        apiCalls.mensajes
          .getAllMensajesByTicket(this.ticketID)
          .then(response => {
              console.log(response)
            this.mensajes = response.data;
          })
          .catch(e => {
            console.log(e.response);
          });
      }
    }
  }
};
</script>
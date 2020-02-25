<template>
  <div v-if="datosTicket">
    <div class="col-md-7 mx-auto">
      <div class="container text-left">
        <b-card
          header-text-variant="white"
          header-tag="header"
          header-bg-variant="dark"
          footer-tag="footer"
          footer-bg-variant="light"
          footer-border-variant="light"
          :title="`Asunto: ${datosTicket.titulo}`"
        >
          <template v-slot:header>
            <div class="row">
              <div class="col-sm-6">
                <h4 class="mb-0">Ticket #{{ticketID}}</h4>
              </div>
              <div v-if="false || $store.getters.getTipoUsuario == 'Agente'" class="col-sm-6 text-right">
                <b-form-select
                  :options="estadosDeTickets"
                  v-model="datosTicket.estado"
                  @change="changeStatusTicket"
                ></b-form-select>
              </div>
            </div>
          </template>
          <b-card-text>
            <strong>Descripción:</strong>
          </b-card-text>
          <b-card-text>{{datosTicket.descripcion}}</b-card-text>
          <template v-slot:footer>
            <h4 class="mb-0">{{datosTicket.usuario.nombre}}</h4>
          </template>
        </b-card>
      </div>
    </div>
  </div>
</template>

<script>
import apiCalls from "../apiCalls";
export default {
  data() {
    return {
      ticketID: null,
      datosTicket: null,
      estadosDeTickets: [
        { text: "Cerrado", value: 0 },
        { text: "Pendiente", value: 1 },
        { text: "En proceso", value: 2 }
      ]
    };
  },
  mounted() {
    this.ticketID = this.$route.params.id;
    this.obtenerDatosTicket();
  },
  methods: {
    obtenerDatosTicket() {
      if (this.ticketID) {
        apiCalls.tickets.getTicketByID(this.ticketID).then(response => {
          this.datosTicket = response.data;
          this.$root.$emit('datos-ticket', response.data); 
        });
      }
    },
    //Falta manejar el response y los errores
    changeStatusTicket() {
      apiCalls.tickets.changeEstadoTicket(this.ticketID, this.datosTicket).then(response => {
        console.log(response)
      }).catch(e => {
        console.log(e.response)
      })
    }
  }
};
</script>
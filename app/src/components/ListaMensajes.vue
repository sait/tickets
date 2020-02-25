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
            <h4 class="mb-0">{{mensaje.emisor == 1 ? nombreUsuario : nombreAgente}}</h4> 
          </template>
        </b-card>
      </div>
      <br />
    </div>
  </div>
</template>

<script>
import apiCalls from "../apiCalls";
import utilerias from '../utils/utilerias';
export default {
  data() {
    return {
      ticketID: null,
      mensajes: [],
      nombreUsuario: null,
      nombreAgente: null,
      isAgenteAsignado: true,
      datosTicket: null
    };
  },
  mounted() {
    this.$root.$on('datos-ticket', data => { 
      this.datosTicket = data
      this.obtenerMensajes();
      console.log(data)
      this.nombreUsuario = data.usuario.nombre
      this.nombreAgente = data.agente.nombre
      this.nombreAgente ? this.isAgenteAsignado = true : this.isAgenteAsignado = false
    }); 
    this.ticketID = this.$route.params.id;
    this.$root.$on('refresh-mensajes', () => { 
    this.obtenerMensajes()
    }); 
  },
  methods: {
    obtenerMensajes() {
      console.log(this.isAgenteAsignado)
      if(!this.isAgenteAsignado && this.$store.getters.getTipoUsuario == "Agente") {
        console.log("Sm")
        this.datosTicket.agente_id = this.$store.getters.getID
        apiCalls.tickets.changeAgenteTicket(this.ticketID, this.datosTicket).then(() => {
          this.$router.go()
        })
      }
      this.mensajes = [];
      if (this.ticketID) {
        apiCalls.mensajes
          .getAllMensajesByTicket(this.ticketID)
          .then(response => {
              console.log(response)
            this.mensajes = response.data;
          })
          .catch(e => {
            if(e.response.data == "Token is expired") {
              utilerias.sesion.cerrar()
              this.$router.push("/")
            }
          });
      }
    }
  }
};
</script>
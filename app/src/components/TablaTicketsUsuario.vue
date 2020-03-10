<template>
  <div>
    <b-table-simple outlined hover responsive>
      <b-thead head-variant="light">
        <b-tr>
          <b-th>ID</b-th>
          <b-th>Asunto</b-th>
          <b-th>Fecha de envío</b-th>
          <b-th>Nombre agente</b-th>
          <b-th>Estado</b-th>
        </b-tr>
      </b-thead>
      <b-tbody>
        <b-tr v-for="ticket in tickets" :key="ticket.id" @click="openTicket(ticket.id)">
          <b-td>{{ticket.id}}</b-td>
          <b-td>{{ticket.titulo}}</b-td>
          <b-td>{{ticket.created_at}}</b-td>
          <b-td>{{ticket.agente.nombre}}</b-td>
          <b-td>{{estadosDeTickets[ticket.estado]}}</b-td>
        </b-tr>
      </b-tbody>
    </b-table-simple>
  </div>
</template>

<script>
import apiCalls from "../apiCalls";
export default {
  data() {
    return {
      rows: 0,
      currentPage: 1,
      tickets: [],
      estadosDeTickets: [
        "Cerrado", //0
        "Pendiente", //1
        "En proceso" //2
      ]
    };
  },
  methods: {
    openTicket(id) {
      this.$router.push(`/tickets/${id}`);
    }
  },
  mounted() {
    apiCalls.initConf()
    if (this.$store.getters.getTipoUsuario == "Usuario" || this.$store.getters.getTipoUsuario == null) {
      apiCalls.tickets
        .getTicketsByUsuario(this.$store.getters.getID)
        .then(response => {
          this.tickets = response.data;
          this.$store.commit("setEmail", response.headers.email);
          this.$store.commit("setTipoUsuario", response.headers.client);
        })
        .catch(e => {
          if (e.response.status == 401) {
            localStorage.removeItem("Token");
          } else if (e.response.status == 400) {
            if (e.response.data == "Token is expired") {
              localStorage.removeItem("Token");
            }
          }
          this.$router.go();
        });
    } else {
      this.$router.push("/");
    }
  }
};
</script>

<style scoped>
</style>

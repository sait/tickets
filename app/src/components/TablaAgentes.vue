<template>
  <b-table-simple outlined hover responsive>
    <b-thead head-variant="light">
      <b-tr>
        <b-th>ID</b-th>
        <b-th>Nombre</b-th>
        <b-th>Correo electrónico</b-th>
        <b-th>Estado</b-th>
      </b-tr>
    </b-thead>
    <b-tbody>
      <b-tr v-for="agente in agentes" :key="agente.id">
        <b-td>{{agente.id}}</b-td>
        <b-td>{{agente.nombre}}</b-td>
        <b-td>{{agente.email}}</b-td>
        <b-td>{{estadosDeAgentes[agente.estado]}}</b-td>
      </b-tr>
    </b-tbody>
  </b-table-simple>
</template>

<script>
import apiCalls from "../apiCalls";
export default {
  data() {
    return {
      agentes: [],
      estadosDeAgentes: [
        "Inactivo", //0
        "Activo" //1
      ]
    };
  },
  mounted() {
    apiCalls.agentes
      .getAgentes()
      .then(response => {
        this.agentes = response.data;
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
  }
};
</script>
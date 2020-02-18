<template>
  <b-table-simple outlined hover responsive>
    <b-thead head-variant="light">
      <b-tr>
        <b-th>ID</b-th>
        <b-th>Nombre</b-th>
        <b-th>Correo electrónico</b-th>
        <b-th>Teléfono</b-th>
        <b-th>Estado</b-th>
      </b-tr>
    </b-thead>
    <b-tbody>
      <b-tr v-for="usuario in usuarios" :key="usuario.id">
        <b-td>{{usuario.id}}</b-td>
        <b-td>{{usuario.nombre}}</b-td>
        <b-td>{{usuario.email}}</b-td>
        <b-td>{{usuario.telefono}}</b-td>
        <b-td>{{estadosDeUsuarios[usuario.estado]}}</b-td>
      </b-tr>
    </b-tbody>
  </b-table-simple>
</template>

<script>
import apiCalls from "../apiCalls";
export default {
  data() {
    return {
      usuarios: [],
      estadosDeUsuarios: [
        "Inactivo", //0
        "Activo" //1
      ]
    };
  },
  mounted() {
    apiCalls.usuarios
      .getUsuarios()
      .then(response => {
        this.usuarios = response.data;
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
<template>
  <div>
    <vue-bootstrap4-table :rows="rows" :columns="columns" :config="config">
      <template slot="estado" slot-scope="props">
        <b-button @click="changeStatus(props.row)">{{estadosDeAgentes[props.cell_value]}}</b-button>
      </template>
      <template slot="pagination-info" slot-scope="props">
        Mostrando {{props.currentPageRowsLength}} resultados |
        {{props.filteredRowsLength}} resultados en total
    </template>
    <template slot="simple-filter-clear-icon">
        <b-icon-x-circle></b-icon-x-circle>
    </template>
    <template slot="refresh-button-text">
        <b-icon-arrow-repeat></b-icon-arrow-repeat>
    </template>
    <template slot="empty-results">
        Registros no encontrados
    </template>
    </vue-bootstrap4-table>
  </div>
</template>

<script>
import VueBootstrap4Table from "vue-bootstrap4-table";
import apiCalls from "../apiCalls";
import utilerias from "../utils/utilerias";
export default {
  data() {
    return {
      estadosDeAgentes: [
        "Inactivo", //0
        "Activo" //1
      ],
      rows: [],
      columns: [
        {
          label: "ID",
          name: "id",
          filter: {
            type: "simple",
            placeholder: "ID"
          },
          uniqueId: true,
          sort: true,
          initial_sort: true,
          initial_sort_order: "desc"
        },
        {
          label: "Nombre",
          name: "nombre",
          filter: {
            type: "simple",
            placeholder: "Nombre"
          }
        },
        {
          label: "Correo electrónico",
          name: "email",
          filter: {
            type: "simple",
            placeholder: "Correo electrónico"
          }
        },
        {
          label: "Estado",
          name: "estado",
          sort: true
        }
      ],
      config: {
        rows_selectable: true,
        global_search: {
          visibility: false
        },
        show_reset_button: false,
        show_refresh_button: false,
        card_mode: false
      }
    };
  },
  components: {
    VueBootstrap4Table
  },
  methods: {
    changeStatus(selectedRow) {
      this.rows.find(element => element.id == selectedRow.id).estado == 1
        ? (this.rows.find(element => element.id == selectedRow.id).estado = 0)
        : (this.rows.find(element => element.id == selectedRow.id).estado = 1);
      apiCalls.agentes
        .modifyAgente(
          selectedRow.id,
          this.rows.find(element => element.id == selectedRow.id)
        )
        .then(() => {})
        .catch(e => {
          console.log(e.response);
        });
    }
  },
  mounted() {
    apiCalls.agentes
      .getAgentes()
      .then(response => {
        this.rows = response.data;
        this.$store.commit("setEmail", response.headers.email);
        this.$store.commit("setTipoUsuario", response.headers.client);
      })
      .catch(e => {
        if (e.response.status == 401) {
          localStorage.removeItem("Token");
        } else if (e.response.status == 400) {
          if (e.response.data == "Token is expired") {
            utilerias.sesion.cerrar()
          }
        }
        this.$router.go();
      });
  }
};
</script>
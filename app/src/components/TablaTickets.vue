<template>
  <div>
    <vue-bootstrap4-table
      :rows="rows"
      :columns="columns"
      :config="config"
      @on-select-row="openTicket($event)"
    >
      <template slot="estado" slot-scope="props">{{estadosDeTickets[props.cell_value]}}</template>
      <template slot="pagination-info" slot-scope="props">
        {{$t('tableMos')}} {{props.currentPageRowsLength}} {{$t('tableRes')}} |
        {{props.filteredRowsLength}} {{$t('tableRes')}} {{$t('tableTot')}}
      </template>
      <template slot="created_at" slot-scope="props">
        <pre id="format-pre"><strong>{{$t('tableFecha')}}: </strong>{{props.cell_value.split("T")[0]+"\n"}}</pre>
        <strong>{{$t('tableHora')}}:</strong> {{props.cell_value.split("T")[1].split("-")[0]}}
      </template>
      <template slot="simple-filter-clear-icon">
        <b-icon-x-circle></b-icon-x-circle>
      </template>
      <template slot="refresh-button-text">
        <b-icon-arrow-repeat></b-icon-arrow-repeat>
      </template>
      <template slot="empty-results">{{$t('tableRecNotFound')}}</template>
    </vue-bootstrap4-table>
  </div>
</template>
 
<script>
import VueBootstrap4Table from "vue-bootstrap4-table";
import apiCalls from "../apiCalls";
import utilerias from "../utils/utilerias";
export default {
  data: function() {
    return {
      estadosDeTickets: [
        this.$t('tableCerrado'), //0
        this.$t('tablePend'), //1
        this.$t('tableEnProc') //2
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
          label: "Asunto",
          name: "titulo"
        },
        {
          label: "Fecha de envío",
          name: "created_at",
          sort: true
        },
        {
          label: "Nombre usuario",
          name: "usuario.nombre",
          filter: {
            type: "simple",
            placeholder: "Usuario"
          }
        },
        {
          label: "Nombre agente",
          name: "agente.nombre",
          filter: {
            type: "simple",
            placeholder: "Agente"
          }
        },
        {
          label: "Estado",
          name: "estado",
          filter: {
            type: "select",
            mode: "multi",
            closeDropdownOnSelection: true,
            placeholder: "Select options",
            options: [
              {
                name: "Cerrado",
                value: "0"
              },
              {
                name: "Pendiente",
                value: "1"
              },
              {
                name: "En proceso",
                value: "2"
              }
            ],
            init: {
              value: [1, 2]
            }
          },
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
    openTicket(selectedRow) {
      this.$router.push(`/tickets/${selectedRow.selected_item.id}`);
    }
  },
  mounted() {
    apiCalls.initConf()
    if (
      this.$store.getters.getTipoUsuario == "Agente" ||
      this.$store.getters.getTipoUsuario == null
    ) {
      apiCalls.tickets
        .getTickets()
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
              utilerias.sesion.cerrar();
            }
          }
          if(e.response.status != 500){
            this.$router.go();
          }
        });
    } else {
      this.$router.push("/usuarios/inicio");
    }
    this.$store.dispatch('fetchTickets')
    console.log(this.$store.getters.getTickets)
  }
};
</script>

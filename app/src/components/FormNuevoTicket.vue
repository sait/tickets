<template>
  <div class="container">
    <div class="row">
      <div class="col-sm-12">
        <b-card>
          <div class="row">
            <div class="col-sm-6">
              <b-form-group label="Temas de ayuda:*">
                <b-form-select
                  v-model="ticket.tema_ayuda"
                  :options="temasAyuda"
                  :state="clickOnSubmit ? validarTemaAyuda : null"
                ></b-form-select>
              </b-form-group>
            </div>

            <div class="col-sm-6">
              <b-form-group label="Producto:*">
                <b-form-select
                  v-model="ticket.producto"
                  :options="productos"
                  :state="clickOnSubmit ? validarProducto : null"
                ></b-form-select>
              </b-form-group>
            </div>
          </div>
        </b-card>
      </div>
    </div>
    <br />
    <div class="row">
      <div class="col-sm-5">
        <b-card title="Detalle orden de servicio">
          <b-form-group label="Título:*">
            <b-form-input v-model="ticket.titulo" :state="clickOnSubmit ? validarTitulo : null"></b-form-input>
          </b-form-group>

          <b-form-group label="Descripción:">
            <b-form-textarea
              v-model="ticket.descripcion"
              placeholder="Detalles del problema"
              rows="4"
            ></b-form-textarea>
          </b-form-group>
        </b-card>
      </div>

      <div class="col-sm-7">
        <b-card title="Información de contacto">
          <b-form-group label="Correo electrónico:*">
            <b-form-input
              v-model="usuario.email"
              type="email"
              placeholder="ej. example@gmail.com"
              :state="clickOnSubmit ? validarEmail : null"
            ></b-form-input>
          </b-form-group>

          <b-form-group label="Nombre completo:*">
            <b-form-input v-model="usuario.nombre" :state="clickOnSubmit ? validarNombre : null"></b-form-input>
          </b-form-group>

          <div class="row">
            <div class="col-sm-7">
              <b-form-group label="Número de teléfono:*">
                <b-form-input
                  v-model="usuario.telefono"
                  :state="clickOnSubmit ? validarTelefono : null"
                ></b-form-input>
              </b-form-group>
            </div>
            <div class="col-sm-5">
              <b-form-group label="EXT:">
                <b-form-input v-model="usuario.extension" placeholder="Extensión"></b-form-input>
              </b-form-group>
            </div>
          </div>
        </b-card>
      </div>
    </div>

    <br />

    <div class="row">
      <div class="col-sm-12">
        <b-card>
          <div class="row">
            <div class="col-sm-6">
              <b-button variant="danger" @click="restablecer()">Restablecer</b-button>
            </div>
            <div class="col-sm-6">
              <b-button variant="primary" @click="verificarUsuario()">Crear ticket</b-button>
            </div>
          </div>
        </b-card>
      </div>
    </div>
    <br />
  </div>
</template>

<script>
import utilerias from "../utils/utilerias.js";
import apiCalls from "../apiCalls";
export default {
  data() {
    return {
      ticket: {
        usuario_id: 0,
        agente_id: 0,
        titulo: "",
        descripcion: "",
        tema_ayuda: null,
        producto: null
      },
      usuario: {
        email: "",
        nombre: "",
        telefono: "",
        extension: ""
      },
      temasAyuda: [
        { text: "Seleccione un tema de ayuda", value: null },
        "Bóveda",
        "IVA al 8%",
        "CFDI 3.3",
        "Soporte",
        "Solicitar ayuda",
        "Ventas",
        "Cotización",
        "Licencia",
        "Otro asunto"
      ],
      productos: [
        { text: "Seleccione producto", value: null },
        "SAIT Básico",
        "SAIT ERP",
        "SAIT Contabilidad",
        "SAIT Nómina",
        "Bóveda (OCF)",
        "Factura 123",
        "Enlace de sucursales"
      ],
      clickOnSubmit: false
    };
  },
  methods: {
    verificarUsuario() {
      this.clickOnSubmit = true;
      if (this.validarCampos()) {
        apiCalls.usuarios
          .getUsuarioByEmail(this.usuario.email)
          .then(response => {
            this.ticket.usuario_id = response.data.id;
            apiCalls.usuarios
              .modifyUsuario(this.ticket.usuario_id, this.usuario)
              .then(() => {
                this.crearTicket();
              })
              .catch(e => {
                console.log(e);
              });
          })
          .catch(e => {
            if (e.response.data == "record not found") {
              apiCalls.usuarios
                .createUsuario(this.usuario)
                .then(response => {
                  this.ticket.usuario_id = response.data.id;
                  this.crearTicket();
                })
                .catch(e => {
                  console.log(e);
                });
            }
          });
      }
    },
    crearTicket() {
      apiCalls.tickets
        .createTicket(this.ticket)
        .then(response => {
          this.$bvModal
            .msgBoxOk(
              `El ticket ha sido enviado correctamente.\nTicket #${response.data.id}`,
              {
                title: "Enviado",
                size: "sm",
                okTitle: "Aceptar",
                okVariant: "primary",
                noCloseOnEsc: true,
                noCloseOnBackdrop: true,
                centered: true,
                autoFocusButton: "ok"
              }
            )
            .then(value => {
              if(value) {
                this.$router.push("/")
              }
            })
            .catch(err => {
              console.log(err)
            });
          this.restablecer();
        })
        .catch(e => {
          console.log(e);
        });
    },
    restablecer() {
      this.usuario.email = "";
      this.usuario.nombre = "";
      this.usuario.telefono = "";
      this.usuario.extension = "";

      this.ticket.titulo = "";
      this.ticket.descripcion = "";
      this.ticket.tema_ayuda = null;
      this.ticket.producto = null;
      this.ticket.usuario_id = 0;
      this.ticket.agente_id = 0;

      this.clickOnSubmit = false;
    },
    validarCampos() {
      if (
        this.validarNombre &&
        this.validarTelefono &&
        this.validarEmail &&
        this.validarTemaAyuda &&
        this.validarProducto &&
        this.validarTitulo
      ) {
        return true;
      }
      return false;
    }
  },
  computed: {
    validarNombre() {
      return utilerias.validador.isNotNull(this.usuario.nombre)
    },
    validarTelefono() {
      return utilerias.validador.isNumeroTelefono(this.usuario.telefono);
    },
    validarEmail() {
      return utilerias.validador.isEmail(this.usuario.email);
    },
    validarTemaAyuda() {
      return this.ticket.tema_ayuda != null;
    },
    validarProducto() {
      return this.ticket.producto != null;
    },
    validarTitulo() {
      return utilerias.validador.isNotNull(this.ticket.titulo)
    }
  }
};
</script>
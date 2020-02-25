<template>
  <div>
    <div class="container text-right">
      <b-button variant="outline-primary" v-b-modal.modal-nuevo-usuario>Agregar nuevo usuario</b-button>
    </div>
    <br>
    <br>
    <div class="container"><tabla-usuarios/></div>
    <modal-nuevo-usuario></modal-nuevo-usuario>
  </div>
</template>

<script>
import TablaUsuarios from '@/components/TablaUsuarios.vue'
import ModalNuevoUsuario from '@/components/ModalNuevoUsuarioForm.vue'
export default {
  components: {
    ModalNuevoUsuario,
    TablaUsuarios
  }
}
</script>
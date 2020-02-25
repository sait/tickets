<template>
  <div>
    <div class="container text-right">
      <b-button variant="outline-primary" v-b-modal.modal-nuevo-agente>Agregar nuevo agente</b-button>
    </div>
    <br>
    <br>
    <div class="container">
      <tabla-agentes/>
    </div>
    <modal-nuevo-agente />
  </div>
</template>

<script>
import TablaAgentes from '@/components/TablaAgentes.vue'
import ModalNuevoAgente from '@/components/ModalNuevoAgenteForm.vue'
export default {
  components: {
    ModalNuevoAgente,
    TablaAgentes
  }
}
</script>
<style>

</style>
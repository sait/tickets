<template>
  <div class="container">
    <div class="text-right">
      <b-button variant="outline-primary" to="/tickets">Nuevo ticket</b-button>
    </div>
    <br>
    <tabla-tickets-usuario></tabla-tickets-usuario>
  </div>
</template>

<script>
import TablaTicketsUsuario from '@/components/TablaTicketsUsuario.vue'
export default {
  components: {
    TablaTicketsUsuario
  }
};
</script>
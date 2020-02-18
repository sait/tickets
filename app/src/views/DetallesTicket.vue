<template>
  <div>
    <detalle-ticket/>
    <br>
    <lista-mensajes/>
    <responder-ticket-form/>
    <br>
  </div>
</template>

<script>
import DetalleTicket from "@/components/DetalleTicket.vue";
import ResponderTicketForm from "@/components/ResponderTicketForm.vue";
import ListaMensajes from "@/components/ListaMensajes.vue";
export default {
  components: {
    DetalleTicket,
    ResponderTicketForm,
    ListaMensajes
  }
};
</script>
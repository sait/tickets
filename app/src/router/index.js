import Vue from 'vue'
import VueRouter from 'vue-router'
import Home from '../views/Home.vue'
import LoginAgentes from '../views/InicioSesionAgentes.vue'
import HomeAgentes from '../views/HomeAgentes.vue'
import HomeUsuarios from '../views/HomeUsuarios.vue'
import Historial from '../views/Historial.vue'
import Agentes from '../views/CuentasAgentes.vue'
import Usuarios from '../views/CuentasUsuarios.vue'
import Tickets from '../views/NuevosTickets.vue'
import DetallesTicket from '../views/DetallesTicket.vue'

Vue.use(VueRouter)

const routes = [
  {
    path: '/',
    name: 'home',
    component: Home,
    meta: { Auth: false, title: 'Home' },
      beforeEnter: (to, from, next) => {
          if (window.localStorage.getItem('Token')) {
            next({ path: '/inicio' });
          } else {
            next();
          }
      },
  },
  {
    path: '/login/agentes',
    name: 'loginAgentes',
    component: LoginAgentes,
    meta: { Auth: false, title: 'LoginAgentes' },
      beforeEnter: (to, from, next) => {
          if (window.localStorage.getItem('Token')) {
            next({ path: '/inicio' });
          } else {
            next();
          }
      },
  },
  {
    path: '/inicio',
    name: 'inicio',
    component: HomeAgentes,
    meta: { Auth: true, title: 'Inicio' }
  },
  {
    path: '/historial',
    name: 'historial',
    component: Historial,
    meta: { Auth: true, title: 'Historial' }
  },
  {
    path: '/agentes',
    name: 'agentes',
    component: Agentes,
    meta: { Auth: true, title: 'Agentes' }
  },
  {
    path: '/usuarios',
    name: 'usuarios',
    component: Usuarios,
    meta: { Auth: true, title: 'Usuarios' }
  },
  {
    path: '/inicio/2',
    name: 'homeUsuarios',
    component: HomeUsuarios,
    meta: { Auth: true, title: 'Inicio' }
  },
  {
    path: '/tickets',
    name: 'tickets',
    component: Tickets
    // meta: { Auth: false, title: 'Tickets' },
    //   beforeEnter: (to, from, next) => {
    //       if (window.localStorage.getItem('Token')) {
    //         next({ path: '/inicio' });
    //       } else {
    //         next();
    //       }
    //   },
  },
  {
    path: '/tickets/:id',
    name: 'ticket',
    component: DetallesTicket,
    meta: { Auth: true, title: 'Ticket' }
  }
]

const router = new VueRouter({
  mode: 'history',
  base: process.env.BASE_URL,
  routes
})

router.beforeEach((to, from, next) => {
  document.title = to.meta.title;
  if (to.meta.Auth && !window.localStorage.getItem('Token')) {
    next({ path: '/' });
  } else {
    next();
  }
});

export default router

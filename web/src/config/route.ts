import {createMemoryHistory, createRouter} from 'vue-router'
import LoginView from "../page/Login.vue";
import HomeView from "../page/HomeView.vue";


const routes = [
    {path: '/', component: HomeView},
    {path: '/login', component: LoginView},
]

export const router = createRouter({
    history: createMemoryHistory(),
    routes,
})

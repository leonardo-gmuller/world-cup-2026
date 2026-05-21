import { createRouter, createWebHistory } from 'vue-router'

import AppLayout from '@/layouts/AppLayout.vue'

import LoginPage from '@/pages/auth/LoginPage.vue'
import RegisterPage from '@/pages/auth/RegisterPage.vue'

import HomePage from '@/pages/app/HomePage.vue'
import GroupPage from '@/pages/app/GroupPage.vue'
import MatchesPage from '@/pages/app/MatchesPage.vue'
import PredictionsPage from '@/pages/app/PredictionsPage.vue'
import RankingPage from '@/pages/app/RankingPage.vue'
import ProfilePage from '@/pages/app/ProfilePage.vue'
import RulesPage from '@/pages/app/RulesPage.vue'

const routes = [
  { path: '/', redirect: '/app/home' },

  { path: '/login', component: LoginPage },
  { path: '/register', component: RegisterPage },

  {
    path: '/app',
    component: AppLayout,
    children: [
      { path: '', redirect: '/app/home' },
      { path: 'home', component: HomePage, meta: { title: 'Home • Bolão Copa 2026'} },
      { path: 'groups', component: GroupPage, meta: { title: 'Grupos • Bolão Copa 2026'} },
      { path: 'matches', component: MatchesPage, meta: { title: 'Jogos • Bolão Copa 2026'} },
      { path: 'predictions', component: PredictionsPage, meta: { title: 'Palpites • Bolão Copa 2026'} },
      { path: 'ranking', component: RankingPage, meta: { title: 'Ranking • Bolão Copa 2026'} },
      { path: 'profile', component: ProfilePage, meta: { title: 'Perfil • Bolão Copa 2026'} },
      { path: 'rules', component: RulesPage, meta: { title: 'Regras • Bolão Copa 2026'} }
    ],
  },
]



const router = createRouter({
  history: createWebHistory(),
  routes,
})

router.afterEach((to) => {
  document.title = to.meta.title || 'Bolão Copa 2026'
})

export default router
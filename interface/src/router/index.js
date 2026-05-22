import { createRouter, createWebHistory } from 'vue-router'
import { useAuthStore } from '@/stores/auth'

import AppLayout from '@/layouts/AppLayout.vue'

import AuthPage from '@/pages/auth/AuthPage.vue'

import HomePage from '@/pages/app/HomePage.vue'
import GroupPage from '@/pages/app/GroupPage.vue'
import MatchesPage from '@/pages/app/MatchesPage.vue'
import PredictionsPage from '@/pages/app/PredictionsPage.vue'
import RankingPage from '@/pages/app/RankingPage.vue'
import ProfilePage from '@/pages/app/ProfilePage.vue'
import RulesPage from '@/pages/app/RulesPage.vue'

const routes = [
    { path: '/', redirect: '/app/home' },

    {
        path: '/login',
        component: AuthPage,
        meta: { guestOnly: true },
    },
    {
        path: '/register',
        component: AuthPage,
        meta: { guestOnly: true },
    },

    {
        path: '/app',
        component: AppLayout,
        children: [
            { path: '', redirect: '/app/home' },
            { path: 'home', component: HomePage, meta: { title: 'Home • Bolão Copa 2026', requiresAuth: true } },
            { path: 'groups', component: GroupPage, meta: { title: 'Grupos • Bolão Copa 2026', requiresAuth: true } },
            { path: 'matches', component: MatchesPage, meta: { title: 'Jogos • Bolão Copa 2026', requiresAuth: true } },
            { path: 'predictions', component: PredictionsPage, meta: { title: 'Palpites • Bolão Copa 2026', requiresAuth: true } },
            { path: 'ranking', component: RankingPage, meta: { title: 'Ranking • Bolão Copa 2026', requiresAuth: true } },
            { path: 'profile', component: ProfilePage, meta: { title: 'Perfil • Bolão Copa 2026', requiresAuth: true } },
            { path: 'rules', component: RulesPage, meta: { title: 'Regras • Bolão Copa 2026', requiresAuth: true } }
        ],
    },
]



const router = createRouter({
    history: createWebHistory(),
    routes,
    scrollBehavior() {
        return {
            top: 0,
            behavior: 'smooth',
        }
    }
})

router.beforeEach((to) => {
    const auth = useAuthStore()

    if (to.path === '/login' || to.path === '/register') {
        if (auth.isAuthenticated) {
            return '/app/home'
        }
        return true
    }

    if (to.meta.requiresAuth && !auth.isAuthenticated) {
        return '/login'
    }

    if (to.meta.guestOnly && auth.isAuthenticated) {
        return '/app/home'
    }
})

router.afterEach((to) => {
    document.title = to.meta.title || 'Bolão Copa 2026'
})

export default router
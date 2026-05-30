<template>
    <section class="app-page">
        <div>
            <h1 class="app-title">Meus grupos</h1>
            <p class="app-subtitle">
                Crie um bolão ou entre usando um código de convite.
            </p>
        </div>

        <div class="rounded-3xl bg-white p-4 shadow-sm">
            <h2 class="font-semibold text-slate-900">Criar grupo</h2>

            <form class="mt-4 space-y-3" @submit.prevent="submitCreate">
                <InputText v-model="createForm.name" placeholder="Nome do grupo" fluid />

                <Textarea v-model="createForm.description" placeholder="Descrição opcional" rows="2" fluid />

                <Button type="submit" label="Criar grupo" class="w-full" :loading="groupStore.loading" />
            </form>
        </div>

        <div class="rounded-3xl bg-white p-4 shadow-sm">
            <h2 class="font-semibold text-slate-900">Entrar em grupo</h2>

            <form class="mt-4 flex gap-2" @submit.prevent="submitJoin">
                <InputText v-model="inviteCode" placeholder="Código" class="flex-1" />

                <Button type="submit" icon="pi pi-arrow-right" :loading="groupStore.loading" />
            </form>
        </div>

        <div class="space-y-3">
            <GroupCard v-for="group in groupStore.groups" :key="group.id" :group="group" />

            <div v-if="!groupStore.loading && groupStore.groups.length === 0"
                class="rounded-3xl bg-white p-6 text-center text-sm text-slate-500 shadow-sm">
                Você ainda não está em nenhum grupo.
            </div>
        </div>
    </section>
</template>

<script setup>
import { onMounted, reactive, ref } from 'vue'
import Button from 'primevue/button'
import InputText from 'primevue/inputtext'
import Textarea from 'primevue/textarea'
import GroupCard from '@/components/GroupCard.vue'
import { useGroupStore } from '@/stores/group'
import { useAppToast } from '@/utils/toast'

const toast = useAppToast()

const groupStore = useGroupStore()

const inviteCode = ref('')

const createForm = reactive({
    name: '',
    description: '',
})

onMounted(() => {
    groupStore.fetchGroups()
})

async function submitCreate() {
    if (!createForm.name) return

    try {
        await groupStore.create({
            name: createForm.name,
            description: createForm.description || null,
        })

        toast.success('Grupo criado com sucesso')
    } catch (error) {
        toast.error(error, 'Erro ao criar grupo')
    }
}

async function submitJoin() {
    if (!inviteCode.value) return

    await groupStore.join(inviteCode.value)
    toast.success('Entrou no grupo com sucesso')
    inviteCode.value = ''
}
</script>
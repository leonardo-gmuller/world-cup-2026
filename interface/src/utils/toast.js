import { useToast } from 'primevue/usetoast'

export function useAppToast() {
  const toast = useToast()

  function success(detail, summary = 'Sucesso') {
    toast.add({
      severity: 'success',
      summary,
      detail,
      life: 3000,
    })
  }

  function error(detail, summary = 'Erro') {
    toast.add({
      severity: 'error',
      summary,
      detail,
      life: 4000,
    })
  }

  function info(detail, summary = 'Informação') {
    toast.add({
      severity: 'info',
      summary,
      detail,
      life: 3000,
    })
  }

  function warn(detail, summary = 'Atenção') {
    toast.add({
      severity: 'warn',
      summary,
      detail,
      life: 3500,
    })
  }

  return {
    success,
    error,
    info,
    warn,
  }
}
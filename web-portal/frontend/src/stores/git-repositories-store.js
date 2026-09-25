import {ref} from "vue";
import {defineStore} from "pinia";
import GitRepositoriesService from "@/services/backend-api/git-repositories-service.js";

export const useGitRepositoriesStore = defineStore('git-repositories-store', () => {
  // private
  let intervalId = null

  // public
  const repositories = ref([])

  function startPolling() {
    intervalId = setInterval(async () => {
      repositories.value = await GitRepositoriesService.getList()
    }, 2000)
  }

  function stopPolling() {
    clearInterval(intervalId)
  }

  return {repositories, startPolling, stopPolling}
})

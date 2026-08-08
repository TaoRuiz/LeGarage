import {ref} from "vue";
import {defineStore} from "pinia";
import GitReleasesService from "@/services/backend-api/git-releases-service.js";

export const useGitReleasesStore = defineStore('git-releases-store', () => {
  // private
  let intervalId = null

  // public
  const releases = ref([])

  function startPolling() {
    intervalId = setInterval(() => {
      releases.value = GitReleasesService.getList()
    }, 2000)
  }

  function stopPolling() {
    clearInterval(intervalId)
  }

  return {releases, startPolling, stopPolling}
})

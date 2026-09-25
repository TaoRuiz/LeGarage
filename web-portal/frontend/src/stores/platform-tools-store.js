import {ref} from "vue";
import {defineStore} from "pinia";
import PlatformToolsService from "@/services/backend-api/platform-tools-service.js";

export const usePlatformToolsStore = defineStore('platform-tools-store', () => {
  // private
  let intervalId = null

  // public
  const tools = ref([])

  function startPolling() {
    intervalId = setInterval(async () => {
      tools.value = await PlatformToolsService.getList()
    }, 2000)
  }

  function stopPolling() {
    clearInterval(intervalId)
  }

  return {tools, startPolling, stopPolling}
})

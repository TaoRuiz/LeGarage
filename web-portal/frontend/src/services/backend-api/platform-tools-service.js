import {API_BASE_URL} from "@/config/api.js";

export default {
  async getList() {
    const response = await fetch(API_BASE_URL + '/api/platform-tools-services')
    return response.json()
  }
}

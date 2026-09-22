export default {
  async getList() {
    const response = await fetch('http://localhost:8080/api/platform-tools-services')
    return response.json()
  }
}

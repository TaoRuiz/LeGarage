export default {
  async getList() {
    const response = await fetch('http://localhost:8080/api/git-releases')
    return response.json()
  }
}

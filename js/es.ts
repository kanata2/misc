const es = new EventSource('http://localhost:3000')
const node = document.getElementById('app')

es.onopen = () => {
  console.log("connection to server")
}

es.onmessage = (e) => {
  console.log(e.data)
  const p = document.createElement('p')
  p.innerHTML = e.data
  node.appendChild(p)
}

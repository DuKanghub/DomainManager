<template>
  <div id="xterm" />
</template>
<script>
import 'xterm/css/xterm.css'
import '@/style/xterm.css'
import { Terminal } from 'xterm'
import { FitAddon } from 'xterm-addon-fit'
import Base64 from 'crypto-js/enc-base64'
import Utf8 from 'crypto-js/enc-utf8'
const msgData = '1'
const msgResize = '2'
export default {
  name: 'Term',
  mounted() {
    console.log('here')
    const host = this.$route.params.id
    console.log(host)
    console.log(`ws://${process.env.VUE_APP_BASE_DOMAIN}:${process.env.VUE_APP_SERVER_PORT}/ws/${host}`)
    const wsUrl = (location.protocol === 'http:' ? 'ws://' : 'wss://') + location.hostname + ':' + process.env.VUE_APP_SERVER_PORT + '/ws/' + host
    console.log(wsUrl)
    const terminal = new Terminal({})
    const fitAddon = new FitAddon()
    terminal.loadAddon(fitAddon)
    fitAddon.fit()
    const terminalContainer = document.getElementById('xterm')
    const webSocket = new WebSocket(wsUrl)
    webSocket.onmessage = (event) => {
      terminal.write(event.data.toString(Utf8))
    }

    webSocket.onopen = () => {
      terminal.open(terminalContainer)
      fitAddon.fit()
      terminal.write('welcome to WebSSH ☺\r\n')
      terminal.focus()
    }

    webSocket.onclose = () => {
      terminal.write('\r\nWebSSH quit!')
    }

    webSocket.onerror = (event) => {
      console.error(event)
      webSocket.close()
    }

    terminal.onKey((event) => {
      webSocket.send(msgData + Base64.stringify(Utf8.parse(event.key)))
    })

    terminal.onResize(({ cols, rows }) => {
      console.log(cols, rows)
      webSocket.send(msgResize +
            Base64.stringify(
              Utf8.parse(
                JSON.stringify({
                  columns: cols,
                  rows: rows
                })
              )
            )
      )
    })
    // 内容全屏显示-窗口大小发生改变时
    // resizeScreen
    window.addEventListener('resize', () => {
      fitAddon.fit()
    }, false)
  },
  methods: {
  }
}
</script>

<style scoped>
    .active {
        color: #000;
        background: rgb(211, 24, 24);
    }
</style>

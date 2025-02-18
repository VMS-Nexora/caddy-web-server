# Caddyfile
example.com {
    reverse_proxy /tenant1/* localhost:8001
    reverse_proxy /tenant2/* localhost:8002
    reverse_proxy /tenant3/* localhost:8003
}

tenant1.example.com {
    reverse_proxy localhost:8001
}

tenant2.example.com {
    reverse_proxy localhost:8002
}

tenant3.example.com {
    reverse_proxy localhost:8003
}

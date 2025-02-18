bkulab.tech {
    root * /var/www/bkulab.tech
    root * C:i_will_update_path_later
    file_server
    @http {
        protocol http
    }
    redir @http https://{host}{uri}
}

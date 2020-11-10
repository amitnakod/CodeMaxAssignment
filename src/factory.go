package src

func mailprovider(config Config) imailservice {
  switch config.mailprovider {
    case "mailjet" :
          return mailj{publickey:config.publickey,secretkey:config.secretkey}
    default:
        nil
  }
}
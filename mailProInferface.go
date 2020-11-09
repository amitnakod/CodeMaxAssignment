package main

import "mailModels"

type imailservice interface{
  send (msg mailtemplate)
}
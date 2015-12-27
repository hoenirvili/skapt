# Skapt                                                                                                                                                                                                                                                                                      
[![Build Status](https://travis-ci.org/hoenirvili/Skapt.svg)](https://travis-ci.org/hoenirvili/Skapt)                                                                                               
                                                                                                                                                                                                
### Package for building command line apps in Go

![experimental](doc/ref.png)

> I was inspired from other cli frameworks in goLang and for the fun/learning purpose i'm trying to do my own little framework.                                                                       

## This package is still in development

**Note** : This package will support the two main command line patterns.¬                                                                                                                                 
### Sub-Command                                                                                                                                                                                      
**Sub-Command** pattern is the pattern that executable takes sub-command for change its behavior. git command is one example for this pattern or node package manager(npm) It takes push, pull subcommand  s and as for npm init, start, stop, update, upgrade etc.                                                                                                                                            
### Flag                                                                                                                                                                                             
**Flag** pattern is the pattern that executable has flag options for changing its behavior. For example, grep command inherits this pattern. 

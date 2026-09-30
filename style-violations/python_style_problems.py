from os import *


class user_processor:
  def ProcessUser(self,UserName,user_age):
    User_name = UserName.strip()
    print("User: "+User_name+", age: "+str(user_age))

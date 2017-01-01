def get_name_from_user():
    name = raw_input("What's your name ?")
    return name

def make_text( name ) :
    return "Hello, " + name + ", nice to meet you"

def print_text( text ) :
    print text 

if __name__ ==  "__main__" :
    print_text( make_text( get_name_from_user()))









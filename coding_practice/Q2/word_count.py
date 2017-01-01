
def get_name_from_user():
    return raw_input("What's your name ?")

def make_text_with_word_count( name ):
    if len(name) == 0 :
        return "You just pressed enter key with empty string. Please press enter key after typed something"
    else : 
        return name + " has " + str(len(name)) + " characters"

def print_text( text ):
    print text

if __name__ == "__main__":
    print_text( make_text_with_word_count( get_name_from_user()) )



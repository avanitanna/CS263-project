import random

# read random text file 
with open("random_text.txt", "r") as file:
    data = file.read()

# write mutated files
num_of_mutations = 1000
for i in range(num_of_mutations):
    with open(f"random_text_mut_{i}.text", "w") as file:
        # generate random number (representing character index)
        num = random.randint(0,len(data)-1)
        # replace the random character by a
        new_data = data[:num]+"a"+data[num+1:]
        # write new data to file
        file.write(new_data)

    
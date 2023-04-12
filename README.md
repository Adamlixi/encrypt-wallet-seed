# encrypt wallet seed

For now, the mnemonic words is totally random which is hard for people to memorize. 
So, how to store the seed?

I think it would be convenient to store the seed in Cloud. Just apply some encrypt methods to encrypt the seed using a customized password and then upload the encrypt file into Cloud. If you need the seed, just download the file, and use your password to decrypt it.

But it's unsafe. Because, the encrypt file actually contains the information of the seed. And hacker can crack the encrypt file easily. 

I design an algorithm for user to encrypt their seed using customized words while the encrypt file wouldn't contain any informations(just totally random words in it). And this algorithm will help user store their seed in the cloud.


Here's the idea:

We have a word list contains 65536 words.(Almost all the words we are using in daily life). And the order of the words in the word list is fully random.

## Encrypt process:

1. Get the wallet seed which is 128 bits.

2. Users choose 8 - 20 words from the word list. Suppose user choose n words.

3. Then it will use SHA256 to hash the entropy and get a hash string(256bit).

4. Combine the first m bits of the hash string(256 bits) and the seed(128 bits) to get the new entropy(128 + m bits).
the relationship for m and n :
(128 + m) % n = 0

5. Split the new entropy(128 + m bits) into n segments which means there are (128 + m)/n bits for each segment.
 
6. Each segment is corresponding to a word which the user choose. Change each segment(binary number) into decimal number. And this decimal number is the index of the word the user choose.

7. Switch word which the user choose in the word list with the word which the index refer to. And get a new word list file.

The new word list file is the result of encrypted algorithm. And the file doesn't contain any informations. Also, only the customized words user set can get the result. There's no collision in it.

The hacker just have two way to crack it.
1. Brute force. (But the difficulty is as same as brute force the public key).
2. Guessing the password(Hard).

## Decrypt process:
1. Get the customized words(n words) and the word list file.
2. Calculate the length of each segment.(m bits each segment).
3. Get the index of each customized word in word list.
4. Change the index to binary number and combine them together.
5. Remove the part which longer than 128 bits(we add it before).
6. Get the seed.


## Some problems:
1. People may frequently just using small part of the wordlist(maybe 8k words are frequently used), and this makes the entropy for customized words very small.  
We can change the content of the wordlist. Adding more things people like to use. For example, we can add emoji, common Chinese words, common Japanese words,.. etc. We can have a discussion in community. And make all the stuffs in the word list are people frequently use in their life. And this will make the hacker more difficult to guess the password.
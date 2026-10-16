#include "bpe.h"

#include <stdlib.h>
#include <stdio.h>
#include <string.h>

bpe_tokenizer* init_bpe_tokenizer(char* text) {
    bpe_tokenizer *bpe = malloc(sizeof(bpe_tokenizer)); 
    if (bpe == NULL) { return NULL; }

    bpe->text = text;

    return bpe;
}

int pre_tokenization(bpe_tokenizer *bpe) {

    if (bpe == NULL || bpe->text == NULL) { return FAILURE; }

    int word_count = 0;
    int text_length = strlen(bpe->text);

    // break text into unique words
    for (int i = 0; i <= text_length; i++) {
        if(bpe->text[i] == ' ' || bpe->text[i] == '\0') {
            word_count++;
        }
    }

    printf("total words: %d\n", word_count);
}


int main() {
    bpe_tokenizer *bpe = init_bpe_tokenizer("stinky jimmys love to steal");

    if (bpe == NULL) {
        perror("bpe tokenizer failed to init properly");
        return FAILURE;
    }

    pre_tokenization(bpe);
}
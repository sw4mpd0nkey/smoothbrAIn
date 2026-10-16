#ifndef BPE_H
#define BPE_H


typedef enum {
    SUCCESS = 0,
    FAILURE = -1,
} return_t;

typedef struct {
    char *text;
} bpe_tokenizer;

bpe_tokenizer *init_bpe_tokenizer(char *text);

int pre_tokenization(bpe_tokenizer *bpe);



#endif
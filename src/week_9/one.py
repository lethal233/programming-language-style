import numpy as np
from collections import Counter


def read_file(filename):
    return np.array([' '] + list(open(filename, 'r').read()) + [' '])


def normalize(characters):
    characters[~np.char.isalpha(characters)] = ' '
    characters = np.char.upper(characters)
    return characters


def indices(characters):
    spi = np.where(characters == ' ')
    return np.repeat(spi, 2)


def indices_pairs(sp):
    wranges = np.reshape(sp[1:-1], (-1, 2))
    return wranges[np.where(wranges[:, 1] - wranges[:, 0] > 2)]


def recode(characters, w_ranges):
    loop_up_mapping = str.maketrans({
        "A": '4',
        "E": '3',
        "I": '1',
        "O": '0',
    })
    a = zip(w_ranges[:-1, 0], w_ranges[1:, 1])
    w = list(map(lambda z: characters[z[0]:z[1]], a))
    return np.array(list(map(lambda z: ("".join(z).strip().translate(loop_up_mapping)), w)))


def most_common_bigrams(bigrams):
    return Counter(bigrams).most_common(5)


if __name__ == '__main__':
    file = read_file("../pride-and-prejudice.txt")
    file = normalize(file)
    sp = indices(file)
    w_ranges = indices_pairs(sp)
    words = recode(file, w_ranges)

    nps, c, = np.unique(words, axis=0, return_counts=True)
    wf_sorted = sorted(zip(nps, c), key=lambda y: y[1], reverse=True)
    for x in wf_sorted[:5]:
        print(x)

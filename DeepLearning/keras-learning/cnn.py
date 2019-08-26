import numpy as np
from keras.datasets import mnist
from keras.utils import np_utils
from keras.models import Sequential
from keras.optimizers import Adam
from keras.layers import Dense, Activation, Conv2D, MaxPooling2D, Flatten

"""
    website:  https://www.youtube.com/watch?v=zHop6Oq757Y&list=WL&index=19&t=233s
"""


def start():
    np.random.seed(1337)
    (X_train, Y_train), (X_test, Y_test) = mnist.load_data()

    # data  pre-processing
    X_train = X_train.reshape(-1, 1, 28, 28)
    X_test = X_test.reshape(-1, 1, 28, 28)
    Y_train = np_utils.to_categorical(Y_train, nb_classes=10)
    Y_test = np_utils.to_categorical(Y_test, nb_classes=10)

    # build your cnn models
    model = Sequential()

    # conv layer 1 output shape
    model.add(Conv2D(
        nb_filter=32,
        nb_row=5,
        nb_col=5,
        border_model='same',  # padding method
        input_shape=(1, 28, 28)
    ))
    model.add(Activation('relu'))  # 激活一下

    # pooling layer
    model.add(MaxPooling2D(
        pool_size=(2, 2),
        strides=()
    ))

    # another way to define your optimizer
    adam = Adam(lr=1e-4)

    # we add metrics to get more res


if __name__ == '__main__':
    start()

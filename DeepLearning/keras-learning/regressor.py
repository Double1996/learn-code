import numpy as np
from keras.models import Sequential
from keras.layers import Dense
from matplotlib import pyplot as plt


def start():
    # 1.创建数据
    X = np.linspace(-1, 1, 200)
    np.random.shuffle(X)
    Y = 0.5 * X + 2 + np.random.normal(0, 0.05, (200,))

    # 2.展示数据分布
    plt.scatter(X, Y)
    plt.show()

    # 3. 定义训练集和测试集
    X_train, Y_train = X[:160], Y[:160]
    X_test, Y_test = X[160:], Y[160:]

    # 4.build model
    model = Sequential()
    model.add(Dense(output_dim=1, input_dim=1))

    # 5.choose loss function and optimizing method
    model.compile(loss='mse', optimizer='sgd')

    # 6.training
    print('Training--------')

    for step in range(300):
        cost = model.train_on_batch(X_train, Y_train)
        if step % 100 == 0:
            print('train cost:', cost)

    # 7.test
    print('Tesing---------')
    cost = model.evaluate(X_test, Y_test)
    W, b = model.layers[0].get_weights()
    print(W, b)

    # 8.打印测试结果
    Y_pred = model.predict(X_test)
    plt.scatter(X_test, Y_test)
    plt.plot(X_test, Y_test)
    plt.show()


if __name__ == '__main__':
    start()

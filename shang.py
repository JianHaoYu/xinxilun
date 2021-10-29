#-*- coding: UTF-8 -*- 

import numpy as np
import matplotlib.pyplot as plt
import mpl_toolkits.axisartist as axisartist
from mpl_toolkits.mplot3d import Axes3D #画三维图不可少
from matplotlib import cm  #cm 是colormap的简写

def gaussian(x,mu,sigma):
    f_x = 1/(sigma*np.sqrt(2*np.pi))*np.exp(-np.power(x-mu, 2.)/(2*np.power(sigma,2.)))
    return(f_x)

def gaussian_2(x,y):
    f_x_y = x*np.log(1/x) + (y)*np.log(1/y)
    return(f_x_y)
#设置2维表格
x_values = np.linspace(0,1,2000)
x_values = np.delete(x_values, -1,axis = None)
x_values = np.delete(x_values, 0,axis = None)
y_values = np.linspace(0,1,2000)
y_values = np.delete(y_values, -1,axis = None)
y_values = np.delete(y_values, 0,axis = None)
X,Y = np.meshgrid(x_values,y_values)
#高斯函数
F_x_y = gaussian_2(X,Y)
#显示三维图
fig = plt.figure()
ax = plt.gca(projection='3d')
ax.plot_surface(X,Y,F_x_y,cmap='rainbow')
plt.show()


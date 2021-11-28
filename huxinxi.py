#-*- coding: UTF-8 -*- 

import numpy as np
import matplotlib.pyplot as plt
import mpl_toolkits.axisartist as axisartist
from mpl_toolkits.mplot3d import Axes3D #画三维图不可少
from matplotlib import cm  #cm 是colormap的简写

x_values = np.linspace(0,1,2000)
x_values = np.delete(x_values, -1,axis = None)
x_values = np.delete(x_values, 0,axis = None)
yx_values = np.linspace(0,1,2000)
yx_values = np.delete(yx_values, -1,axis = None)
yx_values = np.delete(yx_values, 0,axis = None)
p,q = np.meshgrid(x_values,yx_values)

Hnoise = -p*np.log(p)-(1-p)*np.log(1-p)
x=(1-p)*q +p*(1-q)
I= -x*np.log(x) -(1-x) *np.log((1-x)) - Hnoise


fig = plt.figure()
ax = plt.gca(projection='3d')
ax.plot_surface(p,q,I,alpha= 0.8,cmap='jet')

ax.set_xlabel( 'P(X)')
ax.set_xlim(0, 1) #拉开坐标轴范围显示投影
ax.set_ylabel( 'P(Y|X)')
ax.set_ylim(0, 1)

plt.show()


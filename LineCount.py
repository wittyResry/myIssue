#/usr/bin/python
#-*- coding:utf-8 -*-
from pyspark import SparkContext
sparkContext = SparkContext("local", "Line Count")
lines = sparkContext.textFile("/etc/passwd")
print lines.count()


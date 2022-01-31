import os
os.chdir("/projects/hashmweb/site/")
os.system("gitfolio update")
with open("/projects/hashmweb/site/dist/index.html", "r") as r:
    newcontents = r.readlines()
with open("/projects/hashmweb/old/oldindex.html", "r") as r:
    oldcontents = r.readlines()
for i in range(0,33):
    newcontents[i] = oldcontents[i]
for i in range(len(newcontents)):
    if newcontents[i] == '<a href="https://github.com/imfunniee" target="_blank">made on earth by a human</a>':
        newcontents = '<a>made with </a><a href="https://github.com/imfunniee/gitfolio" target="_blank">gitfolio</a><a> and modified by </a><a href="https://github.com/imfunniee/gitfolio">me</a>'
with open("/projects/hashmweb/site/dist/index.html", "w") as w:
    w.writelines(newcontents)
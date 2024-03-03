#!/usr/bin/env bash
# 添加环境变量

cmd=$1
shell=$2

echo "Installing $cmd completion for $shell..."


function add()
{
  cat << EOF >> ${HOME}/.bashrc
# ${cmd} shell completion
if [ -f \${HOME}/.${cmd}-completion.bash ]; then
    source \${HOME}/.${cmd}-completion.bash
fi
EOF
}
# 生成命令的补全文件
${cmd} completion ${shell} > ${HOME}/.${cmd}-completion.bash
## 判断是否有 .bashrc 文件
if ! grep -q "# ${cmd} shell completion" ${HOME}/.bashrc;then
    echo "Adding completion to ${HOME}/.bashrc"
    add
fi
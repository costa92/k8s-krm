#!/usr/bin/env bash
# 添加环境变量

cmd=$1
shell=$2

echo "Installing $cmd completion for $shell..."
function add_bash()
{
  cat << EOF >> ${HOME}/.bashrc
# ${cmd} shell completion
if [ -f \${HOME}/.${cmd}-completion.bash ]; then
    source \${HOME}/.${cmd}-completion.bash
fi
EOF
}


function add_zsh() {
    cat << EOF >> ${HOME}/.zshrc_completion
# ${cmd} shell completion
if [ -f \${HOME}/.${cmd}-completion.bash ]; then
    source \${HOME}/.${cmd}-completion.bash
fi
EOF
}

# 生成命令的补全文件
${cmd} completion ${shell} > ${HOME}/.${cmd}-completion.bash


if  [ "${shell}" == "bash" ];then
  ## 判断是否有 .bashrc 文件
  if ! grep -q "# ${cmd} shell completion" ${HOME}/.bashrc;then
      echo "Adding completion to ${HOME}/.bashrc"
      add_bash
      echo "Completion added to ${HOME}/.bashrc"
  fi
elif [ "${shell}" == "zsh" ]; then
 if ! grep -q "# ${cmd} shell completion" ${HOME}/.zshrc_completion;then
      echo "Adding completion to ${HOME}/.zshrc"
      add_zsh
 fi
fi



## 判断是否有 .bashrc 文件
if ! grep -q "# ${cmd} shell completion" ${HOME}/.bashrc;then
    echo 11
    echo "Adding completion to ${HOME}/.bashrc"
    if  [ "${shell}" == "bash" ];then
        add_bash
    elif [ "${shell}" == "zsh" ]; then
        add_zsh
    fi
    echo "Completion added to ${HOME}/.bashrc ${shell}"
fi
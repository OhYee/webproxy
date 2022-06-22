#!/bin/bash
BaseImage=node:16-alpine
WebProxyImage=ohyee/proxy

function func_get_commit() {
    if [[ "$(git status | grep 'nothing to commit' | wc -l)" -eq "1" ]]; then
        Commit=$(git rev-parse --short HEAD)
        echo $Commit
    fi
}

function cmd_build() { ### 构建镜像
    Commit=$(func_get_commit)
    if [[ -n $Commit ]]; then
        echo $Commit
        docker build -t ${WebProxyImage}:${Commit} .
        echo "镜像: ${WebProxyImage}:${Commit}"

        backendLatest="$(echo $backendImage | cut -d ":" -f 1):latest"
    else
        echo "存在未提交的更改，请提交后再构建镜像"
    fi
}

function cmd_push() { ### 推送镜像
    Commit=$(func_get_commit)
    echo "最新版本: ${Commit}"
    docker tag  ${WebProxyImage}:${Commit} ${WebProxyImage}:latest
    docker push ${WebProxyImage}:${Commit}
    docker push ${WebProxyImage}:latest
}

##################
#      Core      #
##################

HELP_TEXT="$0\n"
Cmds=()
Helps=()

function cmd_help() { ### 显示帮助
    echo -e "$HELP_TEXT"
}

func_debug() {
    [[ -n ${DEBUG} ]] && echo -e "\e[034;1m[DEBUG]\e[0m \e[34;2m./$BASH_SOURCE:$LINENO$([[ -n $FUNCNAME ]] && echo -n " ${FUNCNAME}()" )\e[0m\n    $@";
}

function func_init() {
    REG='^function\s*cmd_([a-z]+)\(\)\s*\{\s*###\s*(.*)$'
    LINES_CMD="cat ${0} | grep -E '${REG}' | awk '{ match(\$0, /${REG}/, a); printf \"%s %s\\n\", a[1], a[2]; }'"
    LINES=`bash -c "$LINES_CMD"`
    MAX_LEN=0    

    OLDIFS=$IFS
    IFS=$'\n'
    for line in ${LINES[*]}; do
        Cmd=$(echo $line | cut -d " " -f 1)
        Help=$(echo $line | cut -d " " -f 2-)

        Cmds+=($Cmd)
        Helps+=($Help)

        [[ ${#Cmd} -gt $MAX_LEN ]] && MAX_LEN=${#Cmd}
    done
    IFS=$OLDIFS

    for ((i=0; i<${#Cmds[*]}; i++)); do
        Cmd=${Cmds[$i]}
        Help=${Helps[$i]}

        SPACE_LEN=$(expr ${MAX_LEN} - ${#Cmd})
        SPACE=""
        for ((j=0; j<${SPACE_LEN}; j++)); do
            SPACE+=" "
        done

        HELP_TEXT+="    \e[034m$Cmd\e[0m${SPACE}    $Help\n"
    done
}

func_init

for cmd in ${Cmds[*]}; do 
    if [[ "$cmd" == "$1" ]]; then
        func_debug "RUN cmd_${cmd} with args [${@:2}]"
        "cmd_${cmd}" $@
        exit 0
    fi
done

cmd_help


```sh
# params:

deploy_127_0_0_1="true"

# script:

host=127_0_0_1

deploy=deploy_$host
echo '$deploy': $deploy
echo '$deploy_127_0_0_1': $deploy_127_0_0_1
echo '\$$deploy': \$$deploy
echo '$(eval echo \$$deploy)': $(eval echo \$$deploy)
echo "\`eval echo \'\$\'\$deploy\`": `eval echo '$'$deploy`
echo '${!deploy}': ${!deploy}

if [ ${!deploy} = "false" ]; then exit 0; fi

echo start deploying: $host
```

> Ubuntu 20.04 LTS:

```sh
# GNU bash, version 5.0.16(1)-release (x86_64-pc-linux-gnu)

$deploy: deploy_127_0_0_1
$deploy_127_0_0_1: true
\$$deploy: $deploy_127_0_0_1
$(eval echo \$$deploy): true
`eval echo \'$\'$deploy`: true
${!deploy}: true
start deploying: 127_0_0_1

# dash 0.5.10.2-6

➜  test sh ./test.sh
$deploy: deploy_127_0_0_1
$deploy_127_0_0_1: true
\$$deploy: $deploy_127_0_0_1
$(eval echo \$$deploy): true
`eval echo \'$\'$deploy`: true
./test.sh: 15: Bad substitution

# zsh 5.8 (x86_64-ubuntu-linux-gnu)

➜  test zsh ./test.sh
$deploy: deploy_127_0_0_1
$deploy_127_0_0_1: true
\$$deploy: $deploy_127_0_0_1
$(eval echo \$$deploy): true
`eval echo \'$\'$deploy`: true
./test.sh:15: bad substitution
```

> CentOS release 6.5 (Final)

```sh
# GNU bash, version 4.1.2(1)-release (x86_64-redhat-linux-gnu)

$deploy: deploy_127_0_0_1
$deploy_127_0_0_1: true
\$$deploy: $deploy_127_0_0_1
$(eval echo \$$deploy): true
`eval echo \'$\'$deploy`: true
${!deploy}: true
start deploying: 127_0_0_1

# dash-0.5.5.1-4.el6.x86_64

$deploy: deploy_127_0_0_1
$deploy_127_0_0_1: true
\$$deploy: $deploy_127_0_0_1
$(eval echo \$$deploy): true
`eval echo \'$\'$deploy`: true
${!deploy}: true
start deploying: 127_0_0_1
```

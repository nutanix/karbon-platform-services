TEST_GENERATE_FOLDER="generated"
DEBUG_GENERATE_FOLDER="generated-debug"
LOGGING_GENERATE_FOLDER="generated-logging"

DEBUG_NS="zihan"

function create_and_use_context {
    bin_path=$1;shift
    url_name=$1;shift
    if ! $bin_path config create-context --password test -m test@ntnxsherlock.com -u ${url_name}.ntnxsherlock.com ${url_name}-user; then
        echo "failed to set up context"
        return 1
    fi
    if ! $bin_path config use-context ${url_name}-user; then
        return 1
    fi
    return 0
}

function setup {
    bin_path=$1;shift
    name_salt=$1;shift
    mkdir -p $TEST_GENERATE_FOLDER
    create_and_use_context $bin_path test
    if ! create_and_use_context $bin_path test; then
        return 1
    fi
    if [ $($bin_path config get-contexts | grep test-user | wc -l |  xargs) -lt 1 ]; then
        return 1
    fi
}

function teardown {
    bin_path=$1;shift
    name_salt=$1;shift
    
    if ! $bin_path delete application echo-cli-app-$name_salt; then
        return 1
    fi

    if ! $bin_path delete datapipeline pipeline-0-$name_salt pipeline-1-$name_salt pipeline-2-$name_salt; then
        return 1
    fi

    if ! $bin_path delete function echo2-$name_salt detect-$name_salt; then
        return 1
    fi

    if ! $bin_path delete datasource datasource-hls-ifc-deepomatic-$name_salt datasource-cli-test-$name_salt mqtt-test-datasource1-$name_salt mqtt-test-datasource2-$name_salt mqtt-test-datasource-$name_salt; then
        return 1
    fi

    if ! $bin_path delete category mqtt-data1-$name_salt mqtt-data2-$name_salt testyaml-$name_salt; then
        return 1
    fi
}

function create {
    bin_path=$1;shift
    name_salt=$1;shift
    mkdir -p $TEST_GENERATE_FOLDER
    edge_name=`"$bin_path" get edge -c | awk '{print $1}' | grep edge | head -1 |  xargs`
    application_yaml="${TEST_GENERATE_FOLDER}/e2e_with_edge.yaml"
    pipeline_1_yaml="${TEST_GENERATE_FOLDER}/pipeline1.yaml"
    pipeline_2_yaml="${TEST_GENERATE_FOLDER}/pipeline2.yaml"
    sed -e "s/EDGE_NAME/$edge_name/g" -e "s/NAME_SALT/$name_salt/g" application-with-dataifc-create.yaml > $application_yaml
    sed -e "s/EDGE_NAME/$edge_name/g" -e "s/NAME_SALT/$name_salt/g" datapipeline-input-mqtt-output-s3-create.yaml > $pipeline_1_yaml
    sed -e "s/EDGE_NAME/$edge_name/g" -e "s/NAME_SALT/$name_salt/g" datapipeline-concat-output-s3-create.yaml > $pipeline_2_yaml

    if ! $bin_path create -f $application_yaml; then
        return 1
    fi

    if ! $bin_path create -f $pipeline_1_yaml; then
        return 1
    fi

    if ! $bin_path create -f $pipeline_2_yaml; then
        return 1
    fi
}

function assert_output {
    cmd=$1;shift
    expected_cmd=$1;shift
    
    if ! out=$(bash -c "$cmd"); then
        echo "failed to run $cmd"
        return 1
    fi
    
    if ! expected=$(bash -c "$expected_cmd"); then
        echo "failed to run $expected_cmd"
        return 1
    fi

    if [[ $(echo $out | tr -d '[:space:]') != $(echo $expected | tr -d '[:space:]') ]]; then
        printf "************output************\n$out \nis not the same as\n************expected************\n$expected\n"
        return 1
    fi
    return 0
}

function get {
    bin_path=$1;shift
    name_salt=$1;shift
    # expect atleast 1 application to be returned
    if [ $($bin_path get application | wc -l | xargs) -lt 2 ]; then
        return 1
    fi

    if ! assert_output "$bin_path get application echo-cli-app-$name_salt" "cat ./expected_data/echo-app.table | sed -e s/NAME_SALT/$name_salt/g"; then
      return 1
    fi

    if ! assert_output "$bin_path get application echo-cli-app-$name_salt -o yaml" "cat ./expected_data/echo-app.yml | sed -e s/NAME_SALT/$name_salt/g"; then
      return 1
    fi

    # expect atleast 1 project to be returned
    if [ $($bin_path get project | wc -l | xargs) -lt 2 ]; then
        return 1
    fi

    # get project(we can't really check the exact output as the edges might have changed)
    n_proj=`$bin_path get project Starter |  grep Starter | wc -l | xargs`
    if [ $n_proj -eq 0 ]; then
        echo "expected 1 or more projects, but got $n_proj"
        return 1
    fi

     # get project(we can't really check the exact output as the edges might have changed)
    n_proj=`$bin_path get project Starter -o yaml |  grep Starter | wc -l`
    if [ $n_proj -eq 0 ]; then
        echo "expected 1 or more projects, but got 0"
        return 1
    fi

    # expect atleast 1 category to be returned
    if [ $($bin_path get category | wc -l | xargs) -lt 2 ]; then
        return 1
    fi

    cat_count=`$bin_path get category mqtt-data2-$name_salt | grep mqtt-data2 | wc -l | xargs`
    if [ $cat_count -eq 0 ]; then
        echo "expected 1 or more categories, but got 0"
        return 1
    fi

    if ! assert_output "$bin_path get category mqtt-data2-$name_salt -o yaml" "cat ./expected_data/mqtt-data2-cat.yml | sed -e s/NAME_SALT/$name_salt/g"; then
      return 1
    fi

    # expect atleast 1 function to be returned
    if [ $($bin_path get \function | wc -l | xargs) -lt 2 ]; then
        return 1
    fi

    func_count=`$bin_path get \function echo2-$name_salt | wc -l | xargs`
    if [ $func_count -eq 0 ]; then
        echo "expected 1 or more functions, but got 0"
        return 1
    fi

    if ! assert_output "$bin_path get \function echo2-$name_salt -o yaml" "cat ./expected_data/echo2-function.yml | sed -e s/NAME_SALT/$name_salt/g"; then
      return 1
    fi

    if ! assert_output "$bin_path get datapipeline pipeline-0-$name_salt -o yaml" "cat ./expected_data/pipeline_0.yml | sed -e s/NAME_SALT/$name_salt/g"; then
      return 1
    fi

    if ! assert_output "$bin_path get datapipeline pipeline-0-$name_salt" "cat ./expected_data/pipeline_0.table | sed -e s/NAME_SALT/$name_salt/g"; then
      return 1
    fi

    if ! assert_output "$bin_path get datapipeline pipeline-1-$name_salt -o yaml" "cat ./expected_data/pipeline_1.yml | sed -e s/NAME_SALT/$name_salt/g"; then
      return 1
    fi

    if ! assert_output "$bin_path get datapipeline pipeline-1-$name_salt" "cat ./expected_data/pipeline_1.table | sed -e s/NAME_SALT/$name_salt/g"; then
      return 1
    fi

    # expect atleast 1 data pipeline to be returned
    # TODO: FIXME: commented out for now as some pipeline on test.ntnxsherlock.com are supporting Azure as cloud destination. 
    # which the CLI does not support so far. 
    # if [ $($bin_path get datapipeline | wc -l | xargs) -lt 2 ]; then
    #     return 1
    # fi
    
    # check all keys in table format are present
    if ! assert_output "$bin_path get edge | sed -n 1p |  tr '|' '\n' | xargs" "echo NAME IP ADDRESS CONNECTED SERIAL NUMBER"; then
        return 1
    fi

    # check output of a single edge name
    edge_name=$($bin_path get edge | awk '{print $1}' | grep edge | head -1 |  xargs)
    if ! assert_output "$bin_path get edge $edge_name | sed -n 1p |  tr '|' '\n' | xargs" "echo NAME IP ADDRESS CONNECTED SERIAL NUMBER"; then
        return 1
    fi

    # asserting phone datasource to assert data interface output as well
    if ! assert_output "$bin_path get datasource phone -o yaml | grep -v edge:" "cat ./expected_data/phone_ds.yml | grep -v edge:"; then
        return 1
    fi

    if ! assert_output "$bin_path get datasource mqtt-test-datasource1-$name_salt -o yaml | grep -v edge:" "cat ./expected_data/mqtt-test-datasource1.yml | sed -e s/NAME_SALT/$name_salt/g | grep -v edge:"; then
        return 1
    fi


    if ! assert_output "$bin_path get datasource -e $edge_name |  sed -n 1p |  tr '|' '\n' | xargs" "echo NAME EDGE PROTOCOL TOPICS"; then
        return 1
    fi
    
    if ! assert_output "$bin_path get datasource mqtt-test-datasource1-$name_salt |  sed -n 1p |  tr '|' '\n' | xargs" "echo NAME EDGE PROTOCOL TOPICS"; then
        return 1
    fi

    # expect atleast 1 data source to be returned
    if [ $($bin_path get datasource | wc -l | xargs) -lt 2 ]; then
        return 1
    fi

    # sanity check --show-artifacts as the output has edge name which changes.
    if ! $bin_path get datasource mqtt-test-datasource1-$name_salt --show-artifacts; then
        echo "get datasource mqtt-test-datasource1-$name_salt --show-artifacts"
        return 1
    fi

    echo "successfully tested get operations"
}

function debug {
    bin_path=$1;shift
    name_salt=$1;shift
    if ! create_and_use_context $bin_path $DEBUG_NS; then
        return 1
    fi
    mkdir -p $DEBUG_GENERATE_FOLDER
    test_echo2_yaml="${DEBUG_GENERATE_FOLDER}/test-echo2.yaml"
    test_detect_yaml="${DEBUG_GENERATE_FOLDER}/test-detect.yaml"
    test_error_yaml="${DEBUG_GENERATE_FOLDER}/test-error.yaml"
    sed -e "s/NAME_SALT/$name_salt/g" function-echo2-create.yaml > $test_echo2_yaml
    sed -e "s/NAME_SALT/$name_salt/g" function-detect-create.yaml > $test_detect_yaml
    sed -e "s/NAME_SALT/$name_salt/g" function-error-create.yaml > $test_error_yaml
    
    if ! $bin_path delete function test-echo2-$name_salt test-detect-$name_salt test-error-$name_salt test-error-runtime-$name_salt test-error-syntax-$name_salt; then
        return 1
    fi
    if ! $bin_path create -f $test_echo2_yaml; then
        return 1
    fi
    if ! $bin_path create -f $test_detect_yaml; then
        return 1
    fi
    if ! $bin_path create -f $test_error_yaml; then
        return 1
    fi
    # Test debugging feature
    if ! echo "testMsg" | $bin_path debug function -f test-echo2-${name_salt}=param1:dumy -i '-' -o '-'; then
        return 1
    fi

    if [ $("$bin_path" debug function -f test-echo2-${name_salt}=param1:dumy -f test-echo2-${name_salt}=param1:dumy -i './testinput/*.txt'| grep "test message" | wc -l |  xargs) -lt 4 ]; then
        return 1
    fi
    
    if ! $bin_path debug function -f test-detect-${name_salt} -i './testinput/test*' -o './testoutput/*.png' -t 6s -r 3 --repeat-interval 2s; then
        return 1
    fi

    if [ $("$bin_path" debug function -f test-error-${name_salt} -i './testinput/test*' | grep "Traceback (most recent call last):" | wc -l |  xargs) -lt 1 ]; then
        return 1
    fi

    if [ $("$bin_path" debug function -f test-error-runtime-${name_salt} -i './testinput/test*' | grep "TypeError: exceptions must derive from BaseException" | wc -l |  xargs) -lt 1 ]; then
        return 1
    fi

    if [ $("$bin_path" debug function -f test-error-syntax-${name_salt} -i './testinput/test*' | grep "Container exited with 1" | wc -l |  xargs) -lt 1 ]; then
        return 1
    fi
    
    if ! $bin_path delete function test-echo2-$name_salt test-detect-$name_salt test-error-$name_salt test-error-runtime-$name_salt test-error-syntax-$name_salt; then
        return 1
    fi

    # Switch back to test-user
    if ! $bin_path config use-context test-user; then
        return 1
    fi
}

function logging {
    bin_path=$1;shift
    name_salt=$1;shift
    if ! $bin_path config create-context --password 'Sherlock1!' -m sain@nutanix.com -u stage.ntnxsherlock.com logging-user; then
        echo "failed to set up context"
        return 1
    fi
    if ! $bin_path config use-context logging-user; then
        return 1
    fi
    mkdir -p $LOGGING_GENERATE_FOLDER
    log_app_create_yaml="${LOGGING_GENERATE_FOLDER}/application-logging-create.yaml"
    log_app_yaml="${LOGGING_GENERATE_FOLDER}/application-logging.yaml"

    sed -e "s/NAME_SALT/$name_salt/g" -e "s/LOGGING_GENERATE_FOLDER/$LOGGING_GENERATE_FOLDER/g" application-logging-create.yaml > $log_app_create_yaml
    sed -e "s/NAME_SALT/$name_salt/g" application-logging.yaml > $log_app_yaml
    
    if ! $bin_path delete application logging-application-$name_salt; then
        return 1
    fi
    
    if ! $bin_path create -f $log_app_create_yaml; then
        return 1
    fi

    # TODO: for now hard code a time for application to be deployed. 
    # ideally, logging api would let user wait if application is created but the container is initializing, because that only takes several seconds.
    # currently, logging api is throwing 'no container found' for application that is created in cloudmgmt but under initialization. 
    sleep 40
    ( 
        rm ${LOGGING_GENERATE_FOLDER}/log-app-out.txt
        $bin_path log app logging-application-$name_salt -e awslogstreamb289 -c loggingtest-$name_salt > "${LOGGING_GENERATE_FOLDER}/log-app-out.txt"
    )& log_pid=$!

    sleep 10; kill -9 $log_pid; wait $log_pid
    if [ $(cat ${LOGGING_GENERATE_FOLDER}/log-app-out.txt | grep 'SAIN: Iter:' | wc -l |  xargs) -lt 1 ]; then
        return 1
    fi
    
    if ! $bin_path delete application logging-application-$name_salt; then
        return 1
    fi

    # Switch back to test-user
    if ! $bin_path config use-context test-user; then
        return 1
    fi
}

# This is ONLY used to clean up on error. Assuming on debugging context 
function teardown_debug {
    bin_path=$1;shift
    name_salt=$1;shift
    if ! $bin_path delete function test-error-$name_salt test-detect-$name_salt test-echo2-$name_salt test-error-runtime-$name_salt test-error-syntax-$name_salt; then
        return 1
    fi
    
    if ! $bin_path debug purge; then
        return 1
    fi
}

# This is ONLY used to clean up on error. Assuming on logging context 
function teardown_logging {
    bin_path=$1;shift
    name_salt=$1;shift
    if ! $bin_path delete application logging-application-$name_salt; then
        return 1
    fi
}
#!/usr/bin/env python

# Copyright 2016 Attic Labs, Inc. All rights reserved.
# Licensed under the Apache License, version 2.0:
# http://www.apache.org/licenses/LICENSE-2.0

'''
This script builds the JS SDK documentation and puts the generated files
in $WORKSPACE/build.
'''

import copy
import json
import os
import subprocess

def call_with_env_and_cwd(cmd, env, cwd):
    print(cmd)
    proc = subprocess.Popen(cmd, env=env, cwd=cwd, shell=False)
    proc.wait()
    assert proc.returncode == 0

def main():
    # Workspace is the root of the builder, e.g. "/var/lib/jenkins/workspace/builder".
    workspace = os.getenv('WORKSPACE')
    assert workspace

    # Directory where node/npm binaries have been installed.
    node_bin = '/var/lib/jenkins/node-v5.11.1-linux-x64/bin'
    assert os.path.exists(node_bin)

    noms_dir = os.path.join(workspace, 'src/github.com/attic-labs/noms')
    noms_js_dir = os.path.join(noms_dir, 'js/noms')

    env = copy.copy(os.environ)
    env.update({
        'PATH': '%s:%s' % (os.getenv('PATH'), node_bin),
    })

    call_with_env_and_cwd(['npm', 'install'], env, noms_js_dir)
    call_with_env_and_cwd(['npm', 'run' 'build-docs'], env, noms_js_dir)

if __name__ == '__main__':
    main()

# dirToJSON

Small Go package to parse a directory into valid JSON.

## Install

    go get github.com/Dome91/jsonToDir

## Usage

    dirAsJSON, err := dirToJSON.Parse("root")

## Example output
    {
        "name":"test",
        "path":"test/",
        "is_dir":true,
        "children":[
        {
            "name":"another",
            "path":"test/another/",
            "is_dir":true,
            "children":[
                {
                    "name":"one_file.txt",
                    "path":"test/another/one_file.txt",
                    "is_dir":false,
                    "children":null
                },
                {
                    "name":"two_file.txt",
                    "path":"test/another/two_file.txt",
                    "is_dir":false,
                    "children":null
                }
            ]
        },
        {
            "name":"test",
            "path":"test/test/",
            "is_dir":true,
            "children":[
                {
                    "name":"test222",
                    "path":"test/test/test222/",
                    "is_dir":true,
                    "children":[
                    {
                        "name":"hello",
                        "path":"test/test/test222/hello/",
                        "is_dir":true,
                        "children":null
                    }
                    ]
                },
                {
                    "name":"third_file.txt",
                    "path":"test/test/third_file.txt",
                    "is_dir":false,
                    "children":null
                }
            ]
        }
        ]
    }
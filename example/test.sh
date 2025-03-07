curl --location --request POST 'localhost:8000/process' \
--header 'Content-Type: application/json' \
--data-raw '{
    "examples": [
        {
            "input": "Simple: Appropriate design with plain and easy code.",
            "output": "- Simple"
        }
    ],
    "rules": [
        {
            "content": "Extract the Subheading for Each Sentence"
        },
        {
            "content": "Output an unordered list in Markdown format"
        }
    ],
    "input": "Simple: Appropriate design with plain and easy code.\nGeneral: Cover the various utilities for business development.\nHighly efficient: Speeding up the efficiency of businesses upgrading.\nStable: The base libs validated in the production environment have the characteristics of high testability, high coverage as well as high security and reliability.\nRobust: Eliminating misusing through high quality of the base libs.\nHigh-performance: Optimal performance excluding the optimization of hacking in case of unsafe. \nExpandability: Properly designed interfaces where you can expand utilities such as base libs to meet your further requirements.\nFault-tolerance: Designed against failure, enhance the understanding and exercising of SRE within Kratos to achieve more robustness.\nToolchain: Includes an extensive toolchain, such as the code generation of cache, the lint tool, and so forth.\n"
}'

# Output
#{
#    "result": "- Simple\n- General\n- Highly efficient\n- Stable\n- Robust\n- High-performance\n- Expandability\n- Fault-tolerance\n- Toolchain"
#}
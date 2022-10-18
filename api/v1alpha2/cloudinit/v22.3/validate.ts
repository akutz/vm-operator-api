import Ajv from "ajv"
import * as schema from "./schema-cloud-config-v1.json"

import fs from "graceful-fs"
import YAML from "yaml"

const ajv = new Ajv()
const validate = ajv.compile(schema)

const file = fs.readFileSync("./cloudconfig.yaml", "utf8")
const data = YAML.parse(file)

if (validate(data)) {
    console.log(data)
} else {
    console.log(validate.errors)
}

package jsonschema

var schemaDataV1 = `{
  "$schema": "http://json-schema.org/draft-07/schema#",
  "$id": "config_schema_v1.json",
  "type": "object",
  "patternProperties": {
    "jobs": {
      "type": "array",
      "items": {
        "$ref": "#/definitions/job"
      },
      "default": []
    },
    "name": {
      "type": "string"
    }
  },
  "additionalProperties": false,
  "definitions": {
    "job": {
      "id": "#/definitions/job",
      "type": "object",
      "properties": {
        "name": {
          "type": "string"
        },
        "executor": {
          "$ref": "#/definitions/runner"
        },
        "env": {
          "type": "string"
        },
        "steps": {
          "type": "array",
          "items": {
            "oneOf": [
              {
                "type": "string"
              },
              {
                "type": "object"
              }
            ]
          }
        },
        "branches": {
          "type": "string"
        }
      }
    },
    "runner": {
      "id": "#/definitions/runner",
      "type": "object",
      "minProperties": 1,
      "maxProperties": 1,
      "patternProperties": {
        "^docker": {
          "$ref": "#/definitions/docker"
        },
        "^machine": {
          "$ref": "#/definitions/machine"
        }
      },
      "additionalProperties": false
    },
    "docker": {
      "id": "#/definitions/docker",
      "type": "object",
      "properties": {
        "image": {
          "type": "string"
        },
        "tags": {
          "type": "string"
        }
      },
      "additionalProperties": false
    },
    "machine": {
      "id": "#/definitions/machine",
      "type": "object",
      "properties": {
        "os": {
          "type": "string"
        },
        "cpu_cores": {
          "type": [
            "number",
            "string"
          ]
        },
        "mem": {
          "type": [
            "number",
            "string"
          ]
        }
      },
      "additionalProperties": false
    }
  }
}`

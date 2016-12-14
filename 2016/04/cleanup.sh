#!/bin/bash
sed 's:-::g' | sed 's:\[: :' | sed 's:\]::'

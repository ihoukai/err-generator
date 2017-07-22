package main

var csharpFrame = `
// Code generated by err-generator.
// source: @filename
// DO NOT EDIT!

namespace @packagename
{

    public enum @classname
    {
@key-vals
    }

    public static class @classnameHelper
    {
        public static string String(this @classname retCode)
        {
            switch (retCode)
            {
@val-strs
            }
            return "";
        }
    }
}
`

var csharpKeyVals = `	    %s = %d,
`
var csharpValStrs = `               case @classname.%s:
                    return "%s";
`

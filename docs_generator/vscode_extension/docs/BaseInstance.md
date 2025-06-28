# BaseInstance(3) Manual Page

## Name

BaseInstance - First instance being rendered



## [](#_description)Description

`BaseInstance`

Decorating a variable with the `BaseInstance` built-in will make that variable contain the integer value corresponding to the first instance that was passed to the command that invoked the current vertex shader invocation. `BaseInstance` is the `firstInstance` parameter to a *direct drawing command* or the `firstInstance` member of a structure consumed by an *indirect drawing command*.

Valid Usage

- [](#VUID-BaseInstance-BaseInstance-04181)VUID-BaseInstance-BaseInstance-04181  
  The `BaseInstance` decoration **must** be used only within the `Vertex` `Execution` `Model`
- [](#VUID-BaseInstance-BaseInstance-04182)VUID-BaseInstance-BaseInstance-04182  
  The variable decorated with `BaseInstance` **must** be declared using the `Input` `Storage` `Class`
- [](#VUID-BaseInstance-BaseInstance-04183)VUID-BaseInstance-BaseInstance-04183  
  The variable decorated with `BaseInstance` **must** be declared as a scalar 32-bit integer value

## [](#_see_also)See Also

No cross-references are available

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#BaseInstance)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0
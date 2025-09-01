# FragDepth(3) Manual Page

## Name

FragDepth - Application-specified depth for depth testing



## [](#_description)Description

`FragDepth`

To have a shader supply a fragment-depth value, the shader **must** declare the `DepthReplacing` execution mode. Such a shader’s fragment-depth value will come from the variable decorated with the `FragDepth` built-in decoration.

This value will be used for any subsequent depth testing performed by the implementation or writes to the depth attachment. See [fragment shader depth replacement](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#fragops-shader-depthreplacement) for details.

Valid Usage

- [](#VUID-FragDepth-FragDepth-04213)VUID-FragDepth-FragDepth-04213  
  The `FragDepth` decoration **must** be used only within the `Fragment` `Execution` `Model`
- [](#VUID-FragDepth-FragDepth-04214)VUID-FragDepth-FragDepth-04214  
  The variable decorated with `FragDepth` **must** be declared using the `Output` `Storage` `Class`
- [](#VUID-FragDepth-FragDepth-04215)VUID-FragDepth-FragDepth-04215  
  The variable decorated with `FragDepth` **must** be declared as a scalar 32-bit floating-point value
- [](#VUID-FragDepth-FragDepth-04216)VUID-FragDepth-FragDepth-04216  
  If the shader dynamically writes to the variable decorated with `FragDepth`, the `DepthReplacing` `Execution` `Mode` **must** be declared

## [](#_see_also)See Also

No cross-references are available

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#FragDepth).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0
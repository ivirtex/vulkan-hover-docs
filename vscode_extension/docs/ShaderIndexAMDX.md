# ShaderIndexAMDX(3) Manual Page

## Name

ShaderIndexAMDX - Index assigned to the shader within the workgraph



## [](#_description)Description

`ShaderIndexAMDX`

Decorating a variable with the `ShaderIndexAMDX` built-in decoration will make that variable contain the index of the shader specified when it was compiled, either via [VkPipelineShaderStageNodeCreateInfoAMDX](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineShaderStageNodeCreateInfoAMDX.html)::`index` or by the `ShaderIndexAMDX` execution mode.

Valid Usage

- [](#VUID-ShaderIndexAMDX-ShaderIndexAMDX-09175)VUID-ShaderIndexAMDX-ShaderIndexAMDX-09175  
  The variable decorated with `ShaderIndexAMDX` **must** be declared using the `Input` `Storage` `Class`
- [](#VUID-ShaderIndexAMDX-ShaderIndexAMDX-09176)VUID-ShaderIndexAMDX-ShaderIndexAMDX-09176  
  The variable decorated with `ShaderIndexAMDX` **must** be declared as a scalar 32-bit integer value

## [](#_see_also)See Also

No cross-references are available

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#ShaderIndexAMDX)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0
# ShaderIndexAMDX(3) Manual Page

## Name

ShaderIndexAMDX - Index assigned to the shader within the workgraph



## <a href="#_description" class="anchor"></a>Description

`ShaderIndexAMDX`  
Decorating a variable with the `ShaderIndexAMDX` built-in decoration
will make that variable contain the index of the shader specified when
it was compiled, either via
[VkPipelineShaderStageNodeCreateInfoAMDX](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineShaderStageNodeCreateInfoAMDX.html)::`index`
or by the `ShaderIndexAMDX` execution mode.

Valid Usage

- <a href="#VUID-ShaderIndexAMDX-ShaderIndexAMDX-09175"
  id="VUID-ShaderIndexAMDX-ShaderIndexAMDX-09175"></a>
  VUID-ShaderIndexAMDX-ShaderIndexAMDX-09175  
  The variable decorated with `ShaderIndexAMDX` **must** be declared
  using the `Input` `Storage` `Class`

- <a href="#VUID-ShaderIndexAMDX-ShaderIndexAMDX-09176"
  id="VUID-ShaderIndexAMDX-ShaderIndexAMDX-09176"></a>
  VUID-ShaderIndexAMDX-ShaderIndexAMDX-09176  
  The variable decorated with `ShaderIndexAMDX` **must** be declared as
  a scalar 32-bit integer value

## <a href="#_see_also" class="anchor"></a>See Also

No cross-references are available

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#ShaderIndexAMDX"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700

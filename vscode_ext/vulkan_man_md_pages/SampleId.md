# SampleId(3) Manual Page

## Name

SampleId - Sample ID within a fragment



## <a href="#_description" class="anchor"></a>Description

`SampleId`  
Decorating a variable with the `SampleId` built-in decoration will make
that variable contain the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#primsrast-multisampling-coverage-mask"
target="_blank" rel="noopener">coverage index</a> for the current
fragment shader invocation. `SampleId` ranges from zero to the number of
samples in the framebuffer minus one. If a fragment shader entry point’s
interface includes an input variable decorated with `SampleId`, <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#primsrast-sampleshading"
target="_blank" rel="noopener">Sample Shading</a> is considered enabled
with a `minSampleShading` value of 1.0.

Valid Usage

- <a href="#VUID-SampleId-SampleId-04354"
  id="VUID-SampleId-SampleId-04354"></a> VUID-SampleId-SampleId-04354  
  The `SampleId` decoration **must** be used only within the `Fragment`
  `Execution` `Model`

- <a href="#VUID-SampleId-SampleId-04355"
  id="VUID-SampleId-SampleId-04355"></a> VUID-SampleId-SampleId-04355  
  The variable decorated with `SampleId` **must** be declared using the
  `Input` `Storage` `Class`

- <a href="#VUID-SampleId-SampleId-04356"
  id="VUID-SampleId-SampleId-04356"></a> VUID-SampleId-SampleId-04356  
  The variable decorated with `SampleId` **must** be declared as a
  scalar 32-bit integer value

## <a href="#_see_also" class="anchor"></a>See Also

No cross-references are available

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#SampleId"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700

# GlobalInvocationId(3) Manual Page

## Name

GlobalInvocationId - Global invocation ID



## <a href="#_description" class="anchor"></a>Description

`GlobalInvocationId`  
Decorating a variable with the `GlobalInvocationId` built-in decoration
will make that variable contain the location of the current invocation
within the global workgroup. Each component is equal to the index of the
local workgroup multiplied by the size of the local workgroup plus
`LocalInvocationId`.

Valid Usage

- <a href="#VUID-GlobalInvocationId-GlobalInvocationId-04236"
  id="VUID-GlobalInvocationId-GlobalInvocationId-04236"></a>
  VUID-GlobalInvocationId-GlobalInvocationId-04236  
  The `GlobalInvocationId` decoration **must** be used only within the
  `GLCompute`, `MeshEXT`, `TaskEXT`, `MeshNV`, or `TaskNV` `Execution`
  `Model`

- <a href="#VUID-GlobalInvocationId-GlobalInvocationId-04237"
  id="VUID-GlobalInvocationId-GlobalInvocationId-04237"></a>
  VUID-GlobalInvocationId-GlobalInvocationId-04237  
  The variable decorated with `GlobalInvocationId` **must** be declared
  using the `Input` `Storage` `Class`

- <a href="#VUID-GlobalInvocationId-GlobalInvocationId-04238"
  id="VUID-GlobalInvocationId-GlobalInvocationId-04238"></a>
  VUID-GlobalInvocationId-GlobalInvocationId-04238  
  The variable decorated with `GlobalInvocationId` **must** be declared
  as a three-component vector of 32-bit integer values

## <a href="#_see_also" class="anchor"></a>See Also

No cross-references are available

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#GlobalInvocationId"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700

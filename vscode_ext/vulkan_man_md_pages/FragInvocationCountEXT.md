# FragInvocationCountEXT(3) Manual Page

## Name

FragInvocationCountEXT - Number of fragment shader invocations for a
fragment



## <a href="#_description" class="anchor"></a>Description

`FragInvocationCountEXT`  
Decorating a variable with the `FragInvocationCountEXT` built-in
decoration will make that variable contain the maximum number of
fragment shader invocations for the fragment, as determined by
`minSampleShading`.

If <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#primsrast-sampleshading"
target="_blank" rel="noopener">Sample Shading</a> is not enabled,
`FragInvocationCountEXT` will be filled with a value of 1.

Valid Usage

- <a href="#VUID-FragInvocationCountEXT-FragInvocationCountEXT-04217"
  id="VUID-FragInvocationCountEXT-FragInvocationCountEXT-04217"></a>
  VUID-FragInvocationCountEXT-FragInvocationCountEXT-04217  
  The `FragInvocationCountEXT` decoration **must** be used only within
  the `Fragment` `Execution` `Model`

- <a href="#VUID-FragInvocationCountEXT-FragInvocationCountEXT-04218"
  id="VUID-FragInvocationCountEXT-FragInvocationCountEXT-04218"></a>
  VUID-FragInvocationCountEXT-FragInvocationCountEXT-04218  
  The variable decorated with `FragInvocationCountEXT` **must** be
  declared using the `Input` `Storage` `Class`

- <a href="#VUID-FragInvocationCountEXT-FragInvocationCountEXT-04219"
  id="VUID-FragInvocationCountEXT-FragInvocationCountEXT-04219"></a>
  VUID-FragInvocationCountEXT-FragInvocationCountEXT-04219  
  The variable decorated with `FragInvocationCountEXT` **must** be
  declared as a scalar 32-bit integer value

## <a href="#_see_also" class="anchor"></a>See Also

No cross-references are available

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#FragInvocationCountEXT"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700

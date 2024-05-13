# FragSizeEXT(3) Manual Page

## Name

FragSizeEXT - Size of the screen-space area covered by the fragment



## <a href="#_description" class="anchor"></a>Description

`FragSizeEXT`  
Decorating a variable with the `FragSizeEXT` built-in decoration will
make that variable contain the dimensions in pixels of the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#glossary-fragment-area"
target="_blank" rel="noopener">area</a> that the fragment covers for
that invocation.

If fragment density map is not enabled, `FragSizeEXT` will be filled
with a value of (1,1).

Valid Usage

- <a href="#VUID-FragSizeEXT-FragSizeEXT-04220"
  id="VUID-FragSizeEXT-FragSizeEXT-04220"></a>
  VUID-FragSizeEXT-FragSizeEXT-04220  
  The `FragSizeEXT` decoration **must** be used only within the
  `Fragment` `Execution` `Model`

- <a href="#VUID-FragSizeEXT-FragSizeEXT-04221"
  id="VUID-FragSizeEXT-FragSizeEXT-04221"></a>
  VUID-FragSizeEXT-FragSizeEXT-04221  
  The variable decorated with `FragSizeEXT` **must** be declared using
  the `Input` `Storage` `Class`

- <a href="#VUID-FragSizeEXT-FragSizeEXT-04222"
  id="VUID-FragSizeEXT-FragSizeEXT-04222"></a>
  VUID-FragSizeEXT-FragSizeEXT-04222  
  The variable decorated with `FragSizeEXT` **must** be declared as a
  two-component vector of 32-bit integer values

## <a href="#_see_also" class="anchor"></a>See Also

No cross-references are available

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#FragSizeEXT"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700

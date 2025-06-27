# CoreIDARM(3) Manual Page

## Name

CoreIDARM - Core ID on which a shader invocation is running



## <a href="#_description" class="anchor"></a>Description

`CoreIDARM`  
Decorating a variable with the `CoreIDARM` built-in decoration will make
that variable contain the ID of the core on which the current shader
invocation is running. This variable is in the range \[0,
`CoreMaxIDARM`\].

Valid Usage

- <a href="#VUID-CoreIDARM-CoreIDARM-07599"
  id="VUID-CoreIDARM-CoreIDARM-07599"></a>
  VUID-CoreIDARM-CoreIDARM-07599  
  The variable decorated with `CoreIDARM` **must** be declared using the
  `Input` `Storage` `Class`

- <a href="#VUID-CoreIDARM-CoreIDARM-07600"
  id="VUID-CoreIDARM-CoreIDARM-07600"></a>
  VUID-CoreIDARM-CoreIDARM-07600  
  The variable decorated with `CoreIDARM` **must** be declared as a
  scalar 32-bit integer value

## <a href="#_see_also" class="anchor"></a>See Also

No cross-references are available

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#CoreIDARM"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700

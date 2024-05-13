# WarpIDARM(3) Manual Page

## Name

WarpIDARM - Warp ID within a core of a shader invocation



## <a href="#_description" class="anchor"></a>Description

`WarpIDARM`  
Decorating a variable with the `WarpIDARM` built-in decoration will make
that variable contain the ID of the warp on a core on which the current
shader invocation is running. This variable is in the range \[0,
`WarpMaxIDARM`\].

Valid Usage

- <a href="#VUID-WarpIDARM-WarpIDARM-07603"
  id="VUID-WarpIDARM-WarpIDARM-07603"></a>
  VUID-WarpIDARM-WarpIDARM-07603  
  The variable decorated with `WarpIDARM` **must** be declared using the
  `Input` `Storage` `Class`

- <a href="#VUID-WarpIDARM-WarpIDARM-07604"
  id="VUID-WarpIDARM-WarpIDARM-07604"></a>
  VUID-WarpIDARM-WarpIDARM-07604  
  The variable decorated with `WarpIDARM` **must** be declared as a
  scalar 32-bit integer value

## <a href="#_see_also" class="anchor"></a>See Also

No cross-references are available

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#WarpIDARM"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700

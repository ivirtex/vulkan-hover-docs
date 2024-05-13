# CoreCountARM(3) Manual Page

## Name

CoreCountARM - Number of cores on the device



## <a href="#_description" class="anchor"></a>Description

`CoreCountARM`  
Decorating a variable with the `CoreCountARM` built-in decoration will
make that variable contain the number of cores on the device.

Valid Usage

- <a href="#VUID-CoreCountARM-CoreCountARM-07595"
  id="VUID-CoreCountARM-CoreCountARM-07595"></a>
  VUID-CoreCountARM-CoreCountARM-07595  
  The variable decorated with `CoreCountARM` **must** be declared using
  the `Input` `Storage` `Class`

- <a href="#VUID-CoreCountARM-CoreCountARM-07596"
  id="VUID-CoreCountARM-CoreCountARM-07596"></a>
  VUID-CoreCountARM-CoreCountARM-07596  
  The variable decorated with `CoreCountARM` **must** be declared as a
  scalar 32-bit integer value

## <a href="#_see_also" class="anchor"></a>See Also

No cross-references are available

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#CoreCountARM"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700

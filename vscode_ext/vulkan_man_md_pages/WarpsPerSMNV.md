# WarpsPerSMNV(3) Manual Page

## Name

WarpsPerSMNV - Number of warps per SM



## <a href="#_description" class="anchor"></a>Description

`WarpsPerSMNV`  
Decorating a variable with the `WarpsPerSMNV` built-in decoration will
make that variable contain the maximum number of warps executing on a
SM.

Valid Usage

- <a href="#VUID-WarpsPerSMNV-WarpsPerSMNV-04418"
  id="VUID-WarpsPerSMNV-WarpsPerSMNV-04418"></a>
  VUID-WarpsPerSMNV-WarpsPerSMNV-04418  
  The variable decorated with `WarpsPerSMNV` **must** be declared using
  the `Input` `Storage` `Class`

- <a href="#VUID-WarpsPerSMNV-WarpsPerSMNV-04419"
  id="VUID-WarpsPerSMNV-WarpsPerSMNV-04419"></a>
  VUID-WarpsPerSMNV-WarpsPerSMNV-04419  
  The variable decorated with `WarpsPerSMNV` **must** be declared as a
  scalar 32-bit integer value

## <a href="#_see_also" class="anchor"></a>See Also

No cross-references are available

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#WarpsPerSMNV"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700

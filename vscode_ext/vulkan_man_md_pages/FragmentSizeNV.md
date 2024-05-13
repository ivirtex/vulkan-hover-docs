# FragmentSizeNV(3) Manual Page

## Name

FragmentSizeNV - Size of the screen-space area covered by the fragment



## <a href="#_description" class="anchor"></a>Description

`FragmentSizeNV`  
Decorating a variable with the `FragmentSizeNV` built-in decoration will
make that variable contain the width and height of the fragment.

Valid Usage

- <a href="#VUID-FragmentSizeNV-FragmentSizeNV-04226"
  id="VUID-FragmentSizeNV-FragmentSizeNV-04226"></a>
  VUID-FragmentSizeNV-FragmentSizeNV-04226  
  The `FragmentSizeNV` decoration **must** be used only within the
  `Fragment` `Execution` `Model`

- <a href="#VUID-FragmentSizeNV-FragmentSizeNV-04227"
  id="VUID-FragmentSizeNV-FragmentSizeNV-04227"></a>
  VUID-FragmentSizeNV-FragmentSizeNV-04227  
  The variable decorated with `FragmentSizeNV` **must** be declared
  using the `Input` `Storage` `Class`

- <a href="#VUID-FragmentSizeNV-FragmentSizeNV-04228"
  id="VUID-FragmentSizeNV-FragmentSizeNV-04228"></a>
  VUID-FragmentSizeNV-FragmentSizeNV-04228  
  The variable decorated with `FragmentSizeNV` **must** be declared as a
  two-component vector of 32-bit integer values

## <a href="#_see_also" class="anchor"></a>See Also

No cross-references are available

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#FragmentSizeNV"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700

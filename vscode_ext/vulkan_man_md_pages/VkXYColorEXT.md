# VkXYColorEXT(3) Manual Page

## Name

VkXYColorEXT - Specify X,Y chromaticity coordinates



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkXYColorEXT` structure is defined as:

``` c
// Provided by VK_EXT_hdr_metadata
typedef struct VkXYColorEXT {
    float    x;
    float    y;
} VkXYColorEXT;
```

## <a href="#_members" class="anchor"></a>Members

- `x` is the x chromaticity coordinate.

- `y` is the y chromaticity coordinate.

## <a href="#_description" class="anchor"></a>Description

Chromaticity coordinates are as specified in CIE 15:2004 “Calculation of
chromaticity coordinates” (Section 7.3) and are limited to between 0 and
1 for real colors.

## <a href="#_see_also" class="anchor"></a>See Also

[VK_EXT_hdr_metadata](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_hdr_metadata.html),
[VkHdrMetadataEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkHdrMetadataEXT.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkXYColorEXT"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700

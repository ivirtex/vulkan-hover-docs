# VkXYColorEXT(3) Manual Page

## Name

VkXYColorEXT - Specify X,Y chromaticity coordinates



## [](#_c_specification)C Specification

The `VkXYColorEXT` structure is defined as:

```c++
// Provided by VK_EXT_hdr_metadata
typedef struct VkXYColorEXT {
    float    x;
    float    y;
} VkXYColorEXT;
```

## [](#_members)Members

- `x` is the x chromaticity coordinate.
- `y` is the y chromaticity coordinate.

## [](#_description)Description

Chromaticity coordinates are as specified in CIE 15:2004 “Calculation of chromaticity coordinates” (Section 7.3) and are limited to between 0 and 1 for real colors.

## [](#_see_also)See Also

[VK\_EXT\_hdr\_metadata](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_hdr_metadata.html), [VkHdrMetadataEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkHdrMetadataEXT.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkXYColorEXT)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0
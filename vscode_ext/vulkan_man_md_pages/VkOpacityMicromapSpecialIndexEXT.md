# VkOpacityMicromapSpecialIndexEXT(3) Manual Page

## Name

VkOpacityMicromapSpecialIndexEXT - Enum for special indices in the
opacity micromap



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkOpacityMicromapSpecialIndexEXT` enumeration is defined as:

``` c
// Provided by VK_EXT_opacity_micromap
typedef enum VkOpacityMicromapSpecialIndexEXT {
    VK_OPACITY_MICROMAP_SPECIAL_INDEX_FULLY_TRANSPARENT_EXT = -1,
    VK_OPACITY_MICROMAP_SPECIAL_INDEX_FULLY_OPAQUE_EXT = -2,
    VK_OPACITY_MICROMAP_SPECIAL_INDEX_FULLY_UNKNOWN_TRANSPARENT_EXT = -3,
    VK_OPACITY_MICROMAP_SPECIAL_INDEX_FULLY_UNKNOWN_OPAQUE_EXT = -4,
} VkOpacityMicromapSpecialIndexEXT;
```

## <a href="#_description" class="anchor"></a>Description

- `VK_OPACITY_MICROMAP_SPECIAL_INDEX_FULLY_TRANSPARENT_EXT` specifies
  that the entire triangle is fully transparent.

- `VK_OPACITY_MICROMAP_SPECIAL_INDEX_FULLY_OPAQUE_EXT` specifies that
  the entire triangle is fully opaque.

- `VK_OPACITY_MICROMAP_SPECIAL_INDEX_FULLY_UNKNOWN_TRANSPARENT_EXT`
  specifies that the entire triangle is unknown-transparent.

- `VK_OPACITY_MICROMAP_SPECIAL_INDEX_FULLY_UNKNOWN_OPAQUE_EXT` specifies
  that the entire triangle is unknown-opaque.

## <a href="#_see_also" class="anchor"></a>See Also

[VK_EXT_opacity_micromap](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_opacity_micromap.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkOpacityMicromapSpecialIndexEXT"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700

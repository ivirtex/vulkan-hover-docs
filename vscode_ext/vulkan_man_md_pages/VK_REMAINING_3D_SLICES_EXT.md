# VK_REMAINING_3D_SLICES_EXT(3) Manual Page

## Name

VK_REMAINING_3D_SLICES_EXT - Sentinel for all remaining 3D slices



## <a href="#_c_specification" class="anchor"></a>C Specification

`VK_REMAINING_3D_SLICES_EXT` is a special constant value used for
[VkImageViewSlicedCreateInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageViewSlicedCreateInfoEXT.html)::`sliceCount`
to indicate that all remaining 3D slices in an image after the first
slice offset specified should be included in the view.

``` c
#define VK_REMAINING_3D_SLICES_EXT        (~0U)
```

## <a href="#_see_also" class="anchor"></a>See Also

[VK_EXT_image_sliced_view_of_3d](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_image_sliced_view_of_3d.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VK_REMAINING_3D_SLICES_EXT"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700

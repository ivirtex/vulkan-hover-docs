# VK\_REMAINING\_3D\_SLICES\_EXT(3) Manual Page

## Name

VK\_REMAINING\_3D\_SLICES\_EXT - Sentinel for all remaining 3D slices



## [](#_c_specification)C Specification

`VK_REMAINING_3D_SLICES_EXT` is a special constant value used for [VkImageViewSlicedCreateInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageViewSlicedCreateInfoEXT.html)::`sliceCount` to indicate that all remaining 3D slices in an image after the first slice offset specified should be included in the view.

```c++
#define VK_REMAINING_3D_SLICES_EXT        (~0U)
```

## [](#_see_also)See Also

[VK\_EXT\_image\_sliced\_view\_of\_3d](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_image_sliced_view_of_3d.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VK_REMAINING_3D_SLICES_EXT).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0
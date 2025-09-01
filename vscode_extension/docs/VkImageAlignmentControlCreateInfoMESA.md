# VkImageAlignmentControlCreateInfoMESA(3) Manual Page

## Name

VkImageAlignmentControlCreateInfoMESA - Specify image alignment



## [](#_c_specification)C Specification

If the `pNext` list of [VkImageCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageCreateInfo.html) includes a `VkImageAlignmentControlCreateInfoMESA` structure, then that structure describes desired alignment for this image.

The `VkImageAlignmentControlCreateInfoMESA` structure is defined as:

```c++
// Provided by VK_MESA_image_alignment_control
typedef struct VkImageAlignmentControlCreateInfoMESA {
    VkStructureType    sType;
    const void*        pNext;
    uint32_t           maximumRequestedAlignment;
} VkImageAlignmentControlCreateInfoMESA;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `maximumRequestedAlignment` specifies the maximum alignment for the image.

## [](#_description)Description

If `maximumRequestedAlignment` is not 0, the implementation **should** choose an image memory layout that requires an alignment no larger than `maximumRequestedAlignment` as reported in [VkMemoryRequirements](https://registry.khronos.org/vulkan/specs/latest/man/html/VkMemoryRequirements.html)::`alignment`. If such a layout does not exist for the given image creation parameters, the implementation **should** return the smallest alignment which is supported in [VkMemoryRequirements](https://registry.khronos.org/vulkan/specs/latest/man/html/VkMemoryRequirements.html).

If an implementation needs to disable image compression for `maximumRequestedAlignment` to be honored - where a larger alignment would enable image compression - the implementation **should** not use `maximumRequestedAlignment`, and **should** return the smallest alignment which does not compromise compression. If the [`imageCompressionControl`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-imageCompressionControl) feature is enabled, the application **can** chain a [VkImageCompressionControlEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageCompressionControlEXT.html) with `VK_IMAGE_COMPRESSION_DISABLED_EXT`. In this case, image compression considerations **should** not apply when implementation decides alignment.

Valid Usage

- [](#VUID-VkImageAlignmentControlCreateInfoMESA-maximumRequestedAlignment-09655)VUID-VkImageAlignmentControlCreateInfoMESA-maximumRequestedAlignment-09655  
  If `maximumRequestedAlignment` is not 0, `maximumRequestedAlignment` **must** be a power of two
- [](#VUID-VkImageAlignmentControlCreateInfoMESA-maximumRequestedAlignment-09656)VUID-VkImageAlignmentControlCreateInfoMESA-maximumRequestedAlignment-09656  
  If `maximumRequestedAlignment` is not 0, the bitwise-and of `maximumRequestedAlignment` and [`supportedImageAlignmentMask`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#limits-supportedImageAlignmentMask) **must** be non-zero
- [](#VUID-VkImageAlignmentControlCreateInfoMESA-imageAlignmentControl-09657)VUID-VkImageAlignmentControlCreateInfoMESA-imageAlignmentControl-09657  
  [`imageAlignmentControl`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-imageAlignmentControl) **must** be enabled

Valid Usage (Implicit)

- [](#VUID-VkImageAlignmentControlCreateInfoMESA-sType-sType)VUID-VkImageAlignmentControlCreateInfoMESA-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_IMAGE_ALIGNMENT_CONTROL_CREATE_INFO_MESA`

## [](#_see_also)See Also

[VK\_MESA\_image\_alignment\_control](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_MESA_image_alignment_control.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkImageAlignmentControlCreateInfoMESA).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0
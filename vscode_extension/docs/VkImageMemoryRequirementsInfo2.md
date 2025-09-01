# VkImageMemoryRequirementsInfo2(3) Manual Page

## Name

VkImageMemoryRequirementsInfo2 - (None)



## [](#_c_specification)C Specification

The `VkImageMemoryRequirementsInfo2` structure is defined as:

```c++
// Provided by VK_VERSION_1_1
typedef struct VkImageMemoryRequirementsInfo2 {
    VkStructureType    sType;
    const void*        pNext;
    VkImage            image;
} VkImageMemoryRequirementsInfo2;
```

or the equivalent

```c++
// Provided by VK_KHR_get_memory_requirements2
typedef VkImageMemoryRequirementsInfo2 VkImageMemoryRequirementsInfo2KHR;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `image` is the image to query.

## [](#_description)Description

Valid Usage

- [](#VUID-VkImageMemoryRequirementsInfo2-image-01589)VUID-VkImageMemoryRequirementsInfo2-image-01589  
  If `image` was created with a *multi-planar* format and the `VK_IMAGE_CREATE_DISJOINT_BIT` flag, there **must** be a [VkImagePlaneMemoryRequirementsInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImagePlaneMemoryRequirementsInfo.html) included in the `pNext` chain of the [VkImageMemoryRequirementsInfo2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageMemoryRequirementsInfo2.html) structure
- [](#VUID-VkImageMemoryRequirementsInfo2-image-02279)VUID-VkImageMemoryRequirementsInfo2-image-02279  
  If `image` was created with `VK_IMAGE_CREATE_DISJOINT_BIT` and with `VK_IMAGE_TILING_DRM_FORMAT_MODIFIER_EXT`, then there **must** be a [VkImagePlaneMemoryRequirementsInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImagePlaneMemoryRequirementsInfo.html) included in the `pNext` chain of the [VkImageMemoryRequirementsInfo2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageMemoryRequirementsInfo2.html) structure
- [](#VUID-VkImageMemoryRequirementsInfo2-image-01590)VUID-VkImageMemoryRequirementsInfo2-image-01590  
  If `image` was not created with the `VK_IMAGE_CREATE_DISJOINT_BIT` flag, there **must** not be a [VkImagePlaneMemoryRequirementsInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImagePlaneMemoryRequirementsInfo.html) included in the `pNext` chain of the [VkImageMemoryRequirementsInfo2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageMemoryRequirementsInfo2.html) structure
- [](#VUID-VkImageMemoryRequirementsInfo2-image-02280)VUID-VkImageMemoryRequirementsInfo2-image-02280  
  If `image` was created with a single-plane format and with any `tiling` other than `VK_IMAGE_TILING_DRM_FORMAT_MODIFIER_EXT`, then there **must** not be a [VkImagePlaneMemoryRequirementsInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImagePlaneMemoryRequirementsInfo.html) included in the `pNext` chain of the [VkImageMemoryRequirementsInfo2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageMemoryRequirementsInfo2.html) structure
- [](#VUID-VkImageMemoryRequirementsInfo2-image-01897)VUID-VkImageMemoryRequirementsInfo2-image-01897  
  If `image` was created with the `VK_EXTERNAL_MEMORY_HANDLE_TYPE_ANDROID_HARDWARE_BUFFER_BIT_ANDROID` external memory handle type, then `image` **must** be bound to memory
- [](#VUID-VkImageMemoryRequirementsInfo2-image-08961)VUID-VkImageMemoryRequirementsInfo2-image-08961  
  If `image` was created with the `VK_EXTERNAL_MEMORY_HANDLE_TYPE_SCREEN_BUFFER_BIT_QNX` external memory handle type, then `image` **must** be bound to memory

Valid Usage (Implicit)

- [](#VUID-VkImageMemoryRequirementsInfo2-sType-sType)VUID-VkImageMemoryRequirementsInfo2-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_IMAGE_MEMORY_REQUIREMENTS_INFO_2`
- [](#VUID-VkImageMemoryRequirementsInfo2-pNext-pNext)VUID-VkImageMemoryRequirementsInfo2-pNext-pNext  
  `pNext` **must** be `NULL` or a pointer to a valid instance of [VkImagePlaneMemoryRequirementsInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImagePlaneMemoryRequirementsInfo.html)
- [](#VUID-VkImageMemoryRequirementsInfo2-sType-unique)VUID-VkImageMemoryRequirementsInfo2-sType-unique  
  The `sType` value of each structure in the `pNext` chain **must** be unique
- [](#VUID-VkImageMemoryRequirementsInfo2-image-parameter)VUID-VkImageMemoryRequirementsInfo2-image-parameter  
  `image` **must** be a valid [VkImage](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImage.html) handle

## [](#_see_also)See Also

[VK\_VERSION\_1\_1](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_1.html), [VkImage](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImage.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html), [vkGetImageMemoryRequirements2](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetImageMemoryRequirements2.html), [vkGetImageMemoryRequirements2](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetImageMemoryRequirements2.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkImageMemoryRequirementsInfo2).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0
# VkPhysicalDeviceImageAlignmentControlPropertiesMESA(3) Manual Page

## Name

VkPhysicalDeviceImageAlignmentControlPropertiesMESA - Structure describing supported image alignments for a physical device



## [](#_c_specification)C Specification

The `VkPhysicalDeviceImageAlignmentControlPropertiesMESA` structure is defined as:

```c++
// Provided by VK_MESA_image_alignment_control
typedef struct VkPhysicalDeviceImageAlignmentControlPropertiesMESA {
    VkStructureType    sType;
    void*              pNext;
    uint32_t           supportedImageAlignmentMask;
} VkPhysicalDeviceImageAlignmentControlPropertiesMESA;
```

## [](#_members)Members

The members of the `VkPhysicalDeviceImageAlignmentControlPropertiesMESA` structure describe the following:

## [](#_description)Description

- []()`supportedImageAlignmentMask` is a bitwise-or of all potentially supported image alignments for a given physical device when using `VK_IMAGE_TILING_OPTIMAL`. If a given alignment is supported, the application **can** request an image to have that alignment. A given set of image creation parameters **may** support a subset of these alignments. To determine if a particular alignment is supported for a given set of image creation parameters, check [VkMemoryRequirements](https://registry.khronos.org/vulkan/specs/latest/man/html/VkMemoryRequirements.html)::`alignment` after chaining in [VkImageAlignmentControlCreateInfoMESA](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageAlignmentControlCreateInfoMESA.html).

If the `VkPhysicalDeviceImageAlignmentControlPropertiesMESA` structure is included in the `pNext` chain of the [VkPhysicalDeviceProperties2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceProperties2.html) structure passed to [vkGetPhysicalDeviceProperties2](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPhysicalDeviceProperties2.html), it is filled in with each corresponding implementation-dependent property.

Valid Usage (Implicit)

- [](#VUID-VkPhysicalDeviceImageAlignmentControlPropertiesMESA-sType-sType)VUID-VkPhysicalDeviceImageAlignmentControlPropertiesMESA-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_IMAGE_ALIGNMENT_CONTROL_PROPERTIES_MESA`

## [](#_see_also)See Also

[VK\_MESA\_image\_alignment\_control](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_MESA_image_alignment_control.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkPhysicalDeviceImageAlignmentControlPropertiesMESA).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0
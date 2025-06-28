# VkPhysicalDeviceImageAlignmentControlFeaturesMESA(3) Manual Page

## Name

VkPhysicalDeviceImageAlignmentControlFeaturesMESA - Structure describing features supported by VK\_MESA\_image\_alignment\_control



## [](#_c_specification)C Specification

The `VkPhysicalDeviceImageAlignmentControlFeaturesMESA` structure is defined as:

```c++
// Provided by VK_MESA_image_alignment_control
typedef struct VkPhysicalDeviceImageAlignmentControlFeaturesMESA {
    VkStructureType    sType;
    void*              pNext;
    VkBool32           imageAlignmentControl;
} VkPhysicalDeviceImageAlignmentControlFeaturesMESA;
```

## [](#_members)Members

This structure describes the following feature:

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- []()`imageAlignmentControl` specifies that [VkImageAlignmentControlCreateInfoMESA](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageAlignmentControlCreateInfoMESA.html) **can** be chained in [VkImageCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageCreateInfo.html)

## [](#_description)Description

If the `VkPhysicalDeviceImageAlignmentControlFeaturesMESA` structure is included in the `pNext` chain of the [VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceFeatures2.html) structure passed to [vkGetPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPhysicalDeviceFeatures2.html), it is filled in to indicate whether each corresponding feature is supported. If the application wishes to use a [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html) with any features described by `VkPhysicalDeviceImageAlignmentControlFeaturesMESA`, it **must** add an instance of the structure, with the desired feature members set to `VK_TRUE`, to the `pNext` chain of [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceCreateInfo.html) when creating the [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html).

Valid Usage (Implicit)

- [](#VUID-VkPhysicalDeviceImageAlignmentControlFeaturesMESA-sType-sType)VUID-VkPhysicalDeviceImageAlignmentControlFeaturesMESA-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_IMAGE_ALIGNMENT_CONTROL_FEATURES_MESA`

## [](#_see_also)See Also

[VK\_MESA\_image\_alignment\_control](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_MESA_image_alignment_control.html), [VkBool32](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBool32.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkPhysicalDeviceImageAlignmentControlFeaturesMESA)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0
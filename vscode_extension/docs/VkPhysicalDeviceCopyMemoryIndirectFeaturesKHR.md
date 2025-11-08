# VkPhysicalDeviceCopyMemoryIndirectFeaturesKHR(3) Manual Page

## Name

VkPhysicalDeviceCopyMemoryIndirectFeaturesKHR - Structure describing indirect copy features supported by an implementation



## [](#_c_specification)C Specification

The `VkPhysicalDeviceCopyMemoryIndirectFeaturesKHR` structure is defined as:

```c++
// Provided by VK_KHR_copy_memory_indirect
typedef struct VkPhysicalDeviceCopyMemoryIndirectFeaturesKHR {
    VkStructureType    sType;
    void*              pNext;
    VkBool32           indirectMemoryCopy;
    VkBool32           indirectMemoryToImageCopy;
} VkPhysicalDeviceCopyMemoryIndirectFeaturesKHR;
```

## [](#_members)Members

This structure describes the following features:

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- []()`indirectMemoryCopy` indicates whether [indirect memory to memory copies](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#indirect-copies) are supported.
- []()`indirectMemoryToImageCopy` indicates whether [indirect memory to image copies](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#indirect-copies) are supported.

## [](#_description)Description

If the `VkPhysicalDeviceCopyMemoryIndirectFeaturesKHR` structure is included in the `pNext` chain of the [VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceFeatures2.html) structure passed to [vkGetPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPhysicalDeviceFeatures2.html), it is filled in to indicate whether each corresponding feature is supported. If the application wishes to use a [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html) with any features described by `VkPhysicalDeviceCopyMemoryIndirectFeaturesKHR`, it **must** add an instance of the structure, with the desired feature members set to `VK_TRUE`, to the `pNext` chain of [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceCreateInfo.html) when creating the [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html).

Valid Usage (Implicit)

- [](#VUID-VkPhysicalDeviceCopyMemoryIndirectFeaturesKHR-sType-sType)VUID-VkPhysicalDeviceCopyMemoryIndirectFeaturesKHR-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_COPY_MEMORY_INDIRECT_FEATURES_KHR`

## [](#_see_also)See Also

[VK\_KHR\_copy\_memory\_indirect](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_copy_memory_indirect.html), [VkBool32](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBool32.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkPhysicalDeviceCopyMemoryIndirectFeaturesKHR).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0
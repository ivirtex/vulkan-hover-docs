# VkPhysicalDeviceZeroInitializeDeviceMemoryFeaturesEXT(3) Manual Page

## Name

VkPhysicalDeviceZeroInitializeDeviceMemoryFeaturesEXT - Structure describing whether the implementation supports cleared allocation functionality



## [](#_c_specification)C Specification

The `VkPhysicalDeviceZeroInitializeDeviceMemoryFeaturesEXT` structure is defined as:

```c++
// Provided by VK_EXT_zero_initialize_device_memory
typedef struct VkPhysicalDeviceZeroInitializeDeviceMemoryFeaturesEXT {
    VkStructureType    sType;
    void*              pNext;
    VkBool32           zeroInitializeDeviceMemory;
} VkPhysicalDeviceZeroInitializeDeviceMemoryFeaturesEXT;
```

## [](#_members)Members

This structure describes the following features:

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- []()`zeroInitializeDeviceMemory` indicates that the implementation supports zeroing memory allocations using a user-specified flag.

## [](#_description)Description

If the `VkPhysicalDeviceZeroInitializeDeviceMemoryFeaturesEXT` structure is included in the `pNext` chain of the [VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceFeatures2.html) structure passed to [vkGetPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPhysicalDeviceFeatures2.html), it is filled in to indicate whether each corresponding feature is supported. If the application wishes to use a [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html) with any features described by `VkPhysicalDeviceZeroInitializeDeviceMemoryFeaturesEXT`, it **must** add an instance of the structure, with the desired feature members set to `VK_TRUE`, to the `pNext` chain of [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceCreateInfo.html) when creating the [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html).

Valid Usage (Implicit)

- [](#VUID-VkPhysicalDeviceZeroInitializeDeviceMemoryFeaturesEXT-sType-sType)VUID-VkPhysicalDeviceZeroInitializeDeviceMemoryFeaturesEXT-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_ZERO_INITIALIZE_DEVICE_MEMORY_FEATURES_EXT`

## [](#_see_also)See Also

[VK\_EXT\_zero\_initialize\_device\_memory](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_zero_initialize_device_memory.html), [VkBool32](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBool32.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkPhysicalDeviceZeroInitializeDeviceMemoryFeaturesEXT).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0
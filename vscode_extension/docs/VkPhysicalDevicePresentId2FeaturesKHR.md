# VkPhysicalDevicePresentId2FeaturesKHR(3) Manual Page

## Name

VkPhysicalDevicePresentId2FeaturesKHR - Structure indicating support for present id 2



## [](#_c_specification)C Specification

The `VkPhysicalDevicePresentId2FeaturesKHR` structure is defined as:

```c++
// Provided by VK_KHR_present_id2
typedef struct VkPhysicalDevicePresentId2FeaturesKHR {
    VkStructureType    sType;
    void*              pNext;
    VkBool32           presentId2;
} VkPhysicalDevicePresentId2FeaturesKHR;
```

## [](#_members)Members

This structure describes the following feature:

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- []()`presentId2` indicates that the implementation supports specifying present ID values in the `VkPresentId2KHR` extension to the `VkPresentInfoKHR` struct.

## [](#_description)Description

If the `VkPhysicalDevicePresentId2FeaturesKHR` structure is included in the `pNext` chain of the [VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceFeatures2.html) structure passed to [vkGetPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPhysicalDeviceFeatures2.html), it is filled in to indicate whether each corresponding feature is supported. If the application wishes to use a [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html) with any features described by `VkPhysicalDevicePresentId2FeaturesKHR`, it **must** add an instance of the structure, with the desired feature members set to `VK_TRUE`, to the `pNext` chain of [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceCreateInfo.html) when creating the [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html).

Valid Usage (Implicit)

- [](#VUID-VkPhysicalDevicePresentId2FeaturesKHR-sType-sType)VUID-VkPhysicalDevicePresentId2FeaturesKHR-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_PRESENT_ID_2_FEATURES_KHR`

## [](#_see_also)See Also

[VK\_KHR\_present\_id2](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_present_id2.html), [VkBool32](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBool32.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkPhysicalDevicePresentId2FeaturesKHR)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0
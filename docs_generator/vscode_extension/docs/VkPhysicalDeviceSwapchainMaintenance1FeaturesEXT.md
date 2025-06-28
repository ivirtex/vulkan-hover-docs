# VkPhysicalDeviceSwapchainMaintenance1FeaturesEXT(3) Manual Page

## Name

VkPhysicalDeviceSwapchainMaintenance1FeaturesEXT - Structure describing whether implementation supports swapchain maintenance1 functionality



## [](#_c_specification)C Specification

The `VkPhysicalDeviceSwapchainMaintenance1FeaturesEXT` structure is defined as:

```c++
// Provided by VK_EXT_swapchain_maintenance1
typedef struct VkPhysicalDeviceSwapchainMaintenance1FeaturesEXT {
    VkStructureType    sType;
    void*              pNext;
    VkBool32           swapchainMaintenance1;
} VkPhysicalDeviceSwapchainMaintenance1FeaturesEXT;
```

## [](#_members)Members

This structure describes the following feature:

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- []()`swapchainMaintenance1` indicates that the implementation supports the following:
  
  - [VkSwapchainPresentFenceInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSwapchainPresentFenceInfoEXT.html), specifying a fence that is signaled when the resources associated with a present operation **can** be safely destroyed.
  - [VkSwapchainPresentModesCreateInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSwapchainPresentModesCreateInfoEXT.html) and [VkSwapchainPresentModeInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSwapchainPresentModeInfoEXT.html), allowing the swapchain to switch present modes without a need for recreation.
  - [VkSwapchainPresentScalingCreateInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSwapchainPresentScalingCreateInfoEXT.html), specifying the scaling behavior of the swapchain in presence of window resizing.
  - The `VK_SWAPCHAIN_CREATE_DEFERRED_MEMORY_ALLOCATION_BIT_EXT` flag, allowing the implementation to defer the allocation of swapchain image memory until first acquisition.
  - [vkReleaseSwapchainImagesEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/vkReleaseSwapchainImagesEXT.html), allowing acquired swapchain images to be released without presenting them.

## [](#_description)Description

If the `VkPhysicalDeviceSwapchainMaintenance1FeaturesEXT` structure is included in the `pNext` chain of the [VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceFeatures2.html) structure passed to [vkGetPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPhysicalDeviceFeatures2.html), it is filled in to indicate whether each corresponding feature is supported. If the application wishes to use a [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html) with any features described by `VkPhysicalDeviceSwapchainMaintenance1FeaturesEXT`, it **must** add an instance of the structure, with the desired feature members set to `VK_TRUE`, to the `pNext` chain of [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceCreateInfo.html) when creating the [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html).

Valid Usage (Implicit)

- [](#VUID-VkPhysicalDeviceSwapchainMaintenance1FeaturesEXT-sType-sType)VUID-VkPhysicalDeviceSwapchainMaintenance1FeaturesEXT-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_SWAPCHAIN_MAINTENANCE_1_FEATURES_EXT`

## [](#_see_also)See Also

[VK\_EXT\_swapchain\_maintenance1](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_swapchain_maintenance1.html), [VkBool32](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBool32.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkPhysicalDeviceSwapchainMaintenance1FeaturesEXT)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0
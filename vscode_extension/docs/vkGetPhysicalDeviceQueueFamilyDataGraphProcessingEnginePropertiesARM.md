# vkGetPhysicalDeviceQueueFamilyDataGraphProcessingEnginePropertiesARM(3) Manual Page

## Name

vkGetPhysicalDeviceQueueFamilyDataGraphProcessingEnginePropertiesARM - Query the properties of a data graph processing engine for a specific queue family of a physical device



## [](#_c_specification)C Specification

To query the properties of a data graph processing engine for a specific queue family of a physical device, call:

```c++
// Provided by VK_ARM_data_graph
void vkGetPhysicalDeviceQueueFamilyDataGraphProcessingEnginePropertiesARM(
    VkPhysicalDevice                            physicalDevice,
    const VkPhysicalDeviceQueueFamilyDataGraphProcessingEngineInfoARM* pQueueFamilyDataGraphProcessingEngineInfo,
    VkQueueFamilyDataGraphProcessingEnginePropertiesARM* pQueueFamilyDataGraphProcessingEngineProperties);
```

## [](#_parameters)Parameters

- `physicalDevice` is the physical device to query.
- `pQueueFamilyDataGraphProcessingEngineInfo` is a pointer to a [VkPhysicalDeviceQueueFamilyDataGraphProcessingEngineInfoARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceQueueFamilyDataGraphProcessingEngineInfoARM.html) structure that specifies the data graph processing engine and queue family to query.
- `pQueueFamilyDataGraphProcessingEngineProperties` is a pointer to a [VkQueueFamilyDataGraphProcessingEnginePropertiesARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkQueueFamilyDataGraphProcessingEnginePropertiesARM.html) structure in which the queries properties are returned.

## [](#_description)Description

Valid Usage (Implicit)

- [](#VUID-vkGetPhysicalDeviceQueueFamilyDataGraphProcessingEnginePropertiesARM-physicalDevice-parameter)VUID-vkGetPhysicalDeviceQueueFamilyDataGraphProcessingEnginePropertiesARM-physicalDevice-parameter  
  `physicalDevice` **must** be a valid [VkPhysicalDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDevice.html) handle
- [](#VUID-vkGetPhysicalDeviceQueueFamilyDataGraphProcessingEnginePropertiesARM-pQueueFamilyDataGraphProcessingEngineInfo-parameter)VUID-vkGetPhysicalDeviceQueueFamilyDataGraphProcessingEnginePropertiesARM-pQueueFamilyDataGraphProcessingEngineInfo-parameter  
  `pQueueFamilyDataGraphProcessingEngineInfo` **must** be a valid pointer to a valid [VkPhysicalDeviceQueueFamilyDataGraphProcessingEngineInfoARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceQueueFamilyDataGraphProcessingEngineInfoARM.html) structure
- [](#VUID-vkGetPhysicalDeviceQueueFamilyDataGraphProcessingEnginePropertiesARM-pQueueFamilyDataGraphProcessingEngineProperties-parameter)VUID-vkGetPhysicalDeviceQueueFamilyDataGraphProcessingEnginePropertiesARM-pQueueFamilyDataGraphProcessingEngineProperties-parameter  
  `pQueueFamilyDataGraphProcessingEngineProperties` **must** be a valid pointer to a [VkQueueFamilyDataGraphProcessingEnginePropertiesARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkQueueFamilyDataGraphProcessingEnginePropertiesARM.html) structure

## [](#_see_also)See Also

[VK\_ARM\_data\_graph](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_ARM_data_graph.html), [VkPhysicalDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDevice.html), [VkPhysicalDeviceQueueFamilyDataGraphProcessingEngineInfoARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceQueueFamilyDataGraphProcessingEngineInfoARM.html), [VkQueueFamilyDataGraphProcessingEnginePropertiesARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkQueueFamilyDataGraphProcessingEnginePropertiesARM.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkGetPhysicalDeviceQueueFamilyDataGraphProcessingEnginePropertiesARM).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0
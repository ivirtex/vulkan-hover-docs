# VkPhysicalDeviceLayeredApiVulkanPropertiesKHR(3) Manual Page

## Name

VkPhysicalDeviceLayeredApiVulkanPropertiesKHR - Structure describing physical device properties of a layered Vulkan implementation underneath the Vulkan physical device



## [](#_c_specification)C Specification

The `VkPhysicalDeviceLayeredApiVulkanPropertiesKHR` structure is defined as:

```c++
// Provided by VK_KHR_maintenance7
typedef struct VkPhysicalDeviceLayeredApiVulkanPropertiesKHR {
    VkStructureType                sType;
    void*                          pNext;
    VkPhysicalDeviceProperties2    properties;
} VkPhysicalDeviceLayeredApiVulkanPropertiesKHR;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `properties` is a [VkPhysicalDeviceProperties2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceProperties2.html) in which properties of the underlying layered Vulkan implementation are returned.

## [](#_description)Description

The implementation **must** zero-fill the contents of `properties.properties.limits` and `properties.properties.sparseProperties`.

Valid Usage

- [](#VUID-VkPhysicalDeviceLayeredApiVulkanPropertiesKHR-pNext-10011)VUID-VkPhysicalDeviceLayeredApiVulkanPropertiesKHR-pNext-10011  
  Only [VkPhysicalDeviceDriverProperties](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceDriverProperties.html) and [VkPhysicalDeviceIDProperties](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceIDProperties.html) are allowed in the `pNext` chain of `properties`

Valid Usage (Implicit)

- [](#VUID-VkPhysicalDeviceLayeredApiVulkanPropertiesKHR-sType-sType)VUID-VkPhysicalDeviceLayeredApiVulkanPropertiesKHR-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_LAYERED_API_VULKAN_PROPERTIES_KHR`

## [](#_see_also)See Also

[VK\_KHR\_maintenance7](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_maintenance7.html), [VkPhysicalDeviceProperties2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceProperties2.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkPhysicalDeviceLayeredApiVulkanPropertiesKHR)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0
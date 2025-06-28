# VkPhysicalDeviceMaintenance9PropertiesKHR(3) Manual Page

## Name

VkPhysicalDeviceMaintenance9PropertiesKHR - Structure describing various implementation-defined properties introduced with VK\_KHR\_maintenance9



## [](#_c_specification)C Specification

The `VkPhysicalDeviceMaintenance9PropertiesKHR` structure is defined as:

```c++
// Provided by VK_KHR_maintenance9
typedef struct VkPhysicalDeviceMaintenance9PropertiesKHR {
    VkStructureType                     sType;
    void*                               pNext;
    VkBool32                            image2DViewOf3DSparse;
    VkDefaultVertexAttributeValueKHR    defaultVertexAttributeValue;
} VkPhysicalDeviceMaintenance9PropertiesKHR;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- []()If the [`image2DViewOf3D`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-image2DViewOf3D) feature is enabled, `image2DViewOf3DSparse` indicates whether the implementation supports binding a slice of a sparse 3D image to a 2D image view.
- []()`defaultVertexAttributeValue` is a [VkDefaultVertexAttributeValueKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDefaultVertexAttributeValueKHR.html) that indicates what value the implementation will return when the vertex shader reads unbound vertex attributes. Unbound attributes are those that have no corresponding [VkVertexInputAttributeDescription](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVertexInputAttributeDescription.html)::`location` defined in the bound graphics pipeline or no corresponding [VkVertexInputAttributeDescription2EXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVertexInputAttributeDescription2EXT.html)::`location` set by the most recent call to [vkCmdSetVertexInputEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdSetVertexInputEXT.html) when the state is dynamic .

## [](#_description)Description

If the `VkPhysicalDeviceMaintenance9PropertiesKHR` structure is included in the `pNext` chain of the [VkPhysicalDeviceProperties2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceProperties2.html) structure passed to [vkGetPhysicalDeviceProperties2](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPhysicalDeviceProperties2.html), it is filled in with each corresponding implementation-dependent property.

Valid Usage (Implicit)

- [](#VUID-VkPhysicalDeviceMaintenance9PropertiesKHR-sType-sType)VUID-VkPhysicalDeviceMaintenance9PropertiesKHR-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_MAINTENANCE_9_PROPERTIES_KHR`

## [](#_see_also)See Also

[VK\_KHR\_maintenance9](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_maintenance9.html), [VkBool32](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBool32.html), [VkDefaultVertexAttributeValueKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDefaultVertexAttributeValueKHR.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkPhysicalDeviceMaintenance9PropertiesKHR)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0
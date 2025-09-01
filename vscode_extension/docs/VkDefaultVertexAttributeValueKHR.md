# VkDefaultVertexAttributeValueKHR(3) Manual Page

## Name

VkDefaultVertexAttributeValueKHR - Values returned for unbound vertex attributes



## [](#_c_specification)C Specification

The possible values returned by the implementation when the vertex shader reads an unbound vertex attribute are:

```c++
// Provided by VK_KHR_maintenance9
typedef enum VkDefaultVertexAttributeValueKHR {
    VK_DEFAULT_VERTEX_ATTRIBUTE_VALUE_ZERO_ZERO_ZERO_ZERO_KHR = 0,
    VK_DEFAULT_VERTEX_ATTRIBUTE_VALUE_ZERO_ZERO_ZERO_ONE_KHR = 1,
} VkDefaultVertexAttributeValueKHR;
```

## [](#_description)Description

- `VK_DEFAULT_VERTEX_ATTRIBUTE_VALUE_ZERO_ZERO_ZERO_ZERO_KHR` - the value read for an unbound vertex attribute is (0,0,0,0).
- `VK_DEFAULT_VERTEX_ATTRIBUTE_VALUE_ZERO_ZERO_ZERO_ONE_KHR` - the value read for an unbound vertex attribute is (0,0,0,1).

## [](#_see_also)See Also

[VK\_KHR\_maintenance9](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_maintenance9.html), [VkPhysicalDeviceMaintenance9PropertiesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceMaintenance9PropertiesKHR.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkDefaultVertexAttributeValueKHR).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0
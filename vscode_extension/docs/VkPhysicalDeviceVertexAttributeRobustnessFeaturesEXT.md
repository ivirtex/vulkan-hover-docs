# VkPhysicalDeviceVertexAttributeRobustnessFeaturesEXT(3) Manual Page

## Name

VkPhysicalDeviceVertexAttributeRobustnessFeaturesEXT - Structure describing whether the vertex attribute robustness feature is supported by an implementation



## [](#_c_specification)C Specification

The `VkPhysicalDeviceVertexAttributeRobustnessFeaturesEXT` structure is defined as:

```c++
// Provided by VK_EXT_vertex_attribute_robustness
typedef struct VkPhysicalDeviceVertexAttributeRobustnessFeaturesEXT {
    VkStructureType    sType;
    void*              pNext;
    VkBool32           vertexAttributeRobustness;
} VkPhysicalDeviceVertexAttributeRobustnessFeaturesEXT;
```

## [](#_members)Members

This structure describes the following feature:

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- []()`vertexAttributeRobustness` indicates that vertex shaders **can** read vertex attribute locations that have no vertex attribute description and the value returned is (0,0,0,0) or (0,0,0,1).

## [](#_description)Description

If the `VkPhysicalDeviceVertexAttributeRobustnessFeaturesEXT` structure is included in the `pNext` chain of the [VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceFeatures2.html) structure passed to [vkGetPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPhysicalDeviceFeatures2.html), it is filled in to indicate whether each corresponding feature is supported. If the application wishes to use a [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html) with any features described by `VkPhysicalDeviceVertexAttributeRobustnessFeaturesEXT`, it **must** add an instance of the structure, with the desired feature members set to `VK_TRUE`, to the `pNext` chain of [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceCreateInfo.html) when creating the [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html).

Valid Usage (Implicit)

- [](#VUID-VkPhysicalDeviceVertexAttributeRobustnessFeaturesEXT-sType-sType)VUID-VkPhysicalDeviceVertexAttributeRobustnessFeaturesEXT-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_VERTEX_ATTRIBUTE_ROBUSTNESS_FEATURES_EXT`

## [](#_see_also)See Also

[VK\_EXT\_vertex\_attribute\_robustness](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_vertex_attribute_robustness.html), [VkBool32](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBool32.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkPhysicalDeviceVertexAttributeRobustnessFeaturesEXT)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0
# VkPhysicalDeviceVertexAttributeDivisorPropertiesEXT(3) Manual Page

## Name

VkPhysicalDeviceVertexAttributeDivisorPropertiesEXT - Structure describing max value of vertex attribute divisor that can be supported by an implementation



## [](#_c_specification)C Specification

The `VkPhysicalDeviceVertexAttributeDivisorPropertiesEXT` structure is defined as:

```c++
// Provided by VK_EXT_vertex_attribute_divisor
typedef struct VkPhysicalDeviceVertexAttributeDivisorPropertiesEXT {
    VkStructureType    sType;
    void*              pNext;
    uint32_t           maxVertexAttribDivisor;
} VkPhysicalDeviceVertexAttributeDivisorPropertiesEXT;
```

## [](#_members)Members

## [](#_description)Description

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- []()`maxVertexAttribDivisor` is the maximum value of the number of instances that will repeat the value of vertex attribute data when instanced rendering is enabled.

If the `VkPhysicalDeviceVertexAttributeDivisorPropertiesEXT` structure is included in the `pNext` chain of the [VkPhysicalDeviceProperties2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceProperties2.html) structure passed to [vkGetPhysicalDeviceProperties2](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPhysicalDeviceProperties2.html), it is filled in with each corresponding implementation-dependent property.

Valid Usage (Implicit)

- [](#VUID-VkPhysicalDeviceVertexAttributeDivisorPropertiesEXT-sType-sType)VUID-VkPhysicalDeviceVertexAttributeDivisorPropertiesEXT-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_VERTEX_ATTRIBUTE_DIVISOR_PROPERTIES_EXT`

## [](#_see_also)See Also

[VK\_EXT\_vertex\_attribute\_divisor](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_vertex_attribute_divisor.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkPhysicalDeviceVertexAttributeDivisorPropertiesEXT)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0
# VkPhysicalDeviceVertexAttributeDivisorProperties(3) Manual Page

## Name

VkPhysicalDeviceVertexAttributeDivisorProperties - Structure describing max value of vertex attribute divisor that can be supported by an implementation



## [](#_c_specification)C Specification

The `VkPhysicalDeviceVertexAttributeDivisorProperties` structure is defined as:

```c++
// Provided by VK_VERSION_1_4
typedef struct VkPhysicalDeviceVertexAttributeDivisorProperties {
    VkStructureType    sType;
    void*              pNext;
    uint32_t           maxVertexAttribDivisor;
    VkBool32           supportsNonZeroFirstInstance;
} VkPhysicalDeviceVertexAttributeDivisorProperties;
```

or the equivalent

```c++
// Provided by VK_KHR_vertex_attribute_divisor
typedef VkPhysicalDeviceVertexAttributeDivisorProperties VkPhysicalDeviceVertexAttributeDivisorPropertiesKHR;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.

## [](#_description)Description

- []()`maxVertexAttribDivisor` is the maximum value of the number of instances that will repeat the value of vertex attribute data when instanced rendering is enabled.
- []()`supportsNonZeroFirstInstance` specifies whether a non-zero value for the `firstInstance` parameter of [drawing commands](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#drawing) is supported when [VkVertexInputBindingDivisorDescription](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVertexInputBindingDivisorDescription.html)::`divisor` is not `1`.

If the `VkPhysicalDeviceVertexAttributeDivisorProperties` structure is included in the `pNext` chain of the [VkPhysicalDeviceProperties2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceProperties2.html) structure passed to [vkGetPhysicalDeviceProperties2](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPhysicalDeviceProperties2.html), it is filled in with each corresponding implementation-dependent property.

Valid Usage (Implicit)

- [](#VUID-VkPhysicalDeviceVertexAttributeDivisorProperties-sType-sType)VUID-VkPhysicalDeviceVertexAttributeDivisorProperties-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_VERTEX_ATTRIBUTE_DIVISOR_PROPERTIES`

## [](#_see_also)See Also

[VK\_KHR\_vertex\_attribute\_divisor](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_vertex_attribute_divisor.html), [VK\_VERSION\_1\_4](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_4.html), [VkBool32](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBool32.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkPhysicalDeviceVertexAttributeDivisorProperties)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0
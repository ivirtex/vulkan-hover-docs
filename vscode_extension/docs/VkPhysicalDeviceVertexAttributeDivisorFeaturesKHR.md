# VkPhysicalDeviceVertexAttributeDivisorFeatures(3) Manual Page

## Name

VkPhysicalDeviceVertexAttributeDivisorFeatures - Structure describing if fetching of vertex attribute may be repeated for instanced rendering



## [](#_c_specification)C Specification

The `VkPhysicalDeviceVertexAttributeDivisorFeatures` structure is defined as:

```c++
// Provided by VK_VERSION_1_4
typedef struct VkPhysicalDeviceVertexAttributeDivisorFeatures {
    VkStructureType    sType;
    void*              pNext;
    VkBool32           vertexAttributeInstanceRateDivisor;
    VkBool32           vertexAttributeInstanceRateZeroDivisor;
} VkPhysicalDeviceVertexAttributeDivisorFeatures;
```

or the equivalent

```c++
// Provided by VK_KHR_vertex_attribute_divisor
typedef VkPhysicalDeviceVertexAttributeDivisorFeatures VkPhysicalDeviceVertexAttributeDivisorFeaturesKHR;
```

or the equivalent

```c++
// Provided by VK_EXT_vertex_attribute_divisor
typedef VkPhysicalDeviceVertexAttributeDivisorFeatures VkPhysicalDeviceVertexAttributeDivisorFeaturesEXT;
```

## [](#_members)Members

This structure describes the following features:

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.

## [](#_description)Description

- []()`vertexAttributeInstanceRateDivisor` specifies whether vertex attribute fetching may be repeated in the case of instanced rendering.
- []()`vertexAttributeInstanceRateZeroDivisor` specifies whether a zero value for [VkVertexInputBindingDivisorDescriptionEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVertexInputBindingDivisorDescriptionEXT.html)::`divisor` is supported.

If the `VkPhysicalDeviceVertexAttributeDivisorFeatures` structure is included in the `pNext` chain of the [VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceFeatures2.html) structure passed to [vkGetPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPhysicalDeviceFeatures2.html), it is filled in to indicate whether each corresponding feature is supported. If the application wishes to use a [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html) with any features described by `VkPhysicalDeviceVertexAttributeDivisorFeatures`, it **must** add an instance of the structure, with the desired feature members set to `VK_TRUE`, to the `pNext` chain of [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceCreateInfo.html) when creating the [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html).

Valid Usage (Implicit)

- [](#VUID-VkPhysicalDeviceVertexAttributeDivisorFeatures-sType-sType)VUID-VkPhysicalDeviceVertexAttributeDivisorFeatures-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_VERTEX_ATTRIBUTE_DIVISOR_FEATURES`

## [](#_see_also)See Also

[VK\_EXT\_vertex\_attribute\_divisor](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_vertex_attribute_divisor.html), [VK\_KHR\_vertex\_attribute\_divisor](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_vertex_attribute_divisor.html), [VK\_VERSION\_1\_4](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_4.html), [VkBool32](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBool32.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkPhysicalDeviceVertexAttributeDivisorFeatures).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0
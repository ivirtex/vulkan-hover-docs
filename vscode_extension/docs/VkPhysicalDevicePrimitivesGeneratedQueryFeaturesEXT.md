# VkPhysicalDevicePrimitivesGeneratedQueryFeaturesEXT(3) Manual Page

## Name

VkPhysicalDevicePrimitivesGeneratedQueryFeaturesEXT - Structure describing support for primitives generated query



## [](#_c_specification)C Specification

The `VkPhysicalDevicePrimitivesGeneratedQueryFeaturesEXT` structure is defined as:

```c++
// Provided by VK_EXT_primitives_generated_query
typedef struct VkPhysicalDevicePrimitivesGeneratedQueryFeaturesEXT {
    VkStructureType    sType;
    void*              pNext;
    VkBool32           primitivesGeneratedQuery;
    VkBool32           primitivesGeneratedQueryWithRasterizerDiscard;
    VkBool32           primitivesGeneratedQueryWithNonZeroStreams;
} VkPhysicalDevicePrimitivesGeneratedQueryFeaturesEXT;
```

## [](#_members)Members

This structure describes the following features:

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- []()`primitivesGeneratedQuery` indicates whether the implementation supports the `VK_QUERY_TYPE_PRIMITIVES_GENERATED_EXT` query type.
- []()`primitivesGeneratedQueryWithRasterizerDiscard` indicates whether the implementation supports this query when [rasterization discard](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#primsrast-discard) is enabled.
- []()`primitivesGeneratedQueryWithNonZeroStreams` indicates whether the implementation supports this query with a non-zero index in [vkCmdBeginQueryIndexedEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdBeginQueryIndexedEXT.html).

## [](#_description)Description

If the `VkPhysicalDevicePrimitivesGeneratedQueryFeaturesEXT` structure is included in the `pNext` chain of the [VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceFeatures2.html) structure passed to [vkGetPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPhysicalDeviceFeatures2.html), it is filled in to indicate whether each corresponding feature is supported. If the application wishes to use a [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html) with any features described by `VkPhysicalDevicePrimitivesGeneratedQueryFeaturesEXT`, it **must** add an instance of the structure, with the desired feature members set to `VK_TRUE`, to the `pNext` chain of [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceCreateInfo.html) when creating the [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html).

Valid Usage (Implicit)

- [](#VUID-VkPhysicalDevicePrimitivesGeneratedQueryFeaturesEXT-sType-sType)VUID-VkPhysicalDevicePrimitivesGeneratedQueryFeaturesEXT-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_PRIMITIVES_GENERATED_QUERY_FEATURES_EXT`

## [](#_see_also)See Also

[VK\_EXT\_primitives\_generated\_query](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_primitives_generated_query.html), [VkBool32](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBool32.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkPhysicalDevicePrimitivesGeneratedQueryFeaturesEXT)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0
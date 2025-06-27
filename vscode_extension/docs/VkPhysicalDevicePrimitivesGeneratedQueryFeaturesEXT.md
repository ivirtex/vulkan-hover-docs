# VkPhysicalDevicePrimitivesGeneratedQueryFeaturesEXT(3) Manual Page

## Name

VkPhysicalDevicePrimitivesGeneratedQueryFeaturesEXT - Structure
describing support for primitives generated query



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkPhysicalDevicePrimitivesGeneratedQueryFeaturesEXT` structure is
defined as:

``` c
// Provided by VK_EXT_primitives_generated_query
typedef struct VkPhysicalDevicePrimitivesGeneratedQueryFeaturesEXT {
    VkStructureType    sType;
    void*              pNext;
    VkBool32           primitivesGeneratedQuery;
    VkBool32           primitivesGeneratedQueryWithRasterizerDiscard;
    VkBool32           primitivesGeneratedQueryWithNonZeroStreams;
} VkPhysicalDevicePrimitivesGeneratedQueryFeaturesEXT;
```

## <a href="#_members" class="anchor"></a>Members

This structure describes the following features:

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- <span id="features-primitivesGeneratedQuery"></span>
  `primitivesGeneratedQuery` indicates whether the implementation
  supports the `VK_QUERY_TYPE_PRIMITIVES_GENERATED_EXT` query type.

- <span id="features-primitivesGeneratedQueryWithRasterizerDiscard"></span>
  `primitivesGeneratedQueryWithRasterizerDiscard` indicates whether the
  implementation supports this query when <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#primsrast-discard"
  target="_blank" rel="noopener">rasterization discard</a> is enabled.

- <span id="features-primitivesGeneratedQueryWithNonZeroStreams"></span>
  `primitivesGeneratedQueryWithNonZeroStreams` indicates whether the
  implementation supports this query with a non-zero index in
  [vkCmdBeginQueryIndexedEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdBeginQueryIndexedEXT.html).

## <a href="#_description" class="anchor"></a>Description

If the `VkPhysicalDevicePrimitivesGeneratedQueryFeaturesEXT` structure
is included in the `pNext` chain of the
[VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceFeatures2.html) structure
passed to
[vkGetPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceFeatures2.html), it is
filled in to indicate whether each corresponding feature is supported.
`VkPhysicalDevicePrimitivesGeneratedQueryFeaturesEXT` **can** also be
used in the `pNext` chain of
[VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceCreateInfo.html) to selectively enable
these features.

Valid Usage (Implicit)

- <a
  href="#VUID-VkPhysicalDevicePrimitivesGeneratedQueryFeaturesEXT-sType-sType"
  id="VUID-VkPhysicalDevicePrimitivesGeneratedQueryFeaturesEXT-sType-sType"></a>
  VUID-VkPhysicalDevicePrimitivesGeneratedQueryFeaturesEXT-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_PRIMITIVES_GENERATED_QUERY_FEATURES_EXT`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_EXT_primitives_generated_query](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_primitives_generated_query.html),
[VkBool32](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBool32.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkPhysicalDevicePrimitivesGeneratedQueryFeaturesEXT"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
